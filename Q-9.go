package main

import "fmt"

func findMinMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return
}

func main() {
	var n int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Println("Enter numbers:")
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	min, max := findMinMax(nums)
	fmt.Printf("Min: %d\nMax: %d\n", min, max)
}
