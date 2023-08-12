package data

import (
	"github.com/ofdl/ofdl/model"
	"github.com/ofdl/ofdl/model/query"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"gorm.io/gorm"
)

type OFDLDataAPI interface {
	GetSubscriptions() ([]*model.Subscription, error)
	GetEnabledSubscriptions() ([]*model.Subscription, error)

	SaveSubscription(onlyfans.Subscription) error
	SaveMediaPost(onlyfans.MediaPost) error
	SaveMessage(uint, onlyfans.Message) error
}

func NewOFDLData(db *gorm.DB) OFDLDataAPI {
	return &GormOFDLData{
		DB:    db,
		Query: query.Use(db),
	}
}
