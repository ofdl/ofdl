package downloader

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/siku2/arigo"
	"github.com/spf13/viper"
)

type Aria2Downloader struct {
	ctx  context.Context
	rpc  *arigo.Client
	root string
}

func NewAria2Downloader(ctx context.Context, address, secret, root string) (Downloader, error) {
	ag, err := arigo.Dial(address, secret)
	if err != nil {
		return nil, err
	}

	return &Aria2Downloader{
		ctx:  ctx,
		rpc:  &ag,
		root: root,
	}, nil
}

// DownloadMany implements Downloader.
func (d *Aria2Downloader) DownloadMany(mm []Downloadable) <-chan error {
	d1 := make(chan error)

	go func() {
		sem := make(chan struct{}, 5)
		for _, m := range mm {
			sem <- struct{}{}
			go func(m Downloadable) {
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
func (d *Aria2Downloader) DownloadOne(m Downloadable) (<-chan float64, <-chan error) {
	progress := make(chan float64)
	done := make(chan error)

	go func() {
		defer close(progress)

		if m.GetFull() == "" {
			done <- m.MarkDownloaded(d.ctx)
			return
		}

		dir := path.Join(d.root, m.Directory())
		if viper.GetString("downloads.aria2.platform") == "windows" {
			dir = fmt.Sprintf("%s%s", d.root, strings.ReplaceAll(m.Directory(), "/", `\`))
		}

		_, err := d.rpc.AddURI([]string{m.GetFull()}, &arigo.Options{
			Out: m.Filename(),
			Dir: dir,
		})
		if err != nil {

			done <- err
			return
		}

		done <- m.MarkDownloaded(d.ctx)
	}()

	return progress, done
}
