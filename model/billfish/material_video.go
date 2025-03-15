package billfish

// MaterialVideo represents the 'bf_material_video' table.
type MaterialVideo struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	FileID   int64  `gorm:"not null;unique" json:"file_id"`
	Duration int    `json:"duration"`
	Rotation int    `json:"rotation"`
	BfExtend string `gorm:"type:text" json:"bf_extend"`
}

func (MaterialVideo) TableName() string {
	return "bf_material_video"
}
