package main

//Go is statically typed
import "fmt"

func variable() {
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)
	//all types have a zero value if nothing assigned to it yet
	//boolean false, int 0, float 0.0, string "", others like pointers, slices, interfaces, etc. are nil
	//Golang statically typed
}
