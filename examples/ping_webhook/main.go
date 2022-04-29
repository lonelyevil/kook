package main

import (
	"github.com/lonelyevil/khl"
	"github.com/lonelyevil/khl/log_adapter/plog"
	"github.com/phuslu/log"
	"net/http"
	"os"
	"strings"
)

func main() {
	l := &log.Logger{
		Level:  log.TraceLevel,
		Writer: &log.ConsoleWriter{},
	}
	s := khl.New(os.Getenv("BOTAPI"),
		plog.NewLogger(l),
		khl.SessionWithEncryptKey([]byte(os.Getenv("BOTKEY"))),
		khl.SessionWithVerifyToken(os.Getenv("BOTTOKEN")))
	http.HandleFunc("/endpoint", s.WebhookHandler())
	l.Trace().Msg("Bot is running now")
	s.AddHandler(messageHan)
	http.ListenAndServe(":8000", nil)
}
func messageHan(ctx *khl.KmarkdownMessageContext) {
	if ctx.Common.Type != khl.MessageTypeKMarkdown || ctx.Extra.Author.Bot {
		return
	}
	if strings.Contains(ctx.Common.Content, "ping") {
		ctx.Session.MessageCreate(&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  "pong",
				Quote:    ctx.Common.MsgID,
				Type:     khl.MessageTypeKMarkdown,
			},
		})
	}

}
