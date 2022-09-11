package kook

import (
	"encoding/json"
)

const unsupportedCardType = "Unsupported type."

// CardMessage is the type for a message of cards called 卡片消息.
type CardMessage []*CardMessageCard

// BuildMessage is a helper function to marshal card message for sending.
func (c CardMessage) BuildMessage() (s string, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(b), err
}

// MustBuildMessage is a helper function to marshal card message for sending.
func (c CardMessage) MustBuildMessage() string {
	s, err := c.BuildMessage()
	if err != nil {
		panic(`kook.CardMessage.BuildMessage:` + err.Error())
	}
	return s
}

// CardTheme is the type for card theme.
type CardTheme string

// These are predefined usable card themes.
const (
	CardThemePrimary   CardTheme = "primary"
	CardThemeSuccess   CardTheme = "success"
	CardThemeDanger    CardTheme = "danger"
	CardThemeWarning   CardTheme = "warning"
	CardThemeInfo      CardTheme = "info"
	CardThemeSecondary CardTheme = "secondary"
)

// CardSize is the type for card size.
type CardSize string

// These are predefined usable card sizes.
const (
	CardSizeSm CardSize = "sm"
	CardSizeLg CardSize = "lg"
)

// CardMessageCard is the type for 卡片.
type CardMessageCard struct {
	Theme   CardTheme     `json:"theme,omitempty"`
	Color   string        `json:"color,omitempty"`
	Size    CardSize      `json:"size,omitempty"`
	Modules []interface{} `json:"modules"`
}

type fakeCardMessageCard CardMessageCard

// MarshalJSON adds additional type field when marshaling
func (c CardMessageCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageCard
	}{
		"card", fakeCardMessageCard(c),
	})
}

// AddModule adds Modules to a card and provides a runtime type check for CardMessageCard Modules.
//
// Allowed Modules: *CardMessageHeader, *CardMessageSection, *CardMessageImageGroup, *CardMessageContainer,
// *CardMessageActionGroup, *CardMessageContext, *CardMessageDivider, *CardMessageFile, *CardMessageCountdown.
func (c *CardMessageCard) AddModule(i ...interface{}) *CardMessageCard {
	for _, item := range i {
		switch v := item.(type) {
		case *CardMessageHeader,
			*CardMessageSection,
			*CardMessageImageGroup,
			*CardMessageContainer,
			*CardMessageActionGroup,
			*CardMessageContext,
			*CardMessageDivider,
			*CardMessageFile,
			*CardMessageCountdown:
			c.Modules = append(c.Modules, v)
		default:
			panic(unsupportedCardType)
		}
	}
	return c
}

// CardMessageHeader is the type for 模块-标题模块.
type CardMessageHeader struct {
	Text CardMessageElementText `json:"text"`
}

type fakeCardMessageHeader CardMessageHeader

// MarshalJSON adds additional type field when marshaling
func (c CardMessageHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageHeader
	}{
		"header", fakeCardMessageHeader(c),
	})
}

// CardMessageSectionMode is the type of mode for CardMessageSection
type CardMessageSectionMode string

// These are predefined usable CardMessageSectionModes
const (
	CardMessageSectionModeLeft  CardMessageSectionMode = "left"
	CardMessageSectionModeRight CardMessageSectionMode = "right"
)

// CardMessageSection is the type for 模块-内容模块.
type CardMessageSection struct {
	Mode      CardMessageSectionMode `json:"mode,omitempty"`
	Text      interface{}            `json:"text"`
	Accessory interface{}            `json:"accessory,omitempty"`
}

type fakeCardMessageSection CardMessageSection

// MarshalJSON adds additional type field when marshaling.
func (c CardMessageSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageSection
	}{
		"section", fakeCardMessageSection(c),
	})
}

