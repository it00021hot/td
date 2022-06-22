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

// AvailableReactions represents TL type `availableReactions#58fdcd12`.
type AvailableReactions struct {
	// List of reactions
	Reactions []AvailableReaction
}

// AvailableReactionsTypeID is TL type id of AvailableReactions.
const AvailableReactionsTypeID = 0x58fdcd12

// Ensuring interfaces in compile-time for AvailableReactions.
var (
	_ bin.Encoder     = &AvailableReactions{}
	_ bin.Decoder     = &AvailableReactions{}
	_ bin.BareEncoder = &AvailableReactions{}
	_ bin.BareDecoder = &AvailableReactions{}
)

func (a *AvailableReactions) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Reactions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AvailableReactions) String() string {
	if a == nil {
		return "AvailableReactions(nil)"
	}
	type Alias AvailableReactions
	return fmt.Sprintf("AvailableReactions%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AvailableReactions) TypeID() uint32 {
	return AvailableReactionsTypeID
}

// TypeName returns name of type in TL schema.
func (*AvailableReactions) TypeName() string {
	return "availableReactions"
}

// TypeInfo returns info about TL type.
func (a *AvailableReactions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "availableReactions",
		ID:   AvailableReactionsTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Reactions",
			SchemaName: "reactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AvailableReactions) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode availableReactions#58fdcd12 as nil")
	}
	b.PutID(AvailableReactionsTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AvailableReactions) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode availableReactions#58fdcd12 as nil")
	}
	b.PutInt(len(a.Reactions))
	for idx, v := range a.Reactions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare availableReactions#58fdcd12: field reactions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AvailableReactions) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode availableReactions#58fdcd12 to nil")
	}
	if err := b.ConsumeID(AvailableReactionsTypeID); err != nil {
		return fmt.Errorf("unable to decode availableReactions#58fdcd12: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AvailableReactions) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode availableReactions#58fdcd12 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode availableReactions#58fdcd12: field reactions: %w", err)
		}

		if headerLen > 0 {
			a.Reactions = make([]AvailableReaction, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value AvailableReaction
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare availableReactions#58fdcd12: field reactions: %w", err)
			}
			a.Reactions = append(a.Reactions, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AvailableReactions) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode availableReactions#58fdcd12 as nil")
	}
	b.ObjStart()
	b.PutID("availableReactions")
	b.Comma()
	b.FieldStart("reactions")
	b.ArrStart()
	for idx, v := range a.Reactions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode availableReactions#58fdcd12: field reactions element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AvailableReactions) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode availableReactions#58fdcd12 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("availableReactions"); err != nil {
				return fmt.Errorf("unable to decode availableReactions#58fdcd12: %w", err)
			}
		case "reactions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value AvailableReaction
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode availableReactions#58fdcd12: field reactions: %w", err)
				}
				a.Reactions = append(a.Reactions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode availableReactions#58fdcd12: field reactions: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReactions returns value of Reactions field.
func (a *AvailableReactions) GetReactions() (value []AvailableReaction) {
	if a == nil {
		return
	}
	return a.Reactions
}
