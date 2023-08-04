package ofdl

import (
	"github.com/ofdl/ofdl/model"
)

func (o *OFDL) GetMissingMedia(limit int) ([]model.DownloadableMedia, error) {
	q := o.Query.Media
	return ToDownloadableMedia(
		q.Preload(
			q.Post,
			q.Post.Subscription,
		).Missing(limit),
	)
}

func (o *OFDL) GetMissingMessageMedia(limit int) ([]model.DownloadableMedia, error) {
	q := o.Query.MessageMedia
	return ToDownloadableMedia(
		q.Preload(
			q.Message,
			q.Message.Subscription,
		).Missing(limit),
	)
}

func ToDownloadableMedia[T model.DownloadableMedia](v []T, e error) ([]model.DownloadableMedia, error) {
	r := make([]model.DownloadableMedia, len(v))
	for i := range v {
		r[i] = v[i]
	}
	return r, e
}
