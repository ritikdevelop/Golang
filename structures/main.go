// Note: Structures in the Go programming language are user-defined types that group together related data. They are similar to classes in other programming languages but do not support inheritance. Structures allow you to create complex data types that can hold multiple fields of different types.

// Advantages of using structures in Go:
// 1. Encapsulation: Structures allow you to encapsulate related data together, making it easier to manage and organize your code.

// 2. Reusability: You can define a structure once and reuse it throughout your code, reducing redundancy and improving maintainability.

// 3. Custom Data Types: Structures enable you to create custom data types that can represent real-world entities more accurately.

// 4. Methods: You can define methods on structures, allowing you to associate behavior with the data they hold.

//Note: Example of a struct definition and usage
// package main

// import "fmt"

// type person struct {
// 	name string
// }

// func (p *person) updateName(newName string) {
// 	p.name = newName
// }

// func main() {
// 	// Creating an instance of the person struct
// 	p := person{name: "Alice"}
// 	fmt.Println("Original Name: ", p.name)
// 	// Updating the name using the method with pointer receiver
// 	p.updateName("Bob")
// 	fmt.Println("Updated Name: ", p.name)
// }

// Note: Golang program to show how to declare and define the struct
// package main

// import "fmt"

// // Defining a struct type
// type Address struct {
// 	Name    string
// 	city    string
// 	Pincode int
// }

// func main() {

// 	// Declaring a variable of a `struct` type
// 	// All the struct fields are initialized
// 	// with their zero value
// 	var a Address
// 	fmt.Println(a)

// 	// Declaring and initializing a
// 	// struct using a struct literal
// 	a1 := Address{"Akshay", "Dehradun", 3623572}

// 	fmt.Println("Address1: ", a1)

// 	// Naming fields while
// 	// initializing a struct
// 	a2 := Address{Name: "Anikaa", city: "Ballia", Pincode: 277001}

// 	fmt.Println("Address2: ", a2)

// 	// Uninitialized fields are set to
// 	// their corresponding zero-value
// 	a3 := Address{Name: "Delhi"}
// 	fmt.Println("Address3: ", a3)
// }

// Note: Access individual fields of a struct using the dot operator

// Golang program to show how to
// access the fields of struct
package main

import "fmt"

// defining the struct
type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

// Main Function
func main() {
	c := Car{Name: "Ferrari", Model: "GTC4",
		Color: "Red", WeightInKg: 1920}

	// Accessing struct fields
	// using the dot operator
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	// Assigning a new value
	// to a struct field
	c.Color = "Black"

	// Displaying the result
	fmt.Println("Car: ", c)
}
