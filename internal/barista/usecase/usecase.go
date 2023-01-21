package usecase

type BaristaUsecase struct {
	barista.PgRepository
}

// GetName returns the name of the barista.
func (u *BaristaUsecase) GetName() string {
	return u.PgRepository.GetName()
}
