package data

import (
	"errors"

	"github.com/ofdl/ofdl/model"
	"github.com/ofdl/ofdl/model/query"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"gorm.io/gorm"
)

type GormOFDLData struct {
	DB    *gorm.DB
	Query *query.Query
}

var _ OFDLDataAPI = &GormOFDLData{}

func (o *GormOFDLData) GetEnabledSubscriptions() ([]model.Subscription, error) {
	return o.Query.Subscription.GetEnabled()
}

func (o *GormOFDLData) SaveSubscription(v onlyfans.Subscription) error {
	var sub model.Subscription
	if err := o.DB.First(&sub, uint(v.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sub = model.Subscription{
				Model: gorm.Model{
					ID: uint(v.ID),
				},
			}
		} else {
			return err
		}
	}

	sub.Avatar = v.Avatar
	sub.Header = v.Header
	sub.Name = v.Name
	sub.Username = v.Username

	if err := o.DB.Save(&sub).Error; err != nil {
		return err
	}

	return nil
}

func (o *GormOFDLData) SaveMediaPost(v onlyfans.MediaPost) error {
	var mp model.Post
	if err := o.DB.First(&mp, uint(v.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			mp = model.Post{
				Model: gorm.Model{
					ID: uint(v.ID),
				},
			}
		} else {
			return err
		}
	}

	mp.SubscriptionID = uint(v.Author.ID)
	mp.Text = v.Text
	mp.PostedAt = v.PostedAt

	// only saves during a create operation
	for _, m := range v.Media {
		mp.Medias = append(mp.Medias, model.Media{
			Model: gorm.Model{
				ID: uint(m.ID),
			},
			Type: m.Type,
			Full: m.Full,
		})
	}

	if err := o.DB.Save(&mp).Error; err != nil {
		return err
	}

	return nil
}
