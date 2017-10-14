/*
  Access substring of a string using slice syntax
*/

package main

import "fmt"

func main10() {
	/*
		    In Go, no substring func is available. Instead we access parts
		    of strings (substrings) with slice syntax.

			  The slice end index is exclusive—this means a character at
			  that position is not included.

			  We can omit the end index. This takes the substring from a
			  start index to the end of the string. This is a clearer way
			  of using the length as the end.
	*/
	/////////012345678910
	name := "Hello World"

	//Print substring "Hello"
	fmt.Println(name[0:5])
	//Print substring "Hello"
	fmt.Println(name[:5])

	//Print substring "World"
	fmt.Println(name[6:11])
	//Print substring "World"
	fmt.Println(name[6:])
	//Print substring "World"
	fmt.Println(name[6:len(name)])

	fmt.Println("Length of name: ", len(name))
}
