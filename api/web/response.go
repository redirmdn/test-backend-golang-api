package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func ResponseUnauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, Response{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	})
}

func ResponseBadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Message: message,
	})
}
