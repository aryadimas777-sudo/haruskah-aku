package main

import (
	"fmt"
	"math/rand"
	"time"
)

// function random jawaban
func Jawab() string {
	// list pilihan
	jawaban := []string{"ya", "tidak"}

	// random index -> ambil angka random dari 0 sampai panjang slice - 1
	index := rand.Intn(len(jawaban))

	return jawaban[index]
}

func main() {
	// biar random beda tiap run
	rand.Seed(time.Now().UnixNano())

	fmt.Println("jawaban:", Jawab())
}
