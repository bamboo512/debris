package pixcall

import (
	"time"

	"gorm.io/gorm"
)

// Entry represents a general entry (e.g., file, folder, note).
type Entry struct {
	ID          int64  `gorm:"type:bigint;primaryKey;autoIncrement" json:"id"`
	Revision    int    `gorm:"not null;default:1" json:"revision"`
	ParentID    int64  `gorm:"type:bigint;not null" json:"parent_id"`
	Name        string `gorm:"type:text;not null" json:"name"`
	Kind        int    `gorm:"type:int;not null" json:"kind"`
	Description string `gorm:"type:text" json:"description"`
	Link        string `gorm:"type:varchar(255)" json:"link"`
	/* Tag Ids
	Multiple tag ids are seperated by  `|`.
	For instance: `426606124943284224|426606229880576000`
	*/
	Tags string `gorm:"type:text" json:"tags"`
	/* User rating of a file */
	Rating      int       `gorm:"type:int" json:"rating"`
	Size        int64     `gorm:"type:bigint;not null;default:0" json:"size"`
	ContentType string    `gorm:"type:varchar(100);not null;default:''" json:"content_type"`
	ContentHash string    `gorm:"type:varchar(64)" json:"content_hash"`
	Metadata    string    `gorm:"type:text" json:"metadata"`
	Ranking     int64     `gorm:"type:bigint;default:0" json:"ranking"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	SourcePath  string    `gorm:"type:text" json:"source_path"`
	NamePinyin  string    `gorm:"type:text" json:"name_pinyin"`
	Status      int64     `gorm:"type:bigint;not null;default:0" json:"status"`
	Inode       int64     `gorm:"type:bigint" json:"inode"`
	Mtime       int64     `gorm:"type:bigint" json:"mtime"`
	IsHidden    int       `gorm:"type:smallint;not null;default:0" json:"is_hidden"`
	IsDeleted   int       `gorm:"type:int;not null;default:0" json:"is_deleted"`
}

// BeforeCreate will set a default value for the created_at,updated_at fields.
func (entity *Entry) BeforeCreate(tx *gorm.DB) (err error) {
	if entity.CreatedAt.IsZero() {
		entity.CreatedAt = time.Now()
	}

	if entity.UpdatedAt.IsZero() {
		entity.UpdatedAt = time.Now()
	}
	return
}

func (entity *Entry) BeforeUpdate(tx *gorm.DB) (err error) {
	entity.UpdatedAt = time.Now()
	return
}
