package validators

import (
	"test-backend-golang-api/api/web"

	"github.com/gin-gonic/gin"
)

func handleErrorValidation(ctx *gin.Context, err error) {
	if err != nil {
		web.ResponseBadRequest(ctx, err.Error())

		ctx.Abort()
		return
	}

	ctx.Next()
}
