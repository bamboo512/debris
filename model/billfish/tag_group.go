package billfish

// TagGroup represents the 'bf_tag_group' table.
type TagGroup struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:text" json:"name"`
	Color int    `json:"color"`
	Born  int64  `json:"born"`
	Pid   int64  `json:"pid"`
}

func (TagGroup) TableName() string {
	return "bf_tag_group"
}
