package main

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("websocket.db?_journal_mode=WAL"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	// 初始化时执行迁移
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to migrate")
	}

	return db
}
