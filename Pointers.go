package main

import "fmt"

//We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr.
// zeroval has an int parameter, so arguments will be passed to it by value.
// zeroval will get a copy of ival distinct from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}
//zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the value at the referenced address.

func zeroprt(iptr *int) {
	*iptr = 0
}
func main() {

	//The &i syntax gives the memory address of i, i.e. a pointer to i.

	i := 1
	fmt.Println("Inital :", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroprt(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
