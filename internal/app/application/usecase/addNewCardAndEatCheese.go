package usecase

import (
	"github.com/google/uuid"
	"payday-savings/m/internal/app/domain"
	"payday-savings/m/internal/app/domain/repository"
	"payday-savings/m/internal/app/domain/valueobject"
)

// AddNewCardAndEatCheese updates payment card and jerry's weight
func AddNewCardAndEatCheese(o repository.IOrder) domain.Order {
	order := o.Get()
	newCardBrand := valueobject.VISA
	if order.Payment.Card.Brand == valueobject.VISA {
		newCardBrand = valueobject.AMEX
	}
	newCard := valueobject.Card{
		ID:    uuid.New().String(),
		Brand: newCardBrand,
	}
	order.Person.Weight++
	order.Payment.Card = newCard
	o.Update(order)
	return order
}
