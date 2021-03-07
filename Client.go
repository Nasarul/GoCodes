package main

import (
	"fmt"
	"net"
	"time"
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

	for {
		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}

		bs2 := make([]byte, 1024)
		m, err := conn.Read(bs2)
		if err != nil {
			fmt.Println(e.Error())
		}

		reqstr := string(bs2[:m])
		fmt.Println(reqstr)
		recvTime := time.Now().Format("2006-01-02 15:04:05")
		msg := fmt.Sprintf(`Your message: %s, received at %s`, reqstr, recvTime)
		conn.Write([]byte(msg))

	}

}
