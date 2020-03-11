package actions

import "github.com/labstack/echo"

func Register(e *echo.Echo) {
	e.GET("/login", actions.Login)
}
