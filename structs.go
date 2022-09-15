package kook

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// RolePermission is the type for the permission of a user in guilds or channels.
type RolePermission int64

// HasPermission checks if having the provided permission.
func (r RolePermission) HasPermission(p RolePermission) bool {
	return r&(p|RolePermissionAdmin) != 0
}

// These are the permission defined in system.
const (
	RolePermissionAdmin RolePermission = 1 << iota
	RolePermissionManageGuild
	RolePermissionViewAuditLog
	RolePermissionCreateInvite
	RolePermissionManageInvite
	RolePermissionManageChannel
	RolePermissionKickUser
	RolePermissionBanUser
	RolePermissionManageGuildEmoji
	RolePermissionChangeNickname
	RolePermissionManageRolePermission
	RolePermissionViewChannel
	RolePermissionSendMessage
	RolePermissionManageMessage
	RolePermissionUploadFile
	RolePermissionConnectVoice
	RolePermissionManageVoice
	RolePermissionMentionEveryone
	RolePermissionCreateReaction
	RolePermissionFollowReaction
	RolePermissionInvitedToVoice
	RolePermissionForceManualVoice
	RolePermissionFreeVoice
	RolePermissionVoice
	RolePermissionManageUserVoiceReceive
	RolePermissionManageUserVoiceCreate
	RolePermissionManageNickname
	RolePermissionPlayMusic
)

// UserStatus is the type for user status(banned or not).
type UserStatus int8

// IsBanned checks if the user is banned.
func (r UserStatus) IsBanned() bool {
	return r == UserStatusBanned
}

// These are all the status of a user.
const (
	UserStatusNormal UserStatus = 1
	UserStatusBanned UserStatus = 10
)

// GuildNotifyType is the type of the notify type of a guild.
type GuildNotifyType int8

// These are all notify types.
const (
	GuildNotifyTypeDefault GuildNotifyType = iota
	GuildNotifyTypeAll
	GuildNotifyTypeMention
	GuildNotifyTypeDisable
)

// ChannelType is the type of a channel.
type ChannelType int8

// These are all channel types.
const (
	ChannelTypeText ChannelType = 1 + iota
	ChannelTypeVoice
)

// User is the struct for a user. Some property may missing.
type User struct {
	ID             string         `json:"id"`
	Username       string         `json:"username"`
	IdentifyNum    string         `json:"identify_num"`
	Online         bool           `json:"online"`
	Status         UserStatus     `json:"status"`
	Avatar         string         `json:"avatar"`
	VipAvatar      string         `json:"vip_avatar"`
	Bot            bool           `json:"bot"`
	MobileVerified bool           `json:"mobile_verified"`
	System         bool           `json:"system"`
	MobilePrefix   string         `json:"mobile_prefix"`
	Mobile         string         `json:"mobile"`
	InviteCount    int64          `json:"invite_count"`
	Nickname       string         `json:"nickname"`
	Roles          []int64        `json:"roles"`
	FullName       string         `json:"full_name"`
	IsVip          bool           `json:"is_vip"`
	JoinedAt       MilliTimeStamp `json:"joined_at"`
	ActiveTime     MilliTimeStamp `json:"active_time"`
}

// Guild is the struct for a server/guild(服务器).
type Guild struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Topic    string `json:"topic"`
	MasterID string `json:"master_id"`
	// For compatibility reason
	UserID           string          `json:"user_id"`
	Icon             string          `json:"icon"`
	NotifyType       GuildNotifyType `json:"notify_type"`
	Region           string          `json:"region"`
	EnableOpen       IntBool         `json:"enable_open"`
	OpenID           string          `json:"open_id"`
	DefaultChannelID string          `json:"default_channel_id"`
	WelcomeChannelID string          `json:"welcome_channel_id"`
	Roles            []Role          `json:"roles"`
	Channels         []Channel       `json:"channels"`
}

// GetMasterID returns the master id of the guild, not sure if it is necessary as may doc is wrong.
func (g Guild) GetMasterID() string {
	if g.MasterID != "" {
		return g.MasterID
	}
	return g.UserID
}

