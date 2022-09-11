package main

import (
	"github.com/lonelyevil/kook"
	"github.com/lonelyevil/kook/log_adapter/plog"
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
	s := kook.New(os.Getenv("BOTAPI"),
		plog.NewLogger(l),
		kook.SessionWithEncryptKey([]byte(os.Getenv("BOTKEY"))),
		kook.SessionWithVerifyToken(os.Getenv("BOTTOKEN")))
	http.HandleFunc("/endpoint", s.WebhookHandler())
	l.Trace().Msg("Bot is running now")
	s.AddHandler(messageHan)
	http.ListenAndServe(":8000", nil)
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
