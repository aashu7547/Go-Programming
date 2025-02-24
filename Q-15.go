package main

import "fmt"

func main() {
	ages := map[string]int{
		"Aashutosh": 19,
		"Parth":     25,
		"Sanjay":    21,
	}

	ages["Aashu"] = 20
	fmt.Println("Aashutosh's age:", ages["Aashutosh"])
	delete(ages, "Parth")
	fmt.Println("Updated map:", ages)

	if age, exists := ages["Sanjay"]; exists {
		fmt.Println("Sanjay's age:", age)
	} else {
		fmt.Println("Sanjay is not in the map.")
	}
}
