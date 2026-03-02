// This is the Blank or underscore function in Go. It is used to ignore the value of a variable or a return value from a function. The underscore is a special identifier in Go that can be used to discard values that are not needed.
package main

import "fmt"

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	// Example of using the underscore to ignore a return value from a function
	result, _ := div(10, 2) // We only care about the result, not the error
	fmt.Println("Result of division is: ", result)

	// Example of using the underscore to ignore a variable
	_, err := div(10, 0) // We only care about the error, not the result
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
