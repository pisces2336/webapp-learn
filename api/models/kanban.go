package models

type Kanban struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Category int    `json:"category"`
}
