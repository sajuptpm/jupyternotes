/*
Hello World Program
*/

/*When you build reusable pieces of code, you will develop a package as a shared library.
But when you develop executable programs, you will use the package “main” for making the
package as an executable program.
The package “main” tells the Go compiler that the package should compile as an executable
program instead of a shared library.
*/
package main

/*The keyword “import” is used for importing a package into other packages.
When you import packages, the Go compiler will look on the locations specified
by the environment variable GOROOT and GOPATH. Packages from the standard library
are available in the GOROOT location. The packages that are created by yourself,
and third-party packages which you have imported, are available in the GOPATH location.
*/
import "fmt"

//The main function in the package “main” will be the entry point of our executable program.
func main1() {

	//https://golang.org/pkg/fmt/#Println
	fmt.Println("Hello world!")

}
