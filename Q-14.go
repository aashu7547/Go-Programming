package main

import "fmt"

type Employee struct {
	Name   string
	ID     int
	Salary float64
}

func displayEmployee(emp Employee) {
	fmt.Printf("Employee Name: %s\n", emp.Name)
	fmt.Printf("Employee ID: %d\n", emp.ID)
	fmt.Printf("Employee Salary: %.2f\n", emp.Salary)
}

func main() {
	emp1 := Employee{Name: "Aashutosh Kumar", ID: 47, Salary: 50000.00}
	displayEmployee(emp1)
}
