package irc_client

func ProcessCommand(client *IRCClient, msg IRCMessage) IRCMessage {
	//println("from server", msg.String())
	var response IRCMessage
	switch msg.Command {
	case "PING":
		response = client.CreatePong(msg)
	}
	return response
}
