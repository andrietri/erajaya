package schema

import (
	"time"
)

type Base struct {
	Id        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"type:datetime;notNull;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;notNull;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
