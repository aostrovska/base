package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var s string
	fmt.Scan(&s)

	conn.Write([]byte(s))
	switch s {
	case "delete":
		var id int
		fmt.Scan(&id)
		conn.Write(id)
	case "add":
		var person interface{}
		fmt.Scan(&person)
		conn.Write(person)
	}
	buf := make([]byte, 2000)

	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf[:n]))
}
