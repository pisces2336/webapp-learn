package database

import (
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() {
	sqlConnect()
}

func sqlConnect() {
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

	DB, err = gorm.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
}
