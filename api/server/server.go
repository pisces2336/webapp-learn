package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(port string) {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/kanbans", kanbanAll)

	e.Logger.Fatal(e.Start(port))
}
