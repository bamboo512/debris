package billfish

// MaterialUserData represents the 'bf_material_userdata' table.
type MaterialUserData struct {
	ID              int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	FileID          int64  `gorm:"not null;unique"  json:"file_id"`
	CommentsSummary string `gorm:"type:text" json:"comments_summary"`
	CommentsCount   int    `json:"comments_count"`
	CommentsDetail  string `gorm:"type:text"  json:"comments_detail"`
	Note            string `gorm:"type:text" json:"note"`
	Origin          string `gorm:"type:text"  json:"origin"`
	Score           int    `json:"score"`
	Rotation        int    `json:"rotation"`
	Hflip           int    `json:"hflip"`
	Vflip           int    `json:"vflip"`
	CoverTid        string `gorm:"type:text" json:"cover_tid"`
}

func (MaterialUserData) TableName() string {
	return "bf_material_userdata"
}
