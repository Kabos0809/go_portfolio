package main

import (
	"gorm.io/gorm"
    "gorm.io/driver/postgres"
	"github.com/gin-gonic/gin"

    "github.com/kabos0809/go_portfolio/backend/Models"
    "github.com/kabos0809/go_portfolio/backend/Routes"
    "github.com/kabos0809/go_portfolio/backend/Config"
)

func main() {
    gin.SetMode(gin.DebugMode)

    dsn := Config.DBUrl()
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    
    sqlDB, _ := db.DB()
    defer sqlDB.Close()

    err = db.AutoMigrate(&Models.Contact{}, &Models.User{})
    if err != nil {
        panic(err)
    }
    err = db.AutoMigrate(&Models.Blog{})
    if err != nil {
        panic(err)
    }
    err = db.AutoMigrate(&Models.Work{})
    if err != nil {
        panic(err)
    }

    m := Models.Model{Db: db}

    r := Routes.SetRoutes(m)

    r.Run(":8080")
}