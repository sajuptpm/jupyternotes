/*
  Boolean Data Type and Logical Operators
*/

package main

import "fmt"

func main11() {

	///////Boolean data type has two values (usually denoted true and false)///////
	p := true
	q := false

	//print variables
	fmt.Println(p)
	fmt.Println(q)

	//Print type of Variables
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", q)

	fmt.Println("p == q: ", p == q)
	fmt.Println("1 == 1: ", 1 == 1)
	fmt.Println("1 == '1': ", 1 == '1')

	//fmt.Println("1 == '1': ", true == 1)

	///////Logical Operators AND (&&), OR (||), NOT (!)///////
	//Logical operators are typically used with Boolean (logical) values
	a := true
	b := false
	c := true
	d := false

	//AND
	fmt.Println("a && c: ", a && c)
	fmt.Println("1==1 && 2==2: ", 1 == 1 && 2 == 2)
	//OR
	fmt.Println("a || b: ", a || b)
	fmt.Println("b || d: ", b || d)
	//NOT
	fmt.Println("!a: ", !a)
	fmt.Println("!b: ", !b)

}
