package downloader

import (
	"path/filepath"

	"github.com/ofdl/ofdl/model"
	"github.com/siku2/arigo"
)

type Aria2Downloader struct {
	rpc  *arigo.Client
	root string
}

func NewAria2Downloader(address, secret, root string) (*Aria2Downloader, error) {
	ag, err := arigo.Dial(address, secret)
	if err != nil {
		return nil, err
	}

	return &Aria2Downloader{
		rpc: &ag,
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

				if m.URL() == "" {
					d1 <- nil
					return
				}

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

		_, err := d.rpc.AddURI([]string{m.URL()}, &arigo.Options{
			Out: m.Filename(),
			Dir: filepath.Join(d.root, m.Directory()),
		})

		done <- err
	}()

	return progress, done
}
