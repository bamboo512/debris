package billfish

// TagJoinGroup represents the 'bf_tag_join_group' table.
type TagJoinGroup struct {
	ID    int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Gid   int64 `json:"gid"`
	TagID int64 `json:"tag_id"`
	Born  int64 `json:"born"`
}

func (TagJoinGroup) TableName() string {
	return "bf_tag_join_group"
}
