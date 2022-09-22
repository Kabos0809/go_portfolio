package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    // 1. Routerを作成
    router := mux.NewRouter().StrictSlash(true)
    // 2. URLと処理を紐付ける
    router.HandleFunc("/", home)
    // 3. ポートを指定して起動
    log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}