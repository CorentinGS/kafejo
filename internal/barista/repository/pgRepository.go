package repository

import "gorm.io/gorm"

type PgRepository struct {
	DBConn *gorm.DB
}

// GetName returns the name of the repository.
func (r *PgRepository) GetName() string {
	return "PgRepository"
}
