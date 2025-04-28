package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func setupTestServer() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default() // 使用Default而不是New，以获取默认中间件

	// 确保数据库已初始化
	if db == nil {
		db = initDB()
	}

	// 设置数据库中间件
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	addRouter(router)
	return router
}

// 测试发送消息
func TestSendMessage(t *testing.T) {
	router := setupTestServer()

	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("收到请求: %s %s", r.Method, r.URL.Path)
		t.Logf("请求头: %+v", r.Header)
		router.ServeHTTP(w, r)
	}))
	defer server.Close()

	// 将 HTTP URL 转换为 WebSocket URL
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/message"
	t.Logf("尝试连接到 WebSocket URL: %s", wsURL)

	// 设置WebSocket dialer
	dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  5 * time.Second,
		EnableCompression: true,
	}

	// 连接 WebSocket
	ws, resp, err := dialer.Dial(wsURL, nil)
	if err != nil {
		if resp != nil {
			t.Fatalf("无法建立WebSocket连接: %v\n响应状态码: %d\n响应头: %+v\n响应体: %s",
				err,
				resp.StatusCode,
				resp.Header,
				resp.Status)
		} else {
			t.Fatalf("无法建立WebSocket连接: %v\n没有收到响应", err)
		}
	}
	defer ws.Close()

	// 发送测试消息
	testMsg := Message{
		Content:  "测试消息",
		SenderID: 1,
	}

	err = ws.WriteJSON(testMsg)
	if err != nil {
		t.Fatalf("发送WebSocket消息失败: %v", err)
	}

	// 等待消息处理完成
	time.Sleep(100 * time.Millisecond)

	// 验证消息是否被保存到数据库
	var savedMsg Message
	err = db.Where("content = ?", testMsg.Content).First(&savedMsg).Error
	assert.NilError(t, err, "消息应该已保存到数据库")
	assert.Equal(t, savedMsg.Content, testMsg.Content)
	assert.Equal(t, savedMsg.SenderID, testMsg.SenderID)
}

// 测试修改消息
func TestUpdateMessage(t *testing.T) {
	// 先创建一条消息用于测试修改
	initialMsg := Message{
		Content:  "initial message",
		SenderID: 1,
	}

	result := db.Create(&initialMsg)
	if result.Error != nil {
		t.Fatalf("failed to create initial message: %v", result.Error)
	}

	// 模拟修改消息请求数据
	updateBody := fmt.Sprintf(`{"ID": %d, "Content": "updated message"}`, initialMsg.ID)
	req, err := http.NewRequest("PUT", "/message", strings.NewReader(updateBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 执行请求
	w := httptest.NewRecorder()
	router := setupTestServer()
	router.ServeHTTP(w, req)

	// 检查响应状态码
	assert.Equal(t, w.Code, http.StatusOK)

	// 验证消息是否被更新
	var updatedMsg Message
	result = db.First(&updatedMsg, initialMsg.ID)
	if result.Error != nil {
		t.Fatalf("failed to fetch updated message: %v", result.Error)
	}

	assert.Equal(t, updatedMsg.Content, "updated message")
}
