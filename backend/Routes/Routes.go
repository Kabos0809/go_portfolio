package Routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/kabos0809/go_portfolio/backend/Models"
	"github.com/kabos0809/go_portfolio/backend/Controllers"
	"github.com/kabos0809/go_portfolio/backend/MiddleWare"
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
			"http://kabos-official.com",
			"http://kabos-official.net",
		},

		AllowCredentials: true,
	}))

	v1 := e.Group("/v1")
	{	
		v1.GET("/blog/list", Controllers.BlogController{Model: controller}.GetBlog)
		v1.GET("/blog/detail/:id", Controllers.BlogController{Model: controller}.GetBlogByID)
		v1.GET("/work/list", Controllers.WorkController{Model: controller}.GetWork)
		v1.GET("/work/detail/:id", Controllers.WorkController{Model: controller}.GetWorkByID)
		v1.POST("/contact", Controllers.ContactController{Model: controller}.CreateContact)
	}

	admin := e.Group("/admin")
	{
		//デプロイ時は/signupは無効に
		admin.POST("/signup", Controllers.UserController{Model: controller}.SignUp)
		admin.POST("/login", Controllers.UserController{Model: controller}.Login)
		admin.GET("/logout", MiddleWare.CheckJWT(), Controllers.UserController{Model: controller}.Logout)
		admin.GET("/contact/list", MiddleWare.CheckJWT(), Controllers.ContactController{Model: controller}.GetContact)
		admin.GET("/contact/detail/:id", MiddleWare.CheckJWT(), Controllers.ContactController{Model: controller}.GetContactByID)
		admin.DELETE("/contact/delete/:id", MiddleWare.CheckJWT(), Controllers.ContactController{Model: controller}.DeleteContact)
		admin.POST("/blog/create", MiddleWare.CheckJWT(), Controllers.BlogController{Model: controller}.CreateBlog)
		admin.PUT("/blog/update/:id", MiddleWare.CheckJWT(), Controllers.BlogController{Model: controller}.UpdateBlog)
		admin.DELETE("/blog/delete/:id", MiddleWare.CheckJWT(), Controllers.BlogController{Model: controller}.DeleteBlog)
		admin.PUT("/blog/changeactive/:id", MiddleWare.CheckJWT(), Controllers.BlogController{Model: controller}.ChangeBlogIsActive)
		admin.POST("/work/create", MiddleWare.CheckJWT(), Controllers.WorkController{Model: controller}.CreateWork)
		admin.PUT("/work/update/:id", MiddleWare.CheckJWT(), Controllers.WorkController{Model: controller}.UpdateWork)
		admin.DELETE("/work/delete/:id", MiddleWare.CheckJWT(), Controllers.WorkController{Model: controller}.DeleteWork)
		admin.PUT("/work/changeactive/:id", MiddleWare.CheckJWT(), Controllers.WorkController{Model: controller}.ChangeWorkIsActive)
		admin.GET("/refresh", MiddleWare.CheckRefresh(), Controllers.UserController{Model: controller}.GetRefresh)
	}
	
	return e
}