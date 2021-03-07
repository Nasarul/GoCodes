package main

import (
	"fmt"
)

func main() {
	var x, y, z, n int
	fmt.Print("\nEnter the value of the fibonacci step : ")
	fmt.Scanf("%d", &n)

	if n >= 0 {
		x = 0
		y = 1
		fmt.Print("\nThe positive fibonacci serise is :\n", x, y)
		n = n - 2
		for i := 1; i <= n; i++ {
			z = x + y
			fmt.Printf(" %v", z)
			x = y
			y = z
		}

	} else {
		x = 0
		y = -1
		fmt.Print("\nThe negative fibonacci serise is :\n", x, y)
		n = n + 2
		for i := -1; i >= n; i-- {
			z = x + y
			fmt.Printf(" %v", z)
			x = y
			y = z
		}
	}

}
