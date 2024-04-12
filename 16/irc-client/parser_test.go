package irc_client

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectedMsg IRCMessage
	}{
		{
			name:  "digits as command",
			input: ":*.freenode.net 002 CCIRC :Your host is *.freenode.net, running version InspIRCd-3",
			expectedMsg: IRCMessage{
				raw:     []byte(":*.freenode.net 002 CCIRC :Your host is *.freenode.net, running version InspIRCd-3"),
				Command: "002",
				Prefix:  "*.freenode.net",
				Args:    []string{"CCIRC", ":Your host is *.freenode.net, running version InspIRCd-3"},
			},
		},
		{
			name:  "type 2",
			input: ":.freenode.net NOTICE * :* Looking up your ident...",
			expectedMsg: IRCMessage{
				raw:     []byte(":.freenode.net NOTICE * :* Looking up your ident..."),
				Prefix:  ".freenode.net",
				Command: "NOTICE",
				Args:    []string{"*", ":* Looking up your ident..."},
			},
		},
		{
			name:  "PING",
			input: "PING :ARl|vaTgY",
			expectedMsg: IRCMessage{
				raw:     []byte("PING :ARl|vaTgY"),
				Prefix:  "",
				Command: "PING",
				Args:    []string{":ARl|vaTgY"},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			msg, err := ParseMessage([]byte(testCase.input))
			if err != nil {
				t.Error(err)
			}
			if msg.Prefix != testCase.expectedMsg.Prefix {
				t.Errorf("Prefix() = %v, want %v", (*msg).Prefix, testCase.expectedMsg.Prefix)
			}
			if msg.Command != testCase.expectedMsg.Command {
				t.Errorf("Command() = %v, want %v", (*msg).Command, testCase.expectedMsg.Command)
			}
			if !reflect.DeepEqual((*msg).Args, testCase.expectedMsg.Args) {
				t.Errorf("Args() = %v, want %v", (*msg).Args, testCase.expectedMsg.Args)
			}
		})
	}
}
