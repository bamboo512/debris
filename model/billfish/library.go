package billfish

// Library represents the 'library' table.
type Library struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Version  int    `json:"version"`
	Platform string `gorm:"type:text" json:"platform"`
}

func (Library) TableName() string {
	return "library"
}
