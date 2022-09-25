package main

import (
    "fmt"
    "log"
    "net/http"
	"github.com/jinzhu/gorm"
	"os"

	_ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

func main() {
    // 1. Routerを作成
    router := mux.NewRouter().StrictSlash(true)
    // 2. URLと処理を紐付ける
    router.HandleFunc("/", home)
	router.HandleFunc("/check_db", check_cnct_db)
    // 3. ポートを指定して起動
    log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "World")
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