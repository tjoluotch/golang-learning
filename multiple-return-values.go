package main

import "fmt"

// The (int, int) in this function shows that the function returns two ints
func vals() (int, int) {
	return 3, 7
}

func main() {

	// two different return values used from the function
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// use black identifier _, for subset of return values
	_, c := vals()
	fmt.Println(c)

}
