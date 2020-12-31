package main

import "fmt"

func main() {
	fmt.Println(Calc(1, 2))
}

// Calc is calcurate two numbers.
func Calc(a, b int) int {
	return a + b
}
