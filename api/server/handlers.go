package server

import (
	"fmt"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type kanbansResponse struct {
	Kanbans []models.Kanban `json:"kanbans"`
}

func kanbanAll(c echo.Context) error {
	kanbans := models.KanbanAll()
	fmt.Println(kanbans)
	return c.JSON(http.StatusOK, kanbansResponse{Kanbans: kanbans})
}
