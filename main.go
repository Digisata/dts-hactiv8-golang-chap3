package main

import (
	"github.com/Digisata/dts-hactiv8-golang-chap3/database"
	"github.com/Digisata/dts-hactiv8-golang-chap3/router"
)

// @title           API Specification
// @description     Mygram API documentation.
// @termsOfService  http://swagger.io/terms/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	database.StartDB()

	router.New().Run(":3000")
}
