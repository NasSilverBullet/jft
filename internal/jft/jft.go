package jft

import (
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
		Long:  ``,
	}
	return cmd
}

func Add() *cobra.Command {
	var description string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add today's new plans",
		// TODO: 時間があれば、説明を充実する,
		Long: ``,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			db.AutoMigrate(&model.Plan{})
			defer func() {
				err = db.Close()
			}()
			p, err := model.NewPlan(db, args[0], args[1], args[2], description)
			if err != nil {
				return err
			}
			fmt.Println("Create a new plan!!")
			fmt.Println(p)
			return err
		},
	}
	cmd.Flags().StringVarP(&description, "description", "d", "", "detailed description")
	return cmd
}
