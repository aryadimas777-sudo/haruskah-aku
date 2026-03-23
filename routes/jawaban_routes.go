package routes

import (
	"haruskah-aku/handlers"

	"github.com/gin-gonic/gin"
)

func JawabanRoutes(r *gin.Engine, h *handlers.JawabanHandler) {
	r.GET("tanya", h.GetJawaban)
	r.GET("/tanya-pasti", h.GetJawabPasti)
}
