package ofdl

import (
	"path/filepath"
	"time"

	"github.com/ofdl/ofdl/model"
	"github.com/siku2/arigo"
	"github.com/spf13/viper"
)

func (o *OFDL) GetMissingMedia(limit int) ([]model.Media, error) {
	q := o.Query.Media
	return q.Preload(
		q.Post,
		q.Post.Subscription,
	).Missing(limit)
}

func (o *OFDL) DownloadMedia(m model.Media) (*arigo.GID, error) {
	var gid *arigo.GID

	if m.Full != "" {
		g, err := o.Aria2.AddURI([]string{m.Full}, &arigo.Options{
			Out: m.Filename(),
			Dir: filepath.Join(viper.GetString("aria2.root"), m.Directory()),
		})
		if err != nil {
			return nil, err
		}

		gid = &g
	}

	now := time.Now()
	m.DownloadedAt = &now
	if err := o.DB.Save(&m).Error; err != nil {
		return nil, err
	}

	return gid, nil
}
