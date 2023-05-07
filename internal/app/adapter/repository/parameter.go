package repository

import (
	"payday-savings/m/internal/app/adapter/postgresql"
	"payday-savings/m/internal/app/adapter/postgresql/model"
	"payday-savings/m/internal/app/domain"
)

// Parameter is the repository of domain.Parameter
type Parameter struct{}

// Get gets parameter
func (r Parameter) Get() domain.Parameter {
	db := postgresql.Connection()
	var param model.Parameter
	result := db.First(&param, 1)
	if result.Error != nil {
		panic(result.Error)
	}
	return domain.Parameter{
		Funds: param.Funds,
		Btc:   param.Btc,
	}
}
