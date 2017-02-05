package main

import (
	"fmt"
	"net"
	"bufio"
	"time"
)

func main() {
	fmt.Println("I am the client")

	conn, _ := net.Dial("tcp", ":8080")

	fmt.Print(conn)

	writer := bufio.NewWriter(conn)

	reader := bufio.NewScanner(bufio.NewReader(conn))

	time.Sleep(5000)

	writer.WriteString("me")

	str := reader.Scan()

	fmt.Println(str)
}