//go:generate go run ./gen
package main

import (
	"fmt"
	"os"

	"github.com/defval/di"
	"github.com/ofdl/ofdl/cmd"
	"github.com/ofdl/ofdl/cmd/gui"
	"github.com/ofdl/ofdl/model"
	"github.com/ofdl/ofdl/model/query"
	"github.com/ofdl/ofdl/ofdl"
	"github.com/ofdl/ofdl/ofdl/data"
	"github.com/ofdl/ofdl/ofdl/downloader"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"gorm.io/gorm"
)

func main() {
	app, err := di.New(
		di.Provide(onlyfans.NewOnlyFans),
		di.Provide(data.NewOFDLData),
		di.Provide(downloader.NewDownloader),
		di.Provide(ofdl.NewStash),
		di.Provide(ofdl.NewOFDL),
		di.Provide(model.NewDB),
		di.Provide(func(db *gorm.DB) *query.Query {
			return query.Use(db)
		}),

		di.Provide(gui.NewSubsGui),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Execute(app)
}
