package model

import (
	"fmt"
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

func (s Subscription) FilterValue() string {
	return fmt.Sprintf("%s %s", s.Name, s.Username)
}

func (s Subscription) Title() string {
	enabled := " "
	if s.Enabled {
		enabled = "X"
	}
	return fmt.Sprintf("[%s] %s", enabled, s.Name)
}

func (s Subscription) Description() string {
	return fmt.Sprintf("    %s", s.Username)
}
