package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/roham96/go-basic/server/api"
	"github.com/roham96/go-basic/server/configuration"
	"github.com/roham96/go-basic/server/database"
)

func init() {
	// read env
	configuration.ReadConfiguration()
}

func RunServer() {
	// Connect to Database
	db := database.ConnectDatabase()

	log.Println("Run Server")
	e := echo.New()

	// Initial Beans
	api.InitialBeans(e, db)

	e.Logger.Fatal(e.Start(":8585"))
}
