// Package weather provides a Forecast API.
package weather

// CurrentCondition contains the current condition.
var CurrentCondition string

// CurrentLocation contains the current location.
var CurrentLocation string

// Forecast weather for current location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
