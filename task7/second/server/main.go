package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	exitString = "exit"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		//input := make([]byte, (1024)) // read to buffer size 1024
		//n, err := conn.Read(input)
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		msg = msg[:len(msg) - 1]
		if msg == exitString {
			return
		}

		num, err := strconv.Atoi(msg)
		if err != nil {
			conn.Write([]byte(strings.ToUpper(msg) + "\n"))
		} else {
			conn.Write([]byte(fmt.Sprintf("%d\n", num * 2)))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	log.Println("Start server at :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}
		go handleConn(conn)
	}
}
