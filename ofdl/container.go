package ofdl

import (
	"context"

	"github.com/defval/di"
	"github.com/ofdl/ofdl/model"
	"github.com/ofdl/ofdl/model/query"
	"github.com/ofdl/ofdl/ofdl/data"
	"github.com/ofdl/ofdl/ofdl/downloader"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"gorm.io/gorm"
)

type Container struct {
	*di.Container
	ctx context.Context
}

func NewContainer(ctx context.Context) (*Container, error) {
	container, err := di.New(
		di.Provide(func() context.Context { return ctx }),
		di.Provide(NewStash),
		di.Provide(model.NewDB),
		di.Provide(func(db *gorm.DB) *query.Query {
			return query.Use(db)
		}),
		di.Provide(onlyfans.NewOnlyFans),
		di.Provide(data.NewOFDLData),
		di.Provide(downloader.NewDownloader),
		di.Provide(NewOFDL),
	)
	if err != nil {
		return nil, err
	}

	return &Container{
		Container: container,
		ctx:       ctx,
	}, nil
}
