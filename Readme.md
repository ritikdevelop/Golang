# Golang Tutorial

This comprehensive tutorial will guide you through learning Go (Golang), a powerful and efficient programming language developed by Google. Whether you're a beginner or have some programming experience, this guide covers everything from installation to advanced concepts like concurrency.

## Table of Contents

1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Your First Go Program](#your-first-go-program)
5. [Go Basics](#go-basics)
   - [Variables and Constants](#variables-and-constants)
   - [Data Types](#data-types)
   - [Operators](#operators)
6. [Control Flow](#control-flow)
7. [Functions](#functions)
8. [Packages and Modules](#packages-and-modules)
9. [Pointers](#pointers)
10. [Structs and Methods](#structs-and-methods)
11. [Interfaces](#interfaces)
12. [Concurrency](#concurrency)
13. [Error Handling](#error-handling)
14. [Working with Files](#working-with-files)
15. [Building and Running](#building-and-running)
16. [Next Steps](#next-steps)

## Introduction

Go, also known as Golang, is a statically typed, compiled programming language designed by Google engineers Robert Griesemer, Rob Pike, and Ken Thompson. It was first released in 2009 and has since gained popularity for its simplicity, efficiency, and strong support for concurrent programming.

### Why Learn Go?

- **Simplicity**: Go has a clean syntax similar to C but with modern features.
- **Performance**: Compiled language with fast execution and low memory footprint.
- **Concurrency**: Built-in support for goroutines and channels makes concurrent programming easy.
- **Standard Library**: Rich standard library for networking, HTTP, cryptography, and more.
- **Cross-Platform**: Compiles to native binaries for Windows, macOS, Linux, and more.
- **Growing Ecosystem**: Used by companies like Google, Uber, Dropbox, and Netflix.

## Prerequisites

Before starting this tutorial, you should have:

- A computer running Windows, macOS, or Linux
- Basic understanding of programming concepts (variables, loops, functions)
- Familiarity with command-line interfaces
- A text editor (we recommend Visual Studio Code with Go extensions)

## Installation

### Windows

1. Download the Go installer from the official website: https://golang.org/dl/
2. Run the installer and follow the setup wizard
3. Verify installation by opening Command Prompt and running:
   
```bash
   go version
   
```

### macOS

1. Install Homebrew if you haven't already:
   
```bash
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   
```
2. Install Go:
   
```bash
   brew install go
   
```
3. Verify installation:
   
```bash
   go version
   
```

### Linux

1. Download the Go tarball from https://golang.org/dl/
2. Extract it to `/usr/local`:
   
```bash
   sudo tar -C /usr/local -xzf go1.XX.X.linux-amd64.tar.gz
   
```
3. Add Go to your PATH by adding this line to your `~/.bashrc` or `~/.profile`:
   
```bash
   export PATH=$PATH:/usr/local/go/bin
   
```
4. Reload your profile and verify:
   
```bash
   source ~/.bashrc
   go version
   
```

## Your First Go Program

Let's start with the classic "Hello, World!" program. Create a file called `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

To run this program:

```bash
go run hello.go
```

You should see `Hello, World!` printed to the console.

### Explanation

- `package main`: Defines this as an executable program
- `import "fmt"`: Imports the format package for printing
- `func main()`: The entry point of the program
- `fmt.Println()`: Prints text to the console

## Go Basics

### Variables and Constants

Variables store data that can change:

```go
package main

import "fmt"

func main() {
    var name string = "Go"
    age := 10  // Short variable declaration
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
```

Constants are immutable:

```go
const pi = 3.14159
const (
    version = "1.0"
    author  = "Google"
)
```

### Data Types

Go has several built-in data types:

```go
// Integers
var i int = 42
var ui uint = 100

// Floats
var f float64 = 3.14

// Strings
var s string = "Hello, Go!"

// Booleans
var b bool = true

// Arrays and Slices
var arr [3]int = [3]int{1, 2, 3}
var slice []int = []int{1, 2, 3, 4, 5}

// Maps
var m map[string]int = map[string]int{"one": 1, "two": 2}
```

### Operators

Go supports various operators:

```go
// Arithmetic
a := 10 + 5  // 15
b := 10 - 5  // 5
c := 10 * 5  // 50
d := 10 / 5  // 2
e := 10 % 3  // 1

// Comparison
fmt.Println(10 > 5)   // true
fmt.Println(10 == 5)  // false

// Logical
fmt.Println(true && false)  // false
fmt.Println(true || false)  // true
fmt.Println(!true)          // false
```

## Control Flow

### If/Else Statements

```go
age := 18
if age >= 18 {
    fmt.Println("You are an adult")
} else {
    fmt.Println("You are a minor")
}

// With else if
if age < 13 {
    fmt.Println("Child")
} else if age < 18 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Adult")
}
```

### Switch Statements

```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("TGIF!")
default:
    fmt.Println("Regular day")
}
```

### Loops

Go has only one type of loop: `for`

```go
// Basic for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-like loop
j := 0
for j < 5 {
    fmt.Println(j)
    j++
}

// Range loop
fruits := []string{"apple", "banana", "cherry"}
for index, fruit := range fruits {
    fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
}
```

## Functions

Functions are defined with the `func` keyword:

```go
func add(a int, b int) int {
    return a + b
}

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func main() {
    result := add(5, 3)
    fmt.Println(result)  // 8
    greet("World")
}
```

### Multiple Return Values

Go functions can return multiple values:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

## Packages and Modules

Go code is organized into packages. Create a new directory for your project:

```bash
mkdir myproject
cd myproject
go mod init myproject
```

Create a file `math.go`:

```go
package math

func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}
```

And `main.go`:

```go
package main

import (
    "fmt"
    "myproject/math"
)

func main() {
    fmt.Println(math.Add(5, 3))      // 8
    fmt.Println(math.Multiply(5, 3)) // 15
}
```

## Pointers

Pointers store memory addresses:

```go
func main() {
    x := 42
    p := &x  // p points to x
    fmt.Println(*p)  // 42 (dereference)

    *p = 50  // Change value through pointer
    fmt.Println(x)  // 50
}
```

## Structs and Methods

Structs are custom data types:

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    person.Greet()
}
```

## Interfaces

Interfaces define behavior:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area())
    fmt.Println("Perimeter:", rect.Perimeter())
}
```

## Concurrency

Go's concurrency model uses goroutines and channels:

```go
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")  // Start goroutine
    say("hello")
}
```

### Channels

Channels allow communication between goroutines:

```go
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum  // Send sum to channel
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c  // Receive from channel
    fmt.Println(x, y, x+y)
}
```

## Error Handling

Go uses explicit error handling:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

## Working with Files

Reading and writing files:

```go
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Writing to a file
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    writer.WriteString("Hello, Go!\n")
    writer.Flush()

    // Reading from a file
    file, err = os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
```

## Building and Running

### Running Go Programs

```bash
go run main.go
```

### Building Executables

```bash
go build main.go
```

This creates an executable file you can run directly.

### Go Modules

Initialize a module:

```bash
go mod init example.com/myproject
```

Add dependencies:

```bash
go get github.com/some/package
```

## Next Steps

Congratulations! You've completed the basics of Go programming. Here are some next steps:

1. **Practice**: Build small projects like a command-line calculator or a simple web server
2. **Read the Documentation**: Visit https://golang.org/doc/ for official docs
3. **Join the Community**: Participate in Go forums and meetups
4. **Explore Frameworks**: Learn Gin for web development or Cobra for CLI apps
5. **Contribute**: Check out open-source Go projects on GitHub
6. **Advanced Topics**: Study testing, profiling, and performance optimization

### Resources

- [Official Go Website](https://golang.org/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)
- [Awesome Go](https://awesome-go.com/) - Curated list of Go libraries

Happy coding with Go!
