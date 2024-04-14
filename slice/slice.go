package main

import (
	"fmt"
	"strings"
)

func main() {
	gap := strings.Repeat("-", 50)
	fmt.Println(gap)

	months := [...]string{
		"JANUARY",
		"FEBRUARY",
		"MARCH",
		"APRIL",
		"MAY",
		"JUNE",
		"JULY",
		"AUGUST",
		"SEPTEMBER",
		"OCTOBER",
		"NOVEMBER",
		"DECEMBER",
	}
	/*
		if we change the value of array it will change the value of slice
		if we change the value of slice it will change the value of array
	*/
	slice1 := months[0:2] // slice the array from start index to end index
	fmt.Println("my slice 3: ", slice1)
	fmt.Println("my slice 3 length: ", len(slice1))
	fmt.Println("my slice 3 cap: ", len(slice1))
	fmt.Println("example array index from first to specific index: ", months[:2])
	fmt.Println("example array index from specific index to last index: ", months[2:])

	//	create slice with make function
	/*
		syntax
		slice_name := make([]type, length, capacity)
		len() function - returns the length of the slice (the number of elements in the slice)
		cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	*/
	fmt.Println(gap)
	mySlice := make([]int, 5, 10)
	fmt.Printf("my slice %#v \n", mySlice)
	fmt.Printf("my length %d \n", len(mySlice))
	fmt.Printf("my capacity %d \n", cap(mySlice))
	fmt.Println(gap)

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

	//	change the length of slice by append the array
	fmt.Println(gap)
	mySlice4 := append(slice1, "4", "5")
	fmt.Println("my slice 4: ", mySlice4)
	fmt.Println("my slice 4 length: ", len(mySlice4))
	fmt.Println("my slice 4 capacity: ", len(mySlice4))
	fmt.Println(gap)

	/*
		if slicing index is inside the length it will replace the original array
		if slicing index is equals last index it will create new array
	*/

	//monthSlices := months[0:]
	monthSlices := months[11:]
	newMonthSlices := append(monthSlices, "juli")
	fmt.Println("slicing array months", newMonthSlices)

	// change value in slice
	newMonthSlices[0] = "februari"
	fmt.Println("original months: ", months)
	fmt.Println("new slices months: ", newMonthSlices)

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

	fmt.Println(gap)

	intArray := [...]int{1, 2, 3, 4, 5}
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(intArray)
	fmt.Println(intSlice)

}
