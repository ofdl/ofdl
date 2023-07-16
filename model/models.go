package model

import (
	"fmt"
	"net/url"
	"path/filepath"
	"time"

	"gorm.io/gen"
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

type Post struct {
	gorm.Model
	Subscription   Subscription `json:"subscription"`
	SubscriptionID uint         `json:"subscription_id"`

	Text     string `json:"text"`
	PostedAt string `json:"posted_at"`

	Medias []Media `json:"medias"`
}

type Media struct {
	gorm.Model
	Post   Post `json:"post"`
	PostID uint `json:"post_id"`

	Type string `json:"type"`
	Full string `json:"full"`

	DownloadedAt *time.Time
	StashID      string `json:"stash_id"`
	OrganizedAt  *time.Time
}

func (m Media) Directory() string {
	return fmt.Sprintf("/%s/posts/%d", m.Post.Subscription.Username, m.Post.ID)
}

func (m Media) Filename() string {
	u, _ := url.Parse(m.Full)
	return filepath.Base(u.Path)
}

type DownloadableLookup interface {
	// select * from @@table WHERE downloaded_at IS NULL LIMIT @limit
	Missing(limit int) ([]gen.T, error)
}

type OrganizableLookup interface {
	// select * from @@table WHERE organized_at IS NULL LIMIT @limit
	Unorganized(limit int) ([]gen.T, error)

	// select * from @@table WHERE stash_id = $stashID
	FindByStashID(stashID string) (gen.T, error)
}

type EnableableLookup interface {
	// select * from @@table WHERE enabled = true
	GetEnabled() ([]gen.T, error)
}
