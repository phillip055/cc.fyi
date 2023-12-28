package cmd

import (
	"blueis/store"
	"strconv"
)

func RunDecr(cache *store.Cache, args ...string) string {
	item, found := cache.Get(args[0])
	if !found {
		return "0"
	}
	if intVal, err := strconv.Atoi(item.Value); err == nil {
		item.Value = strconv.Itoa(intVal - 1)
		if err = cache.Set(item.Key, item); err != nil {
			return "ERR " + err.Error()
		}
		return item.Value
	}
	return "0"
}
