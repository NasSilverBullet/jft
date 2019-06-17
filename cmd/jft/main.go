package main

import (
	"os"

	"github.com/NasSilverBullet/jft/internal/cmd"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	c := cmd.New()
	err := c.Execute()
	return err
}
