package repository

import "payday-savings/m/internal/app/domain"

// IOrder is interface of order repository
type IOrder interface {
	Get() domain.Order
	Update(domain.Order)
}
