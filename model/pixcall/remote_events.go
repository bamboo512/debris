package pixcall

import (
	"time"

	"gorm.io/gorm"
)

// RemoteEvents represents a remote event log.
type RemoteEvents struct {
	ID        uint      `gorm:"type:bigint;primaryKey" json:"id"`
	Revision  int       `gorm:"not null;default:1" json:"revision"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	Data      string    `gorm:"type:text;not null" json:"data"`
}

// BeforeCreate will set a default value for the created_at field.
func (entity *RemoteEvents) BeforeCreate(tx *gorm.DB) (err error) {
	if entity.CreatedAt.IsZero() {
		entity.CreatedAt = time.Now()
	}
	return
}
