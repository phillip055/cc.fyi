package cmd

import "blueis/store"

func RunGet(cache *store.Cache, args ...string) string {
	item, found := cache.Get(args[0])
	if found {
		return item.Value
	}
	return "(nil)"
}
