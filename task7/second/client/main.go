package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n') // read msg from stdin
		if err != nil {
			if err == io.EOF { // when user type Ctrl+D to get EOF
				break
			} else {
				log.Fatal(err) // something went wrong
				return
			}
		}
		fmt.Fprintf(conn, text+"\n") // send msg
		msg := make([]byte, 100)     // buffer where we will read msg from server
		n, err := conn.Read(msg)
		if n == 0 || err != nil {
			fmt.Println(err.Error())
			break
		}

		fmt.Print(string(msg))
	}
}
