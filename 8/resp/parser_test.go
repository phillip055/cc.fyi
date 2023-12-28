package resp

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func padTokens(tokens []string) string {
	var data []string
	for _, token := range tokens {
		tok := []rune(token)
		tok = append(tok, '\r', '\n')
		data = append(data, string(tok))
	}
	return strings.Join(data, "")
}

func getReader(tokens []string) *bufio.Reader {
	return bufio.NewReader(
		strings.NewReader(
			padTokens(tokens),
		),
	)
}

func TestParseRequest(t *testing.T) {
	type testCase struct {
		name  string
		input []string
		want  [][]string
	}
	testCases := []testCase{
		{
			name:  "single item",
			input: []string{"*1", "$3", "foo"},
			want: [][]string{
				{"$3", "foo"},
			},
		},
		{
			name: "multiple items",
			input: []string{
				"*2", "$3", "foo", "$3", "bar",
			},
			want: [][]string{
				{"$3", "foo"},
				{"$3", "bar"},
			},
		},
	}
	for _, tCase := range testCases {
		buff := getReader(tCase.input)
		tokens, _ := ReadTokens(buff)
		actual, _ := ParseTokenArray(tokens, buff)
		if !reflect.DeepEqual(actual, tCase.want) {
			t.Errorf("ParseArray() = %v, want %v", actual, tCase.want)
		}
	}
}
