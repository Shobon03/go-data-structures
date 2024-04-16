package main

import "fmt"

var (
	s string = "Olá, mundo!"
	b byte = s[5] // Gets the utf-8 value of the character on index 5 ("m")
)

func main() {
	fmt.Println(s)
	fmt.Println(b)

	s1 := s[3:]
	fmt.Println(s1)
}