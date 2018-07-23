package main

import "fmt"

/*
Variadic Functions can be called with any number of trailing arguements
*/

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, number := range nums {
		total += number
	}

	fmt.Println(total)
}
func main() {
	sum(4, 5, 9, 11)
	sum(1, 2)

	//  multipple arguements in a slice
	// applied to a variadic function
	nums := []int{20, 44, 1, 67}
	sum(nums...)

}
