/*
  Local and Global variables
*/
package main

import "fmt"

var gname = "My Global name"

//name := "My Global name"

func main4() {
	var lname = "My Local name"

	//print variables
	fmt.Println(lname)
	fmt.Println(gname)

	//Print type of Variables
	fmt.Printf("%T\n", lname)
	fmt.Printf("%T\n", gname)

	//call another function
	newfun1()
}

func newfun1() {
	//fmt.Println(lname)
	fmt.Println(gname)
}
