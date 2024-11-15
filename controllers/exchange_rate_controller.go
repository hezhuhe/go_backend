package controllers

import (
	"backend/global"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateExchangeRate(ctx *gin.Context) {
	var ExchangeRate models.ExchangeRate
	if err := ctx.ShouldBindJSON(&ExchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "json error",
		})
		return
	}
	ExchangeRate.Date = time.Now()

	if err := global.Db.AutoMigrate(&ExchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error inititalzate model",
		})
		return
	}

	if err := global.Db.Create(&ExchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error create exchangeRate",
		})
		return
	}
	ctx.JSON(http.StatusOK, ExchangeRate)
}

func GetExchangeRate(ctx *gin.Context) {

}