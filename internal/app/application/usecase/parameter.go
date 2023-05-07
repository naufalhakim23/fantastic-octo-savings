package usecase

import (
	"payday-savings/m/internal/app/domain"
	"payday-savings/m/internal/app/domain/repository"
)

// Parameter is the usecase of getting parameter
func Parameter(r repository.IParameter) domain.Parameter {
	return r.Get()
}
