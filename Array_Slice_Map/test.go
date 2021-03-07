package main

import (
	"fmt"
)

func main() {
	var number [5]int
	var city [5]string
	//var matrix [3][3] int

	fmt.Println("<========== Insurt data =======>")
	for i := 0; i < 5; i++ {
		number[i] = i
		city[i] = fmt.Sprint(i)
		fmt.Println(i)
		fmt.Println(i)
	}

}
