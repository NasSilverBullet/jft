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
