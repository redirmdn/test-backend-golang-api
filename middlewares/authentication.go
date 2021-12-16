package middlewares

import (
	"strings"
	"test-backend-golang-api/api/web"
	"test-backend-golang-api/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")

		token := strings.Split(authorization, " ")

		user, err := helpers.ParseJWT(token[1])

		if err != nil {
			web.ResponseUnauthorized(ctx)

			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
