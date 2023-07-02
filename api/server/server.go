package server

import (
	"github.com/labstack/echo/v4"
)

func Start(port string) {
	e := echo.New()

	e.GET("/kanbans", kanbanAll)

	e.Logger.Fatal(e.Start(port))
}
