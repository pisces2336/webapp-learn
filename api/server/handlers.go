package server

import (
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func kanbanAll(c echo.Context) error {
	kanbans := models.KanbanAll()
	return c.JSON(http.StatusOK, kanbans)
}
