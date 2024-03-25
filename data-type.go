package main

import (
	"fmt"
	s "strings"
)

func main() {
	// number data type
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)

	// boolean data type
	fmt.Println("Benar ", true)
	fmt.Println("Salah ", false)

	// string data type
	fmt.Println("Alvindo")
	fmt.Println("Tri")
	fmt.Println("Jatmiko")

	// function of string
	fmt.Println(len("Alvindo"))
	fmt.Println("Alvindo"[0])

	// variable
	var name string

	name = "Alvindo"
	fmt.Println(name)

	fullName := "Alvindo Tri Jatmiko"
	fmt.Println(fullName)

	var age = 24
	fmt.Println(age)

	isMarried := false
	fmt.Println(isMarried)

	var (
		firstName = "Alvindo"
		lastName  = "Jatmiko"
	)

	fmt.Println(firstName + " " + lastName)

	// constant variable
	const regency string = "Nganjuk"
	const country = "Indonesia"
	const weight = 65

	fmt.Println(regency, country, weight)

	// declaration multiple const variable
	const (
		village string = "Werungotok"
		height         = 170
	)

	fmt.Println(village, height)

	// data type conversion
	var int32Value = 100000
	var int64Value = int64(int32Value)
	var int8Value = int8(int32Value)

	fmt.Println(int32Value)
	fmt.Println(int64Value)
	fmt.Println(int8Value)

	// string to char conversion
	var newName = "alvindo"
	e := newName[0]
	eToString := string(e)

	fmt.Println(newName)
	fmt.Println(e)
	fmt.Println(eToString)

	//	type declarations
	type nik string
	type married bool

	var myNik nik = "1923304495"
	var marriedStatus married = true

	fmt.Println(myNik)
	fmt.Println(marriedStatus)

	// arithmetic operator
	var a = 10
	var b = 2

	sum := a + b
	fmt.Println(sum)

	sub := a - b
	fmt.Println(sub)

	multiply := a * b
	fmt.Println(multiply)

	divide := a / b
	fmt.Println(divide)

	mod := a % b
	fmt.Println(mod)

	// augmented assignment
	fmt.Println(s.Repeat("=", 50))
	a += 10 // a = a + 10
	fmt.Println(a)

	a -= 50 // a = a - 50
	fmt.Println(a)

	a *= 5 // a = a * 5
	fmt.Println(a)

	a /= 2 // a = a / 2
	fmt.Println(a)

	a %= 4 // a = a % 4
	fmt.Println(a)

	// unary operator
	fmt.Println(s.Repeat("=", 50))
	i := 50
	i++ // i = i + 1
	fmt.Println(i)

	i-- // i = i - 1
	fmt.Println(i)

	negativeValue := -10
	positiveValue := +10
	fmt.Println(negativeValue, positiveValue)

	isMarried = true
	isMarried = !isMarried
	fmt.Println(isMarried)

	// comparing operator
	fmt.Println(s.Repeat("+", 50))
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a == b)
	fmt.Println(a != b)

	name1 := "alvindo"
	name2 := "budi"

	result := name1 == name2
	fmt.Println(result)

	// boolean operator => &&, || ,!
	fmt.Println(s.Repeat("=", 50))

	examValue := 90
	presentFrequent := 70

	graduate := examValue > 70 && presentFrequent > 70
	fmt.Println(graduate)
}
