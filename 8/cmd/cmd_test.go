package cmd

import "testing"

func padArgs(args []string) [][]string {
	var data [][]string
	for i := range args {
		data = append(data, []string{"", args[i]})
	}
	return data
}

func TestCmd(t *testing.T) {
	type testCase struct {
		name string
		args []string
		want string
	}
	testCases := []testCase{
		{
			name: "SET",
			args: []string{"SET", "key", "123"},
			want: "OK",
		},
		{
			name: "GET",
			args: []string{"GET", "key"},
			want: "123",
		},
		{
			name: "PING",
			args: []string{"PING"},
			want: "PONG",
		},
		{
			name: "ECHO",
			args: []string{"ECHO", "hello"},
			want: "hello",
		},
		{
			name: "EXISTS",
			args: []string{"EXISTS", "key"},
			want: "1",
		},
		{
			name: "INCR",
			args: []string{"INCR", "key"},
			want: "124",
		},
		{
			name: "DECR",
			args: []string{"DECR", "key"},
			want: "123",
		},
		{
			name: "DEL",
			args: []string{"DEL", "key"},
			want: "1",
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			response := Execute(padArgs(tCase.args))
			if response != tCase.want {
				t.Errorf("Execute() = %v, want %v", response, tCase.want)
			}
		})
	}
}
