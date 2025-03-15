package billfish

// TagJoinFile represents the 'bf_tag_join_file' table.
type TagJoinFile struct {
	ID     int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	FileID int64 `json:"file_id"`
	TagID  int64 `json:"tag_id"`
	Born   int64 `json:"born"`
}

func (TagJoinFile) TableName() string {
	return "bf_tag_join_file"
}
