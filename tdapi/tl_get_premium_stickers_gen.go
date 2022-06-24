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

// GetPremiumStickersRequest represents TL type `getPremiumStickers#8a2b9a2`.
type GetPremiumStickersRequest struct {
}

// GetPremiumStickersRequestTypeID is TL type id of GetPremiumStickersRequest.
const GetPremiumStickersRequestTypeID = 0x8a2b9a2

// Ensuring interfaces in compile-time for GetPremiumStickersRequest.
var (
	_ bin.Encoder     = &GetPremiumStickersRequest{}
	_ bin.Decoder     = &GetPremiumStickersRequest{}
	_ bin.BareEncoder = &GetPremiumStickersRequest{}
	_ bin.BareDecoder = &GetPremiumStickersRequest{}
)

func (g *GetPremiumStickersRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPremiumStickersRequest) String() string {
	if g == nil {
		return "GetPremiumStickersRequest(nil)"
	}
	type Alias GetPremiumStickersRequest
	return fmt.Sprintf("GetPremiumStickersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPremiumStickersRequest) TypeID() uint32 {
	return GetPremiumStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPremiumStickersRequest) TypeName() string {
	return "getPremiumStickers"
}

// TypeInfo returns info about TL type.
func (g *GetPremiumStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPremiumStickers",
		ID:   GetPremiumStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPremiumStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumStickers#8a2b9a2 as nil")
	}
	b.PutID(GetPremiumStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPremiumStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumStickers#8a2b9a2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPremiumStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumStickers#8a2b9a2 to nil")
	}
	if err := b.ConsumeID(GetPremiumStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPremiumStickers#8a2b9a2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPremiumStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumStickers#8a2b9a2 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetPremiumStickersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumStickers#8a2b9a2 as nil")
	}
	b.ObjStart()
	b.PutID("getPremiumStickers")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetPremiumStickersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumStickers#8a2b9a2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getPremiumStickers"); err != nil {
				return fmt.Errorf("unable to decode getPremiumStickers#8a2b9a2: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPremiumStickers invokes method getPremiumStickers#8a2b9a2 returning error if any.
func (c *Client) GetPremiumStickers(ctx context.Context) (*Stickers, error) {
	var result Stickers

	request := &GetPremiumStickersRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}