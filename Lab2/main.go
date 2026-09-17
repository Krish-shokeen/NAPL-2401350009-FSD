// package main

// import (
// 	"fmt"

// 	"MyProject/mathutil"
// 	"MyProject/strop"
// )

// func main() {

// 	var text string
// 	var num int
// 	var base, exponent int

// 	fmt.Print("Enter a string: ")
// 	fmt.Scanln(&text)

// 	fmt.Println("Reversed String:", strop.Reverse(text))
// 	fmt.Println("Number of Vowels:", strop.CountVowels(text))

// 	fmt.Print("\nEnter a number for factorial: ")
// 	fmt.Scan(&num)

// 	fmt.Println("Factorial:", mathutil.Factorial(num))

// 	fmt.Print("\nEnter the base: ")
// 	fmt.Scan(&base)

// 	fmt.Print("Enter the exponent: ")
// 	fmt.Scan(&exponent)

// 	fmt.Println("Power:", mathutil.Power(base, exponent))
// }



package main

import "fmt"

func removeByIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {

	fmt.Println("===== SLICE OPERATIONS =====")

	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	students := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter student %d name: ", i+1)
		fmt.Scan(&students[i])
	}

	fmt.Println("Initial Slice:", students)

	var newStudent string
	fmt.Print("\nEnter student name to add: ")
	fmt.Scan(&newStudent)

	students = append(students, newStudent)

	fmt.Println("After Adding:", students)

	var removeIndex int
	fmt.Print("\nEnter index to remove: ")
	fmt.Scan(&removeIndex)

	if removeIndex >= 0 && removeIndex < len(students) {
		students = removeByIndex(students, removeIndex)
		fmt.Println("After Removing:", students)
	} else {
		fmt.Println("Invalid index!")
	}

	var updateIndex int
	var updatedName string

	fmt.Print("\nEnter index to update: ")
	fmt.Scan(&updateIndex)

	if updateIndex >= 0 && updateIndex < len(students) {
		fmt.Print("Enter new student name: ")
		fmt.Scan(&updatedName)

		students[updateIndex] = updatedName

		fmt.Println("After Updating:", students)
	} else {
		fmt.Println("Invalid index!")
	}

	fmt.Println("\n===== MAP OPERATIONS =====")

	marks := make(map[string]int)

	var subjects int
	fmt.Print("\nEnter number of subjects: ")
	fmt.Scan(&subjects)

	for i := 0; i < subjects; i++ {
		var subject string
		var mark int

		fmt.Printf("\nEnter subject %d name: ", i+1)
		fmt.Scan(&subject)

		fmt.Print("Enter marks: ")
		fmt.Scan(&mark)

		marks[subject] = mark

		fmt.Println("Current Map:", marks)
	}

	var newSubject string
	var newMarks int

	fmt.Print("\nEnter another subject to insert: ")
	fmt.Scan(&newSubject)

	fmt.Print("Enter marks: ")
	fmt.Scan(&newMarks)

	marks[newSubject] = newMarks

	fmt.Println("After Inserting:", marks)

	var deleteSubject string

	fmt.Print("\nEnter subject to delete: ")
	fmt.Scan(&deleteSubject)

	if _, exists := marks[deleteSubject]; exists {
		delete(marks, deleteSubject)
		fmt.Println("After Deleting:", marks)
	} else {
		fmt.Println("Subject not found!")
	}

	var lookupSubject string

	fmt.Print("\nEnter subject to lookup: ")
	fmt.Scan(&lookupSubject)

	value, exists := marks[lookupSubject]

	if exists {
		fmt.Println("Marks:", value)
	} else {
		fmt.Println("Subject not found!")
	}

	fmt.Println("\nFinal Map:", marks)
}