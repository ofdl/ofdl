package ofdl

import (
	"context"

	"github.com/ofdl/ofdl/model/query"
	"github.com/ofdl/ofdl/ofdl/data"
	"github.com/ofdl/ofdl/ofdl/downloader"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"github.com/shurcooL/graphql"
	"gorm.io/gorm"
)

type OFDL struct {
	ctx context.Context

	Stash *graphql.Client
	DB    *gorm.DB
	Query *query.Query

	OnlyFans   onlyfans.OnlyFansAPI
	Data       data.OFDLDataAPI
	Downloader downloader.Downloader
}

func NewOFDL(
	ctx context.Context,
	stash *graphql.Client,
	db *gorm.DB,
	query *query.Query,
	of onlyfans.OnlyFansAPI,
	d data.OFDLDataAPI,
	dl downloader.Downloader,
) *OFDL {
	return &OFDL{
		ctx:        ctx,
		Stash:      stash,
		DB:         db,
		Query:      query,
		OnlyFans:   of,
		Data:       d,
		Downloader: dl,
	}
}
