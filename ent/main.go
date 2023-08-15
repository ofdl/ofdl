package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema

import (
	"database/sql"

	"github.com/spf13/viper"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/glebarez/go-sqlite"
)

func NewEntClient() (*Client, error) {
	db, err := sql.Open("sqlite", viper.GetString("database"))
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB("sqlite3", db)
	return NewClient(Driver(drv)), nil
}
