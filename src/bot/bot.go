package bot

import (
	"config"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/thoj/go-ircevent"
)

// Bot is the wrapper around ircobj.
type Bot struct {
	Info      *config.Info
	Network   *config.Network
	Channels  []string
	Signal    chan *irc.Event
	IrcObj    *irc.Connection
	Connected bool

	parent     chan *irc.Event
	log        *log.Logger
	callbackid string
	nick       string
	user       string
}

// New makes a new Bot.
//
// TODO: make this less fucko.
func New(info config.Info, net config.Network, channels []string, parent chan *irc.Event) (bot *Bot) {
	bot = &Bot{
		nick:     info.Nick,
		user:     info.User,
		Channels: channels,
		Network:  &net,
		Info:     &info,
		parent:   parent,
		Signal:   make(chan *irc.Event),
	}

	bot.log = log.New(
		os.Stdout,
		fmt.Sprintf("(%s) ", bot.Network.Name),
		log.LstdFlags,
	)

	bot.IrcObj = bot.connect()
	if bot.IrcObj == nil {
		log.Fatal("connect() failed")
	}

	bot.seed()

	// edge trigger
	go func() {
		for e := range bot.Signal {
			switch e.Code {
			case "PRIVMSG":
				bot.log.Printf("%s <%s> %s", e.Arguments[0], bot.Info.Nick, e.Arguments[1])
				bot.IrcObj.Privmsg(e.Arguments[0], e.Arguments[1])
			}
		}
	}()

	return
}

func (bot Bot) connect() *irc.Connection {
	irco := irc.IRC(bot.nick, bot.user)
	irco.UseTLS = bot.Network.UseSSL

	irco.Log = bot.log

	bot.log.Printf("Attempting to connect to %s (%s:%d)", bot.Network.Name, bot.Network.Host, bot.Network.Port)
	err := irco.Connect(fmt.Sprintf("%s:%d", bot.Network.Host, bot.Network.Port))
	if err != nil {
		bot.log.Fatal(err)
	}

	return irco
}

// Send sends an irc.Event to the server.
func (bot *Bot) Message(target, body string) {
	bot.IrcObj.Privmsg(target, body)
}

func (bot *Bot) seed() {
	bot.IrcObj.AddCallback("001", func(e *irc.Event) {
		bot.log.Println("Identifying to NickServ...")
		bot.IrcObj.Privmsg("NickServ", "IDENTIFY "+bot.Network.ServicesPass)

		for _, ch := range bot.Channels {
			bot.parent <- &irc.Event{
				Code: "PRIVMSG",
				Arguments: []string{
					ch,
					fmt.Sprintf(
						"I have connected to %s!",
						bot.Network.Name,
					),
				},
			}
		}

		time.Sleep(5 * time.Second)

		for _, channel := range bot.Channels {
			bot.log.Printf("Joining %s", channel)
			bot.IrcObj.Join(channel)
		}

		bot.Connected = true
	})

	bot.IrcObj.AddCallback("CTCP_ACTION", func(e *irc.Event) {
		go func() {
			bot.parent <- &irc.Event{
				Code:       "PRIVMSG",
				Connection: e.Connection,
				Nick:       e.Nick,
				Arguments: []string{
					e.Arguments[0],
					fmt.Sprintf("* %s", e.Arguments[1]),
				},
			}
		}()
	})

	bot.IrcObj.AddCallback("PRIVMSG", func(e *irc.Event) {
		if strings.HasPrefix(e.Arguments[0], "#") {
			bot.parent <- e
		} else {
			bot.log.Printf("got private message from %s", e.Nick)
			bot.IrcObj.Notice(e.Nick, "I do not yet understand commands in PM")
		}
	})

	bot.IrcObj.AddCallback("ERROR", func(e *irc.Event) {
		bot.log.Printf("wtf i was killed! %s", e.Message())

		bot.Connected = false

		for _, ch := range bot.Channels {
			bot.parent <- &irc.Event{
				Code: "PRIVMSG",
				Arguments: []string{
					ch,
					fmt.Sprintf(
						"I was just killed off of %s! %s",
						bot.Network.Name,
						e.Raw,
					),
				},
			}
		}

		go func() {
			bot.IrcObj = bot.connect()
			bot.seed()
		}()
	})
}
