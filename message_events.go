package khl

// EventTextMessage represents the payload of event '文字消息'
//
// 1
type EventTextMessage struct {
	Type         MessageType `json:"type"`
	GuildID      string      `json:"guild_id"`
	ChannelName  string      `json:"channel_name"`
	Mention      []string    `json:"mention"`
	MentionAll   bool        `json:"mention_all"`
	MentionRoles []string    `json:"mention_roles"`
	MentionHere  bool        `json:"mention_here"`
	Author       User        `json:"author"`
}

// EventImageMessage represents the payload of event '图片消息'
//
// 2
type EventImageMessage struct {
	MessageWithAttachment
}

// EventVideoMessage represents the payload of event '视频消息'
//
// 3
type EventVideoMessage struct {
	MessageWithAttachment
}

// EventFileMessage represents the payload of event '文件消息'
//
// 4
type EventFileMessage struct {
	MessageWithAttachment
}

// EventKMarkdownMessage represents the payload of event 'KMarkdown消息'
//
// 9
type EventKMarkdownMessage struct {
	EventTextMessage
	NavChannels []string `json:"nav_channels"`
	Code        string   `json:"code"`
	KMarkdown   struct {
		RawContent      string   `json:"raw_content"`
		MentionPart     []string `json:"mention_part"`
		MentionRolePart []string `json:"mention_role_part"`
	} `json:"kmarkdown"`
}

// EventCardMessage represents the payload of event 'Card消息'
//
// 10
type EventCardMessage struct {
	EventKMarkdownMessage
}
