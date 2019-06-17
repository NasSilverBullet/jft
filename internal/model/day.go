package model

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Day struct {
	Date  *time.Time
	Plans []Plan
}

func FindDays(db *gorm.DB, month string) ([]Day, error) {
	begin, end, err := getMonthEndAndBeginning(month)
	if err != nil {
		return nil, err
	}
	ps := []Plan{}
	db.Where("start BETWEEN ? AND ?", begin, end).Find(&ps)
	return nil, nil
}

func getMonthEndAndBeginning(dateString string) (*time.Time, *time.Time, error) {
	const dateformat = `^(19[0-9]{2}|20[0-9]{2})/(0?[1-9]|1[0-2])$`
	if dateString == "" {
		n := time.Now()
		dateString = strconv.Itoa(n.Year()) + "/" + strconv.Itoa(int(n.Month()))
	}
	ok, err := regexp.MatchString(dateformat, dateString)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("[%s] is not matched to month format", dateString))
	}
	yearAndMonth := strings.Split(dateString, "/")
	year, err := strconv.Atoi(yearAndMonth[0])
	if err != nil {
		return nil, nil, err
	}
	month, err := strconv.Atoi(yearAndMonth[1])
	if err != nil {
		return nil, nil, err
	}
	nextYear, nextMonth := year, month+1
	if nextMonth > 12 {
		nextMonth = 1
		nextYear++
	}
	begin := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	end := time.Date(nextYear, time.Month(nextMonth), 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
	return &begin, &end, err
}
