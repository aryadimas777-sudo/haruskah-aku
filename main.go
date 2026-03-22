package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Response struct {
	Jawaban string
}

// function random jawaban
func Jawab() Response {
	// list pilihan
	jawaban := []string{"ya", "tidak", "mungkin"}

	// random index -> ambil angka random dari 0 sampai panjang slice - 1
	index := rand.Intn(len(jawaban))

	return Response{
		Jawaban: jawaban[index],
	}
}

func main() {
	// biar random beda tiap run
	rand.Seed(time.Now().UnixNano())

	result := Jawab()

	fmt.Println("jawaban:", result)
}
