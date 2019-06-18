package model

import (
	"time"

	"github.com/NasSilverBullet/jft/internal/util"
	"github.com/jinzhu/gorm"
)

type Day struct {
	Date  *time.Time
	Plans []Plan
}

func FindDays(db *gorm.DB, month string) ([]Day, error) {
	begin, end, err := util.GetMonthEndAndBeginning(month)
	if err != nil {
		return nil, err
	}
	ps := []Plan{}
	db.Where("start BETWEEN ? AND ?", begin, end).Find(&ps)
	return nil, nil
}
