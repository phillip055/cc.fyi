package command

import (
	"bufio"
	"memcache/store"
)

type SetCommand CommandInput

func ParseSetCommand(buffer *bufio.Reader) (Command, error) {
	input, err := parseCommandInput(buffer)
	if err != nil {
		return nil, err
	}
	line, _, _ := buffer.ReadLine()
	return &SetCommand{
		Key:       input.Key,
		Name:      input.Name,
		Data:      string(line),
		NoReply:   input.NoReply,
		Flags:     input.Flags,
		ByteCount: input.ByteCount,
		ExpTime:   input.ExpTime,
	}, nil
}

func (cmd *SetCommand) Execute() Response {
	err := store.GetCacheInstance().Set(cmd.Key, store.Item{
		Key:   cmd.Key,
		Value: cmd.Data,
		TTL:   int64(cmd.ExpTime),
	})
	if err != nil {
		return Response{
			Err:     err,
			Message: "NOT_STORED",
		}
	}
	return Response{
		Err:     nil,
		Message: "STORED",
	}
}
