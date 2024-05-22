package command

import (
	"bufio"
	"bytes"
	"errors"
	"memcache/store"
)

type GetCommand struct {
	Key string
}

func ParseGetCommand(buffer *bufio.Reader) (Command, error) {
	getCmdLine, _, _ := buffer.ReadLine()
	getCmd := bytes.Split(getCmdLine, []byte{' '})
	if len(getCmd) != 2 {
		return nil, errors.New("invalid get command")
	}
	return &GetCommand{
		Key: string(getCmd[1]),
	}, nil
}

func (cmd *GetCommand) Execute() Response {
	item, found := store.GetCacheInstance().Get(cmd.Key)
	if found {
		return Response{
			Message: item.Value,
		}
	}
	return Response{
		Message: "END",
	}
}
