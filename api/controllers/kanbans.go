package controllers

import (
	"fmt"
	"main/database"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateKanban(c echo.Context) error {
	db := database.DB
	kanban := new(models.Kanban)
	err := c.Bind(kanban)
	if err != nil {
		return err
	}
	db.Create(&kanban)
	return c.JSON(http.StatusOK, kanban)
}

func GetAllKanbans(c echo.Context) error {
	db := database.DB
	kanbans := new([]models.Kanban)
	result := db.Find(&kanbans)
	if result.RecordNotFound() {
		fmt.Println("Record not found.")
	}
	return c.JSON(http.StatusOK, kanbans)
}

func UpdateKanban(c echo.Context) error {
	db := database.DB
	newKanban := new(models.Kanban)
	err := c.Bind(&newKanban)
	if err != nil {
		return err
	}
	kanban_id := c.Param("id")
	if kanban_id != "" {
		kanban := models.Kanban{}
		db.First(&kanban, "id = ?", kanban_id).Update(newKanban)
		return c.JSON(http.StatusOK, kanban)
	} else {
		return c.JSON(http.StatusNotFound, nil)
	}
}

func DeleteKanban(c echo.Context) error {
	db := database.DB
	kanban_id := c.Param("id")
	if kanban_id != "" {
		db.Delete(&models.Kanban{}, kanban_id)
		return c.JSON(http.StatusOK, nil)
	} else {
		return c.JSON(http.StatusNotFound, nil)
	}
}
