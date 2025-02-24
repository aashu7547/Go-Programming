package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	var x, y int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&y)

	fmt.Printf("Before swapping: x = %d, y = %d\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("After swapping: x = %d, y = %d\n", x, y)
}
