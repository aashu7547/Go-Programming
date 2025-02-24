package main

import "fmt"

func nCr(n, r int) int {
	if r == 0 || r == n {
		return 1
	}
	return nCr(n-1, r-1) + nCr(n-1, r)
}

func main() {
	var n, r int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)
	fmt.Print("Enter r: ")
	fmt.Scan(&r)
	fmt.Printf("nCr(%d, %d) = %d\n", n, r, nCr(n, r))
}
