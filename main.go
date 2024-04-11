package main

import "fmt"

func Sub(a, b int) int {
	if a == b {
		return 0
	}
	return a - b
}

func Sum(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	if b == 0 {
		return -1
	}
	return a / b
}

func main() {
	fmt.Println("mains")
}
