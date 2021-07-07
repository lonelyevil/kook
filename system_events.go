//go:generate go run tools/cmd/eventhandler/eventhandler_gen.go

package khl

// ChannelMessageReactionAdd represents the event '频道内用户添加 reaction'
//
// added_reaction
type ChannelMessageReactionAdd struct {
	*MessageReaction
}

// ChannelMessageReactionRemove represents the event '频道内用户取消 reaction'
//
// deleted_reaction
type ChannelMessageReactionRemove struct {
	*MessageReaction
}

// ChannelMessageUpdate represents the event '频道消息更新'
//
// updated_message
type ChannelMessageUpdate struct {
	*ChannelMessage
}

// ChannelMessageRemove represents the event '频道消息被删除'
//
// deleted_message
type ChannelMessageRemove struct {
	*ChannelMessage
}

// ChannelAdd represents the event '新增频道'
//
// added_channel
type ChannelAdd struct {
	*Channel
}

// ChannelUpdate represents the event '修改频道信息'
//
// updated_channel
type ChannelUpdate struct {
	*Channel
}

// ChannelRemove represents the event '删除频道'
//
// deleted_channel
type ChannelRemove struct {
	ID        string         `json:"id"`
	DeletedAt MilliTimeStamp `json:"deleted_at"`
}

// ChannelStickyMessageAdd represents the event '新的频道置顶消息'
//
// pinned_message
type ChannelStickyMessageAdd struct {
	ChannelID  string `json:"channel_id"`
	OperatorID string `json:"operator_id"`
	MsgID      string `json:"msg_id"`
}

// ChannelStickyMessageRemove represents the event '取消频道置顶消息'
//
// unpinned_message
type ChannelStickyMessageRemove struct {
	ChannelID  string `json:"channel_id"`
	OperatorID string `json:"operator_id"`
	MsgID      string `json:"msg_id"`
}

// PrivateMessageUpdate represents the event '私聊消息更新'
//
// updated_private_message
type PrivateMessageUpdate struct {
	*PrivateMessage
}

// PrivateMessageRemove represents the event '私聊消息被删除'
//
// deleted_private_message
type PrivateMessageRemove struct {
	*PrivateMessage
}

// PrivateMessageReactionAdd represents the event '私聊内用户添加 reaction'
//
// private_added_reaction
type PrivateMessageReactionAdd struct {
	*MessageReaction
}

// PrivateMessageReactionRemove represents the event '私聊内用户取消 reaction'
//
// private_deleted_reaction
type PrivateMessageReactionRemove struct {
	*MessageReaction
}

// GuildMemberAdd represents the event '新成员加入服务器'
//
// joined_guild
type GuildMemberAdd struct {
	UserID   string         `json:"user_id"`
	JoinedAt MilliTimeStamp `json:"joined_at"`
}

// GuildMemberRemove represents the event '服务器成员退出'
//
// joined_guild
type GuildMemberRemove struct {
	UserID   string         `json:"user_id"`
	ExitedAt MilliTimeStamp `json:"exited_at"`
}

// GuildMemberUpdate represents the event '服务器成员信息更新'
//
// updated_guild_member
type GuildMemberUpdate struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
}

// GuildMemberOnline represents the event '服务器成员上线'
//
// guild_member_online
type GuildMemberOnline struct {
	UserID    string         `json:"user_id"`
	EventTime MilliTimeStamp `json:"event_time"`
	Guilds    []string       `json:"guilds"`
}

// GuildMemberOffline represents the event '服务器成员下线'
//
// guild_member_offline
type GuildMemberOffline struct {
	UserID    string         `json:"user_id"`
	EventTime MilliTimeStamp `json:"event_time"`
	Guilds    []string       `json:"guilds"`
}

// GuildRoleAdd represents the event '服务器角色增加'
//
// added_role
type GuildRoleAdd struct {
	*Role
}

// GuildRoleRemove represents the event '服务器角色删除'
//
// deleted_role
type GuildRoleRemove struct {
	*Role
}

// GuildRoleUpdate represents the event '服务器角色更新'
//
// updated_role
type GuildRoleUpdate struct {
	*Role
}

// GuildUpdate represents the event '服务器信息更新'
//
// updated_guild
type GuildUpdate struct {
	*Guild
}

// GuildDelete represents the event '服务器删除'
//
// deleted_guild
type GuildDelete struct {
	*Guild
}

// GuildBlocklistAdd represents the event '服务器封禁用户'
//
// added_block_list
type GuildBlocklistAdd struct {
	OperatorID string   `json:"operator_id"`
	Remark     string   `json:"remark"`
	UserID     []string `json:"user_id"`
}

// GuildBlocklistRemove represents the event '服务器取消封禁用户'
//
// deleted_block_list
type GuildBlocklistRemove struct {
	OperatorID string   `json:"operator_id"`
	UserID     []string `json:"user_id"`
}

// VoiceChannelMemberAdd represents the event '用户加入语音频道'
//
// joined_channel
type VoiceChannelMemberAdd struct {
	UserID    string         `json:"user_id"`
	ChannelID string         `json:"channel_id"`
	JoinedAt  MilliTimeStamp `json:"joined_at"`
}

// VoiceChannelMemberRemove represents the event '用户退出语音频道'
//
// exited_channel
type VoiceChannelMemberRemove struct {
	UserID    string         `json:"user_id"`
	ChannelID string         `json:"channel_id"`
	ExitedAt  MilliTimeStamp `json:"exited_at"`
}

// UserUpdate represents the event '用户信息更新'
//
// user_updated
type UserUpdate struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// BotGuildAdd represents the event '自己新加入服务器'
//
// self_joined_guild
type BotGuildAdd struct {
	GuildID string `json:"guild_id"`
}

// BotGuildRemove represents the event '自己退出服务器'
//
// self_exited_guild
type BotGuildRemove struct {
	GuildID string `json:"guild_id"`
}

// CardButtonClick represents the event 'Card消息中的Button点击事件'
//
// message_btn_click
type CardButtonClick struct {
	MsgID    string `json:"msg_id"`
	UserID   string `json:"user_id"`
	Value    string `json:"value"`
	TargetID string `json:"target_id"`
}
