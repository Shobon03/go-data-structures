package main

import "fmt"

func main() {
	users := map[int][]string {
		1: {"Lucas", "25"},
		2: {"Matheus", "21"},
		3: {"Rudolph", "100"},
	}

	fmt.Println(users)
}