package main

import (
	"fmt"
	)

func main(){

	b1:= struct{ 
	Name string
	F_Name string
	Roll int
	DoB int
	CGP float32
	}{
		Name:"Md. Nasarul Hasan",
		F_Name:"Late Fazlur Rahman",
		Roll:1118,
		DoB:1977,
		CGP:4.80,
	}
fmt.Println(b1)
fmt.Println(b1.Name)
fmt.Println(b1.F_Name)
fmt.Println(b1.Roll)
fmt.Println(b1.DoB)
fmt.Println(b1.CGP)
}





/*type Student struct { 
Name string
F_Name string
Roll int
DoB int
CGP float32
}

func main(){

	var b1 Student 
	b1.Name = "Md. Nasarul Hasan"  
	b1.F_Name="Late Fazlur Rahman"
	b1.Roll=1118
	b1.DoB=1977
	b1.CGP=4.80

fmt.Println(b1)          
fmt.Println(b1.Name)  
fmt.Println(b1.F_Name)
fmt.Println(b1.Roll)
fmt.Println(b1.DoB)
fmt.Println(b1.CGP)
}
*/