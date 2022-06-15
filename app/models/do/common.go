package do

import (
	"time"
)

type BasicModel struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PageCondition struct {
	Limit  int
	Offset int
}
