package main

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Hello World")
}