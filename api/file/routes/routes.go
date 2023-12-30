package routes

import (
	"uploader-service/api/file/handlers"

	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {

	uploadFile(e)
	downloadFile(e)
	renameFile(e)
	deleteFile(e)
	deleteFile(e)
}

func uploadFile(g *echo.Echo) {
	g.POST("/file", handlers.UploadFile)
}

func downloadFile(g *echo.Echo) {
	g.GET("/file", handlers.UploadFile)
}

func renameFile(g *echo.Echo) {
	g.PATCH("/file", handlers.RenameFile)
}

func deleteFile(g *echo.Echo) {
	g.DELETE("/file", handlers.DeleteFile)
}

func deleteFiles(g *echo.Echo) {
	g.DELETE("/file/group", handlers.DeleteFiles)
}
