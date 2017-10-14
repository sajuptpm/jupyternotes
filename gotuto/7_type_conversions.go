/*
  Type conversions
*/

package main

import "fmt"

func main7() {
	//The expression T(v) converts the value v to the type T.
	var i = 20
	//print variable
	fmt.Println(i)
	//Print type of Variable
	fmt.Printf("%T\n", i)

	var f = float64(i)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

	var u = uint(f)
	fmt.Println(u)
	fmt.Printf("%T\n", u)

}
