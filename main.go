package main

import (
	"github.com/gin-gonic/gin"

	"github.com/khayrultw/go-judge/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/test/python", controllers.TestPython)
	r.GET("/test/kotlin", controllers.TestKotlin)

	codeRepo := controllers.New()
	r.POST("/code", codeRepo.PostCode)
	r.GET("/code/:id", codeRepo.GetCode)
	r.Run() // listen and serve on 0.0.0.0:8080
}
