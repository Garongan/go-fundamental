/*
int- stores integers (whole numbers), such as 123 or -123
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
string - stores text, such as "Hello World". String values are surrounded by double quotes
bool- stores values with two states: true or false
*/

package main

import "fmt"

var a int
var b int = 2

func main() {
	var name string = "alvindo"
	var age int = 24
	var weight float32 = 65.6
	var isMarried bool = false

	fmt.Println(name, age, weight, isMarried)

	// variable withour initial value
	var newName string // default value is ""
	var newAge int // default valu is 0
	var newWeight float32 // default value is 0
	var newIsMarried bool // default value is false
	var score int = -10

	fmt.Println(newName, newAge, newWeight, newIsMarried, score)

	// variable split declaration
	var address string
	address = "Jl. Ahmad Dahlan"

	fmt.Println(address)

	// difference between var and :=
	// var can declare sparately, inside or outside function
	// := must declare initial value, and can't declare outside function
	c := 5
	fmt.Println(a, b, c)

	// variable with multiple declaration
	var a, b, d, e int = 1, 2, 4, 5
	var height, houseName = 170, "My Sweet House"
	j, l := 5, 6

	fmt.Println(a, b, d, e)
	fmt.Println(height, houseName)
	fmt.Println(j, l)

	// variable declaration block
	var (
		x int
		jackpotValue int = 200
		jackpotInformation = "max win"
	)

	fmt.Println(x, jackpotValue, jackpotInformation)
}
