package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// v1 := e.Group("/v1")
	// v1.GET("/healthcheck", Healthcheck)
	// v1.POST("/users", api.CreateUser)
	// v1.GET("/users/:email", api.GetUser)

	return e
}
