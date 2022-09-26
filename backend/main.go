package main

import (
    "fmt"
    "log"
    "net/http"
	"github.com/jinzhu/gorm"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    router := http.Server{
		Addr: ":8080"
	}

	router.HandleFunc("/check_db", check_cnct_db)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func check_cnct_db(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, database)
	db, err := gorm.Open("mysql", connection)

	defer db.Close()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}