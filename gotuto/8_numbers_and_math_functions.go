/*
  Numbers and math package
*/

package main

import "fmt"
import "math"

func main8() {

	var a = 5.0
	var b = 10.0

	//add Numbers
	var r = a + b
	fmt.Println("Sum: ", r)

	//Modulus
	var m = 10 % 6
	fmt.Println("Modulus: ", m)

	//Max
	var max = math.Max(a, b)
	fmt.Println("Max: ", max)

	//main
	var min = math.Min(a, b)
	fmt.Println("Min: ", min)

	//print variable
	//fmt.Println(a)

	//Print type of Variable
	//fmt.Printf("%T\n", b)

}
