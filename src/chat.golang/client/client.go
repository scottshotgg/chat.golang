// TODO: need to make asynchronous read and write

package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	//"time"
)

func input(r *bufio.Reader) {
	for {
		str, _ := r.ReadString('\n')
		fmt.Println(str)
	}
}

func output(w *bufio.Writer) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		w.WriteString(text)
		// This is what is sending the weird newline I think
		// have to also do the error checking that we did in the server side
		w.Flush()
	}
}

func main() {
	fmt.Println("I am the client")

	conn, _ := net.Dial("tcp", ":8080")

	fmt.Print(conn)

	writer := bufio.NewWriter(conn)

	reader := bufio.NewReader(conn)

	//time.Sleep(5000)

	go input(reader)
	go output(writer)

	// Solve this somehow ... ??
	for {}
}