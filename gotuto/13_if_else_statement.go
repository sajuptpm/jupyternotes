/*
  IF ELSE conditional statement
*/
package main

import "fmt"

func main13() {

	/////////Simple IF statement/////////
	/*
		Syntax
		if(boolean_expression) {
				//do something
		}
	*/

	if 1 < 2 {
		fmt.Println("Condition is TRUE")
	}

	condition := 1 < 2

	if condition {
		fmt.Println("Condition is TRUE")
	}

	if !condition {
		fmt.Println("Condition is FALSE")
	}

	/////////IF ELSE statement/////////
	/*
				Syntax
				if(boolean_expression) {
						//do something
				} else {
					//do something
		}
	*/

	num := 10

	if num%2 != 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	/////////Logical Operators in IF statement /////////
	// AND
	a := 10 < 20
	b := 10 > 5

	if a && b {
		fmt.Println("10 is less than 20 AND grater than 5")
	}

	if 10 < 20 && 10 > 5 {
		fmt.Println("10 is less than 20 AND grater than 5")
	}

	// OR
	if 10 < 20 || 10 < 5 {
		fmt.Println("10 is less than 20 OR less than 5")
	}

}
