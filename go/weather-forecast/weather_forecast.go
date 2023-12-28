// Package weather was created to show current weather condition for certain regions in Goblinocus.
package weather

// CurrentCondition shows a weather condition, including temperature (C), atmospheric pressure, 
// wind, humidity, precipitation, and cloudiness.
var CurrentCondition string
// CurrentLocation is a certain city in Goblinocus.
var CurrentLocation string

// Forecast returns a string that shows a current weather condition at certain location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
