package cmd

import "blueis/store"

func RunDelete(cache *store.Cache, args ...string) string {
	_ = cache.Del(args[0])
	return "1"
}
