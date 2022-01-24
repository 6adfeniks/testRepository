package otherfiles

//import (
//	"strings"
//)
//
//// WelcomeMessage returns a welcome message for the customer.
//func WelcomeMessage(customer string) string {
//	return 	"Welcome to the Tech Palace, " + strings.ToUpper(customer)
//}
//
//// AddBorder adds a border to a welcome message.
//func AddBorder(welcomeMsg string, numStarsPerLine int) string {
//	startEnd := ""
//	for i:=0;i<numStarsPerLine;i++{
//		startEnd += "*"
//	}
//	return startEnd + "\n" + welcomeMsg + "\n" + startEnd
//}
//
//// CleanupMessage cleans up an old marketing message.
//func CleanupMessage(oldMsg string) string {
//	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
//	oldMsg = strings.TrimSpace(oldMsg)
//	oldMsg = strings.ReplaceAll(oldMsg, "\n", "")
//
//	return oldMsg
//}
//
///////////////////////////////////////////////////////////////////////////////////////////

//type Car struct{
//	battery int
//	batteryDrain int
//	speed int
//	distance int
//}
//// NewCar creates a new remote controlled car with full battery and given specifications.
//// Output: Car{speed: 5, batteryDrain: 2, battery:100, distance: 0}
//func NewCar(speed, batteryDrain int) Car {
//	return Car{100, batteryDrain, speed, 0}
//}
//
//type Track struct{
//	distance int
//}
//
//// NewTrack created a new track
//func NewTrack(distance int) Track {
//	return Track{ distance}
//}
//
//// Drive drives the car one time. If there is not enough battery to drive on more time,
//// the car will not move.
//func Drive(car Car) Car {
//	car.battery -= car.batteryDrain
//	car.distance += car.speed
//	return car
//}
//
//// CanFinish checks if a car is able to finish a certain track.
//func CanFinish(car Car, track Track) bool {
//	if car.battery/car.batteryDrain*car.speed > track.distance{
//		return true
//	}
//	return false
//}

///////////////////////////////////////////////////////////

//// GetItem retrieves an item from a slice at given position. The second return value indicates whether
//// the given index exists in the slice or not.
//func GetItem(slice []int, index int) (int, bool) {
//	if index < 0 || index > len(slice)-1{
//		return 0, false
//	}
//	return slice[index], true
//}
//
//// SetItem writes an item to a slice at given position overwriting an existing value.
//// If the index is out of range the value needs to be appended.
//func SetItem(slice []int, index, value int) []int {
//	if index < 0 || index > len(slice)-1{
//		slice = append(slice, value)
//		return slice
//	}
//	slice[index] = value
//	return slice
//}
//
//// PrefilledSlice creates a slice of given length and prefills it with the given value.
//func PrefilledSlice(value, length int) []int {
//	slice := make([]int,0)
//	for i:=0; i <length; i++ {
//		slice = append(slice, value)
//	}
//	return slice
//}
//
//// RemoveItem removes an item from a slice by modifying the existing slice.
//func RemoveItem(slice []int, index int) []int {
//
//	if index < 0 || index > len(slice)-1{
//		return slice
//	}
//	sliceNew := slice[:index]
//	sliceNew = append(sliceNew, slice[index+1:]...)
//	return sliceNew
//}
////////////////////////////////////////////////////////////
