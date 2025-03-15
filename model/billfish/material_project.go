package billfish

// MaterialProject represents the 'bf_material_project' table.
type MaterialProject struct {
	ID         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	FileID     int64  `gorm:"not null;unique" json:"file_id"`
	PreviewTid int64  `json:"preview_tid"`
	BfExtend   string `gorm:"type:text" json:"bf_extend"`
}

func (MaterialProject) TableName() string {
	return "bf_material_project"
}
