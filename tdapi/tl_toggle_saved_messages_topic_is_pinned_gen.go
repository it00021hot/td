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

// ToggleSavedMessagesTopicIsPinnedRequest represents TL type `toggleSavedMessagesTopicIsPinned#3f8a8c2d`.
type ToggleSavedMessagesTopicIsPinnedRequest struct {
	// Saved Messages topic to pin or unpin
	SavedMessagesTopic SavedMessagesTopicClass
	// Pass true to pin the topic; pass false to unpin it
	IsPinned bool
}

// ToggleSavedMessagesTopicIsPinnedRequestTypeID is TL type id of ToggleSavedMessagesTopicIsPinnedRequest.
const ToggleSavedMessagesTopicIsPinnedRequestTypeID = 0x3f8a8c2d

// Ensuring interfaces in compile-time for ToggleSavedMessagesTopicIsPinnedRequest.
var (
	_ bin.Encoder     = &ToggleSavedMessagesTopicIsPinnedRequest{}
	_ bin.Decoder     = &ToggleSavedMessagesTopicIsPinnedRequest{}
	_ bin.BareEncoder = &ToggleSavedMessagesTopicIsPinnedRequest{}
	_ bin.BareDecoder = &ToggleSavedMessagesTopicIsPinnedRequest{}
)

func (t *ToggleSavedMessagesTopicIsPinnedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SavedMessagesTopic == nil) {
		return false
	}
	if !(t.IsPinned == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) String() string {
	if t == nil {
		return "ToggleSavedMessagesTopicIsPinnedRequest(nil)"
	}
	type Alias ToggleSavedMessagesTopicIsPinnedRequest
	return fmt.Sprintf("ToggleSavedMessagesTopicIsPinnedRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleSavedMessagesTopicIsPinnedRequest) TypeID() uint32 {
	return ToggleSavedMessagesTopicIsPinnedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleSavedMessagesTopicIsPinnedRequest) TypeName() string {
	return "toggleSavedMessagesTopicIsPinned"
}

// TypeInfo returns info about TL type.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleSavedMessagesTopicIsPinned",
		ID:   ToggleSavedMessagesTopicIsPinnedRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SavedMessagesTopic",
			SchemaName: "saved_messages_topic",
		},
		{
			Name:       "IsPinned",
			SchemaName: "is_pinned",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSavedMessagesTopicIsPinned#3f8a8c2d as nil")
	}
	b.PutID(ToggleSavedMessagesTopicIsPinnedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSavedMessagesTopicIsPinned#3f8a8c2d as nil")
	}
	if t.SavedMessagesTopic == nil {
		return fmt.Errorf("unable to encode toggleSavedMessagesTopicIsPinned#3f8a8c2d: field saved_messages_topic is nil")
	}
	if err := t.SavedMessagesTopic.Encode(b); err != nil {
		return fmt.Errorf("unable to encode toggleSavedMessagesTopicIsPinned#3f8a8c2d: field saved_messages_topic: %w", err)
	}
	b.PutBool(t.IsPinned)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSavedMessagesTopicIsPinned#3f8a8c2d to nil")
	}
	if err := b.ConsumeID(ToggleSavedMessagesTopicIsPinnedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleSavedMessagesTopicIsPinned#3f8a8c2d: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSavedMessagesTopicIsPinned#3f8a8c2d to nil")
	}
	{
		value, err := DecodeSavedMessagesTopic(b)
		if err != nil {
			return fmt.Errorf("unable to decode toggleSavedMessagesTopicIsPinned#3f8a8c2d: field saved_messages_topic: %w", err)
		}
		t.SavedMessagesTopic = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSavedMessagesTopicIsPinned#3f8a8c2d: field is_pinned: %w", err)
		}
		t.IsPinned = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSavedMessagesTopicIsPinned#3f8a8c2d as nil")
	}
	b.ObjStart()
	b.PutID("toggleSavedMessagesTopicIsPinned")
	b.Comma()
	b.FieldStart("saved_messages_topic")
	if t.SavedMessagesTopic == nil {
		return fmt.Errorf("unable to encode toggleSavedMessagesTopicIsPinned#3f8a8c2d: field saved_messages_topic is nil")
	}
	if err := t.SavedMessagesTopic.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode toggleSavedMessagesTopicIsPinned#3f8a8c2d: field saved_messages_topic: %w", err)
	}
	b.Comma()
	b.FieldStart("is_pinned")
	b.PutBool(t.IsPinned)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSavedMessagesTopicIsPinned#3f8a8c2d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleSavedMessagesTopicIsPinned"); err != nil {
				return fmt.Errorf("unable to decode toggleSavedMessagesTopicIsPinned#3f8a8c2d: %w", err)
			}
		case "saved_messages_topic":
			value, err := DecodeTDLibJSONSavedMessagesTopic(b)
			if err != nil {
				return fmt.Errorf("unable to decode toggleSavedMessagesTopicIsPinned#3f8a8c2d: field saved_messages_topic: %w", err)
			}
			t.SavedMessagesTopic = value
		case "is_pinned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSavedMessagesTopicIsPinned#3f8a8c2d: field is_pinned: %w", err)
			}
			t.IsPinned = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSavedMessagesTopic returns value of SavedMessagesTopic field.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) GetSavedMessagesTopic() (value SavedMessagesTopicClass) {
	if t == nil {
		return
	}
	return t.SavedMessagesTopic
}

// GetIsPinned returns value of IsPinned field.
func (t *ToggleSavedMessagesTopicIsPinnedRequest) GetIsPinned() (value bool) {
	if t == nil {
		return
	}
	return t.IsPinned
}

// ToggleSavedMessagesTopicIsPinned invokes method toggleSavedMessagesTopicIsPinned#3f8a8c2d returning error if any.
func (c *Client) ToggleSavedMessagesTopicIsPinned(ctx context.Context, request *ToggleSavedMessagesTopicIsPinnedRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}