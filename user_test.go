package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gotest.tools/assert"
)

var (
	router *gin.Engine
	db     *gorm.DB
)

func TestMain(m *testing.M) {
	// 初始化 Gin 引擎
	router = gin.Default()

	// 模拟数据库
	var err error
	db, err = gorm.Open(sqlite.Open("websocket_test.db?_journal_mode=WAL"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移用户模型
	db.AutoMigrate(&User{}, &Message{})

	// 设置数据库到 Gin 上下文
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// 注册路由
	router.POST("/register", register)
	router.POST("/login", login)
	router.POST("/message", message)
	router.PUT("/message", updateMessage)

	// 运行所有测试用例
	code := m.Run()

	// 退出程序
	os.Exit(code)
}

func TestRegister(t *testing.T) {
	// 模拟请求数据
	body := `{"id": 1, "username": "testuser", "password": "testpassword"}`
	req, err := http.NewRequest("POST", "/register", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 执行请求
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 检查响应状态码
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestLogin(t *testing.T) {
	// 先注册一个用户用于测试登录
	registerBody := `{"id": 1, "username": "testuser", "password": "testpassword"}`
	registerReq, err := http.NewRequest("POST", "/register", strings.NewReader(registerBody))
	if err != nil {
		t.Fatal(err)
	}
	registerReq.Header.Set("Content-Type", "application/json")
	registerW := httptest.NewRecorder()
	router.ServeHTTP(registerW, registerReq)

	// 模拟登录请求
	loginBody := `{"username": "testuser", "password": "testpassword"}`
	loginReq, err := http.NewRequest("POST", "/login", strings.NewReader(loginBody))
	if err != nil {
		t.Fatal(err)
	}
	loginReq.Header.Set("Content-Type", "application/json")

	// 执行登录请求
	loginW := httptest.NewRecorder()
	router.ServeHTTP(loginW, loginReq)

	// 检查响应状态码
	assert.Equal(t, loginW.Code, http.StatusOK)
}
