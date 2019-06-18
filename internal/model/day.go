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

func FindDays(db *gorm.DB, monthStr string) ([]Day, error) {
	begin, end, err := util.GetMonthEndAndBeginning(monthStr)
	if err != nil {
		return nil, err
	}
	plans := []Plan{}
	db.Where("start BETWEEN ? AND ?", begin, end).Find(&plans)

	days := []Day{}
	eachDay := begin
	month := begin.Month()
	for month == eachDay.Month() {
		b, e, err := util.GetDayEndAndBeginning(begin.String())

		if err != nil {
			return nil, err
		}
		ps := []Plan{}
		for _, plan := range plans {
			if b.Before(plan.Start.Local()) && e.After(plan.Start.Local()) {
				ps = append(ps, plan)
			}
		}
		days = append(days, Day{
			Date:  b,
			Plans: ps,
		})
	}

	return days, err
}
