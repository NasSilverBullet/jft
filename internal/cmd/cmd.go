package cmd

import (
	"github.com/NasSilverBullet/jft/internal/jft"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cobra.OnInitialize()
	cmd := jft.Exec()
	cmd.AddCommand(jft.Add())
	cmd.AddCommand(jft.Update())
	cmd.AddCommand(jft.Delete())
	cmd.AddCommand(jft.List())
	cmd.AddCommand(jft.Month())
	cmd.AddCommand(jft.Year())
	return cmd
}
