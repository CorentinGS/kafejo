package barista

import "gorm.io/gorm"

type BaristaPgRepository struct {
	DBConn *gorm.DB
}

// GetName returns the name of the repository.
func (r *BaristaPgRepository) GetName() string {
	return "PgRepository"
}
