package route

import (
	fileRoutes "uploader-service/api/file/routes"
	userRoutes "uploader-service/api/user/routes"
)

func SetRoutes() {
	userRoutes.SetRoutes()
	fileRoutes.SetRoutes()
}
