package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/NasSilverBullet/jft/internal/util"
	"github.com/jinzhu/gorm"
)

type Month struct {
	Begin time.Time
	Days  []Day
}

func FindMonths(db *gorm.DB, yearStr string) ([]Month, error) {
	begin, _, err := util.GetYearEndAndBeginning(yearStr)
	if err != nil {
		return nil, err
	}
	months := []Month{}
	eachDay := *begin
	year := begin.Year()
	for eachDay.Year() == year {
		days, err := FindDays(db, strconv.Itoa(eachDay.Year())+"/"+strconv.Itoa(int(eachDay.Month())))
		if err != nil {
			return nil, err
		}
		months = append(months, Month{
			Begin: eachDay,
			Days:  days,
		})
		eachDay = eachDay.AddDate(0, 1, 0)
	}
	return months, err
}

func (m *Month) count() int {
	var count int
	for _, day := range m.Days {
		if len(day.Plans) != 0 {
			count++
		}
	}
	return count
}

func (m *Month) ratio() float64 {
	return float64(m.count()) / float64(len(m.Days))
}

func (m Month) String() string {
	const layout = "2006/01"
	if m.Begin.After(time.Now()) {
		return fmt.Sprintf("%v  >>>  --------", m.Begin.Format(layout))
	}
	return fmt.Sprintf("%v  >>>  %d(%.2f%%)", m.Begin.Format(layout), m.count(), m.ratio())
}
