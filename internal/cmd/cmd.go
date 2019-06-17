package cmd

import (
	"github.com/NasSilverBullet/jft/internal/jft"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cobra.OnInitialize()
	cmd := jft.Exec()
	cmd.AddCommand(jft.Hello())
	return cmd
}
