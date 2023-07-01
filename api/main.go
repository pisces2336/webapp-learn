package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	_, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("db connection success!")
	}
}

func sqlConnect() (*gorm.DB, error) {
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

	fmt.Println(c.FormatDSN())

	return gorm.Open("mysql", c.FormatDSN())
}
