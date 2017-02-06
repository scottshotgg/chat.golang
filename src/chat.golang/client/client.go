// TODO: need to make asynchronous read and write

package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"time"
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
	protocol := "tcp"
	host := "127.0.0.1"
	port := "8080"
	
	conn, err := net.Dial(protocol, host + ":" + port)

	if err != nil {
		for i := 0; i < 10; i++ {
			conn, err = net.Dial(protocol, host + ":" + port)

			fmt.Println("Could not get connection, trying again in 2 seconds...")
			time.Sleep(2 * time.Second)
		}
		fmt.Println("Could not connect to the server after 10 tries. Try again later.")
	}




	fmt.Print(conn)

	writer := bufio.NewWriter(conn)

	reader := bufio.NewReader(conn)

	//time.Sleep(5000)

	go input(reader)
	go output(writer)

	// Solve this somehow ... ??
	for {}
}