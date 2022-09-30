package Routes

import (
	"github.com/gin-gonic/gin"

	"github.com/kabos0809/go_portfolio/backend/Models"
	"github.com/kabos0809/go_portfolio/backend/Controller"
)

func SetRoutes(controller Models.Model) *gin.Engine {
	e := gin.Default()
	
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