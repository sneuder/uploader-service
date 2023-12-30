package main

import (
	"uploader-service/config/database"
	echoConfig "uploader-service/config/echo"
	"uploader-service/config/middlewares"
	"uploader-service/config/route"
	"uploader-service/config/server"
	"uploader-service/config/tools"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := database.ConnectDB()
	g := e.Group("")

	tools.SetGlobalTools(e, g, db)

	database.SetModels()
	echoConfig.SetCustomes()

	middlewares.SetMiddlewares()
	route.SetRoutes()
	server.RunServer(e)
}
