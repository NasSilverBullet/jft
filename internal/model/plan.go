package model

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/NasSilverBullet/jft/internal/util"
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
	start, err := util.ToTime(startStr)
	if err != nil {
		return nil, err
	}
	end, err := util.ToTime(endStr)
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
		start, err = util.ToTime(startStr)
		if err != nil {
			return nil, err
		}
	}

	if endStr == "" {
		end = p.End
	} else {
		end, err = util.ToTime(endStr)
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

func FindPlans(db *gorm.DB, dateStr string) ([]Plan, error) {
	begin, end, err := util.GetDayEndAndBeginning(dateStr)
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
