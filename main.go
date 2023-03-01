package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/khayrultw/go-judge/controllers"
	"github.com/khayrultw/go-judge/models"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

		var code models.Code
		c.BindJSON(&code)
	})

	codeRepo := controllers.New()
	r.POST("/code", codeRepo.PostCode)
	r.GET("/code/:id", codeRepo.GetCode)
	r.Run() // listen and serve on 0.0.0.0:8080
}
