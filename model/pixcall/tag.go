package pixcall

import "gorm.io/gorm"

// Tag represents a tag.
type Tag struct {
	ID       int64  `gorm:"type:bigint;primaryKey;autoIncrement;not null" json:"id"`
	Revision int    `gorm:"not null;default:1" json:"revision"`
	GroupID  int64  `gorm:"type:bigint;default:null" json:"group_id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Pinyin   string `gorm:"type:varchar(255);not null" json:"pinyin"`
	Category string `gorm:"type:varchar(1);not null" json:"category"`

	// 迁移的时候需要用到的外部标签 id
	ExternalID string `json:"external_id"`
}

/* rename table name */
func (boardEntry *Tag) TableName(tx *gorm.DB) string {
	return "tags"
}
