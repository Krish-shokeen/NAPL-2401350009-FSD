package main

import (
	"fmt"

	"MyProject/mathutil"
	"MyProject/strop"
)

func main() {

	var text string
	var num int
	var base, exponent int

	fmt.Print("Enter a string: ")
	fmt.Scanln(&text)

	fmt.Println("Reversed String:", strop.Reverse(text))
	fmt.Println("Number of Vowels:", strop.CountVowels(text))

	fmt.Print("\nEnter a number for factorial: ")
	fmt.Scan(&num)

	fmt.Println("Factorial:", mathutil.Factorial(num))

	fmt.Print("\nEnter the base: ")
	fmt.Scan(&base)

	fmt.Print("Enter the exponent: ")
	fmt.Scan(&exponent)

	fmt.Println("Power:", mathutil.Power(base, exponent))
}
