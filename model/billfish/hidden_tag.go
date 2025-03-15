package billfish

// HiddenTag represents the 'bf_hidden_tag' table.
type HiddenTag struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:text" json:"name"`
	Type  int    `json:"type"`
	Icon  int    `json:"icon"`
	Color int    `json:"color"`
	Born  int64  `json:"born"`
}

func (HiddenTag) TableName() string {
	return "bf_hidden_tag"
}
