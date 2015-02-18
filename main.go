package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Xe/macarena/bot"
	"github.com/Xe/macarena/config"
	"github.com/thoj/go-ircevent"
)

var (
	cfgFname = flag.String("conf", "./config.json", "config file to use")

	parent chan *irc.Event
	bots   []*bot.Bot
)

func init() {
	parent = make(chan *irc.Event)
}

func main() {
	flag.Parse()

	cfg, err := config.LoadFile(*cfgFname)
	if err != nil {
		log.Fatal(err)
	}

	for _, net := range cfg.Networks {
		mybot := bot.New(cfg.MyInfo, net, cfg.Channels, parent)

		bots = append(bots, mybot)
	}

	for e := range parent {
		sendToAllButOne(e)
	}
}

func sendToAllButOne(e *irc.Event) {
	for _, mybot := range bots {
		if e.Connection == mybot.IrcObj {
			continue
		}

		mybot.Signal <- &irc.Event{
			Code: "PRIVMSG",
			Arguments: []string{
				e.Arguments[0],
				fmt.Sprintf("<-%s> %s", e.Nick, e.Arguments[1]),
			},
		}
	}
}
