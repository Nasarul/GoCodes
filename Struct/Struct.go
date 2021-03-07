package main

import "fmt"

func main(){

st:= struct{
name string
dob int
class int
roll int
}{
st.name:"Hasan",
st.dob:"1977",
st.class:"10",
st.roll:"20",
}
fmt.Println(st)
fmt.Println(st.name)
}

/*
func main(){
type student struct{
name string
dob int
class int
roll int
}

var st student
st.name="Hasan"
st.dob= 1977
st.class= 10
st.roll=20

fmt.Println(st)
fmt.Println(st.name)
}
*/