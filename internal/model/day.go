package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/NasSilverBullet/jft/internal/util"
	"github.com/jinzhu/gorm"
)

type Day struct {
	Date  time.Time
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
		b, e, err := util.GetDayEndAndBeginning(strconv.Itoa(eachDay.Year()) + "/" + strconv.Itoa(int(eachDay.Month())) + "/" + strconv.Itoa(eachDay.Day()))

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
			Date:  *b,
			Plans: ps,
		})
		*eachDay = eachDay.AddDate(0, 0, 1)
	}

	return days, err
}

func (p Day) String() string {
	const layout = "2006/01/02 (Mon)"
	const format = "%v  >>>  %s"
	state := "O"
	if len(p.Plans) == 0 {
		state = "X"
	}
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 59, time.Local)
	switch {
	case p.Date.After(today):
		state = "-"
	case len(p.Plans) == 0:
		state = "X"
	default:
		state = "O"
	}
	return fmt.Sprintf(format, p.Date.Format(layout), state)
}
