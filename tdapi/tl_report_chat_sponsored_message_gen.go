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

// ReportChatSponsoredMessageRequest represents TL type `reportChatSponsoredMessage#cc3e53be`.
type ReportChatSponsoredMessageRequest struct {
	// Chat identifier of the sponsored message
	ChatID int64
	// Identifier of the sponsored message
	MessageID int64
	// Option identifier chosen by the user; leave empty for the initial request
	OptionID []byte
}

// ReportChatSponsoredMessageRequestTypeID is TL type id of ReportChatSponsoredMessageRequest.
const ReportChatSponsoredMessageRequestTypeID = 0xcc3e53be

// Ensuring interfaces in compile-time for ReportChatSponsoredMessageRequest.
var (
	_ bin.Encoder     = &ReportChatSponsoredMessageRequest{}
	_ bin.Decoder     = &ReportChatSponsoredMessageRequest{}
	_ bin.BareEncoder = &ReportChatSponsoredMessageRequest{}
	_ bin.BareDecoder = &ReportChatSponsoredMessageRequest{}
)

func (r *ReportChatSponsoredMessageRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}
	if !(r.MessageID == 0) {
		return false
	}
	if !(r.OptionID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatSponsoredMessageRequest) String() string {
	if r == nil {
		return "ReportChatSponsoredMessageRequest(nil)"
	}
	type Alias ReportChatSponsoredMessageRequest
	return fmt.Sprintf("ReportChatSponsoredMessageRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatSponsoredMessageRequest) TypeID() uint32 {
	return ReportChatSponsoredMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatSponsoredMessageRequest) TypeName() string {
	return "reportChatSponsoredMessage"
}

// TypeInfo returns info about TL type.
func (r *ReportChatSponsoredMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChatSponsoredMessage",
		ID:   ReportChatSponsoredMessageRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "OptionID",
			SchemaName: "option_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatSponsoredMessageRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessage#cc3e53be as nil")
	}
	b.PutID(ReportChatSponsoredMessageRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatSponsoredMessageRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessage#cc3e53be as nil")
	}
	b.PutInt53(r.ChatID)
	b.PutInt53(r.MessageID)
	b.PutBytes(r.OptionID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatSponsoredMessageRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessage#cc3e53be to nil")
	}
	if err := b.ConsumeID(ReportChatSponsoredMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChatSponsoredMessage#cc3e53be: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatSponsoredMessageRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessage#cc3e53be to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatSponsoredMessage#cc3e53be: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatSponsoredMessage#cc3e53be: field message_id: %w", err)
		}
		r.MessageID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatSponsoredMessage#cc3e53be: field option_id: %w", err)
		}
		r.OptionID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportChatSponsoredMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessage#cc3e53be as nil")
	}
	b.ObjStart()
	b.PutID("reportChatSponsoredMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(r.MessageID)
	b.Comma()
	b.FieldStart("option_id")
	b.PutBytes(r.OptionID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportChatSponsoredMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessage#cc3e53be to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportChatSponsoredMessage"); err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessage#cc3e53be: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessage#cc3e53be: field chat_id: %w", err)
			}
			r.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessage#cc3e53be: field message_id: %w", err)
			}
			r.MessageID = value
		case "option_id":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessage#cc3e53be: field option_id: %w", err)
			}
			r.OptionID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *ReportChatSponsoredMessageRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// GetMessageID returns value of MessageID field.
func (r *ReportChatSponsoredMessageRequest) GetMessageID() (value int64) {
	if r == nil {
		return
	}
	return r.MessageID
}

// GetOptionID returns value of OptionID field.
func (r *ReportChatSponsoredMessageRequest) GetOptionID() (value []byte) {
	if r == nil {
		return
	}
	return r.OptionID
}

// ReportChatSponsoredMessage invokes method reportChatSponsoredMessage#cc3e53be returning error if any.
func (c *Client) ReportChatSponsoredMessage(ctx context.Context, request *ReportChatSponsoredMessageRequest) (ReportChatSponsoredMessageResultClass, error) {
	var result ReportChatSponsoredMessageResultBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ReportChatSponsoredMessageResult, nil
}