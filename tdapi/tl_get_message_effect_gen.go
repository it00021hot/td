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

// GetMessageEffectRequest represents TL type `getMessageEffect#9e513d14`.
type GetMessageEffectRequest struct {
	// Unique identifier of the effect
	EffectID int64
}

// GetMessageEffectRequestTypeID is TL type id of GetMessageEffectRequest.
const GetMessageEffectRequestTypeID = 0x9e513d14

// Ensuring interfaces in compile-time for GetMessageEffectRequest.
var (
	_ bin.Encoder     = &GetMessageEffectRequest{}
	_ bin.Decoder     = &GetMessageEffectRequest{}
	_ bin.BareEncoder = &GetMessageEffectRequest{}
	_ bin.BareDecoder = &GetMessageEffectRequest{}
)

func (g *GetMessageEffectRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.EffectID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessageEffectRequest) String() string {
	if g == nil {
		return "GetMessageEffectRequest(nil)"
	}
	type Alias GetMessageEffectRequest
	return fmt.Sprintf("GetMessageEffectRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageEffectRequest) TypeID() uint32 {
	return GetMessageEffectRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageEffectRequest) TypeName() string {
	return "getMessageEffect"
}

// TypeInfo returns info about TL type.
func (g *GetMessageEffectRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessageEffect",
		ID:   GetMessageEffectRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EffectID",
			SchemaName: "effect_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMessageEffectRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageEffect#9e513d14 as nil")
	}
	b.PutID(GetMessageEffectRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageEffectRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageEffect#9e513d14 as nil")
	}
	b.PutLong(g.EffectID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageEffectRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageEffect#9e513d14 to nil")
	}
	if err := b.ConsumeID(GetMessageEffectRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessageEffect#9e513d14: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageEffectRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageEffect#9e513d14 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageEffect#9e513d14: field effect_id: %w", err)
		}
		g.EffectID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessageEffectRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageEffect#9e513d14 as nil")
	}
	b.ObjStart()
	b.PutID("getMessageEffect")
	b.Comma()
	b.FieldStart("effect_id")
	b.PutLong(g.EffectID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMessageEffectRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageEffect#9e513d14 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessageEffect"); err != nil {
				return fmt.Errorf("unable to decode getMessageEffect#9e513d14: %w", err)
			}
		case "effect_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageEffect#9e513d14: field effect_id: %w", err)
			}
			g.EffectID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEffectID returns value of EffectID field.
func (g *GetMessageEffectRequest) GetEffectID() (value int64) {
	if g == nil {
		return
	}
	return g.EffectID
}

// GetMessageEffect invokes method getMessageEffect#9e513d14 returning error if any.
func (c *Client) GetMessageEffect(ctx context.Context, effectid int64) (*MessageEffect, error) {
	var result MessageEffect

	request := &GetMessageEffectRequest{
		EffectID: effectid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
