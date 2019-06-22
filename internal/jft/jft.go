package jft

import (
	"errors"
	"fmt"

	"github.com/NasSilverBullet/jft/internal/db"
	"github.com/NasSilverBullet/jft/internal/model"
	"github.com/spf13/cobra"
)

func Exec() *cobra.Command {
	cmd := &cobra.Command{
		// TODO: 時間があれば、説明を充実する
		Use:   "jft",
		Short: "calendar cli tool, just for today",
		Long: `This tool is a calendar that helps you plan and keep running every day.
If you plan daily, it will surely be for you.`,
	}
	return cmd
}

func Add() *cobra.Command {
	var description string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add today's new plans",
		Long: `You can do this for example with the following command:
  $ jft add 10:00 12:00 'check emails'
  $ jft add 13:00 14:30 'meeting' -d 'on conference room 10' ## You can add detailed description`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			p, err := model.NewPlan(db, args[0], args[1], args[2], description)
			if err != nil {
				return err
			}
			fmt.Println("added a new plan!!")
			fmt.Println(p)
			return err
		},
	}
	cmd.Flags().StringVarP(&description, "description", "d", "", "detailed description")
	return cmd
}

func Update() *cobra.Command {
	var start, end, title, description string
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update today's each plan, please give me id",
		Long: `You can do this for example with the following command:
  $ jft list ## You can check your today's plans
  $ jft list -w 2019/06/01 ## You can check your each day's plans`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			p, err := model.GetPlan(db, args[0])
			if err != nil {
				return err
			}
			p, err = p.Update(db, start, end, title, description)
			if err != nil {
				return err
			}
			fmt.Println("updated the plan!!")
			fmt.Println(p)
			return err
		},
	}
	cmd.Flags().StringVarP(&start, "start", "s", "", "start time")
	cmd.Flags().StringVarP(&end, "end", "e", "", "end time")
	cmd.Flags().StringVarP(&title, "title", "t", "", "short description")
	cmd.Flags().StringVarP(&description, "desc", "d", "", "detailed description")
	return cmd
}

func Delete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete today's each plan, please give me id",
		Long: `You can do this for example with the following command:
  $ jft delete 1`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			p, err := model.GetPlan(db, args[0])
			if err != nil {
				return err
			}
			p, err = p.Delete(db)
			if err != nil {
				return err
			}
			fmt.Println("deleted the plan!!")
			fmt.Println(p)
			return err
		},
	}
	return cmd
}

func List() *cobra.Command {
	var date string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "show plans list",
		Long: `You can do this for example with the following command:
  $ jft list ## You can check your today's plans
  $ jft list -w 2019/06/01 ## You can check your each day's plans`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			ps, err := model.FindPlans(db, date)
			if err != nil {
				return err
			}
			if date == "" {
				date = "today"
			}
			if len(ps) == 0 {
				return errors.New(fmt.Sprintf("no plans on %v", date))
			}
			var lineFeed string
			for _, p := range ps {
				fmt.Print(lineFeed)
				fmt.Println(p)
				lineFeed = "\n"
			}
			return err
		},
	}
	cmd.Flags().StringVarP(&date, "when", "w", "", "choose date")
	return cmd
}

func Month() *cobra.Command {
	var month string
	cmd := &cobra.Command{
		Use:   "month",
		Short: "show monthly calendar",
		Long: `You can do this for example with the following command:
  $ jft month ## You can check your efforts on this month
  $ jft month -w 2019/05 ## You can check your efforts on each month`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			days, err := model.FindDays(db, month)
			if err != nil {
				return err
			}
			if len(days) == 0 {
				return errors.New(fmt.Sprintf("There are no %v", month))
			}
			for _, day := range days {
				fmt.Println(day)
			}
			return err
		},
	}
	cmd.Flags().StringVarP(&month, "when", "w", "", "choose month")
	return cmd
}

func Year() *cobra.Command {
	var year string
	cmd := &cobra.Command{
		Use:   "year",
		Short: "show yearly calendar",
		Long: `You can do this for example with the following command:
  $ jft year ## You can check your efforts on this year
  $ jft year -w 2018 ## You can check your efforts on each year`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			ms, err := model.FindMonths(db, year)
			for _, m := range ms {
				fmt.Println(m)
			}

			return err
		},
	}
	cmd.Flags().StringVarP(&year, "when", "w", "", "choose year")
	return cmd
}
