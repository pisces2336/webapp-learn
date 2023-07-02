package server

import (
	"github.com/labstack/echo/v4"
)

func Start(port string) {
	e := echo.New()

	setRoutes(e)

	e.Logger.Fatal(e.Start(port))
}
