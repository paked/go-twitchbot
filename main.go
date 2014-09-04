package main

import (
	"flag"
	ircevent "github.com/thoj/go-ircevent"
	"log"
)

var stream = flag.String("stream", "paked", "Your (or someone elses) streamname")
var clientID = flag.String("clientID", "9po4ts2jz2niigqq3o9gtt2ntw69njf", "Your client ID (set in settings/connections)")

var ircPassword = flag.String("password", "oauth:jerbrv285soc70ckkhr2qdqgal4bul1", "Twitch oauth password, get yours at http://twitchapps.com/tmi")
var room = flag.String("chatroom", "", "The chat you want to listen in on, leave blank if joining your own")
var server = flag.String("ip", "irc.twitch.tv:6667", "IP")

func main() {
	log.Println("Starting up, connecting to ", server, " as ", *stream)

	// Connect to IRC server
	irc := ircevent.IRC(*stream, *stream)
	irc.Password = *ircPassword
	err := irc.Connect(*server)

	if err != nil {
		log.Fatalln("Could not connect to server")
	}

	if *room == "" {
		*room = "#" + *stream
	}

	// When we've connected to the IRC server, go join the room!
	log.Println("Connected: ", *stream)
	irc.AddCallback("001", func(e *ircevent.Event) {
		irc.Join(*room)
	})

	irc.AddCallback("JOIN", func(e *ircevent.Event) {
		log.Println("[{ME}] Listening to your chat")
	})

	// Check each message to see if it contains a URL, and return the title
	irc.AddCallback("PRIVMSG", func(e *ircevent.Event) {
		log.Printf("[%v] %v", e.Nick, e.Message())
	})

	irc.Loop()
}
