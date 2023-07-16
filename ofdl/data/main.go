package data

import (
	"github.com/ofdl/ofdl/model"
	"github.com/ofdl/ofdl/model/query"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type OFDLDataAPI interface {
	GetEnabledSubscriptions() ([]model.Subscription, error)

	SaveSubscription(onlyfans.Subscription) error
	SaveMediaPost(onlyfans.MediaPost) error
}

func NewOFDLData() OFDLDataAPI {
	db, err := gorm.Open(sqlite.Open(viper.GetString("database")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	// db = db.Debug()
	db.AutoMigrate(&model.Subscription{}, &model.Post{}, &model.Media{})

	return &GormOFDLData{
		DB:    db,
		Query: query.Use(db),
	}
}
