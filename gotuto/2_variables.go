/*
  Variables
*/

package main

import "fmt"

func main2() {
	//Create a variable with type and initlialize it
	var age int = 30

	//Create a variable and let auto detect type
	name := "sam"

	//Declare a variable
	var address string

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(address)

	//Print type of Variables
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", address)

	//Change value of Variables
	age = 40
	name = "john"

	fmt.Println(age)
	fmt.Println(name)

	//Initialize multiple variables in one line
	salary, sex := 1000, "Male"
	fmt.Println(salary)
	fmt.Println(sex)

}
