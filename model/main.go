package model

import (
	"time"

	"gorm.io/gen"
	"gorm.io/gorm"
)

type StorableMedia interface {
	Directory() string
	Filename() string
	URL() string
}

type DownloadableMedia interface {
	StorableMedia
	MarkDownloaded(*gorm.DB) error
}

type OrganizableMedia interface {
	StorableMedia
	Organize(db *gorm.DB, id string) error
	MarkOrganized(db *gorm.DB) error
	GetType() string
	GetTitle() string
	GetPerformerID() string
	GetDate() (time.Time, error)
}

type DownloadableLookup interface {
	// select * from @@table WHERE downloaded_at IS NULL LIMIT @limit
	Missing(limit int) ([]*gen.T, error)
}

type OrganizableLookup interface {
	// select * from @@table WHERE organized_at IS NULL LIMIT @limit
	Unorganized(limit int) ([]gen.T, error)

	// select * from @@table WHERE stash_id = $stashID
	FindByStashID(stashID string) (gen.T, error)
}

type EnableableLookup interface {
	// select * from @@table WHERE enabled = true
	GetEnabled() ([]*gen.T, error)

	// update @@table set enabled = true where id = @id
	Enable(id uint) error

	// update @@table set enabled = false where id = @id
	Disable(id uint) error
}
