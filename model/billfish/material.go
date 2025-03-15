package billfish

// Material represents the 'bf_material_v2' table
type Material struct {
	FileID    int64  `gorm:"primaryKey" json:"file_id"`
	Status    int    `json:"status"`
	ThumbTid  int64  `json:"thumb_tid"`
	ImageTid  int64  `json:"image_tid"`
	W         int    `json:"w"`
	H         int    `json:"h"`
	IsRecycle int    `json:"is_recycle"`
	Colors    string `gorm:"type:text" json:"colors"`
}

func (Material) TableName() string {
	return "bf_material_v2"
}
