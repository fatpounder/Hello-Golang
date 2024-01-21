package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, Marie")
// }

func main() {

	var num1 int
	//if you do not declare value for int, it will default to 0
	//need var when initializing a variable
	//only need to specify type if no value declared for the variable
	var num2 = 3
	//var num2 = 3 can also be written as num2:= 3, is just shorthand

	num1 = 2

	num1, num2 = 2, 3
	//knows variables formatted respectively

	const num3 = 1

	var result = num1 + num2 + num3

	fmt.Print(result)

}
