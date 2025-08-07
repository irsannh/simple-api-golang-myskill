package auth

import (
	"net/http"
	"simple-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	USER="admin"
	PASSWORD="Password123!"
	SECRET="secret"
)

func LoginHandler(ctx *gin.Context) {
	var user models.Credential
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad Request",
		})
	}

	if user.Username != USER {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "User Invalid",
		})
	} else if user.Password != PASSWORD {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Password Invalid",
			})
	} else {
		claim := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		Issuer: "Test",
		IssuedAt: time.Now().Unix(),
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte(SECRET))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token": token,
	})
	}

	
}