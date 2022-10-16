package Controllers

import (
	"net/http"
	"strconv"

	"github.com/kabos0809/go_portfolio/Models"
	"github.com/gin-gonic/gin"
)

//Create Contact
func (c ContactController) CreateContact(ctx *gin.Context) {
	var contact Models.Contact
	if err := ctx.BindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
		return
	}
	if err := c.Model.CreateContact(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

//Fetch all Contact
func (c ContactController) GetContact(ctx *gin.Context) {
	r, err := c.Model.GetContact()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
	} else {
		ctx.JSON(http.StatusOK, r)
	}
}

//Fetch Contact by ID
func (c ContactController) GetContactByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idUint := strconv.ParseUint(id, 10, 64)
	r, err := c.Model.GetContactByID(idUint)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		if err := c.Model.ReadContact(idUint); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
		} else {
			ctx.JSON(http.StatusOK, r)
		}
	}
}

//Delete Contact
func (c ContactController) DeleteContact(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idUint := strconv.ParseUint(id, 10, 64)
	if err := c.Model.DeleteContact(idUint); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id" + id: " was deleted"})
	}
}