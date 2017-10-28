/*
  for Loop with break and continue
*/

package main

import "fmt"

func main22() {

	/*
		//example of break
		for i := 1; i <= 10; i++ {
			fmt.Println(i)

			if i == 5 {
				fmt.Println("Going to break")
				break
			}

		}
	*/

	//example of continue
	for i := 1; i <= 10; i++ {

		if i <= 5 {
			fmt.Println(i)
			continue
		}

		fmt.Println("Going to break")
		break

	}

	fmt.Println("Outside the for loop")

}
