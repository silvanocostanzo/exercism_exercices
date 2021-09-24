// Package weather provide a fonction for weather forecast
package weather

// CurrentCondition is of type 'string'
var CurrentCondition string

// CurrentLocation is of type 'string'
var CurrentLocation string

// Forecast takes two parameters: 'city' and 'conditions', both of type string
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
