package db

import (
	"database/sql"
	"os/exec"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DBMS     = "mysql"
	USER     = "root"
	PASS     = ""
	PROTOCOL = ""
	DBNAME   = "jft"
	OPTION   = "?parseTime=true"
)

func New() (*gorm.DB, error) {
	if err := wakeUpMySQL(); err != nil {
		return nil, err
	}
	if err := createDB(); err != nil {
		return nil, err
	}
	db, err := gormConnect()
	if err != nil {
		return db, err
	}
	//db.LogMode(true)
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	return db, err
}

func wakeUpMySQL() error {
	return exec.Command("mysql.server", "start").Run()
}

func createDB() error {
	SOURCE := USER + ":" + PASS + "@" + PROTOCOL + "/"
	db, err := sql.Open(DBMS, SOURCE)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("USE " + DBNAME)
	if err != nil {
		_, err = db.Exec("CREATE DATABASE " + DBNAME)
		if err != nil {
			return err
		}
		_, err = db.Exec("USE " + DBNAME)
		if err != nil {
			return err
		}
	}
	return err
}

func gormConnect() (*gorm.DB, error) {
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + OPTION
	return gorm.Open(DBMS, CONNECT)
}
