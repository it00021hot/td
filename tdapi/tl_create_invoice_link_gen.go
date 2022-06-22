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

// CreateInvoiceLinkRequest represents TL type `createInvoiceLink#cebe921`.
type CreateInvoiceLinkRequest struct {
	// Information about the invoice of the type inputMessageInvoice
	Invoice InputMessageContentClass
}

// CreateInvoiceLinkRequestTypeID is TL type id of CreateInvoiceLinkRequest.
const CreateInvoiceLinkRequestTypeID = 0xcebe921

// Ensuring interfaces in compile-time for CreateInvoiceLinkRequest.
var (
	_ bin.Encoder     = &CreateInvoiceLinkRequest{}
	_ bin.Decoder     = &CreateInvoiceLinkRequest{}
	_ bin.BareEncoder = &CreateInvoiceLinkRequest{}
	_ bin.BareDecoder = &CreateInvoiceLinkRequest{}
)

func (c *CreateInvoiceLinkRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Invoice == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreateInvoiceLinkRequest) String() string {
	if c == nil {
		return "CreateInvoiceLinkRequest(nil)"
	}
	type Alias CreateInvoiceLinkRequest
	return fmt.Sprintf("CreateInvoiceLinkRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreateInvoiceLinkRequest) TypeID() uint32 {
	return CreateInvoiceLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreateInvoiceLinkRequest) TypeName() string {
	return "createInvoiceLink"
}

// TypeInfo returns info about TL type.
func (c *CreateInvoiceLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createInvoiceLink",
		ID:   CreateInvoiceLinkRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Invoice",
			SchemaName: "invoice",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreateInvoiceLinkRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createInvoiceLink#cebe921 as nil")
	}
	b.PutID(CreateInvoiceLinkRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreateInvoiceLinkRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createInvoiceLink#cebe921 as nil")
	}
	if c.Invoice == nil {
		return fmt.Errorf("unable to encode createInvoiceLink#cebe921: field invoice is nil")
	}
	if err := c.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode createInvoiceLink#cebe921: field invoice: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CreateInvoiceLinkRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createInvoiceLink#cebe921 to nil")
	}
	if err := b.ConsumeID(CreateInvoiceLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createInvoiceLink#cebe921: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreateInvoiceLinkRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createInvoiceLink#cebe921 to nil")
	}
	{
		value, err := DecodeInputMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode createInvoiceLink#cebe921: field invoice: %w", err)
		}
		c.Invoice = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CreateInvoiceLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode createInvoiceLink#cebe921 as nil")
	}
	b.ObjStart()
	b.PutID("createInvoiceLink")
	b.Comma()
	b.FieldStart("invoice")
	if c.Invoice == nil {
		return fmt.Errorf("unable to encode createInvoiceLink#cebe921: field invoice is nil")
	}
	if err := c.Invoice.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode createInvoiceLink#cebe921: field invoice: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CreateInvoiceLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode createInvoiceLink#cebe921 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("createInvoiceLink"); err != nil {
				return fmt.Errorf("unable to decode createInvoiceLink#cebe921: %w", err)
			}
		case "invoice":
			value, err := DecodeTDLibJSONInputMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode createInvoiceLink#cebe921: field invoice: %w", err)
			}
			c.Invoice = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInvoice returns value of Invoice field.
func (c *CreateInvoiceLinkRequest) GetInvoice() (value InputMessageContentClass) {
	if c == nil {
		return
	}
	return c.Invoice
}

// CreateInvoiceLink invokes method createInvoiceLink#cebe921 returning error if any.
func (c *Client) CreateInvoiceLink(ctx context.Context, invoice InputMessageContentClass) (*HTTPURL, error) {
	var result HTTPURL

	request := &CreateInvoiceLinkRequest{
		Invoice: invoice,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
