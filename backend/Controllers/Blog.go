package Controllers

import (
	"net/http"
	"strconv"

	"github.com/kabos0809/go_portfolio/backend/Models"
	"github.com/gin-gonic/gin"
)

//Create Blogs
func (c BlogController) CreateBlog(ctx *gin.Context) {
	var blog Models.Blog
	if err := ctx.BindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400:Bad request"})
		return
	}
	if err := c.Model.CreateBlog(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

//Fetch Blogs
func (c BlogController) GetBlog(ctx *gin.Context) {
	r, err := c.Model.GetAllBlog()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, r)
	}
}

//Fetch Blog by ID
func (c BlogController) GetBlogByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idUint, _ := strconv.ParseUint(id, 10, 64)
	r, err := c.Model.GetBlogByID(idUint)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
	err = c.Model.IncrementSeeBlog(r)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		ctx.JSON(http.StatusOK, r)
	}
}

//Update Blog
func (c BlogController) UpdateBlog(ctx *gin.Context) {
	var blog *Models.Blog
	id := ctx.Params.ByName("id")
	idUint, _ := strconv.ParseUint(id, 10, 64)
	r, err := c.Model.GetBlogByID(idUint)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
		return
	}
	ctx.BindJSON(blog)
	r.Title = blog.Title
	r.Text = blog.Text
	//r.Tag = blog.Tag
	if err := c.Model.UpdateBlog(r); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, r)
	}
}

//Delete Blog
func (c BlogController) DeleteBlog(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idUint, _ := strconv.ParseUint(id, 10, 64)
	if err := c.Model.DeleteBlog(idUint); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id:" + id : "was deleted"})
	}
}

//Change IsActive
func (c BlogController) ChangeBlogIsActive(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idUint, _ := strconv.ParseUint(id, 10, 64)
	if err := c.Model.ChangeBlogIsActive(idUint); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ID:" + id + "'s" : "active was changed"})
	}
}