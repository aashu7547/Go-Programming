package main

import "fmt"

func main() {
	var a int = 42
	var p *int = &a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", p)
	fmt.Println("Value pointed to by p:", *p)

	*p = 100
	fmt.Println("New value of a:", a)
}
