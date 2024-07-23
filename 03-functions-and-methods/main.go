package main

import "fmt"

// creating a type rectangle, which is of type struct
type rectangle struct {
	length, breadth int
}

// function to calculate the sum of integers
func calculateSum(num1, num2 int) int {
	return num1 + num2
}

// area Method with receiver type as rectangle
func (r rectangle) area() int {
	return r.length * r.breadth
}

func main() {
	// Function call
	total := calculateSum(10, 20)
	fmt.Printf("Total: %d\n", total)

	// Method call
	r := rectangle{length: 10, breadth: 20}
	fmt.Printf("Area of rectangle: %d", r.area())
}
