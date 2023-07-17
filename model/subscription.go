package model

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model

	Avatar   string `json:"avatar"`
	Header   string `json:"header"`
	Name     string `json:"name"`
	Username string `json:"username"`

	HeadMarker  string `json:"head_marker"`
	StashID     string `json:"stash_id"`
	OrganizedAt *time.Time
	Enabled     bool `json:"enabled" gorm:"default:true"`

	Posts []Post `json:"posts"`
}
