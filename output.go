package main

import "fmt"

func main() {
	// difference between Print() and Println()
	var i, j string = "Hello", "World"
	var a, b int = 1, 2

	fmt.Print(i)

	// print with new line after
	fmt.Println(j)

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	// add space between text
	fmt.Print(i, " ", j, "\n")

	// Print will add space if neither strings
	fmt.Print(i, j, "\n")
	fmt.Print(a, b, "\n")

	// Println add new line and add whitespace and can add new format
	fmt.Println(i, j, "\n")

	// Printf
	fmt.Printf("this is value of i %s \n", i)
	fmt.Printf("this is value of j %s \n", j)

	// Formatting verbs
	/*
		Verb	Description
		%v	Prints the value in the default format
		%#v	Prints the value in Go-syntax format
		%T	Prints the type of the value
		%%	Prints the % sign
	*/

	exampleNumber, exampleString := 16.5, "Example"

	fmt.Printf("value of a %v \n", exampleNumber)
	fmt.Printf("value of a %#v \n", exampleNumber)
	fmt.Printf("type of a %T \n", exampleNumber)
	fmt.Printf("type of sign %#v%% \n", exampleNumber)

	fmt.Println()
	fmt.Println("===============")
	fmt.Printf("%v \n", exampleString)
	fmt.Printf("%#v \n", exampleString)
	fmt.Printf("%T \n", exampleString)

	// Formatting int number
	/*
		%b	Base 2
		%d	Base 10
		%+d	Base 10 and always show sign
		%o	Base 8
		%O	Base 8, with leading 0o
		%x	Base 16, lowercase
		%X	Base 16, uppercase
		%#x	Base 16, with leading 0x
		%4d	Pad with spaces (width 4, right justified)
		%-4d	Pad with spaces (width 4, left justified)
		%04d	Pad with zeroes (width 4
	*/

	var newExampleNumber int = 10

	fmt.Println("===============")
	fmt.Printf("Base 2 %b \n", newExampleNumber)
	fmt.Printf("Base 10 %d \n", newExampleNumber)
	fmt.Printf("Base 10 and always show sign %+d \n", newExampleNumber)
	fmt.Printf("Base 8 %o \n", newExampleNumber)
	fmt.Printf("Base 8, with leading Go %O \n", newExampleNumber)
	fmt.Printf("Base 16, lowercase %x \n", newExampleNumber)
	fmt.Printf("Base 16, uppercase %X \n", newExampleNumber)
	fmt.Printf("Base 16, with leading 0x %#x \n", newExampleNumber)
	fmt.Printf("Pad with spaces (witdh 4, right justified) %4d \n", newExampleNumber)
	fmt.Printf("Pad wirh spaces (width 4, left justified) %-4d \n", newExampleNumber)
	fmt.Printf("Pad with zeroes (width 4 %04d \n", newExampleNumber)

	// Formating strings
	/*
		%s	Prints the value as plain string
		%q	Prints the value as a double-quoted string
		%8s	Prints the value as plain string (width 8, right justified)
		%-8s	Prints the value as plain string (width 8, left justified)
		%x	Prints the value as hex dump of byte values
		% x	Prints the value as hex dump with spaces
	*/

	fmt.Println("===============")
	fmt.Printf("%s \n", exampleString)
	fmt.Printf("%q \n", exampleString)
	fmt.Printf("%8s \n", exampleString)
	fmt.Printf("%-8s \n", exampleString)
	fmt.Printf("%x \n", exampleString)
	fmt.Printf("% x \n", exampleString)

	// Formatting boolean
	/*
		%t	Value of the boolean operator in true or false format (same as using %v)
	*/

	fmt.Println("===============")
	fmt.Printf("%t \n", true)
	fmt.Printf("%t \n", false)

	// Formatting float
	/*
	%e	Scientific notation with 'e' as exponent
	%f	Decimal point, no exponent
	%.2f	Default width, precision 2
	%6.2f	Width 6, precision 2
	%g	Exponent as needed, only necessary digits
	*/

	fmt.Println("===============")
	fmt.Printf("%e \n", exampleNumber)
	fmt.Printf("%f \n", exampleNumber)
	fmt.Printf("%.2f \n", exampleNumber)
	fmt.Printf("%6.2f \n", exampleNumber)
	fmt.Printf("%g \n", exampleNumber)
}
