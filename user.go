package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`

	// 用户发送的消息
	Messages []Message `gorm:"foreignKey:SenderID"`
}

func register(ctx *gin.Context) {
	u := &User{}

	err := ctx.ShouldBind(u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	db.Create(u)

	ctx.JSON(http.StatusOK, gin.H{
		"id": u.ID,
	})
}

func login(ctx *gin.Context) {
	u := &User{}

	err := ctx.ShouldBind(u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	find := db.Where("username = ?", u.Username).Take(u)

	if find.Error != nil {
		if errors.Is(find.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "user not found",
			})

			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": find.Error.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": u.ID,
	})
}
