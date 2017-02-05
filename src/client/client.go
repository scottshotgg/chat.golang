package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("I am the client")

	conn, _ := net.Dial("tcp", ":8080")

	print(conn)
}