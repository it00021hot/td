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

// ChatInviteLinkMember represents TL type `chatInviteLinkMember#1bce5bb`.
type ChatInviteLinkMember struct {
	// User identifier
	UserID int64
	// Point in time (Unix timestamp) when the user joined the chat
	JoinedChatDate int32
	// True, if the user has joined the chat using an invite link for a chat folder
	ViaChatFolderInviteLink bool
	// User identifier of the chat administrator, approved user join request
	ApproverUserID int64
}

// ChatInviteLinkMemberTypeID is TL type id of ChatInviteLinkMember.
const ChatInviteLinkMemberTypeID = 0x1bce5bb

// Ensuring interfaces in compile-time for ChatInviteLinkMember.
var (
	_ bin.Encoder     = &ChatInviteLinkMember{}
	_ bin.Decoder     = &ChatInviteLinkMember{}
	_ bin.BareEncoder = &ChatInviteLinkMember{}
	_ bin.BareDecoder = &ChatInviteLinkMember{}
)

func (c *ChatInviteLinkMember) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.JoinedChatDate == 0) {
		return false
	}
	if !(c.ViaChatFolderInviteLink == false) {
		return false
	}
	if !(c.ApproverUserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteLinkMember) String() string {
	if c == nil {
		return "ChatInviteLinkMember(nil)"
	}
	type Alias ChatInviteLinkMember
	return fmt.Sprintf("ChatInviteLinkMember%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteLinkMember) TypeID() uint32 {
	return ChatInviteLinkMemberTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteLinkMember) TypeName() string {
	return "chatInviteLinkMember"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteLinkMember) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteLinkMember",
		ID:   ChatInviteLinkMemberTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "JoinedChatDate",
			SchemaName: "joined_chat_date",
		},
		{
			Name:       "ViaChatFolderInviteLink",
			SchemaName: "via_chat_folder_invite_link",
		},
		{
			Name:       "ApproverUserID",
			SchemaName: "approver_user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteLinkMember) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkMember#1bce5bb as nil")
	}
	b.PutID(ChatInviteLinkMemberTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteLinkMember) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkMember#1bce5bb as nil")
	}
	b.PutInt53(c.UserID)
	b.PutInt32(c.JoinedChatDate)
	b.PutBool(c.ViaChatFolderInviteLink)
	b.PutInt53(c.ApproverUserID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteLinkMember) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkMember#1bce5bb to nil")
	}
	if err := b.ConsumeID(ChatInviteLinkMemberTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteLinkMember) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkMember#1bce5bb to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: field joined_chat_date: %w", err)
		}
		c.JoinedChatDate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: field via_chat_folder_invite_link: %w", err)
		}
		c.ViaChatFolderInviteLink = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: field approver_user_id: %w", err)
		}
		c.ApproverUserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatInviteLinkMember) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkMember#1bce5bb as nil")
	}
	b.ObjStart()
	b.PutID("chatInviteLinkMember")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.Comma()
	b.FieldStart("joined_chat_date")
	b.PutInt32(c.JoinedChatDate)
	b.Comma()
	b.FieldStart("via_chat_folder_invite_link")
	b.PutBool(c.ViaChatFolderInviteLink)
	b.Comma()
	b.FieldStart("approver_user_id")
	b.PutInt53(c.ApproverUserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatInviteLinkMember) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkMember#1bce5bb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatInviteLinkMember"); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: field user_id: %w", err)
			}
			c.UserID = value
		case "joined_chat_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: field joined_chat_date: %w", err)
			}
			c.JoinedChatDate = value
		case "via_chat_folder_invite_link":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: field via_chat_folder_invite_link: %w", err)
			}
			c.ViaChatFolderInviteLink = value
		case "approver_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkMember#1bce5bb: field approver_user_id: %w", err)
			}
			c.ApproverUserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *ChatInviteLinkMember) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// GetJoinedChatDate returns value of JoinedChatDate field.
func (c *ChatInviteLinkMember) GetJoinedChatDate() (value int32) {
	if c == nil {
		return
	}
	return c.JoinedChatDate
}

// GetViaChatFolderInviteLink returns value of ViaChatFolderInviteLink field.
func (c *ChatInviteLinkMember) GetViaChatFolderInviteLink() (value bool) {
	if c == nil {
		return
	}
	return c.ViaChatFolderInviteLink
}

// GetApproverUserID returns value of ApproverUserID field.
func (c *ChatInviteLinkMember) GetApproverUserID() (value int64) {
	if c == nil {
		return
	}
	return c.ApproverUserID
}
