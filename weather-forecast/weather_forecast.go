// Package weather provides function to forecast weather.
// according to location.
package weather

// CurrentCondition describes the condition of weather.
var CurrentCondition string

// CurrentLocation describes the location of Goblinocus.
var CurrentLocation string

// Forecast returns the weather condition to given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
