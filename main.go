package main

import (
	"fmt"
	"haruskah-aku/services"
	"math/rand"
	"time"
)

func main() {
	// biar random beda tiap run
	rand.Seed(time.Now().UnixNano())

	result := services.Jawab()

	fmt.Println("jawaban:", result.Jawaban)
}
