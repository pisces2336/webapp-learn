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
	jst, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		panic(err)
	}

	c := mysql.Config{
		DBName:    os.Getenv("MYSQL_DATABASE"),
		User:      "root",
		Passwd:    os.Getenv("MYSQL_ROOT_PASSWORD"),
		Addr:      os.Getenv("DB_ADDR"),
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
