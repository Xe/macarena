package main

import (
	"flag"
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
}
