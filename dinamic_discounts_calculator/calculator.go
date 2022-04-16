package dinamic_discounts_calculator

import (
	"errors"
	"go-unit-tests-course/database"
)

type DiscountCalculator struct {
	MinimumPurchaseAmount int
	DiscountRepository    database.Repository
}

func NewDiscountCalculator(minimumPurchaseAmount int, discountRepository database.Repository) (*DiscountCalculator, error) {

	if minimumPurchaseAmount <= 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount must be greater than zero")
	}

	return &DiscountCalculator{
		MinimumPurchaseAmount: minimumPurchaseAmount,
		DiscountRepository:    discountRepository,
	}, nil
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {

	discount := c.DiscountRepository.FindCurrentDiscount()

	if purchaseAmount > c.MinimumPurchaseAmount {
		return purchaseAmount - discount
	}
	return purchaseAmount
}
