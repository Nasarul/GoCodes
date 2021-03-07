package main

import (
	"fmt"
	"net"
)

func main() {
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
	}
	conn.Write([]byte("Hallo its Hasan"))
	bs := make([]byte, 1024)
	n, _ := conn.Read(bs)
	fmt.Println(string(bs[:n]))
	conn.Close()
}
