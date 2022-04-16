package dinamic_discounts_calculator

import "testing"

type DiscountRepositoryMock struct{}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

// Table driven test
func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{name: "should apply 20", minimumPurchaseAmount: 100, purchaseAmount: 150, expectedAmount: 130},
		{name: "should apply 40", minimumPurchaseAmount: 100, purchaseAmount: 50, expectedAmount: 50},
	}

	t.Log("Only shown when running with -v")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // Creating a sub-test
			// Arrange
			discountRepositoryMock := DiscountRepositoryMock{}
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)
			if err != nil {
				t.Fatalf("Could not create the simple_calculator %v\n", err) // It stops the test immediately
			}

			// Act
			amount := calculator.Calculate(tc.purchaseAmount)

			// Assert
			if tc.expectedAmount != amount {
				//t.Logf("expected 50, got %v\n", amount)
				//t.Fail() // Manual assertion
				t.Errorf("expected %v, got %v\n", tc.expectedAmount, amount) // Error = Log + Fail
			}
		})
	}
}
