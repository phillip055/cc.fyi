package app

import (
	ircclient "irc/irc-client"
	"strings"
)

type Message struct {
	nick string
	text string
}

type App struct {
	client *ircclient.IRCClient
}

func NewApp(client *ircclient.IRCClient) (*App, error) {
	return &App{client: client}, nil
}

func (app *App) Start() {
	ui := UI{
		OnSubmit: func(s string) {
			app.ProcessRawInput(s)
		},
	}
	app.client.Listen(func(client *ircclient.IRCClient, msg ircclient.IRCMessage) error {
		println("Got message")
		return nil
	})
	ui.Start()
}

func (app *App) ProcessRawInput(input string) {
	inputArr := strings.Split(input, " ")
	switch inputArr[0] {
	case "/join":
		_ = app.client.Join(inputArr[1])
		break
	case "/nick":
		_ = app.client.NewNick(inputArr[1])
		break
	case "/part":
		_ = app.client.Part()
		break
	case "/quit":
		_ = app.client.Quit(inputArr[1])
		break
	default:
		_ = app.client.SendMessage(input)
		break
	}
}
