package route

import (
	"uploader-service/api/user/handlers"
	"uploader-service/config/tools"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo
var g *echo.Group

func SetRoutes() {
	e = tools.E
	g = tools.G

	e.POST("/user/signup", handlers.SignUp)
	e.POST("/user/signin", handlers.SignIn)
	g.PATCH("/user/:id", handlers.UpdateUser)
}
