package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening on port :6378")

	// Create a new server
	l, err := net.Listen("tcp", ":6378")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		res := NewResp(conn)
		value, err := res.Read()

		if err != nil {
			fmt.Println("error reading from client: ", err.Error())
			return
		}

		fmt.Println(value)

		// ignore request and send back an RESP OK response
		conn.Write([]byte("+OK\r\n"))
		writer := NewWriter(conn)
		writer.Write(Value{typ: "string", str: "OK"})
	}
}
