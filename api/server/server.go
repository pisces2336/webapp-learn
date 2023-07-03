package server

import (
	"main/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(port string) {
	e := echo.New()

	e.Use(middleware.CORS())

	e.POST("/kanbans", controllers.CreateKanban)
	e.GET("/kanbans", controllers.GetAllKanbans)
	e.PATCH("/kanbans/:id", controllers.UpdateKanban)
	e.DELETE("/kanbans/:id", controllers.DeleteKanban)

	e.Logger.Fatal(e.Start(port))
}
