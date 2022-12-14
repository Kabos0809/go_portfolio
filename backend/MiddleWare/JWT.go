package MiddleWare

import (
	"net/http"
	"os"

	"github.com/kabos0809/go_portfolio/backend/Models"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			ACCESS_TOKEN_SECRETKEY := os.Getenv("ACCESS_TOKEN_SECRETKEY")
			b := []byte(ACCESS_TOKEN_SECRETKEY)
			return b, nil
		})

		if err == nil {
			err := Models.CheckBlackList(token.Raw);
			if err == nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": "Token is invalid"})
				c.Abort()
			} else {
				claims := token.Claims.(jwt.MapClaims)
				c.Set("exp", claims["exp"])
				c.Set("AccessToken", token.Raw)
				c.Next()
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"err":"Token was not authorized"})
			c.Abort()
		}
	}
}

func CheckRefresh() gin.HandlerFunc{
	return func(c *gin.Context) {
		rtoken, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			REFRESH_TOKEN_SECRETKEY := os.Getenv("REFRESH_TOKEN_SECRETKEY")
			b := []byte(REFRESH_TOKEN_SECRETKEY)
			return b, nil
		})

		if err == nil {
			if err := Models.CheckBlackList(rtoken.Raw); err == nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": "Token is invalid"})
				c.Abort()
			} else {
				rclaims := rtoken.Claims.(jwt.MapClaims)
				c.Set("exp", rclaims["exp"])
				c.Set("id", rclaims["id"])
				c.Set("RefreshToken", rtoken.Raw)
				c.Next()
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "Token was not authorized"})
			c.Abort()
		}
	}
}

