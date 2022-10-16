package main

import (
	"hacktiv8-golang-final-project/database"
	"hacktiv8-golang-final-project/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8000")
}
