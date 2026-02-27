package main

// There are two types of variables in Go: package-level variables and function-level variables.
// Package-level variables are declared outside of any function and are accessible throughout the entire package.
// Function-level variables are declared within a function and are only accessible within that function.

// 1 Package level variables examples:
// var x int = 10
// var y string = "Hello, World!"
// var z bool = truepackage main

// import "fmt"

// func main() {
// 	// Package level variable
// 	var x int = 10
// 	fmt.Println("Package level variable x:", x)
// }

// import "fmt"

// func main() {
// 	var message string = "Hello, World!"
// 	fmt.Println(message)
// }

// 2 Function level variables
// Function level variables examples:

import "fmt"

func variable() {
	var a int = 20
	var b string = "Go is awesome!"
	var c bool = false
	fmt.Println(a, b, c)
}

func main() {
	variable()
}
