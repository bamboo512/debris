package pixcall

import "gorm.io/gorm"

// BoardEntry represents the many-to-many relationship between boards and entries.
type BoardEntry struct {
	BoardID   int64 `gorm:"type:bigint;primaryKey;not null" json:"board_id"`
	EntryID   int64 `gorm:"type:bigint;primaryKey;not null" json:"entry_id"`
	EntryKind int   `gorm:"type:int;not null" json:"entry_kind"`
}

/* rename table name */
func (boardEntry *BoardEntry) TableName(tx *gorm.DB) string {
	return "board_entries"
}
