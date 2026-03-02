// NOTE: Loops in Go
// 1. For Loop
// 2. While Loop (using for loop)
// 3. do While Loop (using for loop)

package main

import "fmt"

func main() {

	// NOTE: For Loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i is: ", i)
	// }

	// NOTE: While Loop (using for loop)
	// i := 0
	// for i < 5 {
	// 	fmt.Println("Value of i is: ", i)
	// 	i++
	// }

	// NOTE: do While Loop (using for loop)

	i := 0
	for {
		fmt.Println("Value of i is: ", i)
		i++
		if i >= 5 {
			break
		}
	}
}
