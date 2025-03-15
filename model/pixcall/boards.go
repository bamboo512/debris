package pixcall

import (
	"time"

	"gorm.io/gorm"
)

// Boards represents a board.
type Board struct {
	ID          int64     `gorm:"type:bigint;primaryKey" json:"id"`
	Revision    int       `gorm:"type:int;not null;default:1" json:"revision"`
	ParentID    int64     `gorm:"type:bigint" json:"parent_id"`
	Name        string    `gorm:"type:text;not null" json:"name"`
	NamePinyin  string    `gorm:"type:text" json:"name_pinyin"`
	Filters     string    `gorm:"type:text;not null" json:"filters"`
	FileCount   int       `gorm:"type:int;not null;default:0" json:"file_count"`
	FolderCount int       `gorm:"type:int;not null;default:0" json:"folder_count"`
	FileSize    int64     `gorm:"type:bigint;not null;default:0" json:"file_size"`
	Ranking     int64     `gorm:"type:bigint;not null;default:1" json:"ranking"`
	Layout      string    `gorm:"type:varchar(64)" json:"layout"`
	SortBy      string    `gorm:"type:varchar(64)" json:"sort_by"`
	SortOrder   string    `gorm:"type:varchar(16)" json:"sort_order"`
	Icon        string    `gorm:"type:varchar(32)" json:"icon"`
	IconColor   string    `gorm:"type:varchar(16)" json:"icon_color"`
	CreatedAt   time.Time `gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

/* rename table name */
func (board *Board) TableName(tx *gorm.DB) string {
	return "boards"
}

// BeforeCreate will set a default value for the created_at field.
func (entity *Board) BeforeCreate(tx *gorm.DB) (err error) {
	if entity.CreatedAt.IsZero() {
		entity.CreatedAt = time.Now()
	}
	return
}
