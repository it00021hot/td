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

// AccountInstallWallPaperRequest represents TL type `account.installWallPaper#feed5769`.
// Install wallpaper
//
// See https://core.telegram.org/method/account.installWallPaper for reference.
type AccountInstallWallPaperRequest struct {
	// Wallpaper to install
	Wallpaper InputWallPaperClass
	// Wallpaper settings
	Settings WallPaperSettings
}

// AccountInstallWallPaperRequestTypeID is TL type id of AccountInstallWallPaperRequest.
const AccountInstallWallPaperRequestTypeID = 0xfeed5769

// Encode implements bin.Encoder.
func (i *AccountInstallWallPaperRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.installWallPaper#feed5769 as nil")
	}
	b.PutID(AccountInstallWallPaperRequestTypeID)
	if i.Wallpaper == nil {
		return fmt.Errorf("unable to encode account.installWallPaper#feed5769: field wallpaper is nil")
	}
	if err := i.Wallpaper.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.installWallPaper#feed5769: field wallpaper: %w", err)
	}
	if err := i.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.installWallPaper#feed5769: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *AccountInstallWallPaperRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.installWallPaper#feed5769 to nil")
	}
	if err := b.ConsumeID(AccountInstallWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.installWallPaper#feed5769: %w", err)
	}
	{
		value, err := DecodeInputWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.installWallPaper#feed5769: field wallpaper: %w", err)
		}
		i.Wallpaper = value
	}
	{
		if err := i.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.installWallPaper#feed5769: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountInstallWallPaperRequest.
var (
	_ bin.Encoder = &AccountInstallWallPaperRequest{}
	_ bin.Decoder = &AccountInstallWallPaperRequest{}
)

// AccountInstallWallPaper invokes method account.installWallPaper#feed5769 returning error if any.
// Install wallpaper
//
// See https://core.telegram.org/method/account.installWallPaper for reference.
func (c *Client) AccountInstallWallPaper(ctx context.Context, request *AccountInstallWallPaperRequest) (bool, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
