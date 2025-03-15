package pixcall

// Media represents the metadata for a media entry.
type Media struct {
	EntryID  int64  `gorm:"type:bigint;primaryKey;not null" json:"entry_id"`
	Metadata string `gorm:"type:text;not null" json:"metadata"`
}
