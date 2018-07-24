package main

import "fmt"

// go supports pointers. This allows you to pass references tp values and records
// within a prpgram

// zeroval has an int parameter so arguments will be passed to it by value
/*
zeroval will get a COPY of ival distinc from the one in the calling function
*/
func zeroval(ival int) {
	ival = 0
}

//zeroptr takes an int pointer
func zeroptr(iptr *int) {
	// *iptr code in then function body then DEREFERENCES the pointer from its Memory address
	// to the CURRENT VALUE at THAT address
	*iptr = 0

	/*
		Assigning a value to a dereferenced pointer changes the value at the referenced
		address
	*/
}

func main() {

	i := 1
	fmt.Println("Intitial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The &i syntax gives the memory address of i, i.e. a pointer
	// to i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	// Assigning a value to a dereferenced pointer changes the value at the referenced
	// address. Hence 1 now becomes 0
	// whereas in zeroval a copy is made at separate address so doesn't change anything
	// at i in main which remains 1

	// pointer is printed
	fmt.Println("pointer:", &i)

	// zeroval doesn't change the i in main but zero pointer does
	// because it has a reference to the memory address for that variable
}
