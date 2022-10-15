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
			"http://localhost:3030",
			"http://kabos-portfolio.com",
		},

		AlloCredentials: true,
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
		admin.POST("/login")
		admin.GET("/home", MiddleWare.CheckJWT())
		admin.GET("/contact/list", MiddleWare.CheckJWT())
		admin.GET("/contact/detail/:id", MiddleWare.CheckJWT())
		admin.POST("/blog/create", MiddleWare.CheckJWT())
		admin.PUT("/blog/update/:id", MiddleWare.CheckJWT())
		admin.DELETE("/blog/delete/:id", MiddleWare.CheckJWT())
		admin.POST("/work/create", MiddleWare.CheckJWT())
		admin.PUT("/work/update/:id", MiddleWare.CheckJWT())
		admin.DELETE("/work/delete/:id", MiddleWare.CheckJWT())
		admin.GET("/refresh", MiddleWare.CheckRefresh())
	}
	
	return e
}