package main

import (
	"flag"
	"fmt"
	"hash/fnv"
	"log"

	"bot"
	"config"

	"github.com/thoj/go-ircevent"
)

var (
	cfgFname = flag.String("conf", "./config.json", "config file to use")

	parent chan *irc.Event
	bots   []*bot.Bot
	colors []int
)

func init() {
	parent = make(chan *irc.Event)
	colors = []int{2, 3, 4, 5, 7, 9, 10, 11, 12}
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

func hash(nick string) string {
	myHash := fnv.New32()
	myHash.Write([]byte(nick))

	sum := myHash.Sum32()
	sum = sum % uint32(len(colors))

	return fmt.Sprintf("\x03%d%s\x03", colors[sum], nick)
}

func sendToAllButOne(e *irc.Event) {
	for _, mybot := range bots {
		if e.Connection == mybot.IrcObj {
			continue
		}

		if mybot.Connected {
			mybot.Signal <- &irc.Event{
				Code: "PRIVMSG",
				Arguments: []string{
					e.Arguments[0],
					fmt.Sprintf("{-%s} %s", hash(e.Nick), e.Arguments[1]),
				},
			}
		}
	}
}
