package simple_calculator

import "errors"

type DiscountCalculator struct {
	MinimumPurchaseAmount int
	DiscountAmount        int
}

func NewDiscountCalculator(minimumPurchaseAmount int, discountAmount int) (*DiscountCalculator, error) {

	if minimumPurchaseAmount <= 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount must be greater than zero")
	}

	return &DiscountCalculator{
		MinimumPurchaseAmount: minimumPurchaseAmount,
		DiscountAmount:        discountAmount,
	}, nil
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {

	if purchaseAmount > c.MinimumPurchaseAmount {
		return purchaseAmount - c.DiscountAmount
	}
	return purchaseAmount
}
