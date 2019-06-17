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

func Hello() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "Start today's working",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("create today's record")
			return nil
		},
	}
	return cmd
}
