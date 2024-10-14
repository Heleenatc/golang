package main

import "fmt"

func main() {
	//Strings
	var nameOne string = "Heleena"
	var nameTwo = "Thanka"
	var nameThree string

	nameThree = "Christlet"
	fmt.Println(nameOne, nameTwo, nameThree)

	//integers
	var ageOne int = 40
	var ageTwo = 100
	var ageThree = 12

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 25
	var numTwo int8 = -100
	var numThree uint16 = 256

	fmt.Println(numOne, numTwo, numThree)

	//floats
	var scoreOne float32 = 98.7
	var scoreTwo float64 = 98.70000000022222222444

	fmt.Println(scoreOne, scoreTwo)

}
