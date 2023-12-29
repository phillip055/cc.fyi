package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Process(args []string) string {
	path := fmt.Sprintf("www%s", args[1])
	file, err := os.Open(path)
	if err != nil {
		return "HTTP/1.1 500 Something went wrong\n\n"
	}
	fileInfo, err := file.Stat()
	if strings.HasSuffix(path, "/") || fileInfo.IsDir() {
		path += "/index.html"
	}
	dat, err := os.ReadFile(path)
	if err != nil {
		return "HTTP/1.1 400 Not Found\n\n"
	}
	return fmt.Sprintf("HTTP/1.1 200 OK\n\n%s", string(dat))
}
