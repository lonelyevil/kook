package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/lonelyevil/kook"
	"github.com/lonelyevil/kook/log_adapter/plog"
	"github.com/phuslu/log"
)

func main() {
	l := log.Logger{
		Level:  log.TraceLevel,
		Writer: &log.ConsoleWriter{},
	}
	s := kook.New(os.Getenv("BOTAPI"), plog.NewLogger(&l))
	s.AddHandler(messageHan)
	s.Open()
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Kook session.
	s.Close()
}

func messageHan(ctx *kook.KmarkdownMessageContext) {
	if ctx.Common.Type != kook.MessageTypeKMarkdown || ctx.Extra.Author.Bot {
		return
	}
	if strings.Contains(ctx.Common.Content, "ping") {
		ctx.Session.MessageCreate(&kook.MessageCreate{
			MessageCreateBase: kook.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  "pong",
				Quote:    ctx.Common.MsgID,
				Type:     kook.MessageTypeKMarkdown,
			},
		})
	}

}
