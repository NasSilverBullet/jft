package jft

import (
	"fmt"

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
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add today's new plans",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("add new plans!!")
			return nil
		},
	}
	return cmd
}
