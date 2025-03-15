package billfish

// MaterialImage represents the 'bf_material_image' table.
type MaterialImage struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	FileID   int64  `gorm:"not null;unique" json:"file_id"`
	Rotation int    `json:"rotation"`
	Hflip    int    `json:"hflip"`
	Vflip    int    `json:"vflip"`
	BfExtend string `gorm:"type:text" json:"bf_extend"`
}

func (MaterialImage) TableName() string {
	return "bf_material_image"
}
