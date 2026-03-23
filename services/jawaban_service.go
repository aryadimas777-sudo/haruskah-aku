package services

import (
	"haruskah-aku/repositories"
	"math/rand"
)

type JawabanService struct {
	repo *repositories.JawabanRepository
}

func NewJawabanService(r *repositories.JawabanRepository) *JawabanService {
	return &JawabanService{repo: r}
}

type Response struct {
	Jawaban string
}

// function random jawaban
func (s *JawabanService) Jawab() Response {
	// list pilihan
	list := s.repo.GetJawabanList()

	// random index -> ambil angka random dari 0 sampai panjang slice - 1
	index := rand.Intn(len(list))

	return Response{
		Jawaban: list[index],
	}
}

func (s *JawabanService) JawabPasti() Response {
	return Response{Jawaban: "ya"}
}
