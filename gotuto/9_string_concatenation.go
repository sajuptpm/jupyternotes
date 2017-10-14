/*
   string concatenation
*/

package main

import "fmt"
import "strings"

func main9() {

	var s1 = "Hello"
	s2 := "World"

	/*########Method-1########*/
	res1 := s1 + " " + s2
	//print variable
	fmt.Println(res1)

	/*########Method-2########*/
	//Create an array of strings
	s_array := []string{s1, s2}
	fmt.Println(s_array)
	//Join strings in the array with space
	res2 := strings.Join(s_array, " ")
	fmt.Println(res2)

}
