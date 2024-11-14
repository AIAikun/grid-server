package main

import (
	"server/db"
	"server/router"
)

func main() {
	db.InitMysql()
	r := router.Router()
	r.Run(":8081")
}
