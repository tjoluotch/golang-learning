package main

import "fmt"

func main() {
	/* Note that arrays appear in the form [v1 v2 v3 ...]
	when printed with fmt.Println.
	*/

	// var a is an array that will hold exactl 5 ints
	var a [5]int
	// print contents of a - which is an array
	fmt.Println("emp:", a)

	// set the value at the forth index to 100(last value - remember arrays are zero bound)
	a[4] = 100
	// print array post change
	fmt.Println("set:", a)
	// print int at forth index
	fmt.Println("get:", a[4])

	// print the length of the array
	fmt.Println("len:", len(a))

	// := used to declare and initialize array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// can compose types to build multi dimensional data structures
	var twoD [2][3]int
	// a two dimensional array containing 3 ints in each dimension

	//loop 1:  i will be 0, 1
	for i := 0; i < 2; i++ {
		// loop 2: j will be 0, 1, 2
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	/*
		1st iteration
		[[0 0 0] [0 0 0]]
		2nd iteration
		[[0 1 0] [0 0 0]]
		3rd iteration
		[[0 1 2] [0 0 0]]
		4th iteration
		[[0 1 2] [1 0 0]]
		5th iteration
		[[0 1 2] [1 2 0]]
		6th iteration
		[[0 1 2] [1 2 3]]
	*/
	fmt.Println("2d: ", twoD)
}
