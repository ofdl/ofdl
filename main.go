//go:generate go run ./gen
package main

import (
	"fmt"
	"os"

	"github.com/defval/di"
	"github.com/ofdl/ofdl/cmd"
	"github.com/ofdl/ofdl/cmd/gui"
	"github.com/ofdl/ofdl/ent"
	"github.com/ofdl/ofdl/ofdl"
	"github.com/ofdl/ofdl/ofdl/downloader"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
)

func main() {
	app, err := di.New(
		di.Provide(onlyfans.NewOnlyFans),
		di.Provide(downloader.NewDownloader),
		di.Provide(ofdl.NewStash),
		di.Provide(ofdl.NewOFDL),

		di.Provide(gui.NewSubsGui),
		di.Provide(ent.NewEntClient),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Execute(app)
}
