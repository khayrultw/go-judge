package controllers

import (
	"gorm.io/gorm"

	"github.com/khayrultw/go-judge/database"
	"github.com/khayrultw/go-judge/models"

	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

type CodeRepo struct {
	Db *gorm.DB
}

func New() *CodeRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Code{})
	return &CodeRepo{Db: db}
}

func (repository *CodeRepo) PostCode(c *gin.Context) {
	var code models.Code
	c.BindJSON(&code)
	err := models.PostCode(repository.Db, &code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, code)
}

func (repository *CodeRepo) GetCode(c *gin.Context) {
	var code models.Code
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetCode(repository.Db, &code, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, code)
}
