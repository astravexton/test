package main

import (
	"log"

	"github.com/astravexton/go-ircbot/plugins/command"
	"github.com/goshuirc/irc-go/ircevent"
	"github.com/goshuirc/irc-go/ircmsg"
)

type IRCBot struct {
	conn   *ircevent.Connection
	config BotConfig
}

func main() {

	var bot = new(IRCBot)
	bot.LoadConfig()
	bot.conn = &ircevent.Connection{
		Server:       bot.config.Server,
		UseTLS:       bot.config.SSL,
		Nick:         bot.config.Nick,
		User:         bot.config.User,
		RealName:     bot.config.Gecos,
		RequestCaps:  []string{"extended-join", "account-tag"},
		SASLLogin:    bot.config.SaslUsername,
		SASLPassword: bot.config.SaslPassword,
		Debug:        bot.config.Debug,
	}

	bot.conn.AddCallback(ircevent.RPL_ENDOFMOTD, func(e ircmsg.Message) {
		bot.conn.Join("#Zyrio")
	})

	command.RegisterCallbacks(bot.conn)

	err := bot.conn.Connect()
	if err != nil {
		log.Fatal(err)
	}
	bot.conn.Loop()

}
