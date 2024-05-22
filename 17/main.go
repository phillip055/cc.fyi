package main

import (
	"bufio"
	"fmt"
	"memcache/command"
	"net"
)

func main() {
	fmt.Println("Starting memcache server...")
	listener, err := net.Listen("tcp", ":11212")
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
		cmd, err := command.Parse(buffer)
		if err != nil {
			println("failed parsing", err.Error())
		}
		fmt.Printf("%+v\n", cmd)
		response := cmd.Execute()
		messageString := fmt.Sprintf("%s\r\n", response.Message)
		_, err = conn.Write([]byte(messageString))
		if err != nil {
			println("failed writing")
		}
	}
}
