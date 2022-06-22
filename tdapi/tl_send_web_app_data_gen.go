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

// SendWebAppDataRequest represents TL type `sendWebAppData#ab1fce0c`.
type SendWebAppDataRequest struct {
	// Identifier of the target bot
	BotUserID int64
	// Text of the keyboardButtonTypeWebApp button, which opened the Web App
	ButtonText string
	// Received data
	Data string
}

// SendWebAppDataRequestTypeID is TL type id of SendWebAppDataRequest.
const SendWebAppDataRequestTypeID = 0xab1fce0c

// Ensuring interfaces in compile-time for SendWebAppDataRequest.
var (
	_ bin.Encoder     = &SendWebAppDataRequest{}
	_ bin.Decoder     = &SendWebAppDataRequest{}
	_ bin.BareEncoder = &SendWebAppDataRequest{}
	_ bin.BareDecoder = &SendWebAppDataRequest{}
)

func (s *SendWebAppDataRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.BotUserID == 0) {
		return false
	}
	if !(s.ButtonText == "") {
		return false
	}
	if !(s.Data == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendWebAppDataRequest) String() string {
	if s == nil {
		return "SendWebAppDataRequest(nil)"
	}
	type Alias SendWebAppDataRequest
	return fmt.Sprintf("SendWebAppDataRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendWebAppDataRequest) TypeID() uint32 {
	return SendWebAppDataRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendWebAppDataRequest) TypeName() string {
	return "sendWebAppData"
}

// TypeInfo returns info about TL type.
func (s *SendWebAppDataRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendWebAppData",
		ID:   SendWebAppDataRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "ButtonText",
			SchemaName: "button_text",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendWebAppDataRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendWebAppData#ab1fce0c as nil")
	}
	b.PutID(SendWebAppDataRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendWebAppDataRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendWebAppData#ab1fce0c as nil")
	}
	b.PutInt53(s.BotUserID)
	b.PutString(s.ButtonText)
	b.PutString(s.Data)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendWebAppDataRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendWebAppData#ab1fce0c to nil")
	}
	if err := b.ConsumeID(SendWebAppDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendWebAppData#ab1fce0c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendWebAppDataRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendWebAppData#ab1fce0c to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendWebAppData#ab1fce0c: field bot_user_id: %w", err)
		}
		s.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendWebAppData#ab1fce0c: field button_text: %w", err)
		}
		s.ButtonText = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendWebAppData#ab1fce0c: field data: %w", err)
		}
		s.Data = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendWebAppDataRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendWebAppData#ab1fce0c as nil")
	}
	b.ObjStart()
	b.PutID("sendWebAppData")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(s.BotUserID)
	b.Comma()
	b.FieldStart("button_text")
	b.PutString(s.ButtonText)
	b.Comma()
	b.FieldStart("data")
	b.PutString(s.Data)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendWebAppDataRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendWebAppData#ab1fce0c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendWebAppData"); err != nil {
				return fmt.Errorf("unable to decode sendWebAppData#ab1fce0c: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendWebAppData#ab1fce0c: field bot_user_id: %w", err)
			}
			s.BotUserID = value
		case "button_text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sendWebAppData#ab1fce0c: field button_text: %w", err)
			}
			s.ButtonText = value
		case "data":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sendWebAppData#ab1fce0c: field data: %w", err)
			}
			s.Data = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (s *SendWebAppDataRequest) GetBotUserID() (value int64) {
	if s == nil {
		return
	}
	return s.BotUserID
}

// GetButtonText returns value of ButtonText field.
func (s *SendWebAppDataRequest) GetButtonText() (value string) {
	if s == nil {
		return
	}
	return s.ButtonText
}

// GetData returns value of Data field.
func (s *SendWebAppDataRequest) GetData() (value string) {
	if s == nil {
		return
	}
	return s.Data
}

// SendWebAppData invokes method sendWebAppData#ab1fce0c returning error if any.
func (c *Client) SendWebAppData(ctx context.Context, request *SendWebAppDataRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
