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
	bots   []bot.Bot
)

func init() {
	parent = make(chan bot.Event)
}

func main() {
	flag.Parse()

	cfg, err := config.LoadFile(*cfgFname)
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg
}
