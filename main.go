package main

import (
	"flag"
	ircevent "github.com/thoj/go-ircevent"
	"log"
)

var stream = flag.String("stream", "tehhcwool", "Your (or someone elses) streamname")
var clientID = flag.String("clientID", "9po4ts2jz2niigqq3o9gtt2ntw69njf", "Your client ID (set in settings/connections)")

var ircPassword = flag.String("password", "oauth:c3v70dw7nvznz2amu2wi0ai177cb7sf", "Twitch oauth password, get yours at http://twitchapps.com/tmi")

var server = flag.String("ip", "irc.freenode.net:6667", "IP")

var room = "#lolok"

func main() {
	log.Println("Starting up, connecting to ", *server, " as ", *stream)
	// Connect to IRC server
	irc := ircevent.IRC(*stream, *stream)

	err := irc.Connect(*server)

	if err != nil {
		log.Fatalln("Could not connect to server")
	}
	defer irc.Disconnect()

	// When we've connected to the IRC server, go join the room!
	log.Println("Connected: ", *stream)
	irc.AddCallback("001", func(e *ircevent.Event) {
		irc.Join(room)
		log.Println("Joined room ", room)
	})

	// Say something on arrival
	irc.AddCallback("JOIN", func(e *ircevent.Event) {
		irc.Privmsg(room, "LDLC? NiP")
	})
	// Check each message to see if it contains a URL, and return the title
	irc.AddCallback("PRIVMSG", func(e *ircevent.Event) {
		log.Printf("[%v] %v", e.Nick, e.Message())
	})

	irc.Loop()
}
