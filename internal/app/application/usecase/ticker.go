package usecase

import (
	"payday-savings/m/internal/app/application/service"
	"payday-savings/m/internal/app/domain/valueobject"
)

// Ticker is the usecase of getting ticker
func Ticker(e service.IExchange, p valueobject.Pair) valueobject.Ticker {
	return e.Ticker(p)
}
