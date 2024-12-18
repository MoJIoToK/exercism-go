package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood function determines the number of cows consumed for food. The functions takes the FodderCalculator
// interface and the number of cows. The function returns the amount of food for one cow and an error.
// The amount of food for one cow defines as the amount of food for all cows(method FodderAmount(count)) multiplied
// by the ratio(method FatteningFactor) divided by the count of cows.
func DivideFood(calculator FodderCalculator, count int) (float64, error) {
	fAmount, err := calculator.FodderAmount(count)
	if err != nil {
		return 0, err
	}

	fFactor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return (fAmount * fFactor) / float64(count), nil
}

// ValidateInputAndDivideFood is function to check the correctness of the entered number of cows. Returns the results
// of function DivideFood.
func ValidateInputAndDivideFood(calculator FodderCalculator, count int) (float64, error) {
	if count <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calculator, count)
}

// InvalidCowsError - custom error struct.
type InvalidCowsError struct {
	countCows int
	errorMsg  string
}

// Error method for custom error.
func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %v", err.countCows, err.errorMsg)
}

// ValidateNumberOfCows function checks the correctness of the entered number of cows with custom error.
func ValidateNumberOfCows(count int) *InvalidCowsError {
	switch {
	case count < 0:
		return &InvalidCowsError{
			countCows: count,
			errorMsg:  "there are no negative cows",
		}
	case count == 0:
		return &InvalidCowsError{
			countCows: count,
			errorMsg:  "no cows don't need food",
		}
	default:
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
