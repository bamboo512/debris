package billfish

// MaterialAudio represents the 'bf_material_audio' table.
type MaterialAudio struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	FileID   int64  `gorm:"not null;unique" json:"file_id"`
	Duration int    `json:"duration"`
	BitRate  int    `json:"bit_rate"`
	Bpm      int    `json:"bpm"`
	BfExtend string `gorm:"type:text" json:"bf_extend"`
}

func (MaterialAudio) TableName() string {
	return "bf_material_audio"
}
