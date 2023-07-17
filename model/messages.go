package model

import (
	"fmt"
	"net/url"
	"path/filepath"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Subscription   Subscription `json:"subscription"`
	SubscriptionID uint         `json:"subscription_id"`

	Text     string         `json:"text"`
	PostedAt string         `json:"posted_at"`
	Medias   []MessageMedia `json:"medias"`
}

type MessageMedia struct {
	gorm.Model
	Message   Message `json:"message"`
	MessageID uint    `json:"message_id"`

	Type string `json:"type"`
	Src  string `json:"src"`

	DownloadedAt *time.Time
	StashID      string `json:"stash_id"`
	OrganizedAt  *time.Time
}

var _ OrganizableMedia = &MessageMedia{}

// GetDate implements OrganizableMedia.
func (mm *MessageMedia) GetDate() (time.Time, error) {
	return time.Parse(time.RFC3339, mm.Message.PostedAt)
}

// GetPerformerID implements OrganizableMedia.
func (mm *MessageMedia) GetPerformerID() string {
	return mm.Message.Subscription.StashID
}

// GetTitle implements OrganizableMedia.
func (mm *MessageMedia) GetTitle() string {
	return mm.Message.Text
}

// GetType implements OrganizableMedia.
func (mm *MessageMedia) GetType() string {
	return mm.Type
}

// MarkOrganized implements OrganizableMedia.
func (mm *MessageMedia) MarkOrganized(db *gorm.DB) error {
	now := time.Now()
	mm.OrganizedAt = &now
	return db.Save(mm).Error
}

// Organize implements OrganizableMedia.
func (mm *MessageMedia) Organize(db *gorm.DB, id string) error {
	now := time.Now()
	mm.OrganizedAt = &now
	mm.StashID = id
	return db.Save(mm).Error
}

func (mm MessageMedia) Directory() string {
	return fmt.Sprintf("/%s/messages/%d", mm.Message.Subscription.Username, mm.Message.ID)
}

func (mm MessageMedia) Filename() string {
	u, _ := url.Parse(mm.URL())
	return filepath.Base(u.Path)
}

func (mm MessageMedia) URL() string {
	return mm.Src
}

func (mm *MessageMedia) MarkDownloaded(db *gorm.DB) error {
	now := time.Now()
	mm.DownloadedAt = &now
	return db.Save(mm).Error
}
