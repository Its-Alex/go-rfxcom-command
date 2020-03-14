package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func routerV1(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.POST("/shutters/control", controlShutters)

	v1.GET("/ttys", showTTY)
	v1.POST("/ttys/set", changePort)
}

// Launch start api
func Launch() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routerV1(e)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(
		"%s:%s",
		viper.GetString("addr"),
		viper.GetString("port"),
	)))
}
