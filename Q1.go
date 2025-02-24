// package main

// import (
// 	"fmt"
// )

//	func main() {
//		var n1 int
//		fmt.Println("Enter a number : ")
//		fmt.Scanln(&n1)
//		if n1 >= 18 {
//			fmt.Println("You can drive")
//		} else {
//			fmt.Println("You can not drive")
//		}
//	}

package main

import (
	"fmt"
)

func main() {
	var a1 int
	fmt.Printf("Enter any number : ")
	fmt.Scanln(&a1)
	if a1%2 == 0 {
		fmt.Println("Even : ", a1)
	} else {
		fmt.Println("Odd : ", a1)
	}

}
