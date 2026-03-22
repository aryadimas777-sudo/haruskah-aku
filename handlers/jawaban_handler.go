package handlers

import (
	"haruskah-aku/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJawaban(c *gin.Context) {
	result := services.Jawab()

	c.JSON(http.StatusOK, result)
}

func GetJawabPasti(c *gin.Context) {
	result := services.JawabPasti()

	c.JSON(http.StatusOK, result)
}
