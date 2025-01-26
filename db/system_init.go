package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	username := "root"
	password := "M=qRg_#UJhL?Da>=g9^3"
	host := "101.43.63.11"
	port := "3306"
	Dbname := "86_system_grid"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%s&parseTime=true&loc=Local", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB = db
}
