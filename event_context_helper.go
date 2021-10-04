package khl

// Reply is a helper function for replying to a message.
//
// It accept optional MessageCreateOption, DirectMessageCreateOption and ReplyOption arguments.
func (c *TextMessageContext) Reply(text string, options ...interface{}) (*MessageResp, error) {
	if c.Common.ChannelType == "PERSON" {
		dmc := &DirectMessageCreate{
			ChatCode: c.Common.TargetID,
			MessageCreateBase: MessageCreateBase{
				Quote:   c.Common.MsgID,
				Content: text,
			},
		}
		for _, item := range options {
			switch v := item.(type) {
			case DirectMessageCreateOption:
				v(dmc)
			default:
			}
		}
		return c.Session.DirectMessageCreate(dmc)
	}
	mc := &MessageCreate{
		MessageCreateBase: MessageCreateBase{
			Quote:    c.Common.MsgID,
			Content:  text,
			TargetID: c.Common.TargetID,
		},
	}
	for _, item := range options {
		switch v := item.(type) {
		case MessageCreateOption:
			v(mc)
		case ReplyOption:
			switch v {
			case ReplyOptionTemp:
				mc.TempTargetID = c.Extra.Author.ID
			}
		default:
		}
	}
	return c.Session.MessageCreate(mc)

}

// MessageCreateOption is the type for decorator of MessageCreate.
type MessageCreateOption func(*MessageCreate)

// MessageCreateWithKmarkdown changes message type to Kmarkdown.
func MessageCreateWithKmarkdown() MessageCreateOption {
	return func(mc *MessageCreate) {
		mc.Type = MessageTypeKMarkdown
	}
}

// MessageCreateWithCard changes message type to card.
func MessageCreateWithCard() MessageCreateOption {
	return func(mc *MessageCreate) {
		mc.Type = MessageTypeCard
	}
}

// DirectMessageCreateOption is the type for decorator of DirectMessageCreate.
type DirectMessageCreateOption func(*DirectMessageCreate)

// DirectMessageCreateWithKmarkdown changes message type to Kmarkdown.
func DirectMessageCreateWithKmarkdown() DirectMessageCreateOption {
	return func(mc *DirectMessageCreate) {
		mc.Type = MessageTypeKMarkdown
	}
}

// DirectMessageCreateWithCard changes message type to card.
func DirectMessageCreateWithCard() DirectMessageCreateOption {
	return func(mc *DirectMessageCreate) {
		mc.Type = MessageTypeCard
	}
}

// ReplyOption is the type providing additional options to Message event reply.
type ReplyOption string

const (
	// ReplyOptionTemp let reply temporary.
	ReplyOptionTemp ReplyOption = "reply_option_temp"
)
