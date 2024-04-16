package main

import "fmt"

type Person struct {
	name string
	age int
	petName string
}

func main() {
	 bob := Person{
		"Bob Walker",
		35,
		"Amarildo",
	}

	fmt.Println(bob)
}