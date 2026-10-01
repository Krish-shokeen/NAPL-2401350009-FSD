package main

// 1. Pointer Manipulation — Referencing and Dereferencing (Exp 5)
// Write a Go program demonstrating pointer usage:
// (1) declare a variable, print its address using & and access its value using *;
// (2) write a function that accepts a pointer parameter and modifies the original variable
//     (pass-by-reference) — print before and after values;
// (3) allocate a struct using new() and access/modify its fields through the pointer.

import "fmt"

// Function to modify the value using pointer
func variableDeclare(x *int, newValue int) {
	*x = newValue
}

// Function to modify the original variable using pointer
func modifyValue(a *int, value int) {
	*a = value
}

// Structure for storing person details
type Person struct {
	Name string
	Age  int
	Job  string
}

// Function to modify structure fields using pointer
func modifyStructure(p *Person, newName string, newAge int, newJob string) {
	p.Name = newName
	p.Age = newAge
	p.Job = newJob
}

func main() {

	// Declare a variable and take input from user
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// Print value and address
	fmt.Println("\nValue of number:", number)
	fmt.Println("Address of number:", &number)
	fmt.Println("Value using pointer:", *(&number))

	// Ask user for the new value
	var newNumber int

	fmt.Print("\nEnter new value to modify the number: ")
	fmt.Scan(&newNumber)

	// Print before modification
	fmt.Println("Before modification:", number)

	// Pass pointer to function
	modifyValue(&number, newNumber)

	// Print after modification
	fmt.Println("After modification:", number)

	// Allocate Person structure using new()
	p1 := new(Person)

	// Take initial Person details from user
	fmt.Println("\nEnter Person Details:")

	fmt.Print("Enter name: ")
	fmt.Scan(&p1.Name)

	fmt.Print("Enter age: ")
	fmt.Scan(&p1.Age)

	fmt.Print("Enter job: ")
	fmt.Scan(&p1.Job)

	// Display details before modification
	fmt.Println("\nPerson details before modification:")
	fmt.Println(*p1)

	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	fmt.Println("Job:", p1.Job)

	// Take new values for modification
	var newName string
	var newAge int
	var newJob string

	fmt.Println("\nEnter new details for modification:")

	fmt.Print("Enter new name: ")
	fmt.Scan(&newName)

	fmt.Print("Enter new age: ")
	fmt.Scan(&newAge)

	fmt.Print("Enter new job: ")
	fmt.Scan(&newJob)

	// Modify structure using pointer
	modifyStructure(p1, newName, newAge, newJob)

	// Display details after modification
	fmt.Println("\nPerson details after modification:")
	fmt.Println(*p1)

	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	fmt.Println("Job:", p1.Job)
}