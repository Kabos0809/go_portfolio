package Controllers

import (
	"net/http"
	"strconv"

	"github.com/kabos0809/go_portfolio/Models"
	"github.com/gin-gonic/gin"
)

//Create Work
func (c WorkController) CreateWork(ctx *gin.Context) {
	var work Models.Work
	if err := ctx.BindJSON(&work); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
		return
	}
	if err := c.Model.CreateWork(&work); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

//Fetch all Works
func (c WorkController) GetWork(ctx *gin.Context) {
	r, err := c.Model.GetAllWork()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
	} else {
		ctx.JSON(http.StatusOK, r)
	}
}

//Fetch Work by ID
func (c WorkController) GetWorkByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idUint := strconv.ParseUint(id, 10, 64)
	r, err := c.Model.GetWorkByID(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, r)
	}
}

//Update Work
func (c WorkController) UpdateWork(ctx *gin.Context) {
	var work *Models.Work
	if err := ctx.BindJSON(work); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
	}
	id := ctx.Params.ByName("id")
	idUint := strconv.ParseUint(id, 10, 64)
	r, err := c.Model.GetWorkByID(idUint)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	r.Title = work.Title
	r.Text = work.Text
	r.Thumbnail = work.Thumbnail
	r.Tag = work.Tag
	if err := c.Model.UpdateWork(r); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, r)
	}
}

//Delete Work
func (c WorkController) DeleteWork(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idUint := strconv.ParseUint(id, 10, 64)
	if err := c.Model.DeleteWork(id); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id" + id: "was deleted"})
	}
}