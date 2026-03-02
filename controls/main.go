// Controls in Golang

// if statement
// if-else statement
// if-else if statement
// switch statement

package main

import "fmt"

func main() {

	// NOTE: if statement
	// a := 10
	// if a > 5 {
	// 	fmt.Println("a is greater than 5")
	// }

	// NOTE: if-else statement
	// a := 3
	// if a > 5 {
	// 	fmt.Println("a is greater than 5")
	// } else {
	// 	fmt.Println("a is not greater than 5")
	// }

	// NOTE: if-else if statement
	// a := 10
	// if a > 15 {
	// 	fmt.Println("a is greater than 15")
	// } else if a > 5 {
	// 	fmt.Println("a is greater than 5 but less than or equal to 15")
	// } else {
	// 	fmt.Println("a is less than or equal to 5")
	// }

	// NOTE: switch statement
	a := 2
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is not 1, 2 or 3")
	}
}
