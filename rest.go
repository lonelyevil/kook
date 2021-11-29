package khl

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

// Gateway returns the url for websocket gateway.
// FYI: https://developer.kaiheila.cn/doc/http/gateway#%E8%8E%B7%E5%8F%96%E7%BD%91%E5%85%B3%E8%BF%9E%E6%8E%A5%E5%9C%B0%E5%9D%80
func (s *Session) Gateway() (gateway string, err error) {
	u, _ := url.Parse(EndpointGatewayIndex)
	q := u.Query()
	q.Set("compress", "0")
	if s.Identify.Compress {
		q.Set("compress", "1")
	}
	u.RawQuery = q.Encode()
	response, err := s.Request("GET", u.String(), nil)
	if err != nil {
		return
	}

	temp := struct {
		URL string `json:"url"`
	}{}

	err = json.Unmarshal(response, &temp)
	if err != nil {
		return
	}
	gateway = temp.URL
	return
}

// MessageListOption is the type for optional arguments for MessageList request.
type MessageListOption func(values url.Values)

// MessageListWithMsgID adds optional `msg_id` argument to MessageList request.
func MessageListWithMsgID(msgID string) MessageListOption {
	return func(values url.Values) {
		values.Set("msg_id", msgID)
	}
}

// MessageListWithPin adds optional `pin` argument to MessageList request.
func MessageListWithPin(pin bool) MessageListOption {
	return func(values url.Values) {
		if pin {
			values.Set("pin", "1")
		} else {
			values.Set("pin", "0")
		}
	}
}

// MessageListFlag is the type for the flag of MessageList.
type MessageListFlag string

// These are the usable flags
const (
	MessageListFlagBefore MessageListFlag = "before"
	MessageListFlagAround MessageListFlag = "around"
	MessageListFlagAfter  MessageListFlag = "after"
)

// MessageListWithFlag adds optional `flag` argument to MessageList request.
func MessageListWithFlag(flag MessageListFlag) MessageListOption {
	return func(values url.Values) {
		values.Set("flag", string(flag))
	}
}

// MessageList returns a list of messages of a channel.
// FYI: https://developer.kaiheila.cn/doc/http/message#%E8%8E%B7%E5%8F%96%E9%A2%91%E9%81%93%E8%81%8A%E5%A4%A9%E6%B6%88%E6%81%AF%E5%88%97%E8%A1%A8
func (s *Session) MessageList(targetID string, options ...MessageListOption) (ms []*DetailedChannelMessage, err error) {
	var response []byte
	u, _ := url.Parse(EndpointMessageList)
	q := u.Query()
	q.Set("target_id", targetID)
	for _, item := range options {
		item(q)
	}
	u.RawQuery = q.Encode()
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	data := struct {
		Items []*DetailedChannelMessage `json:"items"`
	}{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return nil, err
	}
	ms = data.Items

	return ms, nil
}

// MessageCreateBase is the common arguments for message creation.
type MessageCreateBase struct {
	Type     MessageType `json:"type,omitempty"`
	TargetID string      `json:"target_id,omitempty"`
	Content  string      `json:"content,omitempty"`
	Quote    string      `json:"quote,omitempty"`
	Nonce    string      `json:"nonce,omitempty"`
}

// MessageCreate is the type for message creation arguments.
type MessageCreate struct {
	MessageCreateBase
	TempTargetID string `json:"temp_target_id,omitempty"`
}

// MessageResp is the type for response for MessageCreate.
type MessageResp struct {
	MsgID        string         `json:"msg_id"`
	MsgTimestamp MilliTimeStamp `json:"msg_timestamp"`
	Nonce        string         `json:"nonce"`
}

// MessageCreate creates a message.
// FYI: https://developer.kaiheila.cn/doc/http/message#%E5%8F%91%E9%80%81%E9%A2%91%E9%81%93%E8%81%8A%E5%A4%A9%E6%B6%88%E6%81%AF
func (s *Session) MessageCreate(m *MessageCreate) (resp *MessageResp, err error) {
	var response []byte
	response, err = s.Request("POST", EndpointMessageCreate, m)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// MessageUpdateBase is the shared arguments for message update related requests.
type MessageUpdateBase struct {
	MsgID   string `json:"msg_id"`
	Content string `json:"content"`
	Quote   string `json:"quote,omitempty"`
}

// MessageUpdate is the request data for MessageUpdate.
type MessageUpdate struct {
	MessageUpdateBase
	TempTargetID string `json:"temp_target_id,omitempty"`
}

// MessageUpdate updates a message.
// FYI: https://developer.kaiheila.cn/doc/http/message#%E6%9B%B4%E6%96%B0%E9%A2%91%E9%81%93%E8%81%8A%E5%A4%A9%E6%B6%88%E6%81%AF
func (s *Session) MessageUpdate(m *MessageUpdate) (err error) {
	_, err = s.Request("POST", EndpointMessageUpdate, m)
	return
}

// MessageDelete deletes a message.
// FYI: https://developer.kaiheila.cn/doc/http/message#%E5%88%A0%E9%99%A4%E9%A2%91%E9%81%93%E8%81%8A%E5%A4%A9%E6%B6%88%E6%81%AF
func (s *Session) MessageDelete(msgID string) (err error) {
	_, err = s.Request("POST", EndpointMessageDelete, struct {
		MsgID string `json:"msg_id"`
	}{msgID})
	return
}

// ReactedUser is the type for every user reacted to a specific message with a specific emoji.
type ReactedUser struct {
	User
	ReactionTime MilliTimeStamp `json:"reaction_time"`
	TagInfo      struct {
		Color string `json:"color"`
		Text  string `json:"text"`
	} `json:"tag_info"`
}

// MessageReactionList returns the list of the reacted users with a specific emoji to a message.
// FYI: https://developer.kaiheila.cn/doc/http/message#%E8%8E%B7%E5%8F%96%E9%A2%91%E9%81%93%E6%B6%88%E6%81%AF%E6%9F%90%E5%9B%9E%E5%BA%94%E7%9A%84%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8
func (s *Session) MessageReactionList(msgID, emoji string) (us []*ReactedUser, err error) {
	u, _ := url.Parse(EndpointMessageReactionList)
	q := u.Query()
	q.Add("msg_id", msgID)
	q.Add("emoji", emoji)
	u.RawQuery = q.Encode()
	var response []byte
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &us)
	if err != nil {
		return nil, err
	}
	return us, nil
}

