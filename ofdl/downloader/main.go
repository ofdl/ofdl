package downloader

import (
	"fmt"

	"github.com/ofdl/ofdl/model"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Downloader interface {
	DownloadOne(model.DownloadableMedia) (<-chan float64, <-chan error)
	DownloadMany([]model.DownloadableMedia) <-chan error
}

func NewDownloader(db *gorm.DB) (dl Downloader, err error) {
	switch viper.GetString("downloads.downloader") {
	case "local":
		dl, err = NewLocalDownloader(db, viper.GetString("downloads.local.root"))
	case "aria2":
		dl, err = NewAria2Downloader(
			db,
			viper.GetString("downloads.aria2.address"),
			viper.GetString("downloads.aria2.secret"),
			viper.GetString("downloads.aria2.root"),
		)
	default:
		return nil, fmt.Errorf("unknown downloader: %s", viper.GetString("downloads.downloader"))
	}
	return
}
