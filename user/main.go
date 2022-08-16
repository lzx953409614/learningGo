package main

import (
	"user/db"
)

func main() {
	defer db.MysqlDB.Close()
	router := initRouter()
	router.Run(":8001")
}