// MessageAddReaction add a reaction to a message as the bot.
// FYI: https://developer.kaiheila.cn/doc/http/message#%E7%BB%99%E6%9F%90%E4%B8%AA%E6%B6%88%E6%81%AF%E6%B7%BB%E5%8A%A0%E5%9B%9E%E5%BA%94
func (s *Session) MessageAddReaction(msgID, emoji string) (err error) {
	_, err = s.Request("POST", EndpointMessageAddReaction, struct {
		MsgID string `json:"msg_id"`
		Emoji string `json:"emoji"`
	}{msgID, emoji})
	return err
}

// MessageDeleteReaction deletes a reaction of a user from a message.
// FYI: https://developer.kaiheila.cn/doc/http/message#%E5%88%A0%E9%99%A4%E6%B6%88%E6%81%AF%E7%9A%84%E6%9F%90%E4%B8%AA%E5%9B%9E%E5%BA%94
func (s *Session) MessageDeleteReaction(msgID, emoji string, userID string) (err error) {
	_, err = s.Request("POST", EndpointMessageDeleteReaction, struct {
		MsgID  string `json:"msg_id"`
		Emoji  string `json:"emoji"`
		UserID string `json:"user_id,omitempty"`
	}{msgID, emoji, userID})
	return err
}

// ChannelList lists all channels from a guild.
// FYI: https://developer.kaiheila.cn/doc/http/channel#%E8%8E%B7%E5%8F%96%E9%A2%91%E9%81%93%E5%88%97%E8%A1%A8
func (s *Session) ChannelList(guildID string, page *PageSetting) (cs []*Channel, meta *PageInfo, err error) {
	var response []byte
	u, _ := url.Parse(EndpointChannelList)
	q := u.Query()
	q.Set("guild_id", guildID)
	u.RawQuery = q.Encode()
	response, meta, err = s.RequestWithPage("GET", u.String(), page)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(response, &cs)
	if err != nil {
		return nil, nil, err
	}
	return cs, meta, nil
}

// ChannelView returns the detailed information for a channel.
// FYI: https://developer.kaiheila.cn/doc/http/channel#%E8%8E%B7%E5%8F%96%E9%A2%91%E9%81%93%E8%AF%A6%E6%83%85
func (s *Session) ChannelView(channelID string) (c *Channel, err error) {
	var response []byte
	u, _ := url.Parse(EndpointChannelView)
	q := u.Query()
	q.Set("target_id", channelID)
	u.RawQuery = q.Encode()
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// ChannelCreate is the arguments for creating a channel.
type ChannelCreate struct {
	GuildID      string      `json:"guild_id"`
	ParentID     string      `json:"parent_id,omitempty"`
	Name         string      `json:"name"`
	Type         ChannelType `json:"type,omitempty"`
	LimitAmount  int         `json:"limit_amount,omitempty"`
	VoiceQuality int         `json:"voice_quality,omitempty"`
}

// ChannelCreate creates a channel.
// FYI: https://developer.kaiheila.cn/doc/http/channel#%E5%88%9B%E5%BB%BA%E9%A2%91%E9%81%93
func (s *Session) ChannelCreate(cc *ChannelCreate) (c *Channel, err error) {
	var response []byte
	response, err = s.Request("POST", EndpointChannelCreate, cc)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &c)
	if err != nil {
		return nil, err
	}
	return c, err
}

// ChannelDelete deletes a channel.
// FYI: https://developer.kaiheila.cn/doc/http/channel#%E5%88%A0%E9%99%A4%E9%A2%91%E9%81%93
func (s *Session) ChannelDelete(channelID string) (err error) {
	_, err = s.Request("POST", EndpointChannelDelete, struct {
		ChannelID string `json:"channel_id"`
	}{channelID})
	return err
}

// ChannelMoveUsers moves users to a channel.
// FYI: https://developer.kaiheila.cn/doc/http/channel#%E8%AF%AD%E9%9F%B3%E9%A2%91%E9%81%93%E4%B9%8B%E9%97%B4%E7%A7%BB%E5%8A%A8%E7%94%A8%E6%88%B7
func (s *Session) ChannelMoveUsers(targetChannelID string, userIDs []string) (err error) {
	_, err = s.Request("POST", EndpointChannelMoveUser, struct {
		TargetID string   `json:"target_id"`
		UserIDs  []string `json:"user_ids"`
	}{targetChannelID, userIDs})
	return err
}

// ChannelRoleIndex is the role and permission list of a channel.
type ChannelRoleIndex struct {
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	PermissionUsers      []struct {
		User  User           `json:"user"`
		Allow RolePermission `json:"allow"`
		Deny  RolePermission `json:"deny"`
	} `json:"permission_users"`
	PermissionSync IntBool `json:"permission_sync"`
}

// ChannelRoleIndex returns the role and permission list of the channel.
// FYI: https://developer.kaiheila.cn/doc/http/channel#%E9%A2%91%E9%81%93%E8%A7%92%E8%89%B2%E6%9D%83%E9%99%90%E8%AF%A6%E6%83%85
func (s *Session) ChannelRoleIndex(channelID string) (cr *ChannelRoleIndex, err error) {
	var response []byte
	u, _ := url.Parse(EndpointChannelRoleIndex)
	q := u.Query()
	q.Set("channel_id", channelID)
	u.RawQuery = q.Encode()
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &cr)
	if err != nil {
		return nil, err
	}
	return cr, err
}

// ChannelRoleBase is the common arguments for channel role requests.
type ChannelRoleBase struct {
	ChannelID string `json:"channel_id"`
	Type      string `json:"type,omitempty"`
	Value     string `json:"value,omitempty"`
}

// ChannelRoleCreate is the request query data for ChannelRoleCreate.
type ChannelRoleCreate ChannelRoleBase

// ChannelRoleCreateResp is the response for ChannelRoleCreate.
type ChannelRoleCreateResp ChannelRoleUpdateResp

// ChannelRoleCreate creates a role for a channel
// FYI: https://developer.kaiheila.cn/doc/http/channel#%E5%88%9B%E5%BB%BA%E9%A2%91%E9%81%93%E8%A7%92%E8%89%B2%E6%9D%83%E9%99%90
func (s *Session) ChannelRoleCreate(crc *ChannelRoleCreate) (crcr *ChannelRoleCreateResp, err error) {
	var resp []byte
	resp, err = s.Request("POST", EndpointChannelRoleCreate, crc)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &crcr)
	if err != nil {
		return nil, err
	}
	return crcr, err
}

