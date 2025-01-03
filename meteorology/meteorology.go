package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
	units := [...]string{"°C", "°F"}
	return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// String method to SpeedUnit
func (su SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// String method to MeteorologyData
func (m MeteorologyData) String() string {
	response := fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",
		m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
	return response
}
