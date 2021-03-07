package main

import (
	"fmt"
)

func main(){
var n int
fmt.Print("Enter the end value of the odd serice : ")
fmt.Scanf("%d",&n) 
fmt.Print("The odd numbers are : ")

for i:=1; i<=n; i=i+2 {
	fmt.Printf("%v ", i)
	}
}

