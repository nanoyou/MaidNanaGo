package model

import (
	"time"

	"github.com/nanoyou/MaidNanaGo/util/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("打开数据库失败")
		panic(err)
	}
	db.Logger = new(logger.VoidLogger)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
}

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
