package main

import "fmt"

func main() {

	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("Original slice:", numbers)

	numbers = append(numbers, 60)
	fmt.Println("Slice after appending 60:", numbers)

	subSlice := numbers[1:4]
	fmt.Println("Sub-slice (1:4):", subSlice)

	numbers[2] = 35
	fmt.Println("Slice after modifying an element:", numbers)

	fmt.Println("Length of the slice:", len(numbers))
	fmt.Println("Capacity of the slice:", cap(numbers))
}
