package billfish

// Type represents the 'bf_type' table.
type Type struct {
	ID         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Tid        int64  `json:"tid"`
	Name       string `gorm:"type:text" json:"name"`
	Gid        int64  `json:"gid"`
	CanAnalyze int    `json:"can_analyze"`
}

func (Type) TableName() string {
	return "bf_type"
}
