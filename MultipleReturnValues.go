package main

import "fmt"

//The (int, int) in this function signature shows that the function returns 2 ints.
func val()(int, int)  {
	return 3,7
}
func main() {
	//Here we use the 2 different return values from the call with multiple assignment.
	a, b := val()
	fmt.Println(a)
	fmt.Println(b)

	//If you only want a subset of the returned values, use the blank identifier _.
	_, c := val()
	fmt.Println(c)

}
