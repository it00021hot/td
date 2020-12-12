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

// AccountGetContactSignUpNotificationRequest represents TL type `account.getContactSignUpNotification#9f07c728`.
// Whether the user will receive notifications when contacts sign up
//
// See https://core.telegram.org/method/account.getContactSignUpNotification for reference.
type AccountGetContactSignUpNotificationRequest struct {
}

// AccountGetContactSignUpNotificationRequestTypeID is TL type id of AccountGetContactSignUpNotificationRequest.
const AccountGetContactSignUpNotificationRequestTypeID = 0x9f07c728

// Encode implements bin.Encoder.
func (g *AccountGetContactSignUpNotificationRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getContactSignUpNotification#9f07c728 as nil")
	}
	b.PutID(AccountGetContactSignUpNotificationRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetContactSignUpNotificationRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getContactSignUpNotification#9f07c728 to nil")
	}
	if err := b.ConsumeID(AccountGetContactSignUpNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getContactSignUpNotification#9f07c728: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetContactSignUpNotificationRequest.
var (
	_ bin.Encoder = &AccountGetContactSignUpNotificationRequest{}
	_ bin.Decoder = &AccountGetContactSignUpNotificationRequest{}
)

// AccountGetContactSignUpNotification invokes method account.getContactSignUpNotification#9f07c728 returning error if any.
// Whether the user will receive notifications when contacts sign up
//
// See https://core.telegram.org/method/account.getContactSignUpNotification for reference.
func (c *Client) AccountGetContactSignUpNotification(ctx context.Context, request *AccountGetContactSignUpNotificationRequest) (bool, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
