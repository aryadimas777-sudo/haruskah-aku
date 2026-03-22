package repositories

type JawabanRepository struct{}

func (r *JawabanRepository) GetJawabanList() []string {
	return []string{"ya", "tidak", "mungkin"}
}
