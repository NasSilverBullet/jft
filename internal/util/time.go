package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ToTime(timeString string) (*time.Time, error) {
	const timeformat = `^(0?[0-9]|1[0-9]|2[0-3]):(0?[0-9]|[1-5][0-9])$`
	ok, err := regexp.MatchString(timeformat, timeString)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New(fmt.Sprintf("[%s] is not matched to time format", timeString))
	}
	hourAndMinStr := strings.Split(timeString, ":")
	hour, err := strconv.Atoi(hourAndMinStr[0])
	if err != nil {
		return nil, err
	}
	min, err := strconv.Atoi(hourAndMinStr[1])
	if err != nil {
		return nil, err
	}
	n := time.Now()
	t := time.Date(n.Year(), n.Month(), n.Day(), hour, min, 0, 0, time.Local)
	return &t, err
}

func GetDayEndAndBeginning(dateString string) (*time.Time, *time.Time, error) {
	const dateformat = `^(19[0-9]{2}|20[0-9]{2})/(0?[1-9]|1[0-2])/(0?[1-9]|[1-2][0-9]|3[0-1])$`
	if dateString == "" {
		n := time.Now()
		dateString = strconv.Itoa(n.Year()) + "/" + strconv.Itoa(int(n.Month())) + "/" + strconv.Itoa(n.Day())
	}
	ok, err := regexp.MatchString(dateformat, dateString)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("[%s] is not matched to date format", dateString))
	}
	yearAndMonthAndDate := strings.Split(dateString, "/")
	year, err := strconv.Atoi(yearAndMonthAndDate[0])
	if err != nil {
		return nil, nil, err
	}
	month, err := strconv.Atoi(yearAndMonthAndDate[1])
	if err != nil {
		return nil, nil, err
	}
	date, err := strconv.Atoi(yearAndMonthAndDate[2])
	if err != nil {
		return nil, nil, err
	}
	begin := time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.Local)
	end := time.Date(year, time.Month(month), date, 23, 59, 59, 59, time.Local)
	return &begin, &end, err
}

func GetMonthEndAndBeginning(monthString string) (*time.Time, *time.Time, error) {
	const dateformat = `^(19[0-9]{2}|20[0-9]{2})/(0?[1-9]|1[0-2])$`
	if monthString == "" {
		n := time.Now()
		monthString = strconv.Itoa(n.Year()) + "/" + strconv.Itoa(int(n.Month()))
	}
	ok, err := regexp.MatchString(dateformat, monthString)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("[%s] is not matched to month format", monthString))
	}
	yearAndMonth := strings.Split(monthString, "/")
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
