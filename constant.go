package main

import "fmt"
import "math"

// go supports const of char, string, boolean, numeric value

/*
const declares a constant value
*/
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	// constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// a numeric constant has no type until it is given one such as
	// by an explicit cast
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call.
	// For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))

}
