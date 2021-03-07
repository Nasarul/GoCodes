package main

import (
	"fmt"
)

func main(){
var res int
fmt.Println("Enter your marks: ")
fmt.Scanf("%d", &res)

//This code write for practic 
switch res {
	case 100,99,98,97,96,95,94,93,92,91,90,89,88,87,86,85,84,83,82,81,80:
		fmt.Println("You got : A+")
	case 79,78,77,76,75,74,73,72,71,70:
		fmt.Println("You got : A")
	case 69,68,67,66,65,64,63,62,61,60:
		fmt.Println("You got : A-")
	case 59,58,57,56,55,54,53,52,51,50:
		fmt.Println("You got : B")
	case 49,48,47,46,45,44,43,42,41,40:
		fmt.Println("You got : C")
	case 39,38,37,36,35,34,33:
		fmt.Println("You got : D")
	default:
		fmt.Println("Sorry to say, you are fail")
	}
}

