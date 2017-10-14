/*
  Declare and Initilize multiple variables at once
*/
package main

import "fmt"

func main6() {

	//Method-1 - Declare and Initilize
	//var name, address string //declare
	var (
		name    string
		address string
	)
	name = "john"
	address = "blabla"

	//Method-2 - Declare and Initilize
	//var a, b int = 1, 2
	var (
		a int = 1
		b int = 2
	)

	//Method-3 - Declare and Initilize
	//x, y, z := 5, "Sam", 6.2
	var (
		x = 5
		y = "Sam"
		z = 6.2
	)

	//Method-4 - Declare and Initilize
	//const p, q = 3, "Foo"
	const (
		p = 3
		q = "Foo"
	)

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(p)
	fmt.Println(q)

	//Print type of Variables
	//fmt.Printf("%T\n", x)
	//fmt.Printf("%T\n", y)
	//fmt.Printf("%T\n", x)

}
