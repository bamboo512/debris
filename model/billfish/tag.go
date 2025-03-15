package billfish

// Tag represents the 'bf_tag_v2' table.
type Tag struct {
	ID    int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string  `gorm:"type:text" json:"name"`
	Pid   int     `json:"pid"`
	Seq   float64 `json:"seq"`
	Icon  int     `json:"icon"`
	Color int     `json:"color"`
	Born  int64   `json:"born"`

	// 迁移的时候需要用到的外部标签 id
	ExternalID string `json:"external_id"`
}

func (Tag) TableName() string {
	return "bf_tag_v2"
}
