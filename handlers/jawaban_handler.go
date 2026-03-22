package handlers

import (
	"haruskah-aku/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JawabanHandler struct {
	service *services.JawabanService
}

func NewJawabanHandler(s *services.JawabanService) *JawabanHandler {
	return &JawabanHandler{service: s}
}

func (h *JawabanHandler) GetJawaban(c *gin.Context) {
	result := h.service.Jawab()

	c.JSON(http.StatusOK, result)
}

func GetJawabPasti(c *gin.Context) {
	result := services.JawabPasti()

	c.JSON(http.StatusOK, result)
}
