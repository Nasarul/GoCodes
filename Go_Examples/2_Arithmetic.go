package main

import "fmt"

func main() {
	// Declaring variable
	var a, b int

	// Assaining variable
	a = 10
	b = 20

	// Arithmatic operations (addition)
	fmt.Println("\n")
	c := a + b
	fmt.Printf("%d + %d = %d", a, b, c)

	fmt.Println("\n")

	// Arithmatic operations (subtraction)
	d := a - b
	fmt.Printf("%d - %d = %d", a, b, d)

	fmt.Println("\n")

	// Arithmatic operations (multiplication)
	e := a * b
	fmt.Printf("%d x %d = %d", a, b, e)

	fmt.Println("\n")

	// Arithmatic operations (division)
	f := float32(a) / float32(b)
	fmt.Printf("%d / %d = %.2f", a, b, f)

	fmt.Println("\n")
}
