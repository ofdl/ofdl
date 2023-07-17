package model

import (
	"fmt"
	"net/url"
	"path/filepath"
	"time"

	"gorm.io/gorm"
)

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

var _ OrganizableMedia = &Media{}

// GetDate implements OrganizableMedia.
func (m *Media) GetDate() (time.Time, error) {
	return time.Parse(time.RFC3339, m.Post.PostedAt)
}

// GetPerformerID implements OrganizableMedia.
func (m *Media) GetPerformerID() string {
	return m.Post.Subscription.StashID
}

// GetTitle implements OrganizableMedia.
func (m *Media) GetTitle() string {
	return m.Post.Text
}

// GetType implements OrganizableMedia.
func (m *Media) GetType() string {
	return m.Type
}

// MarkOrganized implements OrganizableMedia.
func (m *Media) MarkOrganized(db *gorm.DB) error {
	now := time.Now()
	m.OrganizedAt = &now
	return db.Save(m).Error
}

// Organize implements OrganizableMedia.
func (m *Media) Organize(db *gorm.DB, id string) error {
	now := time.Now()
	m.OrganizedAt = &now
	m.StashID = id
	return db.Save(m).Error
}

func (m Media) Directory() string {
	return fmt.Sprintf("/%s/posts/%d", m.Post.Subscription.Username, m.Post.ID)
}

func (m Media) Filename() string {
	u, _ := url.Parse(m.URL())
	return filepath.Base(u.Path)
}

func (m Media) URL() string {
	return m.Full
}

func (m *Media) MarkDownloaded(db *gorm.DB) error {
	now := time.Now()
	m.DownloadedAt = &now
	return db.Save(m).Error
}
