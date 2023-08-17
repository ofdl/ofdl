package downloader

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
)

type Downloadable interface {
	Directory() string
	Filename() string
	GetFull() string
	MarkDownloaded(context.Context) error
}

type Downloader interface {
	DownloadOne(Downloadable) (<-chan float64, <-chan error)
	DownloadMany([]Downloadable) <-chan error
}

func NewDownloader() (dl Downloader, err error) {
	switch viper.GetString("downloads.downloader") {
	case "local":
		dl, err = NewLocalDownloader(viper.GetString("downloads.local.root"))
	case "aria2":
		dl, err = NewAria2Downloader(
			viper.GetString("downloads.aria2.address"),
			viper.GetString("downloads.aria2.secret"),
			viper.GetString("downloads.aria2.root"),
		)
	default:
		return nil, fmt.Errorf("unknown downloader: %s", viper.GetString("downloads.downloader"))
	}
	return
}

func ToDownloadableSlice[T Downloadable](v []T) []Downloadable {
	r := make([]Downloadable, len(v))
	for i := range v {
		r[i] = v[i]
	}
	return r
}
