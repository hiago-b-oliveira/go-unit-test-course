package simple_calculator

import "testing"

// Replaced by test table
func TestDiscountApplied(t *testing.T) {
	// Arrange
	calculator, _ := NewDiscountCalculator(100, 20)

	// Act
	amount := calculator.Calculate(150)

	// Assert
	if 130 != amount {
		t.Fail() // Manual assertion
	}
}

// Replaced by test table
func TestDiscountNotApplied(t *testing.T) {
	// Arrange
	calculator, _ := NewDiscountCalculator(100, 20)

	// Act
	amount := calculator.Calculate(50)

	// Assert
	if 50 != amount {
		t.Errorf("expected 50, got %v\n", amount)
	}
}

// Table driven test
func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{name: "should apply 20", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 150, expectedAmount: 130},
		{name: "should apply 40", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 50, expectedAmount: 50},
	}

	t.Log("Only shown when running with -v")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // Creating a sub-test
			// Arrange
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
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

func TestDiscountCalculatorShouldFailWithZero(t *testing.T) {
	// Arrange / Act
	_, err := NewDiscountCalculator(0, 20)

	// Assert
	if err == nil {
		t.Fatalf("Should not creat the simple_calculator with zero purchase amount")
	}
}
