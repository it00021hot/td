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

// CanPurchaseFromStoreRequest represents TL type `canPurchaseFromStore#3caa9368`.
type CanPurchaseFromStoreRequest struct {
	// Transaction purpose
	Purpose StorePaymentPurposeClass
}

// CanPurchaseFromStoreRequestTypeID is TL type id of CanPurchaseFromStoreRequest.
const CanPurchaseFromStoreRequestTypeID = 0x3caa9368

// Ensuring interfaces in compile-time for CanPurchaseFromStoreRequest.
var (
	_ bin.Encoder     = &CanPurchaseFromStoreRequest{}
	_ bin.Decoder     = &CanPurchaseFromStoreRequest{}
	_ bin.BareEncoder = &CanPurchaseFromStoreRequest{}
	_ bin.BareDecoder = &CanPurchaseFromStoreRequest{}
)

func (c *CanPurchaseFromStoreRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Purpose == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CanPurchaseFromStoreRequest) String() string {
	if c == nil {
		return "CanPurchaseFromStoreRequest(nil)"
	}
	type Alias CanPurchaseFromStoreRequest
	return fmt.Sprintf("CanPurchaseFromStoreRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CanPurchaseFromStoreRequest) TypeID() uint32 {
	return CanPurchaseFromStoreRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CanPurchaseFromStoreRequest) TypeName() string {
	return "canPurchaseFromStore"
}

// TypeInfo returns info about TL type.
func (c *CanPurchaseFromStoreRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "canPurchaseFromStore",
		ID:   CanPurchaseFromStoreRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Purpose",
			SchemaName: "purpose",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CanPurchaseFromStoreRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode canPurchaseFromStore#3caa9368 as nil")
	}
	b.PutID(CanPurchaseFromStoreRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CanPurchaseFromStoreRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode canPurchaseFromStore#3caa9368 as nil")
	}
	if c.Purpose == nil {
		return fmt.Errorf("unable to encode canPurchaseFromStore#3caa9368: field purpose is nil")
	}
	if err := c.Purpose.Encode(b); err != nil {
		return fmt.Errorf("unable to encode canPurchaseFromStore#3caa9368: field purpose: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CanPurchaseFromStoreRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode canPurchaseFromStore#3caa9368 to nil")
	}
	if err := b.ConsumeID(CanPurchaseFromStoreRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode canPurchaseFromStore#3caa9368: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CanPurchaseFromStoreRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode canPurchaseFromStore#3caa9368 to nil")
	}
	{
		value, err := DecodeStorePaymentPurpose(b)
		if err != nil {
			return fmt.Errorf("unable to decode canPurchaseFromStore#3caa9368: field purpose: %w", err)
		}
		c.Purpose = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CanPurchaseFromStoreRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode canPurchaseFromStore#3caa9368 as nil")
	}
	b.ObjStart()
	b.PutID("canPurchaseFromStore")
	b.Comma()
	b.FieldStart("purpose")
	if c.Purpose == nil {
		return fmt.Errorf("unable to encode canPurchaseFromStore#3caa9368: field purpose is nil")
	}
	if err := c.Purpose.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode canPurchaseFromStore#3caa9368: field purpose: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CanPurchaseFromStoreRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode canPurchaseFromStore#3caa9368 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("canPurchaseFromStore"); err != nil {
				return fmt.Errorf("unable to decode canPurchaseFromStore#3caa9368: %w", err)
			}
		case "purpose":
			value, err := DecodeTDLibJSONStorePaymentPurpose(b)
			if err != nil {
				return fmt.Errorf("unable to decode canPurchaseFromStore#3caa9368: field purpose: %w", err)
			}
			c.Purpose = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPurpose returns value of Purpose field.
func (c *CanPurchaseFromStoreRequest) GetPurpose() (value StorePaymentPurposeClass) {
	if c == nil {
		return
	}
	return c.Purpose
}

// CanPurchaseFromStore invokes method canPurchaseFromStore#3caa9368 returning error if any.
func (c *Client) CanPurchaseFromStore(ctx context.Context, purpose StorePaymentPurposeClass) error {
	var ok Ok

	request := &CanPurchaseFromStoreRequest{
		Purpose: purpose,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
