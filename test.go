package main

import "fmt"

func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func sub(x int, y int) int {
	total := 0
	total = x - y
	return total
}

func mul(x int, y int) int {
	total := 0
	total = x * y
	return total
}

func div(x int, y int) int {
	total := 0
	total = x / y
	return total
}

func main() {

	sum = add(20, 30)
	sum = sub(20, 30)
	sum = mul(20, 30)
	sum = div(20, 30)

fmt.Println("Sum of the value : ", sum)
fmt.Println("Subtact of the value : ", sum)
fmt.Println("Multiple of the value : ", sum)
fmt.Println("Divided of the value : ", sum)
}













/*
func main (){

//a := [] int {1, 2, 3, 4, 5}
var x int
var a [5] int
a[0]=10
a[1]=20
a[2]=30
a[3]=40
a[4]=50


fmt.Print("lenth of the array is : ", len(a))
fmt.Print("\nThe serice is : ")
for i=0; i<5; i++{
fmt.Print(a[i]," ")


x = a[0:2]
fmt.Println(x)
}


*/





























/*
package main

import (
	"fmt"
)

func main (){

	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("\n---------------Example 1--------------------\n")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}
}



nameArray:=[4] string {"Md", "Nasarul", "Hasan", "Miran"}
var rollArray [4] int


nameArray[0]="Md."
nameArray[1]="Nasarul"
nameArray[2]="Hasan"
nameArray[3]="(Miran)"

rollArray[0]= 01
rollArray[1]= 02
rollArray[2]= 03
rollArray[3]= 04

fmt.Println("Print Name", nameArray[1])
fmt.Println("Print Roll", rollArray[1])

}
*/