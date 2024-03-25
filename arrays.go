package main

import (
	"fmt"
	"strings"
)

func main() {
	// syntax => [length]datatype{values}
	var exampleArrays = [3]int{1, 2, 3}                            // length is defined
	exampleArrays1 := [...]string{"wanwan", "zilong", "minotaour"} // length is not defined

	fmt.Println(exampleArrays)
	fmt.Println(exampleArrays1)

	//	access element of array
	fmt.Println(exampleArrays[0])
	fmt.Println(exampleArrays1[2])

	//	change element of array
	exampleArrays1[0] = "johnson"
	fmt.Println(exampleArrays1)

	// array init
	//default value array is 0
	exampleArrays2 := [5]int{}              // not init
	exampleArrays3 := [5]int{1, 2}          // partially init
	exampleArrays4 := [5]int{1, 2, 3, 4, 5} // fully init

	fmt.Println(exampleArrays2)
	fmt.Println(exampleArrays3)
	fmt.Println(exampleArrays4)

	exampleArrays5 := [5]int{1: 10, 3: 50, 4: 40} // init with specific index

	fmt.Println(exampleArrays5)

	//	get array length
	fmt.Println(len(exampleArrays))
	fmt.Println(len(exampleArrays1))

	//	go slices
	/*
		len() function - returns the length of the slice (the number of elements in the slice)
		cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	*/
	fmt.Println(len(exampleArrays2))
	fmt.Println(cap(exampleArrays2))
	fmt.Println(exampleArrays2)

	fmt.Println(len(exampleArrays1))
	fmt.Println(cap(exampleArrays1)) // slice array to one
	fmt.Println(exampleArrays1)

	gap := strings.Repeat("-", 50)
	fmt.Println(gap)

	//	create slice with make function
	/*
		syntax
		slice_name := make([]type, length, capacity)
	*/

	mySlice := make([]int, 5, 10)
	fmt.Printf("my slice %#v \n", mySlice)
	fmt.Printf("my length %d \n", len(mySlice))
	fmt.Printf("my capacity %d \n", cap(mySlice))

	/*
		Append element to slice
		slice_name = append(slice_name, element1, element2, ...)
	*/

	mySlice1 := append(mySlice, 30, 50)

	fmt.Println(mySlice1)
	fmt.Println(gap)
	fmt.Println("my slice 1 length: ", len(mySlice1))
	fmt.Println("my slice 1 capacity", cap(mySlice1))
	fmt.Println(mySlice1[len(mySlice1)-1]) // access last data in slice
	mySlice1[len(mySlice1)-1] = 70
	fmt.Println(mySlice1[len(mySlice1)-1]) // changes last data in slice

	/*
	   Append to another slice
	   slice3 = append(slice1, slice2...)
	*/

	mySlice2 := append(mySlice, mySlice1...) // only two value
	fmt.Println(mySlice2)

	//	change the length of slice
	mySlice3 := exampleArrays[0:2] // slice the array from start index to end index
	fmt.Println("my slice 3: ", mySlice3)
	fmt.Println("my slice 3 length: ", len(mySlice3))
	fmt.Println("my slice 3 cap: ", len(mySlice3))

	//	change the length of slice by append the array
	fmt.Println(gap)
	mySlice4 := append(mySlice3, 4, 5)
	fmt.Println("my slice 4: ", mySlice4)
	fmt.Println("my slice 4 length: ", len(mySlice4))
	fmt.Println("my slice 4 capacity: ", len(mySlice4))

	/*
		Memory Efficiency
		When using slices, Go loads all the underlying elements into the memory.
		If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
		The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.
	*/

	/*
		syntax
		copy(dest, src)
		The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.
	*/

	fmt.Println(gap)

	mySliceCopy := make([]int, len(mySlice))
	copy(mySliceCopy, mySlice)

	fmt.Println("my slice copy: ", mySliceCopy)
	fmt.Println("my slice copy length: ", len(mySliceCopy))
	fmt.Println("my slice copy capacity: ", cap(mySliceCopy))

}
