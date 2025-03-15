package billfish

// HiddenTagJoinFile represents the 'bf_hidden_tag_join_file' table.
type HiddenTagJoinFile struct {
	ID     int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	FileID int64 `json:"file_id"`
	TagID  int64 `json:"tag_id"`
	Born   int64 `json:"born"`
}

func (HiddenTagJoinFile) TableName() string {
	return "bf_hidden_tag_join_file"
}
