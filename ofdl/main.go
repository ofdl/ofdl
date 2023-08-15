package ofdl

import (
	"context"

	"github.com/ofdl/ofdl/ent"
	"github.com/ofdl/ofdl/ofdl/downloader"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"github.com/shurcooL/graphql"
)

type OFDL struct {
	ctx context.Context

	Stash      *graphql.Client
	Ent        *ent.Client
	OnlyFans   onlyfans.OnlyFansAPI
	Downloader downloader.Downloader
}

func NewOFDL(
	ctx context.Context,
	stash *graphql.Client,
	ent *ent.Client,
	of onlyfans.OnlyFansAPI,
	dl downloader.Downloader,
) *OFDL {
	return &OFDL{
		ctx:        ctx,
		Stash:      stash,
		Ent:        ent,
		OnlyFans:   of,
		Downloader: dl,
	}
}
