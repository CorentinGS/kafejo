package barista

type BaristaUsecase struct {
	PgRepository
}

// GetName returns the name of the barista.
func (u *BaristaUsecase) GetName() string {
	return u.PgRepository.GetName()
}
