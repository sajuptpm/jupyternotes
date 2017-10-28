/*
  Infinite Loop
*/
//A FOR LOOP without init statement, condition expression and post statement will act as an Infinite Loop.

package main

import "fmt"

func main21() {

	for {
		fmt.Println("Hello infinite")
	}

}

/*
	//Syntax of for loop
			init ; condition; post;
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
*/

//https://tour.golang.org/flowcontrol/4
