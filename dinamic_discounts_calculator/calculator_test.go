package dinamic_discounts_calculator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type DiscountRepositoryMock struct {
	mock.Mock
}

func (drm *DiscountRepositoryMock) FindCurrentDiscount() int {
	args := drm.Called()
	return args.Int(0)
}

// Table driven test
func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		purchaseAmount        int
		discount              int
		expectedAmount        int
	}

	testCases := []testCase{
		{name: "should apply 20", minimumPurchaseAmount: 100, purchaseAmount: 150, discount: 20, expectedAmount: 130},
		{name: "should apply 40", minimumPurchaseAmount: 100, purchaseAmount: 50, discount: 20, expectedAmount: 50},
		{name: "should apply 40", minimumPurchaseAmount: 100, purchaseAmount: 50, discount: 0, expectedAmount: 50},
	}

	t.Log("Only shown when running with -v")

	discountRepositoryMock := &DiscountRepositoryMock{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // Creating a sub-test
			// Arrange
			discountRepositoryMock.On("FindCurrentDiscount").Return(tc.discount)
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

			assert.Equal(t, tc.expectedAmount, amount) // it already gives us a nice error message
		})
	}
	discountRepositoryMock.AssertNumberOfCalls(t, "FindCurrentDiscount", len(testCases))
}
