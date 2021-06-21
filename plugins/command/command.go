package command

import (
	"fmt"

	"github.com/goshuirc/irc-go/ircevent"
	"github.com/goshuirc/irc-go/ircmsg"
)

func RegisterCallbacks(irc *ircevent.Connection) {
	irc.AddCallback("PRIVMSG", func(e ircmsg.Message) {
		fmt.Printf("%+v\n", e)
	})
}
