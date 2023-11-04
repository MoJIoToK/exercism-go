package cars

const (
	//GroupPrice is a price of 10 cars.
	groupPrice int = 95000
	//IndividualPrice is a price of 1 car.
	individualPrice int = 10000
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * successRate / 6000)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	group := carsCount / 10
	individual := carsCount % 10
	return uint(group*groupPrice + individual*individualPrice)
}
