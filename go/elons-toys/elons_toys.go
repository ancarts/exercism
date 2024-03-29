package elon
import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance = car.distance + car.speed*1
		car.battery = car.battery - car.batteryDrain
	}
	
}
// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	messageToDisplay := "Driven " + fmt.Sprint(car.distance) + " meters" 
	return messageToDisplay
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	messageToDisplay := "Battery at " + fmt.Sprint(car.battery) + "%" 
	return messageToDisplay
}
// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	remainingDistance := car.speed*car.battery/car.batteryDrain
	if remainingDistance >= trackDistance {
		return true
	}
	return false
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
