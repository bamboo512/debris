package billfish

// MaterialColor represents the 'bf_material_color' table.
type MaterialColor struct {
	ID      int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	FileID  int64 `gorm:"not null" json:"file_id"`
	Idx     int   `json:"idx"`
	Percent int   `json:"percent"`
	Color   int   `json:"color"`
	R       int   `json:"r"`
	G       int   `json:"g"`
	B       int   `json:"b"`
}

func (MaterialColor) TableName() string {
	return "bf_material_color"
}
