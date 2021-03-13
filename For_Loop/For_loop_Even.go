package main

import (
	"fmt"
)

func main() {
	var n, s int
	fmt.Printf("\nEnter the end value of the even serice: ")
	fmt.Scanf("%d", &n)
	fmt.Println("\nThe even numbers are: ")

	for i := 2; i <= n; i = i + 2 {
		fmt.Printf("%v  ", i)
	s += i
	}
fmt.Print("\n")
fmt.Print("\nThe sum of the serice is :", s)
fmt.Print("\n")
}
