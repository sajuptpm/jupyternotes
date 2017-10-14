/*
  Declaring variables in Go without specifying type.
  Short variable declarations
*/
package main

import "fmt"

func main3() {

	var age = 20
	name := "sam"

	//print variables
	fmt.Println(age)
	fmt.Println(name)

	//Print type of Variables
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", name)

}
