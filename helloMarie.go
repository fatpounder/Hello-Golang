package main

import (
	"fmt"
	//"math"
)

// func main() {
// 	fmt.Println("Hello, Marie")
// }

// func variables() {

// 	var num1 int
// 	//if you do not declare value for int, it will default to 0
// 	//need var when initializing a variable
// 	//only need to specify type if no value declared for the variable
// 	var num2 = 3
// 	//var num2 = 3 can also be written as num2:= 3, is just shorthand

// 	num1 = 2

// 	num1, num2 = 2, 3
// 	//knows variables formatted respectively

// 	const num3 = 1

// 	var result = num1 + num2 + num3

// 	fmt.Print(result)

// }

// func multparam(x int, y int) (int, int) {
// 	//parameters and type defined
// 	var out1 = x + y
// 	var out2 = x - y
// 	return out1, out2
// }

// func adding() {

// 	num1 := 5
// 	num2 := 4

// 	result1, result2 := calc(num1, num2)
// 	fmt.Println(result1, result2)

// }

// func main() {

// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Look at this loop go!", i)

// 	}

// }

//func with capital letter start is accessible outside

// var a = 8

// func demo() {
// 	var a = 9
// 	fmt.Println(a)
// }

// //preference given to most local variable, however need to avoid shadowing anyway for readability

// func main() {
// 	var num float64 = 12
// 	var result = math.Sqrt(float64(num))
// 	fmt.Printf("The output is %.2g Danke!", result)
// }

//keyword defer to run something but call it after the execution block

type Student struct {
	rollno int
	name   string
	marks  int
}

func main() {

	var s1 Student = Student{101, "Marie", 55}
	fmt.Println(s1)
	fmt.Println(s1.name)

	var s2 Student = Student{rollno: 102, marks: 55}
	fmt.Println(s2)
}
