// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// HelpSaveAppLogRequest represents TL type `help.saveAppLog#6f02f748`.
// Saves logs of application on the server.
//
// See https://core.telegram.org/method/help.saveAppLog for reference.
type HelpSaveAppLogRequest struct {
	// List of input events
	Events []InputAppEvent
}

// HelpSaveAppLogRequestTypeID is TL type id of HelpSaveAppLogRequest.
const HelpSaveAppLogRequestTypeID = 0x6f02f748

// Encode implements bin.Encoder.
func (s *HelpSaveAppLogRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.saveAppLog#6f02f748 as nil")
	}
	b.PutID(HelpSaveAppLogRequestTypeID)
	b.PutVectorHeader(len(s.Events))
	for idx, v := range s.Events {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.saveAppLog#6f02f748: field events element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *HelpSaveAppLogRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.saveAppLog#6f02f748 to nil")
	}
	if err := b.ConsumeID(HelpSaveAppLogRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.saveAppLog#6f02f748: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.saveAppLog#6f02f748: field events: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputAppEvent
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode help.saveAppLog#6f02f748: field events: %w", err)
			}
			s.Events = append(s.Events, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpSaveAppLogRequest.
var (
	_ bin.Encoder = &HelpSaveAppLogRequest{}
	_ bin.Decoder = &HelpSaveAppLogRequest{}
)

// HelpSaveAppLog invokes method help.saveAppLog#6f02f748 returning error if any.
// Saves logs of application on the server.
//
// See https://core.telegram.org/method/help.saveAppLog for reference.
func (c *Client) HelpSaveAppLog(ctx context.Context, request *HelpSaveAppLogRequest) (bool, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
