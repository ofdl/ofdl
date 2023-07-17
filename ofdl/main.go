package ofdl

import (
	"context"

	"github.com/ofdl/ofdl/model"
	"github.com/ofdl/ofdl/model/query"
	"github.com/ofdl/ofdl/ofdl/data"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"github.com/shurcooL/graphql"
	"github.com/siku2/arigo"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type OFDL struct {
	ctx context.Context

	Aria2 *arigo.Client
	Stash *graphql.Client
	DB    *gorm.DB
	Query *query.Query

	OnlyFans onlyfans.OnlyFansAPI
	data     data.OFDLDataAPI
}

func NewOFDL() (*OFDL, error) {
	ag, err := arigo.Dial(viper.GetString("aria2.address"), viper.GetString("aria2.secret"))
	if err != nil {
		return nil, err
	}

	sc := graphql.NewClient(viper.GetString("stash.address"), nil)

	db, err := gorm.Open(sqlite.Open(viper.GetString("database")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	// db = db.Debug()
	db.AutoMigrate(&model.Subscription{}, &model.Post{}, &model.Media{}, &model.Message{}, &model.MessageMedia{})

	of, err := onlyfans.NewOnlyFans()
	if err != nil {
		return nil, err
	}

	return &OFDL{
		ctx:      context.Background(),
		Aria2:    &ag,
		Stash:    sc,
		DB:       db,
		Query:    query.Use(db),
		OnlyFans: of,
	}, nil
}

func (o *OFDL) Data() data.OFDLDataAPI {
	return data.NewOFDLData()
}
