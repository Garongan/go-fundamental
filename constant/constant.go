package main

import "fmt"

// untype constant
const PI = 3.14
// type constant
const BASE_URL string = "https://google.com"

func main()  {
	fmt.Println(PI)
	fmt.Println(BASE_URL)

	// const variable can't re assign
	const Z = 1
	// Z = 2 // cannot assign to Z (neither addressable nor a map index expression)

	// multiple const declaration
	const (
		A string = "A"
		B = "B"
	)

	fmt.Println(A, B)
}