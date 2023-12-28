package resp

import (
	"bufio"
	"strconv"
)

func ReadTokens(reader *bufio.Reader) ([]rune, error) {
	var tokens []rune
	prevChar := rune(0)
	for {
		if char, _, err := reader.ReadRune(); err != nil {
			return nil, err
		} else if char == '\n' && prevChar == '\r' {
			break
		} else {
			tokens = append(tokens, char)
			prevChar = char
		}
	}
	return tokens[:len(tokens)-1], nil
}

func ParseTokenArray(tokens []rune, reader *bufio.Reader) ([][]string, error) {
	var tokenArray [][]string
	tokenArraySize, _ := strconv.Atoi(string(tokens[1:]))
	for tokenArraySize > 0 {
		metadata, err := ReadTokens(reader)
		if err != nil {
			return nil, err
		}
		data, err := ReadTokens(reader)
		if err != nil {
			return nil, err
		}
		tokenArray = append(tokenArray, []string{string(metadata), string(data)})
		tokenArraySize--
	}
	return tokenArray, nil
}
