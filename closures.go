package main

import "fmt"

// intSeq function returns anotjer function
// which is defined annonymously in the body of intSeq()
func intSeq() func() int {
	i := 0
	// The returned function closes over the var i to form a closure
	return func() int {
		i++
		return i
	}
}

func main() {

	// IntSeq() called assigning the result (a function) to nextInt
	// this function value captures its own I value which will be updated each time we call
	// nextInt
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Confirms state is unique to that particular function
	// created and tested newInts which is a complete new function and subsequent return val
	newInts := intSeq()
	fmt.Println(newInts())

}
