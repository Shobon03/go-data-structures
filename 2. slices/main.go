package main

import "fmt"

var (
	slice1 = []int{1,2,3} // [] defines a slice; [...] defines an array
	slice2 = [][]int{{1,2},{3,4}}
	slice3 []int // All elements inside an empty slice are nil
)

func main() {
	fmt.Println(slice1)
	fmt.Println(slice2[0][1])
	fmt.Println(len(slice3))

	// Appending items
	slice3 = append(slice3, slice1...) // Expand the slice with ...

	fmt.Println(slice3)
}