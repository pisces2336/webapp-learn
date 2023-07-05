package controllers

import (
	"log"
	"main/database"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateKanban(c echo.Context) error {
	db := database.DB

	kanban := models.Kanban{}
	err := c.Bind(&kanban)
	if err != nil {
		log.Fatal(err)
	}

	result := db.Create(&kanban)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return c.JSON(http.StatusOK, kanban)
}

func GetAllKanbans(c echo.Context) error {
	db := database.DB

	kanbans := []models.Kanban{}

	result := db.Find(&kanbans)
	if result.RecordNotFound() {
		return c.JSON(http.StatusNotFound, kanbans)
	} else if result.Error != nil {
		log.Fatal(result.Error)
	}

	return c.JSON(http.StatusOK, kanbans)
}

func UpdateKanban(c echo.Context) error {
	db := database.DB

	kanban := models.Kanban{}
	err := c.Bind(&kanban)
	if err != nil {
		log.Fatal(err)
	}

	result := db.Save(&kanban)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return c.JSON(http.StatusOK, kanban)
}

func DeleteKanban(c echo.Context) error {
	db := database.DB

	kanban := models.Kanban{}
	result := db.Where("id = ?", c.Param("id")).First(&kanban)
	if result.RecordNotFound() {
		return c.JSON(http.StatusNotFound, kanban)
	} else if result.Error != nil {
		log.Fatal(result.Error)
	}

	db.Delete(&kanban)

	return c.JSON(http.StatusOK, kanban)
}
