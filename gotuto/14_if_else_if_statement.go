/*
  IF ELSE IF statement
*/
package main

import "fmt"

func main14() {

	//IF ELSE IF statement
	/*
		Syntax
		if(boolean_expression1) {
				//do something
		} else if (boolean_expression2){
				//do something
		}else {
			//do something
		}
	*/
	a := 7

	if a == 2 {
		fmt.Println("Number is 2")
	} else if a == 5 {
		fmt.Println("Number is 5")
	} else {
		fmt.Println("Unknown Number")
	}

}
