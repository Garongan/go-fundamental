package main

import (
	"fmt"
	"strings"
)

func main() {
	gap := strings.Repeat("=", 50)
	fmt.Println(gap)
	/*

		syntax of creating map
		variable := map[key]value{
			"key": value
		}

	*/

	// creating map
	person := map[string]string{
		"name":    "alvindo",
		"age":     "24",
		"address": "jl. topaz 7",
	}

	fmt.Println("map of person: ", person)
	fmt.Println(gap)

	// accessing map by key
	fmt.Println("person name: ", person["name"])
	fmt.Println(gap)

	// adding the key and value to map
	person["gender"] = "male"

	// change the value by key
	person["age"] = "23"
	fmt.Println("map of person", person)
	fmt.Println(gap)

	// get map length of key
	fmt.Println("length of map person: ", len(person))
	fmt.Println(gap)

	/*

		creating map with make function
		syntax: mapName := make(map[typeMap]typeValue)

	*/

	// creating empty map using make
	book := make(map[string]string)

	// add value to book map
	book["title"] = "buku bagus"
	book["year"] = "2002"

	fmt.Println(book)
	fmt.Println(gap)

	// delete the value of map
	delete(book, "year")
	delete(person, "address")
	fmt.Println("book map after deleting: ", book)
	fmt.Println("person map after deleting: ", person)

}
