package main

import (
	"fmt"
)

func main(){
var res int
fmt.Println("Enter your marks: ")
fmt.Scanf("%d", &res)

if res <=100 && res>=80 {
		fmt.Println("You got : A+")
	} else if res >=70 && res<=79 {
		fmt.Println("You got : A")
	} else if res >=60 && res<=69 {
		fmt.Println("You got : A-")
	} else if res >=50 && res<=59 {
		fmt.Println("You got : B")
	} else if res >=40 && res<=49 {
		fmt.Println("You got : C")
	} else if res >=33 && res<=39 {
		fmt.Println("You got : D")
	} else {
		fmt.Println("Sorry to say, you are fail”)
	}
}