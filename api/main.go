package main

import (
	"main/database"
	"main/server"
)

func main() {
	database.Init()
	server.Start(":8000")
}
