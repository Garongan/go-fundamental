package main

import (
	"fmt"
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
}
