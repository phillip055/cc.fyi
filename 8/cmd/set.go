package cmd

import (
	"blueis/store"
	"strconv"
	"time"
)

func RunSet(cache *store.Cache, args ...string) string {
	item := store.Item{
		Key:   args[0],
		Value: args[1],
		TTL:   0,
	}
	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "xx":
			// don't set if key already exists
			if _, found := cache.Get(args[0]); found {
				return "OK"
			}
		case "nx":
			// don't set if key doesn't exist
			if _, found := cache.Get(args[0]); !found {
				return "OK"
			}
		case "ex":
			// seconds from now
			x, _ := strconv.Atoi(args[i+1])
			item.TTL = time.Now().Add(time.Duration(x) * time.Second).UnixNano()
			i++
		case "px":
			// milliseconds from now
			x, _ := strconv.Atoi(args[i+1])
			item.TTL = time.Now().Add(time.Duration(x) * time.Millisecond).UnixNano()
			i++
		case "pxat":
			// milliseconds since epoch
			pxat, _ := strconv.Atoi(args[i+1])
			item.TTL = int64(pxat) * 1000000
			i++
		case "exat":
			// seconds since epoch
			exat, _ := strconv.Atoi(args[i+1])
			item.TTL = int64(exat) * 1000000000
			i++
		}
	}
	err := cache.Set(args[0], item)
	if err != nil {
		return "Error"
	}
	return "OK"
}
