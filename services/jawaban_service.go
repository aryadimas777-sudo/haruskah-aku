package services

import "math/rand"

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
