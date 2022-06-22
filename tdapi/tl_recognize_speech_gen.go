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

// RecognizeSpeechRequest represents TL type `recognizeSpeech#67d402b9`.
type RecognizeSpeechRequest struct {
	// Identifier of the chat to which the message belongs
	ChatID int64
	// Identifier of the message
	MessageID int64
}

// RecognizeSpeechRequestTypeID is TL type id of RecognizeSpeechRequest.
const RecognizeSpeechRequestTypeID = 0x67d402b9

// Ensuring interfaces in compile-time for RecognizeSpeechRequest.
var (
	_ bin.Encoder     = &RecognizeSpeechRequest{}
	_ bin.Decoder     = &RecognizeSpeechRequest{}
	_ bin.BareEncoder = &RecognizeSpeechRequest{}
	_ bin.BareDecoder = &RecognizeSpeechRequest{}
)

func (r *RecognizeSpeechRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}
	if !(r.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecognizeSpeechRequest) String() string {
	if r == nil {
		return "RecognizeSpeechRequest(nil)"
	}
	type Alias RecognizeSpeechRequest
	return fmt.Sprintf("RecognizeSpeechRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecognizeSpeechRequest) TypeID() uint32 {
	return RecognizeSpeechRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RecognizeSpeechRequest) TypeName() string {
	return "recognizeSpeech"
}

// TypeInfo returns info about TL type.
func (r *RecognizeSpeechRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recognizeSpeech",
		ID:   RecognizeSpeechRequestTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecognizeSpeechRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recognizeSpeech#67d402b9 as nil")
	}
	b.PutID(RecognizeSpeechRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecognizeSpeechRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recognizeSpeech#67d402b9 as nil")
	}
	b.PutInt53(r.ChatID)
	b.PutInt53(r.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecognizeSpeechRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recognizeSpeech#67d402b9 to nil")
	}
	if err := b.ConsumeID(RecognizeSpeechRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode recognizeSpeech#67d402b9: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecognizeSpeechRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recognizeSpeech#67d402b9 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode recognizeSpeech#67d402b9: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode recognizeSpeech#67d402b9: field message_id: %w", err)
		}
		r.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RecognizeSpeechRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode recognizeSpeech#67d402b9 as nil")
	}
	b.ObjStart()
	b.PutID("recognizeSpeech")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(r.MessageID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RecognizeSpeechRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode recognizeSpeech#67d402b9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("recognizeSpeech"); err != nil {
				return fmt.Errorf("unable to decode recognizeSpeech#67d402b9: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode recognizeSpeech#67d402b9: field chat_id: %w", err)
			}
			r.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode recognizeSpeech#67d402b9: field message_id: %w", err)
			}
			r.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *RecognizeSpeechRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// GetMessageID returns value of MessageID field.
func (r *RecognizeSpeechRequest) GetMessageID() (value int64) {
	if r == nil {
		return
	}
	return r.MessageID
}

// RecognizeSpeech invokes method recognizeSpeech#67d402b9 returning error if any.
func (c *Client) RecognizeSpeech(ctx context.Context, request *RecognizeSpeechRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
