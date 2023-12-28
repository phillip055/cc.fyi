package main

import (
	"blueis/cmd"
	"blueis/resp"
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting redis server...")
	listener, err := net.Listen("tcp", ":6379")
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

	for {
		tokens, err := resp.ReadTokens(buffer)
		if err != nil {
			return
		}
		input, err := resp.ParseTokenArray(tokens, buffer)
		output := cmd.Execute(input)
		output = fmt.Sprintf("+%s\r\n", output)
		_, err = conn.Write([]byte(output))
	}
}