// SetText provides additional type-checking when setting elements to section text.
//
// Allowed elements: *CardMessageElementText, *CardMessageElementKMarkdown, *CardMessageParagraph.
func (c *CardMessageSection) SetText(i interface{}) *CardMessageSection {
	switch v := i.(type) {
	case *CardMessageElementText, *CardMessageElementKMarkdown, *CardMessageParagraph:
		c.Text = v
	default:
		panic(unsupportedCardType)
	}
	return c
}

// SetAccessory provides additional type-checking when setting elements to section accessory.
//
// Allowed elements: *CardMessageElementImage, *CardMessageElementButton.
func (c *CardMessageSection) SetAccessory(i interface{}) *CardMessageSection {
	switch v := i.(type) {
	case *CardMessageElementImage, *CardMessageElementButton:
		c.Accessory = v
	default:
		panic(unsupportedCardType)
	}
	return c
}

// CardMessageImageGroup is the type for 图片模块.
type CardMessageImageGroup []CardMessageElementImage

// MarshalJSON adds additional type field when marshaling
func (c CardMessageImageGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type     string                    `json:"type"`
		Elements []CardMessageElementImage `json:"elements"`
	}{
		"image-group", c,
	})
}

// CardMessageActionGroup is the type for 模块-交互模块.
type CardMessageActionGroup []CardMessageElementButton

// MarshalJSON adds additional type field when marshaling
func (c CardMessageActionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type     string                     `json:"type"`
		Elements []CardMessageElementButton `json:"elements"`
	}{
		"action-group", c,
	})
}

// CardMessageContext is the type for 模块-备注模块.
type CardMessageContext []interface{}

// MarshalJSON adds additional type field when marshaling
func (c CardMessageContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type     string        `json:"type"`
		Elements []interface{} `json:"elements"`
	}{
		"context", c,
	})
}

// AddItem provides additional type-checking when adding elements to context.
//
// Allowed elements: *CardMessageElementText, *CardMessageElementKMarkdown, *CardMessageElementImage
func (c *CardMessageContext) AddItem(i ...interface{}) *CardMessageContext {
	for _, item := range i {
		switch v := item.(type) {
		case *CardMessageElementText,
			*CardMessageElementKMarkdown,
			*CardMessageElementImage:
			*c = append(*c, v)
		default:
			panic("Unsupported type")
		}
	}
	return c
}

// CardMessageDivider is the type for 模块-分割线模块.
type CardMessageDivider struct {
}

// MarshalJSON adds additional type field when marshaling
func (c CardMessageDivider) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
	}{
		"divider",
	})
}

// CardMessageFile is the type for 模块-文件模块.
type CardMessageFile struct {
	Type  CardMessageFileType `json:"type"`
	Src   string              `json:"src"`
	Title string              `json:"title,omitempty"`
	Cover string              `json:"cover,omitempty"`
}

// CardMessageFileType is the type for types of CardMessageFile
type CardMessageFileType string

// These are predefined usable CardMessageFileTypes
const (
	CardMessageFileTypeFile  CardMessageFileType = "file"
	CardMessageFileTypeAudio CardMessageFileType = "audio"
	CardMessageFileTypeVideo CardMessageFileType = "video"
)

// CardMessageCountdown is the type for 模块-倒计时模块.
type CardMessageCountdown struct {
	EndTime   MilliTimeStamp           `json:"endTime"`
	StartTime MilliTimeStamp           `json:"startTime"`
	Mode      CardMessageCountdownMode `json:"mode"`
}

// CardMessageCountdownMode is the type for modes of CardMessageCountdown
type CardMessageCountdownMode string

// These are predefined usable CardMessageCountdownModes
const (
	CardMessageCountdownModeDay    CardMessageCountdownMode = "day"
	CardMessageCountdownModeHour   CardMessageCountdownMode = "hour"
	CardMessageCountdownModeSecond CardMessageCountdownMode = "second"
)

type fakeCardMessageCountdown CardMessageCountdown

// MarshalJSON adds additional type field when marshaling
func (c CardMessageCountdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageCountdown
	}{
		"countdown", fakeCardMessageCountdown(c),
	})
}

