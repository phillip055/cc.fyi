package main

import (
	"irc/app"
	irc_client "irc/irc-client"
)

func main() {
	client := irc_client.IRCClient{
		Host:       "irc.freenode.net",
		Port:       "6667",
		Nick:       "Phillipee",
		ClientName: "Lappie4",
	}

	a, err := app.NewApp(&client)
	if err != nil {
		println(err.Error())
	}
	// go client.Start()
	go a.Start()
	client.Start()
	// defer func() {
	// 	client.Quit("Please")
	// }()
	// go client.Start()
	// a, _ := app.NewApp(&client)
	// a.Start()
}
