package main

import (
	"fmt"
)

func main(){
var x, y, z, n int
fmt.Print("\nEnter the negative value of end step of the fibonacci sequence : ")
fmt.Scanf("%d", &n)
x=0; y=-1
fmt.Print("\nThe febonaki serise is :\n" ,x,y)
n=n+2

for i:=-1; i>=n; i--{
	z=x+y
	fmt.Printf(" %v",z)	
	x=y
	y=z
	}
}


