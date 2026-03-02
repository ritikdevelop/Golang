package main

import "fmt"

// Student struct with an anonymous inner structure for personal details
type Student struct {
	personalDetails struct { // Anonymous inner structure for personal details
		name       string
		enrollment int
	}
	GPA float64 // Standard field
}

func main() {
	// Initializing the anonymous structure inside Student
	student := Student{
		personalDetails: struct {
			name       string
			enrollment int
		}{
			name:       "Ritik",
			enrollment: 12345,
		},
		GPA: 3.8,
	}
	// Displaying values
	fmt.Println("Name:", student.personalDetails.name)
	fmt.Println("Enrollment:", student.personalDetails.enrollment)
	fmt.Println("GPA:", student.GPA)
}
