// Package weather contains function to
// display weather forecast.
package weather

// CurrentCondition represents weather condition in specific location.
var CurrentCondition string

// CurrentLocation represents location where weather was calculated.
var CurrentLocation string

// Forecast returns string with information about weather condition in provided location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
