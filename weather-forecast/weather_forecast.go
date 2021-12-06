// Package weather provides weather conditions.
package weather

// CurrentCondition of weahter.
var CurrentCondition string

// CurrentLocation of weather.
var CurrentLocation string

// Forecast the Weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
