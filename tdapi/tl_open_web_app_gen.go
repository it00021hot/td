// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// OpenWebAppRequest represents TL type `openWebApp#914ac418`.
type OpenWebAppRequest struct {
	// Identifier of the chat in which the Web App is opened
	ChatID int64
	// Identifier of the bot, providing the Web App
	BotUserID int64
	// The URL from an inlineKeyboardButtonTypeWebApp button, a botMenuButton button, or an
	// internalLinkTypeAttachmentMenuBot link, or an empty string otherwise
	URL string
	// Preferred Web App theme; pass null to use the default theme
	Theme ThemeParameters
	// Identifier of the replied message for the message sent by the Web App; 0 if none
	ReplyToMessageID int64
}

// OpenWebAppRequestTypeID is TL type id of OpenWebAppRequest.
const OpenWebAppRequestTypeID = 0x914ac418

// Ensuring interfaces in compile-time for OpenWebAppRequest.
var (
	_ bin.Encoder     = &OpenWebAppRequest{}
	_ bin.Decoder     = &OpenWebAppRequest{}
	_ bin.BareEncoder = &OpenWebAppRequest{}
	_ bin.BareDecoder = &OpenWebAppRequest{}
)

func (o *OpenWebAppRequest) Zero() bool {
	if o == nil {
		return true
	}
	if !(o.ChatID == 0) {
		return false
	}
	if !(o.BotUserID == 0) {
		return false
	}
	if !(o.URL == "") {
		return false
	}
	if !(o.Theme.Zero()) {
		return false
	}
	if !(o.ReplyToMessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (o *OpenWebAppRequest) String() string {
	if o == nil {
		return "OpenWebAppRequest(nil)"
	}
	type Alias OpenWebAppRequest
	return fmt.Sprintf("OpenWebAppRequest%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*OpenWebAppRequest) TypeID() uint32 {
	return OpenWebAppRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*OpenWebAppRequest) TypeName() string {
	return "openWebApp"
}

// TypeInfo returns info about TL type.
func (o *OpenWebAppRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "openWebApp",
		ID:   OpenWebAppRequestTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Theme",
			SchemaName: "theme",
		},
		{
			Name:       "ReplyToMessageID",
			SchemaName: "reply_to_message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (o *OpenWebAppRequest) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode openWebApp#914ac418 as nil")
	}
	b.PutID(OpenWebAppRequestTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *OpenWebAppRequest) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode openWebApp#914ac418 as nil")
	}
	b.PutInt53(o.ChatID)
	b.PutInt53(o.BotUserID)
	b.PutString(o.URL)
	if err := o.Theme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode openWebApp#914ac418: field theme: %w", err)
	}
	b.PutInt53(o.ReplyToMessageID)
	return nil
}

// Decode implements bin.Decoder.
func (o *OpenWebAppRequest) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode openWebApp#914ac418 to nil")
	}
	if err := b.ConsumeID(OpenWebAppRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode openWebApp#914ac418: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *OpenWebAppRequest) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode openWebApp#914ac418 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode openWebApp#914ac418: field chat_id: %w", err)
		}
		o.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode openWebApp#914ac418: field bot_user_id: %w", err)
		}
		o.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode openWebApp#914ac418: field url: %w", err)
		}
		o.URL = value
	}
	{
		if err := o.Theme.Decode(b); err != nil {
			return fmt.Errorf("unable to decode openWebApp#914ac418: field theme: %w", err)
		}
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode openWebApp#914ac418: field reply_to_message_id: %w", err)
		}
		o.ReplyToMessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (o *OpenWebAppRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if o == nil {
		return fmt.Errorf("can't encode openWebApp#914ac418 as nil")
	}
	b.ObjStart()
	b.PutID("openWebApp")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(o.ChatID)
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(o.BotUserID)
	b.Comma()
	b.FieldStart("url")
	b.PutString(o.URL)
	b.Comma()
	b.FieldStart("theme")
	if err := o.Theme.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode openWebApp#914ac418: field theme: %w", err)
	}
	b.Comma()
	b.FieldStart("reply_to_message_id")
	b.PutInt53(o.ReplyToMessageID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (o *OpenWebAppRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if o == nil {
		return fmt.Errorf("can't decode openWebApp#914ac418 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("openWebApp"); err != nil {
				return fmt.Errorf("unable to decode openWebApp#914ac418: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode openWebApp#914ac418: field chat_id: %w", err)
			}
			o.ChatID = value
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode openWebApp#914ac418: field bot_user_id: %w", err)
			}
			o.BotUserID = value
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode openWebApp#914ac418: field url: %w", err)
			}
			o.URL = value
		case "theme":
			if err := o.Theme.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode openWebApp#914ac418: field theme: %w", err)
			}
		case "reply_to_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode openWebApp#914ac418: field reply_to_message_id: %w", err)
			}
			o.ReplyToMessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (o *OpenWebAppRequest) GetChatID() (value int64) {
	if o == nil {
		return
	}
	return o.ChatID
}

// GetBotUserID returns value of BotUserID field.
func (o *OpenWebAppRequest) GetBotUserID() (value int64) {
	if o == nil {
		return
	}
	return o.BotUserID
}

// GetURL returns value of URL field.
func (o *OpenWebAppRequest) GetURL() (value string) {
	if o == nil {
		return
	}
	return o.URL
}

// GetTheme returns value of Theme field.
func (o *OpenWebAppRequest) GetTheme() (value ThemeParameters) {
	if o == nil {
		return
	}
	return o.Theme
}

// GetReplyToMessageID returns value of ReplyToMessageID field.
func (o *OpenWebAppRequest) GetReplyToMessageID() (value int64) {
	if o == nil {
		return
	}
	return o.ReplyToMessageID
}

// OpenWebApp invokes method openWebApp#914ac418 returning error if any.
func (c *Client) OpenWebApp(ctx context.Context, request *OpenWebAppRequest) (*WebAppInfo, error) {
	var result WebAppInfo

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
