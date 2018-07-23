package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	// use the rage to sum all the numbers in a slice
	// arrays work like this too
	// didn't need the index so ignore with blank identifier '_'
	for _, number := range nums {
		sum += number
	}
	fmt.Println("sum:", sum)

	// Range but using the index
	for i, number := range nums {
		if number == 3 {
			fmt.Println("index:", i)
		}
	}

	// range used to iterate over a map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	// range iterating over jsut the keys in map
	for key := range kvs {
		fmt.Println("key:", key)
	}

	/*range on strings iterates over Unicode code points.
	The first value is the starting byte index of the rune and the second the rune itself.
	*/
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
