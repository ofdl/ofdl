package ofdl

import (
	"context"
	"fmt"

	"github.com/ofdl/ofdl/model"
	"github.com/ofdl/ofdl/model/query"
	"github.com/ofdl/ofdl/ofdl/data"
	"github.com/ofdl/ofdl/ofdl/downloader"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"github.com/shurcooL/graphql"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type OFDL struct {
	ctx context.Context

	Stash *graphql.Client
	DB    *gorm.DB
	Query *query.Query

	OnlyFans   onlyfans.OnlyFansAPI
	Data       data.OFDLDataAPI
	Downloader downloader.Downloader
}

func NewOFDL() (*OFDL, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}

	of, err := onlyfans.NewOnlyFans()
	if err != nil {
		return nil, err
	}

	dl, err := getDownloader(db)
	if err != nil {
		return nil, err
	}

	return &OFDL{
		ctx:        context.Background(),
		Stash:      graphql.NewClient(viper.GetString("stash.address"), nil),
		DB:         db,
		Query:      query.Use(db),
		OnlyFans:   of,
		Data:       data.NewOFDLData(db),
		Downloader: dl,
	}, nil
}

func openDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(viper.GetString("database")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	// db = db.Debug()
	db.AutoMigrate(&model.Subscription{}, &model.Post{}, &model.Media{}, &model.Message{}, &model.MessageMedia{})

	return db, nil
}

func getDownloader(db *gorm.DB) (dl downloader.Downloader, err error) {
	switch viper.GetString("downloads.downloader") {
	case "local":
		dl, err = downloader.NewLocalDownloader(db, viper.GetString("downloads.local.root"))
	case "aria2":
		dl, err = downloader.NewAria2Downloader(
			db,
			viper.GetString("downloads.aria2.address"),
			viper.GetString("downloads.aria2.secret"),
			viper.GetString("downloads.aria2.root"),
		)
	default:
		return nil, fmt.Errorf("unsupported downloader: %s", viper.GetString("downloads.downloader"))
	}

	return
}
