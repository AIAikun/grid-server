package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	username := "grid"
	password := "dian@2023"
	host := "47.116.59.2"
	port := "3306"
	Dbname := "grid_ry"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%s&parseTime=true&loc=Local", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB = db
}