// ChannelRoleUpdate is the request query data for ChannelRoleUpdate
type ChannelRoleUpdate struct {
	ChannelRoleBase
	Allow RolePermission `json:"allow,omitempty"`
	Deny  RolePermission `json:"deny,omitempty"`
}

// ChannelRoleUpdateResp is the response of ChannelRoleUpdate
type ChannelRoleUpdateResp struct {
	UserID string         `json:"user_id"`
	RoleID string         `json:"role_id"`
	Allow  RolePermission `json:"allow"`
	Deny   RolePermission `json:"deny"`
}

// ChannelRoleUpdate updates a role from channel setting.
// FYI: https://developer.kaiheila.cn/doc/http/channel#%E6%9B%B4%E6%96%B0%E9%A2%91%E9%81%93%E8%A7%92%E8%89%B2%E6%9D%83%E9%99%90
func (s *Session) ChannelRoleUpdate(cru *ChannelRoleUpdate) (crur *ChannelRoleUpdateResp, err error) {
	var response []byte
	response, err = s.Request("POST", EndpointChannelRoleUpdate, cru)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &crur)
	if err != nil {
		return nil, err
	}
	return crur, nil
}

// ChannelRoleDelete is the type for settings when deleting a role from channel setting.
type ChannelRoleDelete ChannelRoleBase

// ChannelRoleDelete deletes a role form channel setting.
// FYI: https://developer.kaiheila.cn/doc/http/channel#%E5%88%A0%E9%99%A4%E9%A2%91%E9%81%93%E8%A7%92%E8%89%B2%E6%9D%83%E9%99%90
func (s *Session) ChannelRoleDelete(crd *ChannelRoleDelete) (err error) {
	_, err = s.Request("POST", EndpointChannelRoleDelete, crd)
	return err
}

// UserChatList returns a list of user chats that bot owns.
//
// Note: for User in TargetInfo, only ID, Username, Online, Avatar is filled
//
// FYI: https://developer.kaiheila.cn/doc/http/user-chat#%E8%8E%B7%E5%8F%96%E7%A7%81%E4%BF%A1%E8%81%8A%E5%A4%A9%E4%BC%9A%E8%AF%9D%E5%88%97%E8%A1%A8
func (s *Session) UserChatList(page *PageSetting) (ucs []*UserChat, meta *PageInfo, err error) {
	var response []byte
	response, meta, err = s.RequestWithPage("GET", EndpointUserChatList, page)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(response, &ucs)
	if err != nil {
		return nil, nil, err
	}
	return ucs, meta, err
}

// UserChatView returns a detailed user chat.
//
// FYI: https://developer.kaiheila.cn/doc/http/user-chat#%E8%8E%B7%E5%8F%96%E7%A7%81%E4%BF%A1%E8%81%8A%E5%A4%A9%E4%BC%9A%E8%AF%9D%E8%AF%A6%E6%83%85
func (s *Session) UserChatView(chatCode string) (uc *UserChat, err error) {
	var response []byte
	u, _ := url.Parse(EndpointUserChatView)
	q := u.Query()
	q.Set("chat_code", chatCode)
	u.RawQuery = q.Encode()
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	uc = &UserChat{}
	err = json.Unmarshal(response, uc)
	if err != nil {
		return nil, err
	}
	return uc, nil
}

// UserChatCreate creates a direct chat session.
// FYI: https://developer.kaiheila.cn/doc/http/user-chat#%E5%88%9B%E5%BB%BA%E7%A7%81%E4%BF%A1%E8%81%8A%E5%A4%A9%E4%BC%9A%E8%AF%9D
func (s *Session) UserChatCreate(UserID string) (uc *UserChat, err error) {
	var response []byte
	response, err = s.Request("POST", EndpointUserChatCreate, struct {
		TargetID string `json:"target_id"`
	}{UserID})
	if err != nil {
		return nil, err
	}
	uc = &UserChat{}
	err = json.Unmarshal(response, uc)
	if err != nil {
		return nil, err
	}
	return uc, err
}

// UserChatDelete deletes a direct chat session.
// FYI: https://developer.kaiheila.cn/doc/http/user-chat#%E5%88%9B%E5%BB%BA%E7%A7%81%E4%BF%A1%E8%81%8A%E5%A4%A9%E4%BC%9A%E8%AF%9D
func (s *Session) UserChatDelete(ChatCode string) (err error) {
	_, err = s.Request("POST", EndpointUserChatDelete, struct {
		ChatCode string `json:"chat_code"`
	}{ChatCode: ChatCode})
	return err
}

// DirectMessageListOption is the type for optional arguments for DirectMessageList request.
type DirectMessageListOption func(values url.Values)

// DirectMessageListWithChatCode adds optional `chat_code` argument to DirectMessageList request.
func DirectMessageListWithChatCode(chatCode string) DirectMessageListOption {
	return func(values url.Values) {
		values.Set("chat_code", chatCode)
	}
}

// DirectMessageListWithTargetID adds optional `target_id` argument to DirectMessageList request.
func DirectMessageListWithTargetID(targetID string) DirectMessageListOption {
	return func(values url.Values) {
		values.Set("target_id", targetID)
	}
}

// DirectMessageListWithMsgID adds optional `msg_id` argument to DirectMessageList request.
func DirectMessageListWithMsgID(msgID string) DirectMessageListOption {
	return func(values url.Values) {
		values.Set("msg_id", msgID)
	}
}

// DirectMessageListWithFlag adds optional `flag` argument to DirectMessageList request.
func DirectMessageListWithFlag(flag MessageListFlag) DirectMessageListOption {
	return func(values url.Values) {
		values.Set("flag", string(flag))
	}
}

// DirectMessageResp is the type for direct messages.
type DirectMessageResp struct {
	ID          string              `json:"id"`
	Type        MessageType         `json:"type"`
	Content     string              `json:"content"`
	Embeds      []map[string]string `json:"embeds"`
	Attachments []Attachment        `json:"attachments"`
	CreateAt    MilliTimeStamp      `json:"create_at"`
	UpdatedAt   MilliTimeStamp      `json:"updated_at"`
	Reactions   []ReactionItem      `json:"reactions"`
	ImageName   string              `json:"image_name"`
	ReadStatus  bool                `json:"read_status"`
	Quote       *User               `json:"quote"`
	MentionInfo struct {
		MentionPart     []*User `json:"mention_part"`
		MentionRolePart []*Role `json:"mention_role_part"`
	} `json:"mention_info"`
}

