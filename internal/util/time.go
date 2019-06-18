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
