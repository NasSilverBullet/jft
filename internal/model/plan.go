package model

import (
	"fmt"
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
	return fmt.Sprintf("Start : %v\nEnd : %v\nShort description : %v\nDescription : %v", p.Start.Format(layout), p.End.Format(layout), p.ShortDescription, p.Description)
}
