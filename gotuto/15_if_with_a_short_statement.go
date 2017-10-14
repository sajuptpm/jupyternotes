/*
  IF with a Short statement
*/

package main

import "fmt"

func main15() {
	//IF statement can start with a short statement to execute before the condition.
	//It first execute short statement Before checking the condition.

	if a := 10; a == 10 {
		fmt.Println("Number is 10")
	}
	//Variables declared by the statement are only in scope until the end of the if.

	//Try to print "a" - Error undefined: a
	//fmt.Println(a)

	if num := 10; num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

}
