package main

import "fmt"

func main() {

	// slices are like arrays but more flexible with no element clount need be declared
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// another way declare and initialize slices and place values in them
	// all in one go
	lottery := []int32{40, 33, 55, 1, 58}
	fmt.Println("The winning numbers are:", lottery)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len interface always returns length of any data type
	fmt.Println("len:", len(s))

	// slice allows appends which is also diff to array
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// make slice c of length s
	c := make([]string, len(s))
	//copy contents of slice s to slice c
	copy(c, s)
	fmt.Println("This is a copy:", c)

	/*
		Slices support a “slice” operator with the syntax slice[low:high].
		For example, this gets a slice of the elements s[2], s[3],
		and s[4].
	*/
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slices up to but excluding s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// slices including s[2] and onwards
	l = s[2:]
	fmt.Println("sl3:", l)

	// : print everything in the slice
	l = s[:]
	fmt.Println("Sl4:", l)

	//declare and initialize a variable for slice in a single again.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data structures.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
