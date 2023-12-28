package cmd

import "blueis/store"

var ExecMap = map[string]func(*store.Cache, ...string) string{
	"PING":   RunPing,
	"ECHO":   RunEcho,
	"SET":    RunSet,
	"GET":    RunGet,
	"EXISTS": RunExist,
	"DEL":    RunDelete,
	"INCR":   RunIncr,
	"DECR":   RunDecr,
}

func Execute(args [][]string) string {
	cache := store.GetCacheInstance()
	var data []string
	for i := range args {
		data = append(data, args[i][1])
	}
	execFunc, ok := ExecMap[data[0]]
	if !ok {
		return "ERR unknown command '" + data[0] + "'"
	}
	return execFunc(cache, data[1:]...)
}
