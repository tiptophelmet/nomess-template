package doc

import (
	"time"
)

// TODO: implement document orm
type Model struct {
	ID        uint `docorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
