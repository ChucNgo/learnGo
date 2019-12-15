package main

import "fmt"

func main() {
	const a = 5
	// var intVar int = a
	// var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println(float64Var, complex64Var)
}
