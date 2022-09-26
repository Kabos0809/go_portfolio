package main

import (
    "fmt"
    "log"
    "net/http"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	v1 := e.Group("/v1")
	{
		v1.GET("/works", controller.WorkList)
		v1.GET("/blogs", controller.BlogList)
		v1.POST("/work-post", controller.WorkPost)
		v1.POST("/blog-post", controller.BlogPost)
		v1.
	}
}