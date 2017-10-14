/*
  Print Triangle Pattern using NESTED FOR LOOP
*/

package main

import "fmt"

func main18() {

	innerLoopControl := 1

	for i := 1; i <= 5; i++ { //execute 5 times

		fmt.Println("") //print new line

		for j := 1; j <= innerLoopControl; j++ { //number of execution vary based on the value of variable "innerLoopControl".
			fmt.Print("*")
		}

		innerLoopControl++

	}

}
