package main

import (
	"fmt"
	"main/model"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db := sqlConnect()

	db.AutoMigrate(&model.Kanban{})

	new_kanban := model.Kanban{
		Title:    "タイトルテスト",
		Body:     "本文テスト",
		Category: 0,
	}
	model.CreateKanban(db, new_kanban)

	firstKanban := model.GetFirstKanban(db)
	fmt.Println(firstKanban)

	firstKanban.Title = "編集されたタイトル"

	model.UpdateKanban(db, &firstKanban)
	updatedKanban := model.GetFirstKanban(db)
	fmt.Println(updatedKanban)
}

func sqlConnect() *gorm.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	c := mysql.Config{
		DBName:    "webapp_learn",
		User:      "root",
		Passwd:    os.Getenv("MYSQL_ROOT_PASSWORD"),
		Addr:      "db:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := gorm.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("db connection success!")
	}

	return db
}
