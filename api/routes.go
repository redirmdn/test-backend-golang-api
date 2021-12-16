package api

import (
	"test-backend-golang-api/api/controllers"
	"test-backend-golang-api/api/services"
	"test-backend-golang-api/api/validators"
	"test-backend-golang-api/helpers"
	"test-backend-golang-api/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	testService := services.InitialTestService()
	testController := controllers.InitialTestController(testService)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(helpers.LoggerHandler())

	apiv1 := router.Group("/v1")

	apiv1.POST("/login", validators.ValidateAuthLoginRequest(), controllers.Login)

	apiv1.Use(middlewares.Authentication())
	{
		apiv1.GET("/test", testController.GetTest)
		apiv1.POST("/test", testController.PostTest)
	}

	return router
}
