package routes

import (
	"haruskah-aku/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, handler *handlers.JawabanHandler) {
	r.GET("tanya", handler.GetJawaban)
	r.GET("/tanya-pasti", handler.GetJawabPasti)
}
