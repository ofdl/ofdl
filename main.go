//go:generate go run ./gen
package main

import (
	"fmt"
	"os"

	"github.com/defval/di"
	"github.com/ofdl/ofdl/cmd"
	"github.com/ofdl/ofdl/cmd/gui"
	"github.com/ofdl/ofdl/ent"
	"github.com/ofdl/ofdl/ofdl/downloader"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
)

func main() {
	// di.SetTracer(di.StdTracer{})
	app, err := di.New(
		di.Provide(ent.NewEntClient),
		di.Provide(onlyfans.NewOnlyFans),
		di.Provide(downloader.NewDownloader),

		di.Provide(gui.NewSubsGui),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Execute(app)
}
