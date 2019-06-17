package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func New() (*gorm.DB, error) {
	db, err := gormConnect()
	if err != nil {
		return db, err
	}
	db.LogMode(true)
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	return db, err
}

func gormConnect() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := ""
	DBNAME := "jft"
	OPTION := "?parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + OPTION
	return gorm.Open(DBMS, CONNECT)
}
