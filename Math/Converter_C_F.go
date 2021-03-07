package main

import "fmt"

func main(){
var x, y float32
fmt.Print("\nEnter the value of fahrenheit temperature (F): ")
fmt.Scanf("%f", &x)
y= (x-32)*5/9
fmt.Println("Converted value of celsious is : ", y)
}

/*
package main

import "fmt"

func main(){
var x, y float32
fmt.Print("\nEnter the value of celsious temperature (C): ")
fmt.Scanf("%f", &x)
y= (x*9/5)+32
fmt.Println("Converted value of Fahrenheit is (F): ", y)
}

*/