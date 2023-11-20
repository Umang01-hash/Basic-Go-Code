package main

import "fmt"

func main() {
	// declaring variable using var
	var name string
	name = "JIIT"

	// %s prints the string
	fmt.Printf("Welcome to %v\n", name)

	// declaring multiple variables using var
	var (
		i int     = 20
		f float64 = 3.14
	)
	// %v prints the value in default format, %T prints the type of var
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("f: %v %T\n", f, f)

	// declaring variable using shorthand declaration
	flag := true
	fmt.Printf("b: %v %T\n", flag, flag)
}
