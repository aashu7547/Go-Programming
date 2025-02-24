package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello"
	str2 := "World"

	concat := str1 + " " + str2
	fmt.Println("Concatenated string:", concat)

	length := len(concat)
	fmt.Println("Length of the concatenated string:", length)

	substr := concat[0:5]
	fmt.Println("Substring (0:5):", substr)

	upper := strings.ToUpper(concat)
	fmt.Println("Uppercase string:", upper)

	lower := strings.ToLower(concat)
	fmt.Println("Lowercase string:", lower)
}
