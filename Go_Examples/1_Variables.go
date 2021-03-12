package main

import "fmt"

func main() {
	// Declaring variable
	var name string
	var age int
	var weight float32

	// Assaining variable
	name = "Md."
	age = 42
	weight = 71.20

	fmt.Println("\nHello Mr/Mr/Ms.", name)
	fmt.Println("Your age is :", age)
	fmt.Println("Your weight are :", weight)

	// Another way to declaring variable
	fmt.Println("\n=========================================================")
	var (
		yname   string
		yage    int
		yweight float32
	)

	yname = "Nasarul"
	yage = 52
	yweight = 55.56

	fmt.Println("\nHello Mr/Mr/Ms.", yname)
	fmt.Println("Your age is :", yage)
	fmt.Println("Your weight are :", yweight)

	// Declaring and assain value in variable
	fmt.Println("\n=========================================================")

	var n string = "Hasan"
	var a int = 40
	var w float32 = 40.5

	fmt.Println("\nHello Mr/Mr/Ms.", n)
	fmt.Println("Your age is :", a)
	fmt.Println("Your weight are :", w)

	// Another way to declaring and assain value in variable
	fmt.Println("\n=========================================================")

	na := "Miran"
	ag := 41
	we := 61.80

	fmt.Println("\nHello Mr/Mr/Ms.", na)
	fmt.Println("Your age is :", ag)
	fmt.Println("Your weight are :", we)
}