// DirectMessageList returns the messages in direct chat.
//
// FYI: https://developer.kaiheila.cn/doc/http/direct-message#%E8%8E%B7%E5%8F%96%E7%A7%81%E4%BF%A1%E8%81%8A%E5%A4%A9%E6%B6%88%E6%81%AF%E5%88%97%E8%A1%A8
func (s *Session) DirectMessageList(options ...DirectMessageListOption) (dmrs []*DirectMessageResp, err error) {
	var response []byte
	u, _ := url.Parse(EndpointDirectMessageList)
	q := u.Query()
	for _, item := range options {
		item(q)
	}
	u.RawQuery = q.Encode()
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &dmrs)
	if err != nil {
		return nil, err
	}
	return dmrs, nil
}

// DirectMessageCreate is the struct for settings of creating a message in direct chat.
type DirectMessageCreate struct {
	MessageCreateBase
	ChatCode string `json:"chat_code,omitempty"`
}

// DirectMessageCreate creates a message in direct chat.
// FYI: https://developer.kaiheila.cn/doc/http/direct-message#%E5%8F%91%E9%80%81%E7%A7%81%E4%BF%A1%E8%81%8A%E5%A4%A9%E6%B6%88%E6%81%AF
func (s *Session) DirectMessageCreate(create *DirectMessageCreate) (mr *MessageResp, err error) {
	var response []byte
	response, err = s.Request("POST", EndpointDirectMessageCreate, create)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &mr)
	if err != nil {
		return nil, err
	}
	return mr, nil
}

// DirectMessageUpdate is the type for settings of updating a message in direct chat.
type DirectMessageUpdate MessageUpdateBase

// DirectMessageUpdate updates a message in direct chat.
// FYI: https://developer.kaiheila.cn/doc/http/direct-message#%E6%9B%B4%E6%96%B0%E7%A7%81%E4%BF%A1%E8%81%8A%E5%A4%A9%E6%B6%88%E6%81%AF
func (s *Session) DirectMessageUpdate(update *DirectMessageUpdate) (err error) {
	_, err = s.Request("POST", EndpointDirectMessageUpdate, update)
	return err
}

// DirectMessageDelete deletes a message in direct chat.
// FYI: https://developer.kaiheila.cn/doc/http/direct-message#%E5%88%A0%E9%99%A4%E7%A7%81%E4%BF%A1%E8%81%8A%E5%A4%A9%E6%B6%88%E6%81%AF
func (s *Session) DirectMessageDelete(msgID string) (err error) {
	_, err = s.Request("POST", EndpointDirectMessageDelete, struct {
		MsgID string `json:"msg_id"`
	}{msgID})
	return err
}

// DirectMessageReactionList returns the list of the reacted users with a specific emoji to a message.
//
// FYI: https://developer.kaiheila.cn/doc/http/direct-message#%E8%8E%B7%E5%8F%96%E9%A2%91%E9%81%93%E6%B6%88%E6%81%AF%E6%9F%90%E5%9B%9E%E5%BA%94%E7%9A%84%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8
func (s *Session) DirectMessageReactionList(msgID, emoji string) (us []*ReactedUser, err error) {
	u, _ := url.Parse(EndpointDirectMessageReactionList)
	q := u.Query()
	q.Add("msg_id", msgID)
	q.Add("emoji", emoji)
	u.RawQuery = q.Encode()
	var response []byte
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &us)
	if err != nil {
		return nil, err
	}
	return us, nil
}

// DirectMessageAddReaction add a reaction to a message as the bot.
//
// FYI: https://developer.kaiheila.cn/doc/http/direct-message#%E7%BB%99%E6%9F%90%E4%B8%AA%E6%B6%88%E6%81%AF%E6%B7%BB%E5%8A%A0%E5%9B%9E%E5%BA%94
func (s *Session) DirectMessageAddReaction(msgID, emoji string) (err error) {
	_, err = s.Request("POST", EndpointDirectMessageAddReaction, struct {
		MsgID string `json:"msg_id"`
		Emoji string `json:"emoji"`
	}{msgID, emoji})
	return err
}

// AssetCreate uploads attachments to khl server.
//
// FYI: https://developer.kaiheila.cn/doc/http/asset#%E4%B8%8A%E4%BC%A0%E6%96%87%E4%BB%B6/%E5%9B%BE%E7%89%87
func (s *Session) AssetCreate(name string, file []byte) (url string, err error) {
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	var fw io.Writer
	fw, err = w.CreateFormFile("file", name)
	if err != nil {
		return "", err
	}
	_, err = fw.Write(file)
	if err != nil {
		return "", err
	}
	err = w.Close()
	if err != nil {
		return "", err
	}
	var f assetFile
	f.Payload = b.Bytes()
	f.ContentType = w.FormDataContentType()
	var response []byte
	response, err = s.Request("POST", EndpointAssetCreate, &f)
	if err != nil {
		return "", err
	}
	urlStruct := struct {
		URL string `json:"url"`
	}{}
	err = json.Unmarshal(response, &urlStruct)
	if err != nil {
		return "", err
	}
	return urlStruct.URL, nil
}

// DirectMessageDeleteReaction deletes a reaction of a user from a message.
//
// FYI: https://developer.kaiheila.cn/doc/http/direct-message#%E5%88%A0%E9%99%A4%E6%B6%88%E6%81%AF%E7%9A%84%E6%9F%90%E4%B8%AA%E5%9B%9E%E5%BA%94
func (s *Session) DirectMessageDeleteReaction(msgID, emoji string) (err error) {
	_, err = s.Request("POST", EndpointDirectMessageDeleteReaction, struct {
		MsgID string `json:"msg_id"`
		Emoji string `json:"emoji"`
	}{msgID, emoji})
	return err
}

// GuildList returns a list of guild that bot joins.
// FYI: https://developer.kaiheila.cn/doc/http/guild#%E8%8E%B7%E5%8F%96%E5%BD%93%E5%89%8D%E7%94%A8%E6%88%B7%E5%8A%A0%E5%85%A5%E7%9A%84%E6%9C%8D%E5%8A%A1%E5%99%A8%E5%88%97%E8%A1%A8
func (s *Session) GuildList(page *PageSetting) (gs []*Guild, meta *PageInfo, err error) {
	var response []byte
	response, meta, err = s.RequestWithPage("GET", EndpointGuildList, page)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(response, &gs)
	if err != nil {
		return nil, nil, err
	}
	return gs, meta, nil
}

