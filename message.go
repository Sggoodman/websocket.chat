package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net/http"
)

type Message struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Content string `gorm:"type:text;not null"`
	// 发送者ID (外键)
	SenderID uint `gorm:"not null"`
	// 发送者 (属于关系)
	Sender User `gorm:"foreignKey:SenderID"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 开发环境下允许所有来源
	},
}

func message(c *gin.Context) {
	// 获取数据库连接
	db, ok := c.Get("db")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接未找到"})
		return
	}

	// 升级HTTP连接为WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return // WebSocket升级失败，upgrader会自动发送错误响应
	}
	defer conn.Close()

	// 注册客户端
	clients[conn] = true
	defer delete(clients, conn)

	// 处理接收到的消息
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// 记录意外的错误
				println("websocket error:", err.Error())
			}
			break
		}

		// 保存消息到数据库
		if err := db.(*gorm.DB).Create(&msg).Error; err != nil {
			conn.WriteJSON(gin.H{"error": "保存消息失败"})
			continue
		}

		// 广播消息给所有客户端
		broadcast <- msg
	}
}

func updateMessage(ctx *gin.Context) {
	m := &Message{}

	err := ctx.ShouldBind(m)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	db := ctx.MustGet("db").(*gorm.DB)

	newMessage := &Message{}
	find := db.Where("id = ?", m.ID).Take(newMessage)

	if find.Error != nil {
		if errors.Is(find.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "message not found",
			})

			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": find.Error.Error(),
		})

		return
	}

	newMessage.Content = m.Content

	db.Model(newMessage).Select("content").Updates(newMessage)

	// 广播消息更新给所有 websocket 客户端
	broadcast <- *newMessage

	ctx.JSON(http.StatusOK, gin.H{})
}
