/*
  Constant Variables
*/
package main

import "fmt"

func main5() {
	/*
	   Constants are basically variables whose values cannot be changed later.
	   They are created in the same way you create variables but instead of using
	   the var keyword we use the const keyword
	*/
	const name string = "Sam"
	fmt.Println(name)

	//Modify variable name, Error: cannot assign to name
	//name = "John"

	//not specifying type
	const age = 33
	fmt.Println(age)

	//Print type of Variables
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)

}
