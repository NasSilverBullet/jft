package jft

import (
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
