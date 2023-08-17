package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert --template ./template ./schema

import (
	"context"
	"database/sql"

	"github.com/spf13/viper"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/glebarez/go-sqlite"
)

func NewEntClient() (*Client, error) {
	db, err := sql.Open("sqlite", viper.GetString("database")+"?_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB("sqlite3", db)
	c := NewClient(Driver(drv))

	if err := c.Schema.Create(context.TODO()); err != nil {
		return nil, err
	}

	return c, nil
}
