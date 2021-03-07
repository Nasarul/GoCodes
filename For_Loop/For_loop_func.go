package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Enter the end value of the serice :")
	fmt.Scanf("%d", &n)
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func odd() {
	for i := 2; i <= n; i = i + 2 {
		//fmt.Println(i)
		return i
	}
}

func even() {
	for i := 1; i <= n; i = i + 2 {
		fmt.Println(i)
	}
}
