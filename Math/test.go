package main

import (
	"fmt"
)

func add(a, b int) (r, t int){
r= a+b
t=a*b
return 
}


func main() {
	x, y := add(30, 40)
	fmt.Println(x, y)
}
