package main

import (
	"fmt"
)

func main(){
var n int
fmt.Printf("\nEnter the end value of the serice: ")
fmt.Scanf("%d", &n) 
fmt.Printf("\nThe serice are: ")

for i:=1; i<=n; i++ {
	fmt.Printf("%v ",i)
	}
}
