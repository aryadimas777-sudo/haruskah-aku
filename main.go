package main

import (
	"haruskah-aku/handlers"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// biar random beda tiap run
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()

	r.GET("/tanya", handlers.GetJawaban)

	r.Run()
}
