package model

import (
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(viper.GetString("database")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	// db = db.Debug()
	db.AutoMigrate(&Subscription{}, &Post{}, &Media{}, &Message{}, &MessageMedia{})

	return db, nil
}
