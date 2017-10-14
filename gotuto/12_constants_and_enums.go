/*
  Constants and Enums
*/
package main

import "fmt"

func main12() {
	/*
	   Constants are created at compile time, even when defined as locals in functions,
	   and can only be numbers, characters (runes), strings or booleans.
	*/
	//Declare and and initlialize a constant variable
	const name string = "Sam"

	/*####Enums####*/
	//Enums used to define collections of constants
	//Enums come by putting a bunch of Constants in parentheses
	const (
		One   = 1
		Two   = 2
		Three = 3
	)
	//print variables
	fmt.Println("\nOne: ", One)
	fmt.Println("Two: ", Two)
	fmt.Println("Three: ", Three)
	fmt.Println("\n\n")

	//You can use "iota" to create these values for constants.
	//Within a const group or parentheses, it starts at 0, and
	//then increments for each expression.So we dont need to
	//assign value to all Constants in the parentheses.

	//Set values of constants to 0, 1 and 2
	const (
		a = iota
		b
		c
	)
	//print variables
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)

	//Set values of constants to to 5, 6 and 7
	const (
		p = iota + 5 //means 0 + 5
		q
		r
	)
	//print variables
	fmt.Println("p: ", p)
	fmt.Println("q: ", q)
	fmt.Println("r: ", r)

}
