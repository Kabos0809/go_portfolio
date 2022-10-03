package Routes

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/kabos0809/go_portfolio/backend/Models"
	"github.com/kabos0809/go_portfolio/backend/Controller"
)

func SetRoutes(controller Models.Model) *gin.Engine {
	e := gin.Default()
	
	e.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{
			"http://localhost:3030"
		},
		MaxAge: 168 * time.Hour,
	}))

	v1 := e.Group("/v1")
	{	
		v1.GET("/home")
		v1.GET("/blog/list")
		v1.GET("/blog/detail/:id")
		v1.GET("/work/list")
		v1.GET("/work/detail/:id")
		v1.POST("/contact")
	}

	admin := e.Group("/admin")
	{
		admin.GET("/home")
		admin.GET("/contact/list")
		admin.GET("/contact/detail/:id")
		admin.POST("/blog/create")
		admin.PUT("/blog/update/:id")
		admin.DELETE("/blog/delete/:id")
		admin.POST("/work/create")
		admin.PUT("/work/update/:id")
		admin.DELETE("/work/delete/:id")
	}
	
	return e
}