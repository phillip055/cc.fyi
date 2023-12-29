package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"web_server/cmd"
)

func main() {
	fmt.Println("Starting web server...")
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleRequestConnection(conn)
	}
}

func handleRequestConnection(conn net.Conn) {
	defer conn.Close()
	buffer := bufio.NewReader(conn)
	line, _, err := buffer.ReadLine()
	if err != nil {
		panic(err)
	}
	result := cmd.Process(strings.Split(string(line), " "))
	conn.Write([]byte(result))
}
