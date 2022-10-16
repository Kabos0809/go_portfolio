package Controllers

import (
	"net/http"

	"github.com/kabos0809/go_portfolio/backend/Models"
	"github.com/gin-gonic/gin"
)

//Sign Up
func (c UserController) SignUp(ctx *gin.Context) {
	var user Models.User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400: Bad request"})
		return
	} else {
		if err := c.Model.CreateUser(user.UserName, user.Password); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{})
		}
	}
}

//Login
func (c UserController) Login(ctx *gin.Context) {
	var user Models.User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "400: Bad request"})
		return
	} else {
		if err := c.Model.CheckPassword(user.UserName, user.Password); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		} else {
			token, err := Models.CreateJWT(user.UserName, user.ID)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"err":"400: Bad request"})
				return
			}
			ctx.JSON(http.StatusOK, token)
		}
	}
}

//Logout
func (c UserController) Logout(ctx *gin.Context) {
	exp := ctx.MustGet("exp").(float64)
	token := ctx.MustGet("AccessToken").(string)
	if err := Models.SetBlackList(token, int64(exp)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
		ctx.Abort()
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

//Get RefreshToken
func (c UserController) GetRefresh(ctx *gin.Context) {
	ID := ctx.MustGet("ID").(float64)
	rt := ctx.MustGet("RefreshToken").(string)
	exp := ctx.MustGet("exp").(float64)
	token, err := c.Model.CreateRefresh(uint64(ID), rt, int64(exp))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "400: Bad request"})
		ctx.Abort()
	} else {
		ctx.JSON(http.StatusOK, token)
	}
}