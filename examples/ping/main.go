package main

import (
	"fmt"
	"github.com/lonelyevil/khl"
	"github.com/lonelyevil/khl/log_adapter/plog"
	"github.com/phuslu/log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	l := log.Logger{
		Level:  log.TraceLevel,
		Writer: &log.ConsoleWriter{},
	}
	s := khl.New(os.Getenv("BOTAPI"), plog.NewLogger(&l))
	s.AddHandler(messageHan)
	s.Open()
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the KHL session.
	s.Close()
}

func messageHan(s *khl.Session, edg *khl.EventDataGeneral, etm *khl.EventTextMessage) {
	if edg.Type != khl.MessageTypeText || etm.Author.Bot {
		return
	}
	if strings.Contains(edg.Content, "ping") {
		s.MessageCreate(&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				TargetID: edg.TargetID,
				Content:  "pong",
				Quote:    edg.MsgID,
			},
		})
	}

}
