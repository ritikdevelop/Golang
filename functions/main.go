//  Function in Golang

// NOTE: A function is a block of code that performs a specific task. It can take input parameters and return output values. Functions help in breaking down a large program into smaller, manageable pieces and promote code reusability.

// type of function in Golang
// 1. Standard Function
// 2. Variadic Function
// 3. Call by Value Function
// 4. Call by Reference Function

package main

import "fmt"

func main() {
	// NOTE: Standard Function
	// result := add(5, 10)
	// fmt.Println("Result of addition is: ", result)

	// NOTE: Variadic Function
	// result := sum(1, 2, 3, 4, 5)
	// fmt.Println("Result of sum is: ", result)

	// NOTE: Call by Value Function
	a := 10
	b := 20
	fmt.Println("Value of a before swap is: ", a)
	fmt.Println("Value of b before swap is: ", b)
	swaps(a, b) // Pass values to the function
	fmt.Println("Value of a after swap is: ", a)
	fmt.Println("Value of b after swap is: ", b)

	// NOTE: Call by Reference Function
	a = 10
	b = 20
	fmt.Println("Value of a before swap is: ", a)
	fmt.Println("Value of b before swap is: ", b)
	swap(&a, &b) // Pass pointers to the variables
	fmt.Println("Value of a after swap is: ", a)
	fmt.Println("Value of b after swap is: ", b)

}

// Call by Reference Function (using pointers)
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// Call by Value Function (using values)
func swaps(a int, b int) {
	temp := a
	a = b
	b = temp
}

// Standard Function
func add(a int, b int) int {
	return a + b
}

// Variadic Function
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
