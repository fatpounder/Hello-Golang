package main

import (
	"fmt"
)

func conversion() {
	y := 42
	z := 42.0
	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)
	fmt.Printf("%v of type %T \n", m, m)

	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)

}
