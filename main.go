package main

import (
	"flag"
	ircevent "github.com/thoj/go-ircevent"
	"log"
)

var stream = flag.String("stream", "tehhcwool", "Your (or someone elses) streamname")
var clientID = flag.String("clientID", "9po4ts2jz2niigqq3o9gtt2ntw69njf", "Your client ID (set in settings/connections)")

var ircPassword = flag.String("password", "oauth:c3v70dw7nvznz2amu2wi0ai177cb7sf", "Twitch oauth password, get yours at http://twitchapps.com/tmi")

func main() {
	flag.Parse()

	irc := ircevent.IRC(*stream, *stream)
	irc.Debug = true
	irc.Password = *ircPassword

	err := irc.Connect("irc.twitch.tv:6667")

	if err != nil {
		log.Println("IRC CON FAILED")
		return
	}

	irc.AddCallback("001",
		func(e *ircevent.Event) {
			irc.Join(*stream)
		})

	irc.AddCallback("JOIN",
		func(e *ircevent.Event) {
			log.Println("Connected: " + e.Raw)
			irc.Privmsg("#tehhcwool", "Hello")
		})

	irc.AddCallback("GLOBALMSG",
		func(e *ircevent.Event) {
			log.Println("[ " + e.Nick + " ] " + e.Message())
		})

	irc.Loop()
}
