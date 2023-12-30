package routes

import (
	"uploader-service/api/file/handlers"
	"uploader-service/config/tools"

	"github.com/labstack/echo/v4"
)

var g *echo.Group

func SetRoutes() {
	g := tools.G

	g.POST("/file", handlers.UploadFile)
	g.GET("/file/user/:userId", handlers.GetAllFilesByUserId)
	g.GET("/file/:fileId/user/:userId", handlers.GetFilesByUserId)
	g.GET("/file/:fileId/user/:userId/download", handlers.GetFile)
	// g.DELETE("/file", handlers.DeleteFile)
}
