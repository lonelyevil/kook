package khl

import (
	"net/url"
	"path"
)

// All endpoints for http request
var (
	APIVersion = "v3"

	EndpointBase          = "https://www.kaiheila.cn/api"
	EndpointAPI           = urlJoin(EndpointBase, APIVersion)
	EndpointGuild         = urlJoin(EndpointAPI, "guild")
	EndpointGuildMute     = urlJoin(EndpointAPI, "guild-mute")
	EndpointChannel       = urlJoin(EndpointAPI, "channel")
	EndpointChannelRole   = urlJoin(EndpointAPI, "channel-role")
	EndpointMessage       = urlJoin(EndpointAPI, "message")
	EndpointUserChat      = urlJoin(EndpointAPI, "user-chat")
	EndpointDirectMessage = urlJoin(EndpointAPI, "direct-message")
	EndpointGateway       = urlJoin(EndpointAPI, "gateway")
	EndpointUser          = urlJoin(EndpointAPI, "user")
	EndpointAsset         = urlJoin(EndpointAPI, "asset")
	EndpointGuildRole     = urlJoin(EndpointAPI, "guild-role")
	EndpointIntimacy      = urlJoin(EndpointAPI, "intimacy")
	EndpointGuildEmoji    = urlJoin(EndpointAPI, "guild-emoji")
	EndpointInvite        = urlJoin(EndpointAPI, "invite")

	EndpointGuildList       = urlJoin(EndpointGuild, "list")
	EndpointGuildView       = urlJoin(EndpointGuild, "view")
	EndpointGuildUserList   = urlJoin(EndpointGuild, "user-list")
	EndpointGuildNickName   = urlJoin(EndpointGuild, "nickname")
	EndpointGuildLeave      = urlJoin(EndpointGuild, "leave")
	EndpointGuildKickout    = urlJoin(EndpointGuild, "kickout")
	EndpointGuildMuteList   = urlJoin(EndpointGuildMute, "list")
	EndpointGuildMuteCreate = urlJoin(EndpointGuildMute, "create")
	EndpointGuildMuteDelete = urlJoin(EndpointGuildMute, "delete")

	// EndpointChannelMessage is Deprecated.
	EndpointChannelMessage    = urlJoin(EndpointChannel, "message")
	EndpointChannelList       = urlJoin(EndpointChannel, "list")
	EndpointChannelView       = urlJoin(EndpointChannel, "view")
	EndpointChannelCreate     = urlJoin(EndpointChannel, "create")
	EndpointChannelMoveUser   = urlJoin(EndpointChannel, "move-user")
	EndpointChannelDelete     = urlJoin(EndpointChannel, "delete")
	EndpointChannelRoleIndex  = urlJoin(EndpointChannelRole, "index")
	EndpointChannelRoleCreate = urlJoin(EndpointChannelRole, "create")
	EndpointChannelRoleUpdate = urlJoin(EndpointChannelRole, "update")
	EndpointChannelRoleDelete = urlJoin(EndpointChannelRole, "delete")

	EndpointMessageList           = urlJoin(EndpointMessage, "list")
	EndpointMessageCreate         = urlJoin(EndpointMessage, "create")
	EndpointMessageUpdate         = urlJoin(EndpointMessage, "update")
	EndpointMessageDelete         = urlJoin(EndpointMessage, "delete")
	EndpointMessageReactionList   = urlJoin(EndpointMessage, "reaction-list")
	EndpointMessageAddReaction    = urlJoin(EndpointMessage, "add-reaction")
	EndpointMessageDeleteReaction = urlJoin(EndpointMessage, "delete-reaction")

	EndpointUserChatList   = urlJoin(EndpointUserChat, "list")
	EndpointUserChatView   = urlJoin(EndpointUserChat, "view")
	EndpointUserChatCreate = urlJoin(EndpointUserChat, "create")
	EndpointUserChatDelete = urlJoin(EndpointUserChat, "delete")

	EndpointDirectMessageList           = urlJoin(EndpointDirectMessage, "list")
	EndpointDirectMessageCreate         = urlJoin(EndpointDirectMessage, "create")
	EndpointDirectMessageUpdate         = urlJoin(EndpointDirectMessage, "update")
	EndpointDirectMessageDelete         = urlJoin(EndpointDirectMessage, "delete")
	EndpointDirectMessageReactionList   = urlJoin(EndpointDirectMessage, "reaction-list")
	EndpointDirectMessageAddReaction    = urlJoin(EndpointDirectMessage, "add-reaction")
	EndpointDirectMessageDeleteReaction = urlJoin(EndpointDirectMessage, "delete-reaction")

	EndpointGatewayIndex = urlJoin(EndpointGateway, "index")

	EndpointUserMe      = urlJoin(EndpointUser, "me")
	EndpointUserView    = urlJoin(EndpointUser, "view")
	EndpointUserOffline = urlJoin(EndpointUser, "offline")

	EndpointAssetCreate = urlJoin(EndpointAsset, "create")

	EndpointGuildRoleList   = urlJoin(EndpointGuildRole, "list")
	EndpointGuildRoleCreate = urlJoin(EndpointGuildRole, "create")
	EndpointGuildRoleUpdate = urlJoin(EndpointGuildRole, "update")
	EndpointGuildRoleDelete = urlJoin(EndpointGuildRole, "delete")
	EndpointGuildRoleGrant  = urlJoin(EndpointGuildRole, "grant")
	EndpointGuildRoleRevoke = urlJoin(EndpointGuildRole, "revoke")

	EndpointIntimacyIndex  = urlJoin(EndpointIntimacy, "index")
	EndpointIntimacyUpdate = urlJoin(EndpointIntimacy, "update")

	EndpointGuildEmojiList   = urlJoin(EndpointGuildEmoji, "list")
	EndpointGuildEmojiCreate = urlJoin(EndpointGuildEmoji, "create")
	EndpointGuildEmojiUpdate = urlJoin(EndpointGuildEmoji, "update")
	EndpointGuildEmojiDelete = urlJoin(EndpointGuildEmoji, "delete")

	EndpointInviteList   = urlJoin(EndpointInvite, "list")
	EndpointInviteCreate = urlJoin(EndpointInvite, "create")
	EndpointInviteDelete = urlJoin(EndpointInvite, "delete")
)

// Must not be used elsewhere.
func urlJoin(u1, u2 string) string {
	r, err := url.Parse(u1)
	if err != nil {
		panic(err)
	}
	r.Path = path.Join(r.Path, u2)
	return r.String()
}
