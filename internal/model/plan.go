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

type Plan struct {
	gorm.Model
	Start            *time.Time
	End              *time.Time
	ShortDescription string `gorm:"size:255"`
	Description      string `gorm:"size:1024"`
}

func MigratePlan(db *gorm.DB) {
	db.AutoMigrate(Plan{})
}

func NewPlan(db *gorm.DB, startStr string, endStr string, sd string, d string) (*Plan, error) {
	start, err := parseTime(startStr)
	if err != nil {
		return nil, err
	}
	end, err := parseTime(endStr)
	if err != nil {
		return nil, err
	}

	p := &Plan{
		Start:            start,
		End:              end,
		ShortDescription: sd,
		Description:      d,
	}
	db.Create(p)
	return p, err
}

func parseTime(timeString string) (*time.Time, error) {
	ok, err := regexp.MatchString(`^([0-1][0-9]|2[0-3]):([0-5][0-9])$`, timeString)
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

func (p Plan) String() string {
	const layout = "2006-01-02 15:04"
	const format = "ID : %v\nStart : %v\nEnd : %v\nShort description : %v\nDescription : %v"
	return fmt.Sprintf(format, p.ID, p.Start.Format(layout), p.End.Format(layout), p.ShortDescription, p.Description)
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
