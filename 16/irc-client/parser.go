package irc_client

import (
	"errors"
	"strings"
)

type IRCMessage struct {
	Prefix  string
	Command string
	Args    []string
	raw     []byte
}

func ParseMessage(msg []byte) (*IRCMessage, error) {
	skipByte(&msg, ' ')
	raw := msg
	if len(msg) == 0 {
		return nil, errors.New("empty message")
	}
	return &IRCMessage{
		raw:     raw,
		Prefix:  extractPrefix(&msg),
		Command: extractToken(&msg),
		Args:    extractArgs(&msg),
	}, nil
}

func extractArgs(msg *[]byte) (args []string) {
	skipByte(msg, ' ')
	var currentToken []byte
	for idx, b := range *msg {
		if b == ':' {
			currentToken = (*msg)[idx:]
			break
		} else if b != ' ' {
			currentToken = append(currentToken, b)
		} else {
			args = append(args, string(currentToken))
			currentToken = currentToken[:0] //clear arr
		}
	}
	if len(currentToken) > 0 {
		args = append(args, string(currentToken))
	}
	return args
}

func extractPrefix(msg *[]byte) string {
	if skipByte(msg, ':') {
		prefix := extractToken(msg)
		skipByte(msg, ' ')
		return prefix
	}
	return ""
}

func skipByte(msg *[]byte, b rune) bool {
	if len(*msg) > 0 && rune((*msg)[0]) == b {
		*msg = (*msg)[1:]
		return true
	}
	return false
}

func extractToken(msg *[]byte) string {
	var token []byte
	for len(*msg) > 0 && (*msg)[0] != ' ' {
		token = append(token, (*msg)[0])
		*msg = (*msg)[1:]
	}
	return string(token)
}

func (msg *IRCMessage) String() string {
	var result string
	if msg.Prefix != "" {
		result += ":" + msg.Prefix + " "
	}
	if msg.Command != "" {
		result += msg.Command + " "
	}
	if len(msg.Args) > 0 {
		result += strings.Join(msg.Args, " ")
	}
	if result != "" {
		result += "\r\n"
	}
	return result
}
