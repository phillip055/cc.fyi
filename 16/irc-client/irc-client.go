package irc_client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type IRCClient struct {
	Host           string
	Port           string
	currentChannel string
	Nick           string
	ClientName     string
	conn           net.Conn
}

func (client *IRCClient) Start() error {
	if err := client.Connect(); err != nil {
		return err
	}
	if err := client.Logon(); err != nil {
		return err
	}
	err := client.Listen(func(client *IRCClient, msg IRCMessage) error {
		response := ProcessCommand(client, msg)
		return client.send(response)
	})
	if err != nil {
		return err
	}
	return nil
}

func (client *IRCClient) Connect() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", client.Host, client.Port))
	if err != nil {
		return err
	}
	client.conn = conn
	return nil
}

func (client *IRCClient) Logon() error {
	_, err := client.conn.Write([]byte(fmt.Sprintf("NICK %s\n", client.Nick)))
	if err != nil {
		return err
	}
	_, err = client.conn.Write([]byte(fmt.Sprintf("USER guest 0 * :%s\n", client.ClientName)))
	if err != nil {
		return err
	}
	return nil
}

func (client *IRCClient) Listen(handle func(client *IRCClient, msg IRCMessage) error) error {
	for {
		buffer := bufio.NewReader(client.conn)
		if buffer == nil {
			return nil
		}
		line, _, err := buffer.ReadLine()
		println("dude")
		if len(line) == 0 {
			continue
		}
		if err != nil {
			return err
		}
		msg, err := ParseMessage(line)
		if err != nil {
			return err
		}
		_ = handle(client, *msg)
	}
}

func (client *IRCClient) send(msg IRCMessage) error {
	response := msg.String()
	if response == "" {
		return nil
	}
	_, err := client.conn.Write([]byte(response))
	if err != nil {
		return err
	}
	return nil
}

func (client *IRCClient) CreatePong(msg IRCMessage) IRCMessage {
	return IRCMessage{
		Command: "PONG",
		Args:    msg.Args,
	}
}

func (client *IRCClient) Join(channel string) error {
	msg := IRCMessage{
		Command: "JOIN",
		Args:    []string{channel},
	}
	err := client.send(msg)
	if err == nil {
		client.currentChannel = channel
	}
	return err
}

func (client *IRCClient) Part() error {
	msg := IRCMessage{
		Command: "PART",
		Args:    []string{client.currentChannel},
	}
	return client.send(msg)
}

func (client *IRCClient) NewNick(newNick string) error {
	msg := IRCMessage{
		Command: "NICK",
		Args:    []string{newNick},
	}
	return client.send(msg)
}

func (client *IRCClient) SendMessage(input string) error {
	msg := IRCMessage{
		Command: "PRIVMSG",
		Args:    []string{client.currentChannel, ":" + input},
	}
	return client.send(msg)
}

func (client *IRCClient) Quit(message string) error {
	msg := IRCMessage{
		Command: "QUIT",
		Args:    []string{":" + message},
	}
	err := client.send(msg)
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
