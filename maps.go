package main

import (
	"fmt"
)

func main() {

	//To create an empty map,
	//use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	// show all key value pairs
	fmt.Println("string map int:", m)

	// get value for k1
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	// removes key/value pairs from map
	delete(m, "k2")
	fmt.Println("adjusted map:", m)

	// _ , isit := : 2nd value(isit) used to indicate if key was present
	// _ used cause we didn't need the value
	_, isit := m["k2"]
	fmt.Println("is it present:", isit)

	// declare and initialize a new map in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("2nd map:", n)
}
