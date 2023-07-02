package server

import (
	"fmt"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func kanbanAll(c echo.Context) error {
	kanbans := models.KanbanAll()
	fmt.Println(kanbans)
	return c.JSON(http.StatusOK, kanbans)
}
