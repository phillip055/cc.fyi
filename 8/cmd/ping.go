package cmd

import "blueis/store"

func RunPing(cache *store.Cache, args ...string) string {
	return "PONG"
}
