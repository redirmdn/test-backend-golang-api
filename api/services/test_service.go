package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TestService interface {
	GetTest(ctx gin.Context)
	PostTest(ctx gin.Context)
}

type TestServiceImpl struct{}

func InitialTestService() TestService {
	return &TestServiceImpl{}
}

func (service *TestServiceImpl) GetTest(ctx gin.Context) {
	fmt.Println("Berhasil Get")
}

func (service *TestServiceImpl) PostTest(ctx gin.Context) {
	fmt.Println("Berhasil Post")
}
