// Operator in Golang

// NOTE: Arithmetic Operators
//  addition: +
//  subtraction: -
//  multiplication: *
//  division: /
//  modulus: %
//  increment: ++
//  decrement: --

package main

// func main() {

// 	//NOTE: Addition --> (+)

// 	// a := 10
// 	// b := 20

// 	// fmt.Println("Addition of a and b is: ", a+b)

// 	// NOTE: Subtraction --> (-)

// 	// a := 50
// 	// b := 25
// 	// fmt.Println("Subtraction of a and b is: ", a-b)

// 	// NOTE: Multiplication --> (*)
// 	// a := 5
// 	// b := 10
// 	// fmt.Println("Multiplication of a and b is: ", a*b)

// 	// NOTE: Division --> (/)
// 	// a := 100
// 	// b := 4
// 	// fmt.Println("Division of a and b is: ", a/b)

// 	// NOTE: Modulus --> (%)
// 	// a := 10
// 	// b := 3
// 	// fmt.Println("Modulus of a and b is: ", a%b)

// 	// NOTE: Increment --> (++)
// 	// a := 5
// 	// fmt.Println("Value of a before increment is: ", a)
// 	// a++
// 	// fmt.Println("Value of a after increment is: ", a)

// 	// NOTE: Decrement --> (--)
// 	a := 5
// 	fmt.Println("Value of a before decrement is: ", a)
// 	a--
// 	fmt.Println("Value of a after decrement is: ", a)

// }

// Note: In Go, the increment (++) and decrement (--) operators can only be used as statements, not as expressions. This means you cannot use them in a larger expression or assign their result to a variable. They are used solely to increase or decrease the value of a variable by one.

// NOTE: Assignment Operators
//  assignment: =
//  addition assignment: +=
//  subtraction assignment: -=
//  multiplication assignment: *=
//  division assignment: /=
//  modulus assignment: %=

// func main() {
// NOTE: Assignment Operator --> (=)
// a := 10
// fmt.Println("Value of a is: ", a)

// NOTE: Addition Assignment Operator --> (+=)
// a := 5
// fmt.Println("Value of a before addition assignment is: ", a)
// a += 10 // equivalent to a = a + 10
// fmt.Println("Value of a after addition assignment is: ", a)

// NOTE: Subtraction Assignment Operator --> (-=)
// a := 20
// fmt.Println("Value of a before subtraction assignment is: ", a)
// a -= 5 // equivalent to a = a - 5
// fmt.Println("Value of a after subtraction assignment is: ", a)

// NOTE: Multiplication Assignment Operator --> (*=)
// a := 5
// fmt.Println("Value of a before multiplication assignment is: ", a)
// a *= 3 // equivalent to a = a * 3
// fmt.Println("Value of a after multiplication assignment is: ", a)

// NOTE: Division Assignment Operator --> (/=)
// a := 20
// fmt.Println("Value of a before division assignment is: ", a)
// a /= 4 // equivalent to a = a / 4
// fmt.Println("Value of a after division assignment is: ", a)

// NOTE: Modulus Assignment Operator --> (%=)
// a := 10
// fmt.Println("Value of a before modulus assignment is: ", a)
// a %= 3 // equivalent to a = a % 3
// fmt.Println("Value of a after modulus assignment is: ", a)
// }

// NOTE: Relational Operators
//  equal to: ==
//  not equal to: !=
//  greater than: >
//  less than: <
//  greater than or equal to: >=
//  less than or equal to: <=

import "fmt"

// func main() {
// NOTE: Equal to Operator --> (==)
// a := 10
// b := 20
// fmt.Println("Is a equal to b? ", a == b)

// // NOTE: Not Equal to Operator --> (!=)
// fmt.Println("Is a not equal to b? ", a != b)

// // NOTE: Greater than Operator --> (>)
// fmt.Println("Is a greater than b? ", a > b)

// // NOTE: Less than Operator --> (<)
// fmt.Println("Is a less than b? ", a < b)

// // NOTE: Greater than or Equal to Operator --> (>=)
// fmt.Println("Is a greater than or equal to b? ", a >= b)

// // NOTE: Less than or Equal to Operator --> (<=)
// fmt.Println("Is a less than or equal to b? ", a <= b)

// }

// NOTE: Logical Operators
//  and: &&
//  or: ||
//  not: !

func main() {

	// NOTE: And Operator --> (&&)
	a := true
	b := false
	fmt.Println("a AND b is: ", a && b)

	// NOTE: Or Operator --> (||)
	fmt.Println("a OR b is: ", a || b)

	// NOTE: Not Operator --> (!)
	fmt.Println("NOT a is: ", !a)

}
