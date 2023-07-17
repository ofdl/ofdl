package ofdl

import (
	"path/filepath"

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

func (o *OFDL) GetMissingMessageMedia(limit int) ([]model.MessageMedia, error) {
	q := o.Query.MessageMedia
	return q.Preload(
		q.Message,
		q.Message.Subscription,
	).Missing(limit)
}

func (o *OFDL) DownloadMedia(m model.DownloadableMedia) (*arigo.GID, error) {
	var gid *arigo.GID

	if m.URL() != "" {
		g, err := o.Aria2.AddURI([]string{m.URL()}, &arigo.Options{
			Out: m.Filename(),
			Dir: filepath.Join(viper.GetString("aria2.root"), m.Directory()),
		})
		if err != nil {
			return nil, err
		}

		gid = &g
	}

	if err := m.MarkDownloaded(o.DB); err != nil {
		return nil, err
	}

	return gid, nil
}
