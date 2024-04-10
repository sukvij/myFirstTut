package main

import "fmt"

func Sub(a, b int) int {
	if a == b {
		return 0
	}
	return a - b
}

func main() {
	fmt.Println("mains")
}
