package lasagna

import "fmt"

// TODO: define the 'PreparationTime()' function
// PreparationTime returns the cooking time for lasagna depending
// on the number of layers and the average cooking time per layer.
func PreparationTime(layers []string, timeForLayers int) int {

	if timeForLayers == 0 {
		return len(layers) * 2
	}

	return len(layers) * timeForLayers
}

// TODO: define the 'Quantities()' function
// Quantities determines the quantity of noodles and sauce needed to make meal.
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, elem := range layers {
		if elem == "noodles" {
			noodles += 50
		} else if elem == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(slices []float64, count int) []float64 {
	fmt.Println(len(slices))
	quantities := make([]float64, len(slices))

	for _, elem := range slices {
		quantities = append(quantities, (elem/2)*float64(count))
	}

	return quantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
