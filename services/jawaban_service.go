package services

import (
	"haruskah-aku/repostiories"
	"math/rand"
)

type Response struct {
	Jawaban string
}

// function random jawaban
func Jawab() Response {
	// list pilihan
	list := repostiories.GetJawabanList()

	// random index -> ambil angka random dari 0 sampai panjang slice - 1
	index := rand.Intn(len(list))

	return Response{
		Jawaban: list[index],
	}
}

func JawabPasti() Response {
	return Response{Jawaban: "ya"}
}
