package pixcall

// Folders represents a folder.
type Folders struct {
	EntryID      int64  `gorm:"type:bigint;primaryKey;not null" json:"entry_id"`
	Revision     int    `gorm:"type:int;not null;default:1" json:"revision"`
	FileCount    int    `gorm:"type:int;not null;default:0" json:"file_count"`
	FolderCount  int    `gorm:"type:int;not null;default:0" json:"folder_count"`
	FileSize     int64  `gorm:"type:bigint;not null;default:0" json:"file_size"`
	Ranking      int64  `gorm:"type:bigint;not null;default:0" json:"ranking"`
	Layout       string `gorm:"type:varchar(64)" json:"layout"`
	SortBy       string `gorm:"type:varchar(64)" json:"sort_by"`
	SortOrder    string `gorm:"type:varchar(16)" json:"sort_order"`
	Icon         string `gorm:"type:varchar(32)" json:"icon"`
	IconColor    string `gorm:"type:varchar(16)" json:"icon_color"`
	CoverID      int64  `gorm:"type:bigint" json:"cover_id"`
	FixedCoverID int64  `gorm:"type:bigint" json:"fixed_cover_id"`
}
