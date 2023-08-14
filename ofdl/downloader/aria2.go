package downloader

import (
	"fmt"
	"path"
	"strings"

	"github.com/ofdl/ofdl/model"
	"github.com/siku2/arigo"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Aria2Downloader struct {
	db   *gorm.DB
	rpc  *arigo.Client
	root string
}

func NewAria2Downloader(db *gorm.DB, address, secret, root string) (*Aria2Downloader, error) {
	ag, err := arigo.Dial(address, secret)
	if err != nil {
		return nil, err
	}

	return &Aria2Downloader{
		db:   db,
		rpc:  &ag,
		root: root,
	}, nil
}

var _ Downloader = &Aria2Downloader{}

// DownloadMany implements Downloader.
func (d *Aria2Downloader) DownloadMany(mm []model.DownloadableMedia) <-chan error {
	d1 := make(chan error)

	go func() {
		sem := make(chan struct{}, 5)
		for _, m := range mm {
			sem <- struct{}{}
			go func(m model.DownloadableMedia) {
				defer func() { <-sem }()

				p, done := d.DownloadOne(m)
				for {
					select {
					case <-p:
					case err := <-done:
						d1 <- err
						return
					}
				}
			}(m)
		}
		for i := 0; i < cap(sem); i++ {
			sem <- struct{}{}
		}
		close(d1)
	}()

	return d1
}

// DownloadOne implements Downloader.
func (d *Aria2Downloader) DownloadOne(m model.DownloadableMedia) (<-chan float64, <-chan error) {
	progress := make(chan float64)
	done := make(chan error)

	go func() {
		defer close(progress)

		if m.URL() == "" {
			done <- m.MarkDownloaded(d.db)
			return
		}

		dir := path.Join(d.root, m.Directory())
		if viper.GetString("downloads.aria2.platform") == "windows" {
			dir = fmt.Sprintf("%s%s", d.root, strings.ReplaceAll(m.Directory(), "/", `\`))
		}

		_, err := d.rpc.AddURI([]string{m.URL()}, &arigo.Options{
			Out: m.Filename(),
			Dir: dir,
		})
		if err != nil {

			done <- err
			return
		}

		done <- m.MarkDownloaded(d.db)
	}()

	return progress, done
}