// CardMessageContainer is the type for 可拓展图片模块.
type CardMessageContainer []CardMessageElementImage

// MarshalJSON adds additional type field when marshaling.
func (c CardMessageContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type     string                    `json:"type"`
		Elements []CardMessageElementImage `json:"elements"`
	}{"container", c})
}

// AddElements adds elements to container.
func (c *CardMessageContainer) AddElements(i ...CardMessageElementImage) *CardMessageContainer {
	*c = append(*c, i...)
	return c
}

// CardMessageInvite is the type for 邀请模块.
type CardMessageInvite struct {
	Code string `json:"code"`
}

type fakeCardMessageInvite CardMessageInvite

// MarshalJSON adds additional type field when marshaling.
func (c CardMessageInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageInvite
	}{"container", fakeCardMessageInvite(c)})
}

// CardMessageElementText is the type for 元素-普通文本.
type CardMessageElementText struct {
	Content string `json:"content"`
	Emoji   bool   `json:"emoji"`
}

type fakeCardMessageElementText CardMessageElementText

// MarshalJSON adds additional type field when marshaling
func (c CardMessageElementText) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageElementText
	}{
		"plain-text", fakeCardMessageElementText(c),
	})
}

// CardMessageElementKMarkdown is the type for 元素-kmarkdown.
type CardMessageElementKMarkdown struct {
	Content string `json:"content"`
}

type fakeCardMessageElementKMarkdown CardMessageElementKMarkdown

// MarshalJSON adds additional type field when marshaling
func (c CardMessageElementKMarkdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageElementKMarkdown
	}{
		"kmarkdown", fakeCardMessageElementKMarkdown(c),
	})
}

// CardMessageElementImage is the type for 元素-图片.
type CardMessageElementImage struct {
	Src    string `json:"src"`
	Alt    string `json:"alt,omitempty"`
	Size   string `json:"size,omitempty"`
	Circle bool   `json:"circle"`
}

type fakeCardMessageElementImage CardMessageElementImage

// MarshalJSON adds additional type field when marshaling
func (c CardMessageElementImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageElementImage
	}{
		"image", fakeCardMessageElementImage(c),
	})
}

// CardMessageElementButton is the type for 元素-按钮.
type CardMessageElementButton struct {
	Theme CardTheme `json:"theme"`
	Value string    `json:"value"`
	Click string    `json:"click,omitempty"`
	Text  string    `json:"text"`
}

// CardMessageElementButtonClick is the type for click modes of CardMessageElementButton
type CardMessageElementButtonClick string

// These are predefined usable CardMessageElementButtonClicks
const (
	CardMessageElementButtonClickLink      CardMessageElementButtonClick = "link"
	CardMessageElementButtonClickReturnVal CardMessageElementButtonClick = "return-val"
)

type fakeCardMessageElementButton CardMessageElementButton

// MarshalJSON adds additional type field when marshaling
func (c CardMessageElementButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageElementButton
	}{
		"button", fakeCardMessageElementButton(c),
	})
}

// CardMessageParagraph is the type for 结构体-区域文本.
type CardMessageParagraph struct {
	Cols   int           `json:"cols"`
	Fields []interface{} `json:"fields"`
}

type fakeCardMessageParagraph CardMessageParagraph

// MarshalJSON adds additional type field when marshaling
func (c CardMessageParagraph) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		fakeCardMessageParagraph
	}{
		"paragraph", fakeCardMessageParagraph(c),
	})
}

// AddField provides additional type-checking when adding elements to paragraph.
//
// Allowed elements: *CardMessageElementText, *CardMessageElementKMarkdown
func (c *CardMessageParagraph) AddField(i ...interface{}) *CardMessageParagraph {
	for _, item := range i {
		switch v := item.(type) {
		case *CardMessageElementText,
			*CardMessageElementKMarkdown:
			c.Fields = append(c.Fields, v)
		default:
			panic(unsupportedCardType)
		}
	}
	return c
}
