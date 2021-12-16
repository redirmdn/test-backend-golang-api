package controllers

import (
	"test-backend-golang-api/api/web"
	"test-backend-golang-api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Login(ctx *gin.Context) {
	loginRequest := web.AuthLoginRequest{}
	ctx.ShouldBindBodyWith(&loginRequest, binding.JSON)

	user := map[string]string{
		"username": "redi",
		"password": "123456",
	}

	if loginRequest.Username == user["username"] && loginRequest.Password == user["password"] {
		token, err := helpers.GenerateJWT(loginRequest.Username)

		if err != nil {
			web.ResponseUnauthorized(ctx)

			return
		}

		web.ResponseOK(ctx, gin.H{
			"token": token,
		})

		return
	}

	web.ResponseUnauthorized(ctx)
}
