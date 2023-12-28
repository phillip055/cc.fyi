package cmd

import "blueis/store"

func RunExist(cache *store.Cache, args ...string) string {
	_, found := cache.Get(args[0])
	if found {
		return "1"
	}
	return "0"
}
