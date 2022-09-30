package main

import (
    "fmt"
    "log"
    "net/http"
	"github.com/jinzhu/gorm"
    "gorm.io/driver/postgres"
	"github.com/gin-gonic/gin"

    "github.com/kabos0809/go_portfolio/backend/Models"
    "github.com/kabos0809/go_portfolio/backend/Routes"
    "github.com/kabos0809/go_portfolio/backend/Controller"
    "github.com/kabos0809/go_portfolio/backend/Config"
)

func main() {
    gin.SetMode(gin.DebugMode)

    dbconfig := Config.buildDBConfig()
    dsn := DBUrl(dbconfig)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    
    pgs, _ := db.DB()
    defer pgs.Close()

    err := db.AutoMigrate(&Models.Blog{}, &Models.Work{}, &Models.Contact{})
    if err != nil {
        panic(err)
    }

    m := Models.Model{Db: db}

    r := Routes.SetRoutes(m)

    r.Run(":8080")
}