// IntBool is the type for some int value in response which only has two valid values.
type IntBool bool

// MarshalJSON is used to marshal IntBool for reqeust.
func (i *IntBool) MarshalJSON() ([]byte, error) {
	switch *i {
	case true:
		return []byte("1"), nil
	case false:
		return []byte("0"), nil
	default:
		return []byte(""), nil
	}
}

// UnmarshalJSON is used to unmarshal IntBool from response.
func (i *IntBool) UnmarshalJSON(bytes []byte) error {
	switch bytes[0] {
	case '0', 'f':
		*i = false
	case '1', 't':
		*i = true
	default:
		return errors.New("unable to unmarshal int-bool")
	}
	return nil
}

// Role is the struct for a role in the guild.
type Role struct {
	RoleID      int64          `json:"role_id"`
	Name        string         `json:"name,omitempty"`
	Color       int            `json:"color,omitempty"`
	Position    int            `json:"position"`
	Hoist       IntBool        `json:"hoist,omitempty"`
	Mentionable IntBool        `json:"mentionable,omitempty"`
	Permissions RolePermission `json:"permissions,omitempty"`
}

// Channel is the struct for a channel in guild. For different channels, some fields may be empty.
type Channel struct {
	ID                   string                    `json:"id"`
	Name                 string                    `json:"name"`
	UserID               string                    `json:"user_id"`
	MasterID             string                    `json:"master_id"`
	GuildID              string                    `json:"guild_id"`
	Topic                string                    `json:"topic"`
	IsCategory           IntBool                   `json:"is_category"`
	ParentID             string                    `json:"parent_id"`
	Level                int                       `json:"level"`
	SlowMode             int                       `json:"slow_mode"`
	Type                 ChannelType               `json:"type"`
	LimitAmount          int                       `json:"limit_amount"`
	PermissionOverwrites []PermissionOverwrite     `json:"permission_overwrites"`
	PermissionUsers      []UserPermissionOverwrite `json:"permission_users"`
	PermissionSync       IntBool                   `json:"permission_sync"`
	ServerURL            string                    `json:"server_url"`
}

// PermissionOverwrite is the struct for where needs to customize permission for a role in a channel.
type PermissionOverwrite struct {
	RoleID int64          `json:"role_id"`
	Allow  RolePermission `json:"allow"`
	Deny   RolePermission `json:"deny"`
}

// UserPermissionOverwrite is the struct for where needs to customize permission for a user in a channel.
type UserPermissionOverwrite struct {
	User  *User          `json:"user"`
	Allow RolePermission `json:"allow"`
	Deny  RolePermission `json:"deny"`
}

// Event is the struct for every received event.
type Event struct {
	Signal         EventSignal     `json:"s"`
	Data           json.RawMessage `json:"d"`
	SequenceNumber int64           `json:"sn"`
}

// EventSignal is the type for event types.
type EventSignal int8

// All event signal consts.
const (
	EventSignalEvent EventSignal = iota
	EventSignalHello
	EventSignalPing
	EventSignalPong
	EventSignalReconnect EventSignal = iota + 1
	EventSignalResumeAck
)

// EventStatusCode is the type for various event status code.
type EventStatusCode int

// EventStatusCode consts for event status
const (
	EventStatusOk              EventStatusCode = 0
	EventStatusMissingArgument EventStatusCode = 40100 + iota
	EventStatusInvalidToken
	EventStatusTokenAuthFailed
	EventStatusTokenExpired
	EventStatusResumeFailed EventStatusCode = 40100 + 2 + iota
	EventStatusSessionExpired
	EventStatusInvalidSequenceNumber
)

// EventDataHello is the struct for the data of event hello
type EventDataHello struct {
	Code      EventStatusCode `json:"code"`
	SessionID string          `json:"session_id"`
}

// EventDataResumeAck is the struct for the data of event resume ack.
type EventDataResumeAck struct {
	SessionID string `json:"session_id"`
}

// Identify is the struct for the initial settings sent to kookapp.
type Identify struct {
	Token        string
	Compress     bool
	WebsocketKey []byte
	ClientID     string
	ClientSecret string
	VerifyToken  string
}

