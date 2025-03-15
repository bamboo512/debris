package pixcall

// Kvs represents a key-value pair.
type Kvs struct {
	K string `gorm:"type:varchar(100);primaryKey;not null" json:"k"`
	V string `gorm:"type:text;not null" json:"v"`
}
