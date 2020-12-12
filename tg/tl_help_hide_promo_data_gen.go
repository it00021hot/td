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

// HelpHidePromoDataRequest represents TL type `help.hidePromoData#1e251c95`.
// Hide MTProxy/Public Service Announcement information
//
// See https://core.telegram.org/method/help.hidePromoData for reference.
type HelpHidePromoDataRequest struct {
	// Peer to hide
	Peer InputPeerClass
}

// HelpHidePromoDataRequestTypeID is TL type id of HelpHidePromoDataRequest.
const HelpHidePromoDataRequestTypeID = 0x1e251c95

// Encode implements bin.Encoder.
func (h *HelpHidePromoDataRequest) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode help.hidePromoData#1e251c95 as nil")
	}
	b.PutID(HelpHidePromoDataRequestTypeID)
	if h.Peer == nil {
		return fmt.Errorf("unable to encode help.hidePromoData#1e251c95: field peer is nil")
	}
	if err := h.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.hidePromoData#1e251c95: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (h *HelpHidePromoDataRequest) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode help.hidePromoData#1e251c95 to nil")
	}
	if err := b.ConsumeID(HelpHidePromoDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.hidePromoData#1e251c95: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.hidePromoData#1e251c95: field peer: %w", err)
		}
		h.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpHidePromoDataRequest.
var (
	_ bin.Encoder = &HelpHidePromoDataRequest{}
	_ bin.Decoder = &HelpHidePromoDataRequest{}
)

// HelpHidePromoData invokes method help.hidePromoData#1e251c95 returning error if any.
// Hide MTProxy/Public Service Announcement information
//
// See https://core.telegram.org/method/help.hidePromoData for reference.
func (c *Client) HelpHidePromoData(ctx context.Context, request *HelpHidePromoDataRequest) (bool, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