// Session is the struct for a bot session.
type Session struct {
	sync.RWMutex

	Identify          Identify
	LastHeartbeatAck  time.Time
	LastHeartbeatSent time.Time
	Client            *http.Client
	MaxRetry          int
	RetryTimeout      time.Duration
	ContentType       string
	Logger            Logger
	Sync              bool

	wsConn  *websocket.Conn
	wsMutex sync.Mutex
	gateway string
	// sessionID string
	sequence  *int64
	listening chan interface{}

	handlersMu sync.RWMutex
	handlers   map[string][]*eventHandlerInstance

	snStore SnStore
}

// EventDataGeneral is the struct passed to all event handler.
type EventDataGeneral struct {
	ChannelType  string      `json:"channel_type"`
	Type         MessageType `json:"type"`
	TargetID     string      `json:"target_id"`
	AuthorID     string      `json:"author_id"`
	Content      string      `json:"content"`
	MsgID        string      `json:"msg_id"`
	MsgTimestamp int64       `json:"msg_timestamp"`
	Nonce        string      `json:"nonce"`
}

// EventData is the struct for initial parsing event's data payload.
type EventData struct {
	*EventDataGeneral
	Extra json.RawMessage `json:"extra"`
}

// MessageType is the type for messages from events.
type MessageType uint8

// MessageType consts for event message type
const (
	MessageTypeText MessageType = 1 + iota
	MessageTypeImage
	MessageTypeVideo
	MessageTypeFile
	MessageTypeAudio MessageType = 4 + iota
	MessageTypeKMarkdown
	MessageTypeCard
	MessageTypeSystem MessageType = 255
)

// EndpointGeneralResponse is the struct for initial parsing REST requests.
type EndpointGeneralResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

const heartbeatInterval = time.Second * 30

// MessageReaction is the struct for reactions embedded to a message.
type MessageReaction struct {
	MsgID     string    `json:"msg_id"`
	UserID    string    `json:"user_id"`
	ChannelID string    `json:"channel_id"`
	ChatCode  string    `json:"chat_code"`
	Emoji     EmojiItem `json:"emoji"`
}

// EmojiItem is the type for an emoji.
type EmojiItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// IsEqual compares standard emoji string with kook's emoji representation.
func (e *EmojiItem) IsEqual(s string) bool {
	return e.Convert() == s
}

// Convert converts kook's emoji to standard emoji.
func (e *EmojiItem) Convert() string {
	if !strings.HasPrefix(e.ID, "[#") {
		return e.ID
	}
	t := strings.TrimLeft(e.ID, "[#")
	t = strings.TrimRight(t, ";]")
	i, err := strconv.Atoi(t)
	if err != nil {
		return ""
	}
	return string([]rune{int32(i)})
}

// ChannelMessage is the struct for a message in a channel.
type ChannelMessage struct {
	MsgID       string   `json:"msg_id"`
	Content     string   `json:"content"`
	ContentID   string   `json:"content_id"`
	Mention     []string `json:"mention"`
	MentionAll  bool     `json:"mention_all"`
	MentionHere bool     `json:"mention_here"`
	MentionRole []string `json:"mention_role"`
	UpdatedAt   int64    `json:"updated_at"`
}

// DetailedChannelMessage is the struct for a detailed message in a channel.
type DetailedChannelMessage struct {
	ID          string                  `json:"id"`
	Type        MessageType             `json:"type"`
	Author      User                    `json:"author"`
	Content     string                  `json:"content"`
	Mention     []string                `json:"mention"`
	MentionAll  bool                    `json:"mention_all"`
	MentionHere bool                    `json:"mention_here"`
	MentionRole []string                `json:"mention_role"`
	Embeds      []map[string]string     `json:"embeds"`
	Attachments *Attachment             `json:"attachments"`
	Reactions   []ReactionItem          `json:"reactions"`
	Quote       *DetailedChannelMessage `json:"quote"`
	MentionInfo struct {
		MentionPart     []*User `json:"mention_part"`
		MentionRolePart []*Role `json:"mention_role_part"`
	} `json:"mention_info"`
}

