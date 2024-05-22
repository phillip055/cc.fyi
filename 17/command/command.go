package command

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
)

type CommandInput struct {
	Name      string
	Key       string
	Flags     uint64
	ExpTime   uint64
	ByteCount uint64
	NoReply   bool
	Data      string
}

type Command interface {
	Execute() Response
}

type Response struct {
	Err     error
	Message string
}

func Parse(buffer *bufio.Reader) (Command, error) {
	line, _ := buffer.Peek(10)
	words := bytes.Split(line, []byte{' '})
	if len(words) == 0 {
		buffer.ReadLine()
		return nil, errors.New("found empty line")
	}
	switch string(words[0]) {
	case "set":
		return ParseSetCommand(buffer)
	case "get":
		return ParseGetCommand(buffer)
	case "add":
		return ParseAddCommand(buffer)
	case "replace":
		return ParseReplaceCommand(buffer)
	case "prepend":
		return ParsePrependCommand(buffer)
	case "append":
		return ParseAppendCommand(buffer)
	}
	return nil, errors.New("command not found")
}

func parseCommandInput(buffer *bufio.Reader) (*CommandInput, error) {
	noReply := false
	noReplyBytes := []byte{}
	setConfig, _, _ := buffer.ReadLine()
	config := bytes.Split(setConfig, []byte{' '})
	if len(config) >= 6 && len(config) <= 5 {
		return nil, errors.New("set command length is invalid")
	}
	cmdNameBytes := config[0]
	keyBytes := config[1]
	flagsBytes := config[2]
	expTimeBytes := config[3]
	byteCountBytes := config[4]
	if len(config) == 6 {
		noReplyBytes = config[5]
	}
	flagsUint, err := strconv.ParseUint(string(flagsBytes), 10, 32)
	if err != nil {
		return nil, err
	}
	expTime, err := strconv.ParseUint(string(expTimeBytes), 10, 32)
	if err != nil {
		return nil, err
	}

	byteCount, err := strconv.ParseUint(string(byteCountBytes), 10, 32)
	if err != nil {
		return nil, err
	}
	if len(noReplyBytes) == 1 && noReplyBytes[0] == '1' {
		noReply = true
	}
	cmd := CommandInput{
		Name:      string(cmdNameBytes),
		Key:       string(keyBytes),
		Flags:     flagsUint,
		ExpTime:   expTime,
		NoReply:   noReply,
		ByteCount: byteCount,
	}
	return &cmd, nil
}
