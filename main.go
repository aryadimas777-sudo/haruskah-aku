package main

import (
	"haruskah-aku/handlers"
	"haruskah-aku/repositories"
	"haruskah-aku/services"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// biar random beda tiap run
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()

	repo := &repositories.JawabanRepository{}
	service := services.NewJawabanService(repo)
	handler := handlers.NewJawabanHandler(service)

	r.GET("/tanya", handler.GetJawaban)

	r.GET("/tanya-pasti", handlers.GetJawabPasti)

	r.Run()
}
