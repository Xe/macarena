package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Xe/macarena/config"
	"github.com/thoj/go-ircevent"
)

// Bot is the wrapper around ircobj.
type Bot struct {
	Info     config.Info
	Network  config.Network
	Channels []string
	Signal   chan *irc.Event
	IrcObj   *irc.Connection

	parent     chan *irc.Event
	log        *log.Logger
	callbackid string
}

// New makes a new Bot.
//
// TODO: make this less fucko.
func New(info config.Info, net config.Network, channels []string, parent chan *irc.Event) (bot *Bot) {
	bot.IrcObj = irc.IRC(info.Nick, info.User)

	bot.Info = info
	bot.Network = net
	bot.Channels = channels

	bot.IrcObj.UseTLS = net.UseSSL

	bot.parent = parent
	bot.Signal = make(chan *irc.Event)
	bot.log = log.New(
		os.Stdout,
		fmt.Sprintf("%s (%s) ", bot.Info.Nick, bot.Network.Name),
		log.LstdFlags,
	)

	go func() {
		bot.log.Printf("Attempting to connect to %s (%s:%d)", bot.Network.Name, bot.Network.Host, bot.Network.Port)
		err := bot.IrcObj.Connect(fmt.Sprintf("%s:%d", bot.Network.Host, bot.Network.Port))
		if err != nil {
			bot.log.Fatal(err)
		}

		bot.IrcObj.Privmsg("NickServ", "IDENTIFY "+bot.Network.ServicesPass)

		time.Sleep(500 * time.Millisecond)

		for _, channel := range bot.Channels {
			bot.IrcObj.Join(channel)
		}

		bot.callbackid = bot.IrcObj.AddCallback("PRIVMSG", func(e *irc.Event) {
			// Callback hell much but concurrency is nice
			go func() {
				bot.parent <- e
			}()
		})
	}()

	return
}
