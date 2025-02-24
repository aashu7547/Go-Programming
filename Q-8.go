package main

import "fmt"

func main() {
	fmt.Println("Demonstrating break:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("\nDemonstrating continue:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\nDemonstrating return:")
	fmt.Println("Sum:", addNumbers(2, 3))
}

func addNumbers(a int, b int) int {
	return a + b
}
