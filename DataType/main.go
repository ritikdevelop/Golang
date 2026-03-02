// Data Type in Go
// Data types are used to specify the type of data that a variable can hold. Go has several built-in data types, including:
// 1. Basic Data Types:
//    - int: Represents integer values (e.g., 1, 42, -5).
//    - float64: Represents floating-point numbers (e.g., 3.14, -0.001).
//    - string: Represents a sequence of characters (e.g., "Hello, World!").
//    - bool: Represents true or false values (e.g., true, false).

// 2. Composite Data Types:
//    - array: A fixed-size collection of elements of the same type (e.g., [5]int{1, 2, 3, 4, 5}).
//    - slice: A dynamically-sized collection of elements of the same type (e.g., []int{1, 2, 3}).
//    - struct: A collection of fields that can hold different types of data (e.g., type Person struct { Name string; Age int })
//    - map: A collection of key-value pairs (e.g., map[string]int{"Alice": 30, "Bob": 25}).

// 3. Pointer Types:
//    - A pointer is a variable that holds the memory address of another variable (e.g., var p *int = &x).

package main

import "fmt"

func main() {

	//NOTE: Basic Data Types
	// var age int = 30
	// var pi float64 = 3.14
	// var name string = "Alice"
	// var isStudent bool = true

	// fmt.Println("Age:", age)
	// fmt.Println("Pi:", pi)
	// fmt.Println("Name:", name)
	// fmt.Println("Is Student:", isStudent)

	//NOTE: Composite Data Types - Array
	// arr := [5]int{1, 2, 3, 4, 5}

	// fmt.Println("Elements of the array: ")
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// cars := [5]string{"Toyota", "Honda", "Ford", "BMW", "Audi"}
	// fmt.Println("Original Array: ", cars)

	// // Change the array value of index 2
	// cars[2] = "Tesla"
	// fmt.Println("Updated Array: ", cars)

	//NOTE: Composite Data Types - Slice
	// array := []int{1, 2, 3}
	// slice := array[0:2] // Create a slice from the array

	// fmt.Println("Original Slice: ", array)
	// fmt.Printf("Length of the slice array is: %d\n", len(array))

	// fmt.Println("Elements of the slice: ")
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	// fmt.Println("Updated slice array: ", slice)
	// fmt.Println("Updated Slice: ", array)

	// array = append(array, 4, 5, 6)
	// fmt.Println("Updated Slice append array: ", array)

	// fmt.Printf("Length of the updated slice array is: %d", len(array))
	//NOTE: Composite Data Types - Struct
	// type Person struct {
	// 	Name string
	// 	Age  int
	// }

	// person := Person{Name: "Alice", Age: 30}
	// fmt.Println("Person Name:", person.Name)
	// fmt.Println("Person Age:", person.Age)

	//NOTE: Composite Data Types - Map

	// personAge := map[string]int{
	// 	"Alice": 30,
	// 	"Bob":   25,
	// }

	// fmt.Println("Person Age:", personAge["Alice"])
	// fmt.Println("Person Age:", personAge["Bob"])

	//NOTE: Pointer Types

	var x int = 10
	var p *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p (address of x):", p)
	fmt.Println("Value pointed to by p (value of x):", *p)
}
