// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// SMS represents TL type sms#ed8bebfe.
type SMS struct {
	// Text field of SMS.
	Text string
}

// SMSTypeID is TL type id of SMS.
const SMSTypeID = 0xed8bebfe

// Encode implements bin.Encoder.
func (s *SMS) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sms#ed8bebfe as nil")
	}
	b.PutID(SMSTypeID)
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SMS) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sms#ed8bebfe to nil")
	}
	if err := b.ConsumeID(SMSTypeID); err != nil {
		return fmt.Errorf("unable to decode sms#ed8bebfe: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sms#ed8bebfe: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SMS.
var (
	_ bin.Encoder = &SMS{}
	_ bin.Decoder = &SMS{}
)