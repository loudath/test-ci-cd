package main

import "fmt"

// Add sums two integers
func Add(a, b int) int {
	return a + b
}

func main() {
	var num1 int
	var num2 int

	fmt.Println("Enter integer 1:")
	fmt.Scanln(&num1)

	fmt.Println("Enter integer 2:")
	fmt.Scanln(&num2)
	fmt.Printf("%d + %d = %d", num1, num2, Add(num1, num2))
}
