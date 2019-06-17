package main

import (
	"log"

	"github.com/NasSilverBullet/jft/internal/cmd"
	"github.com/NasSilverBullet/jft/internal/db"
	"github.com/NasSilverBullet/jft/internal/model"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := db.New()
	if err != nil {
		return err
	}
	db.AutoMigrate(&model.Plan{})
	defer func() {
		err = db.Close()
	}()
	c := cmd.New()
	err = c.Execute()
	return err
}
