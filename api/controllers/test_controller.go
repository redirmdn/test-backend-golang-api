package controllers

import (
	"test-backend-golang-api/api/services"
	"test-backend-golang-api/api/web"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type TestController interface {
	GetTest(ctx *gin.Context)
	PostTest(ctx *gin.Context)
}

type TestControllerImpl struct {
	TestService services.TestService
}

func InitialTestController(testService services.TestService) TestController {
	return &TestControllerImpl{
		TestService: testService,
	}
}

func (params *TestControllerImpl) GetTest(ctx *gin.Context) {
	params.TestService.GetTest(*ctx)

	web.ResponseOK(ctx, gin.H{
		"user": ctx.MustGet("user"),
	})
}

func (params *TestControllerImpl) PostTest(ctx *gin.Context) {
	testRequest := web.PostTestRequest{}
	ctx.ShouldBindBodyWith(&testRequest, binding.JSON)

	params.TestService.PostTest(*ctx)

	web.ResponseOK(ctx, gin.H{
		"test": testRequest,
	})
}
