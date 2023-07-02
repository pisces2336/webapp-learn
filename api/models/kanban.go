package models

import (
	"log"
	"main/database"
)

type Kanban struct {
	ID       uint
	Title    string
	Body     string
	Category int
}

func KanbanCreate(kanban *Kanban) {
	err := database.DB.Create(kanban).Error
	if err != nil {
		log.Fatal(err)
	}
}

func KanbanFindById(id uint) {
	kanban := Kanban{}
	err := database.DB.Find(&kanban, &Kanban{ID: id}).Error
	if err != nil {
		log.Fatal(err)
	}
}

func KanbanAll() []Kanban {
	kanban_list := []Kanban{}
	err := database.DB.Find(&kanban_list).Error
	if err != nil {
		log.Fatal(err)
	}
	return kanban_list
}

func KanbanUpdate(kanban *Kanban) {
	err := database.DB.Where(&Kanban{ID: kanban.ID}).Update(kanban).Error
	if err != nil {
		log.Fatal(err)
	}
}

func KanbanDelete(id uint) {
	err := database.DB.Where(&Kanban{ID: id}).Delete(&Kanban{}).Error
	if err != nil {
		log.Fatal(err)
	}
}
