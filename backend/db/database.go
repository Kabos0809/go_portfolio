package database

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var err error
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	db_name := os.Getenv("MYSQL_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, db_name)
	
	Db, err := gorm.Open("mysql", connection)
	
	if err != nil {
		panic(err)
	}

	
}