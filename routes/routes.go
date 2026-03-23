// routes/routes.go
package routes

import (
	"haruskah-aku/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handlers.JawabanHandler) {
	JawabanRoutes(r, h)
}
