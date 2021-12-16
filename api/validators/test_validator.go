package validators

import (
	"test-backend-golang-api/api/web"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidatePostTestRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody := web.PostTestRequest{}
		err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON)

		handleErrorValidation(ctx, err)
	}
}
