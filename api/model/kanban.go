package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Kanban struct {
	ID       uint
	Title    string
	Body     string
	Category int
}

func CreateKanban(db *gorm.DB, kanban Kanban) {
	result := db.Create(&kanban)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func GetFirstKanban(db *gorm.DB) Kanban {
	kanban := Kanban{}
	result := db.First(&kanban)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return kanban
}

func GetKanbanListByCategory(db *gorm.DB, category int) []Kanban {
	kanban_list := []Kanban{}
	result := db.Where(&Kanban{Category: category}).Find(&kanban_list)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return kanban_list
}

func UpdateKanban(db *gorm.DB, kanban *Kanban) {
	result := db.Model(&Kanban{}).Where(&Kanban{ID: kanban.ID}).Update(kanban)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func DeleteKanban(db *gorm.DB, id uint) {
	result := db.Where(&Kanban{ID: id}).Delete(&Kanban{})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