// ReactionItem is the reactions for a emoji to a message.
type ReactionItem struct {
	Emoji EmojiItem `json:"emoji"`
	Count int       `json:"count"`
	Me    bool      `json:"me"`
}

// MilliTimeStamp is the timestamp used in kookapp API.
type MilliTimeStamp int64

// ToTime converts the timestamp to golang's time.Time.
func (t *MilliTimeStamp) ToTime() time.Time {
	return time.Unix(int64(*t)/1000, int64(*t)%1000*1000*1000)
}

// MilliTimeStampOfTime converts the time.Time to MillTimeStamp
func MilliTimeStampOfTime(t time.Time) MilliTimeStamp {
	return MilliTimeStamp(t.UnixNano() / 1000000)
}

// PrivateMessage is the struct for messages in direct chat.
type PrivateMessage struct {
	MsgID     string         `json:"msg_id"`
	AuthorID  string         `json:"author_id"`
	TargetID  string         `json:"target_id"`
	Content   string         `json:"content"`
	ChatCode  string         `json:"chat_code"`
	UpdatedAt MilliTimeStamp `json:"updated_at"`
	DeletedAt MilliTimeStamp `json:"deleted_at"`
}

// Attachment is the struct for various attachments, so that according to type, some fields may be empty.
type Attachment struct {
	Type     string  `json:"type"`
	URL      string  `json:"url"`
	Name     string  `json:"name"`
	FileType string  `json:"file_type"`
	Size     int64   `json:"size"`
	Duration float64 `json:"duration"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
}

// MessageWithAttachment is a message with attachment.
type MessageWithAttachment struct {
	Type        MessageType `json:"type"`
	Code        string      `json:"code"`
	GuildID     string      `json:"guild_id"`
	Attachments Attachment  `json:"attachments"`
	Author      User        `json:"author"`
}

// EventDataSystem is the struct for initial parsing system events.
// TODO: Implement a marshaller.
type EventDataSystem struct {
	Type string          `json:"type"`
	Body json.RawMessage `json:"body"`
}

//type RestChannelMessage struct {
//	ID string `json:"id"`
//	Type MessageType `json:"type"`
//	Author User `json:"author"`
//	Content string `json:"content"`
//	Mention []string `json:"mention"`
//	MentionAll bool `json:"mention_all"`
//	MentionRoles []string `json:"mention_roles"`
//	MentionHere bool `json:"mention_here"`
//	Embeds []map[string]string `json:"embeds"`
//	Attachments Attachment `json:"attachments"`
//}

// UserChat is the struct for DirectMessage or UserChat
type UserChat struct {
	Code            string         `json:"code"`
	LastReadTime    MilliTimeStamp `json:"last_read_time"`
	LatestMsgTime   MilliTimeStamp `json:"latest_msg_time"`
	UnreadCount     int            `json:"unread_count"`
	IsFriend        bool           `json:"is_friend"`
	IsBlocked       bool           `json:"is_blocked"`
	IsTargetBlocked bool           `json:"is_target_blocked"`
	TargetInfo      User           `json:"target_info"`
}

// GeneralListData is the struct for list GET responses.
type GeneralListData struct {
	Items json.RawMessage `json:"items"`
	Meta  PageInfo        `json:"meta"`
	Sort  map[string]int  `json:"sort"`
}

// PageInfo is the struct for page info in list GET responses.
type PageInfo struct {
	Page      int `json:"page"`
	PageTotal int `json:"page_total"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
}

// PageSetting is the type for page setting in list GET request arguments.
type PageSetting struct {
	Page     *int    `json:"page"`
	PageSize *int    `json:"page_size"`
	Sort     *string `json:"sort"`
}

// EventHandlerCommonContext is the common context for event handlers.
type EventHandlerCommonContext struct {
	Session *Session
	Common  *EventDataGeneral
}

// Quote is the struct for quotes in message events.
type Quote struct {
	ID       string         `json:"id"`
	Type     MessageType    `json:"type"`
	Content  string         `json:"content"`
	CreateAt MilliTimeStamp `json:"create_at"`
	Author   *User          `json:"author"`
	RongID   string         `json:"rong_id"`
}
