package main

import "fmt"

func main() {
	// declaring variable using var
	var name string
	name = "GoFr Event"

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

	// Byte - alias for uint8
	var x byte = 'A'
	fmt.Printf("x: %v %T\n", x, x)

	// Rune - alias for int32
	var y rune = 'क' // This will overflow, if byte.
	fmt.Printf("y: %v %T\n", y, y)
}
