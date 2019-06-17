package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Day struct {
	Date  *time.Time
	Plans []Plan
}

func FindDays(db *gorm.DB, month string) ([]Day, error) {
	return nil, nil
}
