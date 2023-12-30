package middlewares

import (
	"uploader-service/config/tools"
	middle "uploader-service/middlewares"
)

func SetMiddlewares() {
	tools.G.Use(middle.Auth)
}
