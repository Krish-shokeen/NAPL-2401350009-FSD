// declare variablesof different types and assign them some values and print them
// a> int
// b> float
// c> String
// d> bool

package main

// import "fmt"
// func main()  {
// var a int = 10
// var b float64 = 3.14
// var c string = "Hello, World!"
// var d bool = true
// fmt.Println("Integer value:", a)
// fmt.Println("Float value:", b)
// fmt.Println("String value:", c)
// fmt.Println("Boolean value:", d)
// var age int;
// fmt.Printf("Age: %d", age)
// }

import (
	"fmt"
)

// func main() {
// 	// var a int = 90
// 	// b:=strconv.Itoa(a)
// 	// fmt.Println(b)
// 	// s:= "héllo"
// 	// fmt.Println(len(s))
// 	// fmt.Println(len([]rune(s)))
// 	for i, ch := range "héllo!"{
// 		if i==0{
// 			continue
// 		}
// 		fmt.Println(i, string(ch))
// 	}
// }

// given marks of the students decide the grade of the student using switch case

// func main() {
// 	var marks int
// 	fmt.Print("Enter marks: ")
// 	fmt.Scan(&marks)
// 	switch{
// 		case marks >= 90:
// 			fmt.Println("Grade: A")
// 		case marks >= 80:
// 			fmt.Println("Grade: B")
// 		case marks >= 70:
// 			fmt.Println("Grade: C")
// 		case marks >= 60:
// 			fmt.Println("Grade: D")
// 		case marks >= 50:
// 			fmt.Println("Grade: E")
// 		case marks >= 40:
// 			fmt.Println("Grade: P")
// 		default:
// 			fmt.Println("Grade: F")
// 	}

// }

// func sum(nums ...int) int {
// 	total := 0
// 	for _, n := range nums {
// 		total += n

// 	}

// 	return total
// }
// func main() {
// 	fmt.Println(sum(1, 2, 3, 4, 5))
// 	fmt.Println(sum(10, 20, 30))
// }

// type Celcius float64

// func (c Celcius) ToF() float64 {
// 	return float64(c*1.8 + 32)
// }

// func main() {
// 	var c Celcius = 37.0
// 	fmt.Printf("%.2f°C is %.2f°F\n", c, c.ToF())
// }

// func result(a, b, c int) (float64, bool) {
// 	avg := float64(a+b+c) / 3
// 	return avg, avg >= 40
// }
// func main() {
// 	fmt.Println(result(3, 4, 5))
// }
// package classwork



type student struct {
	name  string
	age   int
	marks float64
}

func main() {
	s := student{
		name:  "krish",
		age:   20,
		marks: 85.5,
	}

	fmt.Printf("Name of the student: %s\n", s.name)
	fmt.Printf("Age of the student: %d\n", s.age)
	fmt.Printf("Marks of the student: %.1f\n", s.marks)
}