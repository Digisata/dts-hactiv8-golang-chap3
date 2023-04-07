package main

import (
	"github.com/Digisata/dts-hactiv8-golang-chap3/database"
	"github.com/Digisata/dts-hactiv8-golang-chap3/router"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}
