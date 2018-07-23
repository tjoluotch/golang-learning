package main

import "fmt"

func main() {
	i := 1

	// basic type of loop with single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// initial/ condition/ after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for loop wothout condition will loop repeatedly until
	// you 'break' out of the loop
	// or 'return' from the enclosing function
	for {
		fmt.Println("loop")
		break
	}

	// you can also 'continue' to next iteration of the loop
	for n := 0; n <= 5; n++ {
		// if n is a multiple of 2 (even number) continue to next loop iteration
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
