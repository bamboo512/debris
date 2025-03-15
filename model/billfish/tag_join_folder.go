package billfish

// TagJoinFolder represents the 'bf_tag_join_folder' table.
type TagJoinFolder struct {
	ID       int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	FolderID int64 `json:"folder_id"`
	TagID    int64 `json:"tag_id"`
}

func (TagJoinFolder) TableName() string {
	return "bf_tag_join_folder"
}
