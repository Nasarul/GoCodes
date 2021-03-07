package main

import "fmt"

/*
func main() {
var b int
fmt.Print("\nEnter value of Factorial number : ")
fmt.Scanf("%d", &b)
Print(g(b))
}

func g(num int) []int {
  var array = make([]int, num)
  var count int = 0
  for num > 0 {
    count += 1
    array[count -1] = count
    num -= 1
  }
  return array
}

func s(num int)  []int {
  var array = make([]int, num)
  for i := 0; i < num; i +=1 {
    array[i] = i + 1
  }
  return array
}

func Print(arr []int) {
var sum int
sum=1
  for _, value := range arr {
    	sum=sum*value
  }
fmt.Println("\nThe factorial sum is : ", sum)
}

func main() {
  var n int
  fmt.Print("Enter n value : ")
  fmt.Scanf("%d", &n)
  set(n)
}

func set(n int) {
  var f int
  a:= make([]int, n)
  for i := 1; i <= n; i++ {
     fmt.Scan(&a[i])

    }
  fmt.Println(f)
 }


i:=0
var a[5] int
for(i<5){
fmt.Print("Enter Input")
var input int
fmt.Scanf("%d",&input)
a[i]=input
i+=1
}
fmt.Print(a)
*/

func main() {
	var n int
	fmt.Println("enter n :")
	fmt.Scanf("%d", &n)
	set(n)
}

func set(n int) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf(&a[i])
	}
	fmt.Println(a)
}

/*
func main() {
	x := [5] int{1,2,3,4,5}
fmt.Println(x)
}
*/
