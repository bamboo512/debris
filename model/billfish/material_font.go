package billfish

// MaterialFont represents the 'bf_material_font' table.
type MaterialFont struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	FileID   int64  `gorm:"not null;unique" json:"file_id"`
	BfExtend string `gorm:"type:text" json:"bf_extend"`
}

func (MaterialFont) TableName() string {
	return "bf_material_font"
}
