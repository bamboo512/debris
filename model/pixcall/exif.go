package pixcall

// Exif represents the EXIF data for a media entry.
type Exif struct {
	EntryID  int64  `gorm:"type:bigint;primaryKey;not null" json:"entry_id"`
	Metadata string `gorm:"type:text;not null" json:"metadata"`
}

func (exif Exif) TableName() string {
	return "exif"
}
