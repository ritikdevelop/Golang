//NOTE: Types of methods in the Go programming language:

// 1. Value Receiver Methods: These methods operate on a copy of the value and do not modify the original value. They are defined with a value receiver (e.g., func (v Type) MethodName() {}).

// 2. Pointer Receiver Methods: These methods operate on a pointer to the value and can modify the original value. They are defined with a pointer receiver (e.g., func (v *Type) MethodName() {}).

// 3. Struct Methods: These methods are associated with a struct type and can be either value receiver or pointer receiver methods. They allow you to define behavior for struct types (e.g., func (s StructType) MethodName() {} or func (s *StructType) MethodName() {}).

// 4. Non-struct Methods: These methods are not associated with any struct type and can be defined for any type, including built-in types (e.g., func (t Type) MethodName() {} or func (t *Type) MethodName() {}).

// package main

// import "fmt"

// type Circle struct {
// 	radius float64
// }

// func (c Circle) area() float64 {
// 	return 3.14 * c.radius * c.radius
// }

// func main() {
// 	// Example of Value Receiver Method
// 	c := Circle{radius: 5}
// 	fmt.Printf("Area of the circle with radius %.2f is: %.2f\n", c.radius, c.area())
// }

//Note: Non-Struct Methods Example

// package main

// import "fmt"

// // Creating a custom type based on int
// type number int

// // Defining a method with a non-struct receiver
// func (n number) square() number {
// 	return n * n
// }

// func main() {
// 	a := number(4)
// 	b := a.square()

// 	fmt.Println("Square of", a, "is", b)
// }

// Pointer Receiver Methods Example

package main

import "fmt"

// Defining a struct
type person struct {
	name string
}

// Method with pointer receiver to modify data
func (p *person) changeName(newName string) {
	p.name = newName
}

func main() {
	a := person{name: "a"}

	fmt.Println("Before:", a.name)

	// Calling the method to change the name
	a.changeName("b")

	fmt.Println("After:", a.name)
}
