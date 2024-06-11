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

// SearchPublicHashtagMessagesRequest represents TL type `searchPublicHashtagMessages#2c6b4c0b`.
type SearchPublicHashtagMessagesRequest struct {
	// Hashtag or cashtag to search for
	Hashtag string
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get the first chunk of results
	Offset string
	// The maximum number of messages to be returned; up to 100. For optimal performance, the
	// number of returned messages is chosen by TDLib and can be smaller than the specified
	// limit
	Limit int32
}

// SearchPublicHashtagMessagesRequestTypeID is TL type id of SearchPublicHashtagMessagesRequest.
const SearchPublicHashtagMessagesRequestTypeID = 0x2c6b4c0b

// Ensuring interfaces in compile-time for SearchPublicHashtagMessagesRequest.
var (
	_ bin.Encoder     = &SearchPublicHashtagMessagesRequest{}
	_ bin.Decoder     = &SearchPublicHashtagMessagesRequest{}
	_ bin.BareEncoder = &SearchPublicHashtagMessagesRequest{}
	_ bin.BareDecoder = &SearchPublicHashtagMessagesRequest{}
)

func (s *SearchPublicHashtagMessagesRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Hashtag == "") {
		return false
	}
	if !(s.Offset == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchPublicHashtagMessagesRequest) String() string {
	if s == nil {
		return "SearchPublicHashtagMessagesRequest(nil)"
	}
	type Alias SearchPublicHashtagMessagesRequest
	return fmt.Sprintf("SearchPublicHashtagMessagesRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchPublicHashtagMessagesRequest) TypeID() uint32 {
	return SearchPublicHashtagMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchPublicHashtagMessagesRequest) TypeName() string {
	return "searchPublicHashtagMessages"
}

// TypeInfo returns info about TL type.
func (s *SearchPublicHashtagMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchPublicHashtagMessages",
		ID:   SearchPublicHashtagMessagesRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hashtag",
			SchemaName: "hashtag",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchPublicHashtagMessagesRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchPublicHashtagMessages#2c6b4c0b as nil")
	}
	b.PutID(SearchPublicHashtagMessagesRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchPublicHashtagMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchPublicHashtagMessages#2c6b4c0b as nil")
	}
	b.PutString(s.Hashtag)
	b.PutString(s.Offset)
	b.PutInt32(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchPublicHashtagMessagesRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchPublicHashtagMessages#2c6b4c0b to nil")
	}
	if err := b.ConsumeID(SearchPublicHashtagMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchPublicHashtagMessages#2c6b4c0b: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchPublicHashtagMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchPublicHashtagMessages#2c6b4c0b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchPublicHashtagMessages#2c6b4c0b: field hashtag: %w", err)
		}
		s.Hashtag = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchPublicHashtagMessages#2c6b4c0b: field offset: %w", err)
		}
		s.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchPublicHashtagMessages#2c6b4c0b: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchPublicHashtagMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchPublicHashtagMessages#2c6b4c0b as nil")
	}
	b.ObjStart()
	b.PutID("searchPublicHashtagMessages")
	b.Comma()
	b.FieldStart("hashtag")
	b.PutString(s.Hashtag)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(s.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchPublicHashtagMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchPublicHashtagMessages#2c6b4c0b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchPublicHashtagMessages"); err != nil {
				return fmt.Errorf("unable to decode searchPublicHashtagMessages#2c6b4c0b: %w", err)
			}
		case "hashtag":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchPublicHashtagMessages#2c6b4c0b: field hashtag: %w", err)
			}
			s.Hashtag = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchPublicHashtagMessages#2c6b4c0b: field offset: %w", err)
			}
			s.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchPublicHashtagMessages#2c6b4c0b: field limit: %w", err)
			}
			s.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetHashtag returns value of Hashtag field.
func (s *SearchPublicHashtagMessagesRequest) GetHashtag() (value string) {
	if s == nil {
		return
	}
	return s.Hashtag
}

// GetOffset returns value of Offset field.
func (s *SearchPublicHashtagMessagesRequest) GetOffset() (value string) {
	if s == nil {
		return
	}
	return s.Offset
}

// GetLimit returns value of Limit field.
func (s *SearchPublicHashtagMessagesRequest) GetLimit() (value int32) {
	if s == nil {
		return
	}
	return s.Limit
}

// SearchPublicHashtagMessages invokes method searchPublicHashtagMessages#2c6b4c0b returning error if any.
func (c *Client) SearchPublicHashtagMessages(ctx context.Context, request *SearchPublicHashtagMessagesRequest) (*FoundMessages, error) {
	var result FoundMessages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