// GuildView returns a detailed info for a guild.
// FYI: https://developer.kaiheila.cn/doc/http/guild#%E8%8E%B7%E5%8F%96%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%AF%A6%E6%83%85
func (s *Session) GuildView(guildID string) (g *Guild, err error) {
	var response []byte
	u, _ := url.Parse(EndpointGuildView)
	q := u.Query()
	q.Add("guild_id", guildID)
	u.RawQuery = q.Encode()
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	g = &Guild{}
	err = json.Unmarshal(response, g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// GuildUserListOption is the type for optional arguments for GuildUserList request.
type GuildUserListOption func(values url.Values)

// GuildUserListWithChannelID adds optional `channel_id` argument to GuildUserList request.
func GuildUserListWithChannelID(id string) GuildUserListOption {
	return func(values url.Values) {
		values.Set("channel_id", id)
	}
}

// GuildUserListWithSearch adds optional `search` argument to GuildUserList request.
func GuildUserListWithSearch(search string) GuildUserListOption {
	return func(values url.Values) {
		values.Set("search", search)
	}
}

// GuildUserListWithRoleID adds optional `role_id` argument to GuildUserList request.
func GuildUserListWithRoleID(roleID int64) GuildUserListOption {
	return func(values url.Values) {
		values.Set("role_id", strconv.FormatInt(roleID, 10))
	}
}

// GuildUserListWithMobileVerified adds optional `mobile_verified` argument to GuildUserList request.
func GuildUserListWithMobileVerified(verified bool) GuildUserListOption {
	return func(values url.Values) {
		if verified {
			values.Set("mobile_verified", "1")
		} else {
			values.Set("mobile_verified", "0")
		}
	}
}

// GuildUserListWithActiveTime adds optional `active_time` argument to GuildUserList request.
func GuildUserListWithActiveTime(activeTime bool) GuildUserListOption {
	return func(values url.Values) {
		if activeTime {
			values.Set("active_time", "1")
		} else {
			values.Set("active_time", "0")
		}
	}
}

// GuildUserListWithJoinedAt adds optional `joined_at` argument to GuildUserList request.
func GuildUserListWithJoinedAt(joinedAt bool) GuildUserListOption {
	return func(values url.Values) {
		if joinedAt {
			values.Set("joined_at", "1")
		} else {
			values.Set("joined_at", "0")
		}
	}
}

// GuildUserList returns the list of users in a guild.
// FYI: https://developer.kaiheila.cn/doc/http/guild#%E8%8E%B7%E5%8F%96%E6%9C%8D%E5%8A%A1%E5%99%A8%E4%B8%AD%E7%9A%84%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8
func (s *Session) GuildUserList(guildID string, page *PageSetting, options ...GuildUserListOption) (us []*User, meta *PageInfo, err error) {
	var response []byte
	u, _ := url.Parse(EndpointGuildUserList)
	q := u.Query()
	q.Set("guild_id", guildID)
	for _, item := range options {
		item(q)
	}
	u.RawQuery = q.Encode()
	response, meta, err = s.RequestWithPage("GET", u.String(), page)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(response, &us)
	if err != nil {
		return nil, nil, err
	}
	return us, meta, err
}

// GuildNickname is the arguments for GuildNickname.
type GuildNickname struct {
	GuildID  string `json:"guild_id"`
	Nickname string `json:"nickname,omitempty"`
	UserID   string `json:"user_id,omitempty"`
}

// GuildNickname changes the nickname of a user in a guild.
// FYI: https://developer.kaiheila.cn/doc/http/guild#%E4%BF%AE%E6%94%B9%E6%9C%8D%E5%8A%A1%E5%99%A8%E4%B8%AD%E7%94%A8%E6%88%B7%E7%9A%84%E6%98%B5%E7%A7%B0
func (s *Session) GuildNickname(gn *GuildNickname) (err error) {
	_, err = s.Request("POST", EndpointGuildNickName, gn)
	return err
}

// GuildLeave let the bot leave a guild.
// FYI: https://developer.kaiheila.cn/doc/http/guild#%E7%A6%BB%E5%BC%80%E6%9C%8D%E5%8A%A1%E5%99%A8
func (s *Session) GuildLeave(guildID string) (err error) {
	_, err = s.Request("POST", EndpointGuildLeave, struct {
		GuildID string `json:"guild_id"`
	}{guildID})
	return err
}

// GuildKickout force a user to leave a guild.
// FYI: https://developer.kaiheila.cn/doc/http/guild#%E8%B8%A2%E5%87%BA%E6%9C%8D%E5%8A%A1%E5%99%A8
func (s *Session) GuildKickout(guildID, targetID string) (err error) {
	_, err = s.Request("POST", EndpointGuildKickout, struct {
		GuildID  string `json:"guild_id"`
		TargetID string `json:"target_id"`
	}{guildID, targetID})
	return err
}

// GuildMuteList is the type for users that got muted in a guild.
type GuildMuteList struct {
	Mic     []string `json:"1"`
	Headset []string `json:"2"`
}

// GuildMuteList returns the list of users got mutes in mic or earphone.
// FYI: https://developer.kaiheila.cn/doc/http/guild#%E6%9C%8D%E5%8A%A1%E5%99%A8%E9%9D%99%E9%9F%B3%E9%97%AD%E9%BA%A6%E5%88%97%E8%A1%A8
func (s *Session) GuildMuteList(guildID string) (gml *GuildMuteList, err error) {
	var response []byte
	u, _ := url.Parse(EndpointGuildMuteList)
	q := u.Query()
	q.Set("guild_id", guildID)
	u.RawQuery = q.Encode()
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	gml = &GuildMuteList{}
	err = json.Unmarshal(response, gml)
	if err != nil {
		return nil, err
	}
	return gml, nil
}

// MuteType is the type for mute status.
type MuteType int8

// These are all mute types.
const (
	MuteTypeMic MuteType = iota + 1
	MuteTypeHeadset
)

// GuildMuteSetting is the type for arguments of GuildMuteSetting.
type GuildMuteSetting struct {
	GuildID string   `json:"guild_id"`
	UserID  string   `json:"user_id"`
	Type    MuteType `json:"type"`
}

// GuildMuteCreate revokes a users privilege of using mic or headset.
// FYI: https://developer.kaiheila.cn/doc/http/guild#%E6%B7%BB%E5%8A%A0%E6%9C%8D%E5%8A%A1%E5%99%A8%E9%9D%99%E9%9F%B3%E6%88%96%E9%97%AD%E9%BA%A6
func (s *Session) GuildMuteCreate(gms *GuildMuteSetting) (err error) {
	_, err = s.Request("POST", EndpointGuildMuteCreate, gms)
	return err
}

// GuildMuteDelete re-grants a users privilege of using mic or headset.
// FYI: https://developer.kaiheila.cn/doc/http/guild#%E5%88%A0%E9%99%A4%E6%9C%8D%E5%8A%A1%E5%99%A8%E9%9D%99%E9%9F%B3%E6%88%96%E9%97%AD%E9%BA%A6
func (s *Session) GuildMuteDelete(gms *GuildMuteSetting) (err error) {
	_, err = s.Request("POST", EndpointGuildMuteDelete, gms)
	return err
}

// GuildRoleList returns the roles in a guild.
// FYI: https://developer.kaiheila.cn/doc/http/guild-role#%E8%8E%B7%E5%8F%96%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%A7%92%E8%89%B2%E5%88%97%E8%A1%A8
func (s *Session) GuildRoleList(guildID string, page *PageSetting) (rs []*Role, meta *PageInfo, err error) {
	var response []byte
	u, _ := url.Parse(EndpointGuildRoleList)
	q := u.Query()
	q.Add("guild_id", guildID)
	u.RawQuery = q.Encode()
	response, meta, err = s.RequestWithPage("GET", u.String(), page)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(response, &rs)
	if err != nil {
		return nil, nil, err
	}
	return rs, meta, err
}

// GuildRoleCreate creates a role for a guild.
//
// FYI: https://developer.kaiheila.cn/doc/http/guild-role#%E5%88%9B%E5%BB%BA%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%A7%92%E8%89%B2
func (s *Session) GuildRoleCreate(name, guildID string) (r *Role, err error) {
	var response []byte
	response, err = s.Request("POST", EndpointGuildRoleCreate, struct {
		Name    string `json:"name,omitempty"`
		GuildID string `json:"guild_id"`
	}{name, guildID})
	if err != nil {
		return nil, err
	}
	r = &Role{}
	err = json.Unmarshal(response, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GuildRoleUpdate updates a role for a guild.
//
// FYI: https://developer.kaiheila.cn/doc/http/guild-role#%E6%9B%B4%E6%96%B0%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%A7%92%E8%89%B2
func (s *Session) GuildRoleUpdate(guildID string, role *Role) (r *Role, err error) {
	var response []byte
	response, err = s.Request("POST", EndpointGuildRoleUpdate, struct {
		*Role
		GuildID string `json:"guild_id"`
	}{
		role, guildID,
	})
	if err != nil {
		return nil, err
	}
	r = &Role{}
	err = json.Unmarshal(response, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GuildRoleDelete deletes a role from a guild.
//
// FYI: https://developer.kaiheila.cn/doc/http/guild-role#%E5%88%A0%E9%99%A4%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%A7%92%E8%89%B2
func (s *Session) GuildRoleDelete(guildID, roleID string) (err error) {
	_, err = s.Request("POST", EndpointGuildRoleDelete, struct {
		GuildID string `json:"guild_id"`
		RoleID  string `json:"role_id"`
	}{guildID, roleID})
	return err
}

// GuildRoleResp is the response of GuildRoleGrant request.
type GuildRoleResp struct {
	GuildID string  `json:"guild_id"`
	UserID  string  `json:"user_id"`
	Roles   []int64 `json:"roles"`
}

// GuildRoleGrant grants a role to a user.
//
// FYI: https://developer.kaiheila.cn/doc/http/guild-role#%E8%B5%8B%E4%BA%88%E7%94%A8%E6%88%B7%E8%A7%92%E8%89%B2
func (s *Session) GuildRoleGrant(guildID, userID string, roleID int64) (grr *GuildRoleResp, err error) {
	return s.guildRoleGrantRevoke(guildID, userID, roleID, true)
}

// GuildRoleRevoke revokes a role from a user.
//
// FYI: https://developer.kaiheila.cn/doc/http/guild-role#%E5%88%A0%E9%99%A4%E7%94%A8%E6%88%B7%E8%A7%92%E8%89%B2
func (s *Session) GuildRoleRevoke(guildID, userID string, roleID int64) (grr *GuildRoleResp, err error) {
	return s.guildRoleGrantRevoke(guildID, userID, roleID, false)
}

func (s *Session) guildRoleGrantRevoke(guildID, userID string, roleID int64, grant bool) (grr *GuildRoleResp, err error) {
	var response []byte
	var endpoint string
	if grant {
		endpoint = EndpointGuildRoleGrant
	} else {
		endpoint = EndpointGuildRoleRevoke
	}
	response, err = s.Request("POST", endpoint, struct {
		GuildID string `json:"guild_id"`
		UserID  string `json:"user_id"`
		RoleID  int64  `json:"role_id"`
	}{guildID, userID, roleID})
	if err != nil {
		return nil, err
	}
	grr = &GuildRoleResp{}
	err = json.Unmarshal(response, grr)
	if err != nil {
		return nil, err
	}
	return grr, err
}

// IntimacyIndexResp is the type for intimacy info.
type IntimacyIndexResp struct {
	ImgURL     string         `json:"img_url"`
	SocialInfo string         `json:"social_info"`
	LastRead   MilliTimeStamp `json:"last_read"`
	ImgList    struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"img_list"`
}

// IntimacyIndex returns the intimacy info for a user.
//
// FYI: https://developer.kaiheila.cn/doc/http/intimacy#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E4%BA%B2%E5%AF%86%E5%BA%A6
func (s *Session) IntimacyIndex(userID string) (iir *IntimacyIndexResp, err error) {
	var response []byte
	u, _ := url.Parse(EndpointIntimacyIndex)
	q := u.Query()
	q.Set("user_id", userID)
	u.RawQuery = q.Encode()
	response, err = s.Request("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	iir = &IntimacyIndexResp{}
	err = json.Unmarshal(response, iir)
	if err != nil {
		return nil, err
	}
	return iir, err
}

// IntimacyUpdate is the type for arguments for IntimacyUpdate request.
type IntimacyUpdate struct {
	UserID     string `json:"user_id"`
	Score      *int   `json:"score,omitempty"`
	SocialInfo string `json:"social_info,omitempty"`
	ImgID      int    `json:"img_id,omitempty"`
}

// IntimacyUpdate updates the intimacy info for a user.
//
// FYI: https://developer.kaiheila.cn/doc/http/intimacy#%E6%9B%B4%E6%96%B0%E7%94%A8%E6%88%B7%E4%BA%B2%E5%AF%86%E5%BA%A6
func (s *Session) IntimacyUpdate(iu *IntimacyUpdate) (err error) {
	_, err = s.Request("POST", EndpointIntimacyUpdate, iu)
	return err
}

// GuildEmojiResp is the type for response of GuildEmojiList request.
type GuildEmojiResp struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	UserInfo User   `json:"user_info"`
}

// GuildEmojiList returns the list of emojis in a guild
//
// FYI: https://developer.kaiheila.cn/doc/http/guild-emoji#%E8%8E%B7%E5%8F%96%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%A1%A8%E6%83%85%E5%88%97%E8%A1%A8
func (s *Session) GuildEmojiList(guildID string, page *PageSetting) (gers []*GuildEmojiResp, meta *PageInfo, err error) {
	var response []byte
	u, _ := url.Parse(EndpointGuildEmojiList)
	q := u.Query()
	q.Set("guild_id", guildID)
	u.RawQuery = q.Encode()
	response, meta, err = s.RequestWithPage("GET", u.String(), page)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(response, &gers)
	if err != nil {
		return nil, nil, err
	}
	return gers, meta, err
}

// GuildEmojiCreate uploads emoji to guild.
//
// FYI: https://developer.kaiheila.cn/doc/http/guild-emoji#%E5%88%9B%E5%BB%BA%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%A1%A8%E6%83%85
func (s *Session) GuildEmojiCreate(name, guildID string, emoji []byte) (ger *GuildEmojiResp, err error) {
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	var fw io.Writer
	fw, err = w.CreateFormFile("emoji", "emoji.png")
	if err != nil {
		return nil, err
	}
	_, err = fw.Write(emoji)
	if err != nil {
		return nil, err
	}
	if len(name) <= 32 && len(name) >= 2 {
		fw, err = w.CreateFormField("name")
		if err != nil {
			return nil, err
		}
		_, err = fw.Write([]byte(name))
		if err != nil {
			return nil, err
		}
	}
	fw, err = w.CreateFormField("guild_id")
	if err != nil {
		return nil, err
	}
	_, err = fw.Write([]byte(guildID))
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	var f assetFile
	f.Payload = b.Bytes()
	f.ContentType = w.FormDataContentType()
	var response []byte
	response, err = s.Request("POST", EndpointGuildEmojiCreate, &f)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &ger)
	if err != nil {
		return nil, err
	}
	return ger, nil
}

// GuildEmojiUpdate updates an emoji's info in a guild.
//
// FYI: https://developer.kaiheila.cn/doc/http/guild-emoji#%E6%9B%B4%E6%96%B0%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%A1%A8%E6%83%85
func (s *Session) GuildEmojiUpdate(name, id string) (err error) {
	_, err = s.Request("POST", EndpointGuildEmojiUpdate, struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}{name, id})
	return err
}

// GuildEmojiDelete deletes an emoji from a guild.
//
// FYI: https://developer.kaiheila.cn/doc/http/guild-emoji#%E5%88%A0%E9%99%A4%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%A1%A8%E6%83%85
func (s *Session) GuildEmojiDelete(id string) (err error) {
	_, err = s.Request("POST", EndpointGuildEmojiDelete, struct {
		ID string `json:"id"`
	}{id})
	return err
}

// InviteListOption is the optional arguments for InviteList requests.
type InviteListOption func(values url.Values)

// InviteListWithGuildID adds optional `guild_id` argument to InviteList request.
func InviteListWithGuildID(guildID string) InviteListOption {
	return func(values url.Values) {
		values.Set("guild_id", guildID)
	}
}

// InviteListWithChannelID adds optional `channel_id` argument to InviteList request.
func InviteListWithChannelID(channelID string) InviteListOption {
	return func(values url.Values) {
		values.Set("channel_id", channelID)
	}
}

// InviteListResp is the type for response of InviteList request.
type InviteListResp struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	URLCode   string `json:"url_code"`
	URL       string `json:"url"`
	User      User   `json:"user"`
}

// InviteList lists invite links of a guild.
//
// FYI: https://developer.kaiheila.cn/doc/http/invite#%E8%8E%B7%E5%8F%96%E9%82%80%E8%AF%B7%E5%88%97%E8%A1%A8
func (s *Session) InviteList(page *PageSetting, options ...InviteListOption) (ilrs []*InviteListResp, meta *PageInfo, err error) {
	u, _ := url.Parse(EndpointInviteList)
	q := u.Query()
	for _, item := range options {
		item(q)
	}
	u.RawQuery = q.Encode()
	var response []byte
	response, meta, err = s.RequestWithPage("GET", u.String(), page)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(response, &ilrs)
	if err != nil {
		return nil, nil, err
	}
	return ilrs, meta, err
}

// InviteCreate is the type for arguments of InviteCreate request.
type InviteCreate struct {
	GuildID      string             `json:"guild_id,omitempty"`
	ChannelID    string             `json:"channel_id,omitempty"`
	Duration     InviteDuration     `json:"duration,omitempty"`
	SettingTimes InviteSettingTimes `json:"setting_times,omitempty"`
}

// InviteDuration is the type for Duration in InviteCreate
//
// You MUST use the enum defined below.
type InviteDuration string

// These are allowed InviteDuration enums. You MUST use them.
const (
	InviteDurationInfinity   InviteDuration = "0"
	InviteDurationHalfHour   InviteDuration = "1800"
	InviteDurationHour       InviteDuration = "3600"
	InviteDurationSixHour    InviteDuration = "21600"
	InviteDurationTwelveHour InviteDuration = "43200"
	InviteDurationDay        InviteDuration = "86400"
	InviteDurationWeek       InviteDuration = "604800"
)

// InviteSettingTimes is the type for SettingTimes in InviteCreate
//
// You SHOULD use the enum defined below.
type InviteSettingTimes int

// These are allowed InviteSettingTimes enums. You SHOULD use them.
const (
	InviteSettingTimesInfinity   InviteSettingTimes = -1
	InviteSettingTimesOne        InviteSettingTimes = 1
	InviteSettingTimesFive       InviteSettingTimes = 5
	InviteSettingTimesTen        InviteSettingTimes = 10
	InviteSettingTimesTwentyFive InviteSettingTimes = 25
	InviteSettingTimesFifty      InviteSettingTimes = 50
	InviteSettingTimesHundred    InviteSettingTimes = 100
)

// InviteCreate creates an invite link.
//
// FYI: https://developer.kaiheila.cn/doc/http/invite#%E5%88%9B%E5%BB%BA%E9%82%80%E8%AF%B7%E9%93%BE%E6%8E%A5
func (s *Session) InviteCreate(ic *InviteCreate) (URL string, err error) {
	var response []byte
	response, err = s.Request("POST", EndpointInviteCreate, ic)
	if err != nil {
		return "", err
	}
	a := &struct {
		URL string `json:"url"`
	}{}
	err = json.Unmarshal(response, a)
	if err != nil {
		return "", err
	}
	return a.URL, nil
}

// InviteDelete is the type for arguments of InviteDelete request.
type InviteDelete struct {
	GuildID   string `json:"guild_id,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
	URLCode   string `json:"url_code"`
}

// InviteDelete deletes an invite link.
//
// FYI: https://developer.kaiheila.cn/doc/http/invite#%E5%88%A0%E9%99%A4%E9%82%80%E8%AF%B7%E9%93%BE%E6%8E%A5
func (s *Session) InviteDelete(id *InviteDelete) (err error) {
	_, err = s.Request("POST", EndpointInviteDelete, id)
	return err
}

// UserMe returns the bot info.
// FYI: https://developer.kaiheila.cn/doc/http/user#%E8%8E%B7%E5%8F%96%E5%BD%93%E5%89%8D%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF
func (s *Session) UserMe() (u *User, err error) {
	var response []byte
	response, err = s.Request("GET", EndpointUserMe, nil)
	if err != nil {
		return nil, err
	}
	u = &User{}
	err = json.Unmarshal(response, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UserView returns a user's info
// FYI: https://developer.kaiheila.cn/doc/http/user#%E8%8E%B7%E5%8F%96%E7%9B%AE%E6%A0%87%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF
func (s *Session) UserView(userID string, options ...UserViewOption) (u *User, err error) {
	var response []byte
	ur, _ := url.Parse(EndpointUserView)
	q := ur.Query()
	q.Set("user_id", userID)
	for _, item := range options {
		item(q)
	}
	ur.RawQuery = q.Encode()
	response, err = s.Request("GET", ur.String(), nil)
	if err != nil {
		return nil, err
	}
	u = &User{}
	err = json.Unmarshal(response, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UserViewOption is the optional arguments for UserView requests.
type UserViewOption func(value url.Values)

// UserViewWithGuildID add the optional `guild_id` arguments to UserView request
func UserViewWithGuildID(guildID string) UserViewOption {
	return func(value url.Values) {
		value.Set("guild_id", guildID)
	}
}

// UserOffline logout the bot.
func (s *Session) UserOffline() error {
	_, err := s.Request("POST", EndpointUserOffline, nil)
	return err
}

// RequestWithPage is the wrapper for internal list GET request, you would prefer to use other method other than this.
func (s *Session) RequestWithPage(method, u string, page *PageSetting) (response []byte, meta *PageInfo, err error) {
	ur, _ := url.Parse(u)
	if page != nil {
		q := ur.Query()
		if page.Page != nil {
			q.Add("page", strconv.Itoa(*page.Page))
		}
		if page.PageSize != nil {
			q.Add("page_size", strconv.Itoa(*page.PageSize))
		}
		if page.Sort != nil {
			q.Add("sort", *page.Sort)
		}
		ur.RawQuery = q.Encode()
	}
	resp, err := s.Request(method, ur.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	g := &GeneralListData{}
	err = json.Unmarshal(resp, g)
	if err != nil {
		return nil, nil, err
	}
	return g.Items, &g.Meta, err
}

// Request is the wrapper for internal request method, you would prefer to use other method other than this.
func (s *Session) Request(method, url string, data interface{}) (response []byte, err error) {
	return s.request(method, url, data, 0)
}

type assetFile struct {
	Payload     []byte
	ContentType string
}

func (s *Session) request(method, url string, data interface{}, sequence int) (response []byte, err error) {
	var body []byte
	var dataMultipart bool
	if data != nil {
		if d, ok := data.(*assetFile); ok {
			body = d.Payload
			dataMultipart = true
		} else {
			body, err = json.Marshal(data)
			if err != nil {
				return
			}
		}
	}
	//s.log(LogTrace, "Api Request %s %s\n", method, url)
	e := s.Logger.Trace().Str("method", method).Str("url", url)
	e = addCaller(e)
	if len(body) != 0 {
		e = e.Bytes("payload", body)
	}
	e.Msg("http api request")
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}
	req.Header.Set("Authorization", s.Identify.Token)
	if len(body) > 0 {
		if dataMultipart {
			req.Header.Set("Content-Type", data.(*assetFile).ContentType)
		} else {
			req.Header.Set("Content-Type", "application/json")
		}
	}
	e = addCaller(s.Logger.Trace())
	for k, v := range req.Header {
		e = e.Strs(k, v)
		//s.log(LogTrace, "Api Request Header %s = %+v\n", k, v)
	}
	e.Msg("http api request headers")
	resp, err := s.Client.Do(req)
	if err != nil {
		addCaller(s.Logger.Error()).Err("err", err).Msg("")
		return
	}
	defer func() {
		err2 := resp.Body.Close()
		if err2 != nil {
			addCaller(s.Logger.Error()).Msg("error closing resp body")
			//s.log(LogError, "error closing resp body")
		}
	}()

	var respByte []byte

	respByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		addCaller(s.Logger.Error()).Err("err", err).Msg("")
		return
	}
	addCaller(s.Logger.Trace()).Int("status_code", resp.StatusCode).
		Str("status", resp.Status).
		Bytes("body", respByte).
		Msg("http response")
	//s.log(LogTrace, "Api Response Status %s\n", resp.Status)
	e = s.Logger.Trace()
	e = addCaller(e)
	for k, v := range resp.Header {
		e = e.Strs(k, v)
		//s.log(LogTrace, "Api Response Header %s = %+v\n", k, v)
	}
	e.Msg("http response headers")
	//s.log(LogTrace, "Api Response Body %s", respByte)
	var r EndpointGeneralResponse
	err = json.Unmarshal(respByte, &r)
	if err != nil {
		addCaller(s.Logger.Error()).Err("err", err).Msg("response unmarshal error")
		//s.log(LogError, "Api Response Unmarshal Error %s", err)
		return
	}
	if r.Code != 0 {
		addCaller(s.Logger.Error()).Int("code", r.Code).Str("error_msg", r.Message).Msg("api response error")
		//s.log(LogError, "Api Response Error Code %d, Message %s", r.Code, r.Message)
		return
	}
	response = r.Data
	return
}
