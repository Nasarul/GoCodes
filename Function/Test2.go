package main

import "fmt"

func add(x, y int) (a, b int) {
	a = x + y
	b = x - y
	return
}

func main() {
	var i, j int
	i, j = add(20, 30)
	fmt.Println(i, j)

}
