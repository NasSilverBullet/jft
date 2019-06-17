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

const rightTimeOrderError = `please set right time order`

type Plan struct {
	gorm.Model
	Start       *time.Time
	End         *time.Time
	Title       string `gorm:"size:255"`
	Description string `gorm:"size:1024"`
}

func MigratePlan(db *gorm.DB) {
	db.AutoMigrate(Plan{})
}

func NewPlan(db *gorm.DB, startStr string, endStr string, title string, description string) (*Plan, error) {
	start, err := parseTime(startStr)
	if err != nil {
		return nil, err
	}
	end, err := parseTime(endStr)
	if err != nil {
		return nil, err
	}
	if !end.After(start.Local()) {
		return nil, errors.New(rightTimeOrderError)
	}
	p := &Plan{
		Start:       start,
		End:         end,
		Title:       title,
		Description: description,
	}
	db.Create(p)
	return p, err
}

func GetPlan(db *gorm.DB, idStr string) (*Plan, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	p := &Plan{}
	db.First(p, id)
	if p.ID != uint(id) {
		err = errors.New(fmt.Sprintf("plan ID [%d] is not found", id))
	}
	return p, err
}

func (p *Plan) Update(db *gorm.DB, startStr string, endStr string, title string, description string) (*Plan, error) {
	var (
		start, end *time.Time
		err        error
	)
	if startStr == "" &&
		endStr == "" &&
		title == "" &&
		description == "" {
		return nil, errors.New("no update target")
	}

	if startStr == "" {
		start = p.Start
	} else {
		start, err = parseTime(startStr)
		if err != nil {
			return nil, err
		}
	}

	if endStr == "" {
		end = p.End
	} else {
		end, err = parseTime(endStr)
		if err != nil {
			return nil, err
		}
	}
	if !end.After(start.Local()) {
		return nil, errors.New(rightTimeOrderError)
	}

	if title == "" {
		title = p.Title
	}
	if description == "" {
		description = p.Description
	}

	db.Model(p).Updates(Plan{
		Start:       start,
		End:         end,
		Title:       title,
		Description: description,
	})
	return p, err
}

func (p *Plan) Delete(db *gorm.DB) (*Plan, error) {
	db.Delete(p)
	return p, nil
}

func FindPlans(db *gorm.DB, monthStr string) ([]Plan, error) {
	begin, end, err := getMonthEndAndBeginning(monthStr)
	if err != nil {
		return nil, err
	}
	ps := []Plan{}
	db.Where("start BETWEEN ? AND ?", begin, end).Find(&ps)
	return ps, err
}

func (p Plan) String() string {
	const layout = "2006-01-02 15:04"
	const format = "ID : %v\nStart : %v\nEnd : %v\nTitle : %v\nDescription : %v"
	return fmt.Sprintf(format, p.ID, p.Start.Format(layout), p.End.Format(layout), p.Title, p.Description)
}

func parseTime(timeString string) (*time.Time, error) {
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

func getMonthEndAndBeginning(monthString string) (*time.Time, *time.Time, error) {
	const monthformat = `^(19[0-9]{2}|20[0-9]{2})/(0?[1-9]|1[0-2])$`
	if monthString == "" {
		n := time.Now()
		monthString = strconv.Itoa(n.Year()) + "/" + strconv.Itoa(int(n.Month()))
	}
	ok, err := regexp.MatchString(monthformat, monthString)
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
	begin := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endYear, endMonth := year, month+1
	if endMonth > 12 {
		endYear++
		endMonth = 1
	}
	end := time.Date(endYear, time.Month(endMonth), 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
	return &begin, &end, err
}
