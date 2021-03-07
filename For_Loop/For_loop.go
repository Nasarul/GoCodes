package main

import (
	"fmt"
)

func add(x, y int) (total int) {
	total = x + y
	return
}

func sub(x, y int) (total int) {
	total = x - y
	return
}

func mul(x, y int) (total int) {
	total = x * y
	return
}

func div(x, y int) (total int) {
	total = x / y
	return
}

func main() {
	c, d := 5, 4
	fmt.Println(add(c, d))
	fmt.Println(sub(c, d))
	fmt.Println(mul(c, d))
	fmt.Println(div(c, d))

	//fmt.Sprintf(add("Sum of the value : ", total))
	//fmt.Println(sub("Subtact of the value : " c, d))
	//fmt.Println(mul("Multiple of the value : " c, d))
	//fmt.Println(div("Divided of the value : "c, d))
}
