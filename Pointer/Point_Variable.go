package main

import "fmt"

func main() {
	var x int
	var y *int

	x = 10
	y = &x

	fmt.Println(&x, "   ", x)
	fmt.Println(&y, "   ", *y)
}
