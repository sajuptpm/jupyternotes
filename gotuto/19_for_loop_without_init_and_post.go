/*
  FOR LOOP without Init and Post statement
*/

package main

import "fmt"

func main19() {

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
