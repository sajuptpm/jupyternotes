/*
  WHILE LOOP using FOR LOOP
*/

//A FOR LOOP without init and post statement is equivalent to a WHILE LOOP in other C-like languages.

package main

import "fmt"

func main20() {

	//init statement
	i := 1

	for i <= 2 {

		fmt.Println(i)

		//post statement
		i++
		//i += 1
	}

	/*
				init ; condition; post;
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
	*/

}
