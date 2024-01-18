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

// SetChatEmojiStatusRequest represents TL type `setChatEmojiStatus#55881912`.
type SetChatEmojiStatusRequest struct {
	// Chat identifier
	ChatID int64
	// New emoji status; pass null to remove emoji status
	EmojiStatus EmojiStatus
}

// SetChatEmojiStatusRequestTypeID is TL type id of SetChatEmojiStatusRequest.
const SetChatEmojiStatusRequestTypeID = 0x55881912

// Ensuring interfaces in compile-time for SetChatEmojiStatusRequest.
var (
	_ bin.Encoder     = &SetChatEmojiStatusRequest{}
	_ bin.Decoder     = &SetChatEmojiStatusRequest{}
	_ bin.BareEncoder = &SetChatEmojiStatusRequest{}
	_ bin.BareDecoder = &SetChatEmojiStatusRequest{}
)

func (s *SetChatEmojiStatusRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.EmojiStatus.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatEmojiStatusRequest) String() string {
	if s == nil {
		return "SetChatEmojiStatusRequest(nil)"
	}
	type Alias SetChatEmojiStatusRequest
	return fmt.Sprintf("SetChatEmojiStatusRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatEmojiStatusRequest) TypeID() uint32 {
	return SetChatEmojiStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatEmojiStatusRequest) TypeName() string {
	return "setChatEmojiStatus"
}

// TypeInfo returns info about TL type.
func (s *SetChatEmojiStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatEmojiStatus",
		ID:   SetChatEmojiStatusRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "EmojiStatus",
			SchemaName: "emoji_status",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatEmojiStatusRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatEmojiStatus#55881912 as nil")
	}
	b.PutID(SetChatEmojiStatusRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatEmojiStatusRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatEmojiStatus#55881912 as nil")
	}
	b.PutInt53(s.ChatID)
	if err := s.EmojiStatus.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setChatEmojiStatus#55881912: field emoji_status: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatEmojiStatusRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatEmojiStatus#55881912 to nil")
	}
	if err := b.ConsumeID(SetChatEmojiStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatEmojiStatus#55881912: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatEmojiStatusRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatEmojiStatus#55881912 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatEmojiStatus#55881912: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		if err := s.EmojiStatus.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setChatEmojiStatus#55881912: field emoji_status: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatEmojiStatusRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatEmojiStatus#55881912 as nil")
	}
	b.ObjStart()
	b.PutID("setChatEmojiStatus")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("emoji_status")
	if err := s.EmojiStatus.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setChatEmojiStatus#55881912: field emoji_status: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatEmojiStatusRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatEmojiStatus#55881912 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatEmojiStatus"); err != nil {
				return fmt.Errorf("unable to decode setChatEmojiStatus#55881912: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatEmojiStatus#55881912: field chat_id: %w", err)
			}
			s.ChatID = value
		case "emoji_status":
			if err := s.EmojiStatus.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setChatEmojiStatus#55881912: field emoji_status: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatEmojiStatusRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetEmojiStatus returns value of EmojiStatus field.
func (s *SetChatEmojiStatusRequest) GetEmojiStatus() (value EmojiStatus) {
	if s == nil {
		return
	}
	return s.EmojiStatus
}

// SetChatEmojiStatus invokes method setChatEmojiStatus#55881912 returning error if any.
func (c *Client) SetChatEmojiStatus(ctx context.Context, request *SetChatEmojiStatusRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}