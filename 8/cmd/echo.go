package cmd

import "blueis/store"

func RunEcho(cache *store.Cache, args ...string) string {
	return args[0]
}
