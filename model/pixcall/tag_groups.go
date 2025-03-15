package pixcall

// TagGroups represents a group of tags.
type TagGroups struct {
	ID       int64  `gorm:"type:bigint increment;primaryKey;autoIncrement;not null" json:"id"`
	Revision int    `gorm:"type:int;not null;default:1" json:"revision"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Pinyin   string `gorm:"type:varchar(100);not null" json:"pinyin"`
	Color    string `gorm:"type:varchar(16)" json:"color"`
	Ranking  int64  `gorm:"type:bigint;not null;default:1" json:"ranking"`
}
