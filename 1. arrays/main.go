package main

import "fmt"

// Defining an array (fixed-length)
var (
	x [3]int // Initializes with all elements being 0
	y = [3]int{1,2,3}
	z = [12]int{567, 7:10, 78, 76} // Places item where is specified before the :
	a = [...]int{0, 1, 1, 2, 3, 5, 8, 13}
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)

	// Length
	fmt.Println(len(a))
}