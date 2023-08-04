package downloader

import "github.com/ofdl/ofdl/model"

type Downloader interface {
	DownloadOne(model.DownloadableMedia) (<-chan float64, <-chan error)
	DownloadMany([]model.DownloadableMedia) <-chan error
}
