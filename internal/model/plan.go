package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Plan struct {
	gorm.Model
	Start            *time.Time
	End              *time.Time
	ShortDescription string `gorm:"size:255"`
	Description      string `gorm:"size:255"`
}
