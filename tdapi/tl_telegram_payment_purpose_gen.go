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

// TelegramPaymentPurposePremiumGiftCodes represents TL type `telegramPaymentPurposePremiumGiftCodes#91009cc`.
type TelegramPaymentPurposePremiumGiftCodes struct {
	// Identifier of the supergroup or channel chat, which will be automatically boosted by
	// the users for duration of the Premium subscription and which is administered by the
	// user; 0 if none
	BoostedChatID int64
	// ISO 4217 currency code of the payment currency
	Currency string
	// Paid amount, in the smallest units of the currency
	Amount int64
	// Identifiers of the users which can activate the gift codes
	UserIDs []int64
	// Number of months the Telegram Premium subscription will be active for the users
	MonthCount int32
}

// TelegramPaymentPurposePremiumGiftCodesTypeID is TL type id of TelegramPaymentPurposePremiumGiftCodes.
const TelegramPaymentPurposePremiumGiftCodesTypeID = 0x91009cc

// construct implements constructor of TelegramPaymentPurposeClass.
func (t TelegramPaymentPurposePremiumGiftCodes) construct() TelegramPaymentPurposeClass { return &t }

// Ensuring interfaces in compile-time for TelegramPaymentPurposePremiumGiftCodes.
var (
	_ bin.Encoder     = &TelegramPaymentPurposePremiumGiftCodes{}
	_ bin.Decoder     = &TelegramPaymentPurposePremiumGiftCodes{}
	_ bin.BareEncoder = &TelegramPaymentPurposePremiumGiftCodes{}
	_ bin.BareDecoder = &TelegramPaymentPurposePremiumGiftCodes{}

	_ TelegramPaymentPurposeClass = &TelegramPaymentPurposePremiumGiftCodes{}
)

func (t *TelegramPaymentPurposePremiumGiftCodes) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.BoostedChatID == 0) {
		return false
	}
	if !(t.Currency == "") {
		return false
	}
	if !(t.Amount == 0) {
		return false
	}
	if !(t.UserIDs == nil) {
		return false
	}
	if !(t.MonthCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TelegramPaymentPurposePremiumGiftCodes) String() string {
	if t == nil {
		return "TelegramPaymentPurposePremiumGiftCodes(nil)"
	}
	type Alias TelegramPaymentPurposePremiumGiftCodes
	return fmt.Sprintf("TelegramPaymentPurposePremiumGiftCodes%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TelegramPaymentPurposePremiumGiftCodes) TypeID() uint32 {
	return TelegramPaymentPurposePremiumGiftCodesTypeID
}

// TypeName returns name of type in TL schema.
func (*TelegramPaymentPurposePremiumGiftCodes) TypeName() string {
	return "telegramPaymentPurposePremiumGiftCodes"
}

// TypeInfo returns info about TL type.
func (t *TelegramPaymentPurposePremiumGiftCodes) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "telegramPaymentPurposePremiumGiftCodes",
		ID:   TelegramPaymentPurposePremiumGiftCodesTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BoostedChatID",
			SchemaName: "boosted_chat_id",
		},
		{
			Name:       "Currency",
			SchemaName: "currency",
		},
		{
			Name:       "Amount",
			SchemaName: "amount",
		},
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
		{
			Name:       "MonthCount",
			SchemaName: "month_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TelegramPaymentPurposePremiumGiftCodes) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode telegramPaymentPurposePremiumGiftCodes#91009cc as nil")
	}
	b.PutID(TelegramPaymentPurposePremiumGiftCodesTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TelegramPaymentPurposePremiumGiftCodes) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode telegramPaymentPurposePremiumGiftCodes#91009cc as nil")
	}
	b.PutInt53(t.BoostedChatID)
	b.PutString(t.Currency)
	b.PutInt53(t.Amount)
	b.PutInt(len(t.UserIDs))
	for _, v := range t.UserIDs {
		b.PutInt53(v)
	}
	b.PutInt32(t.MonthCount)
	return nil
}

// Decode implements bin.Decoder.
func (t *TelegramPaymentPurposePremiumGiftCodes) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode telegramPaymentPurposePremiumGiftCodes#91009cc to nil")
	}
	if err := b.ConsumeID(TelegramPaymentPurposePremiumGiftCodesTypeID); err != nil {
		return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TelegramPaymentPurposePremiumGiftCodes) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode telegramPaymentPurposePremiumGiftCodes#91009cc to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field boosted_chat_id: %w", err)
		}
		t.BoostedChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field currency: %w", err)
		}
		t.Currency = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field amount: %w", err)
		}
		t.Amount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field user_ids: %w", err)
		}

		if headerLen > 0 {
			t.UserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field user_ids: %w", err)
			}
			t.UserIDs = append(t.UserIDs, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field month_count: %w", err)
		}
		t.MonthCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TelegramPaymentPurposePremiumGiftCodes) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode telegramPaymentPurposePremiumGiftCodes#91009cc as nil")
	}
	b.ObjStart()
	b.PutID("telegramPaymentPurposePremiumGiftCodes")
	b.Comma()
	b.FieldStart("boosted_chat_id")
	b.PutInt53(t.BoostedChatID)
	b.Comma()
	b.FieldStart("currency")
	b.PutString(t.Currency)
	b.Comma()
	b.FieldStart("amount")
	b.PutInt53(t.Amount)
	b.Comma()
	b.FieldStart("user_ids")
	b.ArrStart()
	for _, v := range t.UserIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("month_count")
	b.PutInt32(t.MonthCount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TelegramPaymentPurposePremiumGiftCodes) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode telegramPaymentPurposePremiumGiftCodes#91009cc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("telegramPaymentPurposePremiumGiftCodes"); err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: %w", err)
			}
		case "boosted_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field boosted_chat_id: %w", err)
			}
			t.BoostedChatID = value
		case "currency":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field currency: %w", err)
			}
			t.Currency = value
		case "amount":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field amount: %w", err)
			}
			t.Amount = value
		case "user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field user_ids: %w", err)
				}
				t.UserIDs = append(t.UserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field user_ids: %w", err)
			}
		case "month_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiftCodes#91009cc: field month_count: %w", err)
			}
			t.MonthCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBoostedChatID returns value of BoostedChatID field.
func (t *TelegramPaymentPurposePremiumGiftCodes) GetBoostedChatID() (value int64) {
	if t == nil {
		return
	}
	return t.BoostedChatID
}

// GetCurrency returns value of Currency field.
func (t *TelegramPaymentPurposePremiumGiftCodes) GetCurrency() (value string) {
	if t == nil {
		return
	}
	return t.Currency
}

// GetAmount returns value of Amount field.
func (t *TelegramPaymentPurposePremiumGiftCodes) GetAmount() (value int64) {
	if t == nil {
		return
	}
	return t.Amount
}

// GetUserIDs returns value of UserIDs field.
func (t *TelegramPaymentPurposePremiumGiftCodes) GetUserIDs() (value []int64) {
	if t == nil {
		return
	}
	return t.UserIDs
}

// GetMonthCount returns value of MonthCount field.
func (t *TelegramPaymentPurposePremiumGiftCodes) GetMonthCount() (value int32) {
	if t == nil {
		return
	}
	return t.MonthCount
}

// TelegramPaymentPurposePremiumGiveaway represents TL type `telegramPaymentPurposePremiumGiveaway#bfd4a227`.
type TelegramPaymentPurposePremiumGiveaway struct {
	// Giveaway parameters
	Parameters PremiumGiveawayParameters
	// ISO 4217 currency code of the payment currency
	Currency string
	// Paid amount, in the smallest units of the currency
	Amount int64
	// Number of users which will be able to activate the gift codes
	WinnerCount int32
	// Number of months the Telegram Premium subscription will be active for the users
	MonthCount int32
}

// TelegramPaymentPurposePremiumGiveawayTypeID is TL type id of TelegramPaymentPurposePremiumGiveaway.
const TelegramPaymentPurposePremiumGiveawayTypeID = 0xbfd4a227

// construct implements constructor of TelegramPaymentPurposeClass.
func (t TelegramPaymentPurposePremiumGiveaway) construct() TelegramPaymentPurposeClass { return &t }

// Ensuring interfaces in compile-time for TelegramPaymentPurposePremiumGiveaway.
var (
	_ bin.Encoder     = &TelegramPaymentPurposePremiumGiveaway{}
	_ bin.Decoder     = &TelegramPaymentPurposePremiumGiveaway{}
	_ bin.BareEncoder = &TelegramPaymentPurposePremiumGiveaway{}
	_ bin.BareDecoder = &TelegramPaymentPurposePremiumGiveaway{}

	_ TelegramPaymentPurposeClass = &TelegramPaymentPurposePremiumGiveaway{}
)

func (t *TelegramPaymentPurposePremiumGiveaway) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Parameters.Zero()) {
		return false
	}
	if !(t.Currency == "") {
		return false
	}
	if !(t.Amount == 0) {
		return false
	}
	if !(t.WinnerCount == 0) {
		return false
	}
	if !(t.MonthCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TelegramPaymentPurposePremiumGiveaway) String() string {
	if t == nil {
		return "TelegramPaymentPurposePremiumGiveaway(nil)"
	}
	type Alias TelegramPaymentPurposePremiumGiveaway
	return fmt.Sprintf("TelegramPaymentPurposePremiumGiveaway%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TelegramPaymentPurposePremiumGiveaway) TypeID() uint32 {
	return TelegramPaymentPurposePremiumGiveawayTypeID
}

// TypeName returns name of type in TL schema.
func (*TelegramPaymentPurposePremiumGiveaway) TypeName() string {
	return "telegramPaymentPurposePremiumGiveaway"
}

// TypeInfo returns info about TL type.
func (t *TelegramPaymentPurposePremiumGiveaway) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "telegramPaymentPurposePremiumGiveaway",
		ID:   TelegramPaymentPurposePremiumGiveawayTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Parameters",
			SchemaName: "parameters",
		},
		{
			Name:       "Currency",
			SchemaName: "currency",
		},
		{
			Name:       "Amount",
			SchemaName: "amount",
		},
		{
			Name:       "WinnerCount",
			SchemaName: "winner_count",
		},
		{
			Name:       "MonthCount",
			SchemaName: "month_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TelegramPaymentPurposePremiumGiveaway) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode telegramPaymentPurposePremiumGiveaway#bfd4a227 as nil")
	}
	b.PutID(TelegramPaymentPurposePremiumGiveawayTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TelegramPaymentPurposePremiumGiveaway) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode telegramPaymentPurposePremiumGiveaway#bfd4a227 as nil")
	}
	if err := t.Parameters.Encode(b); err != nil {
		return fmt.Errorf("unable to encode telegramPaymentPurposePremiumGiveaway#bfd4a227: field parameters: %w", err)
	}
	b.PutString(t.Currency)
	b.PutInt53(t.Amount)
	b.PutInt32(t.WinnerCount)
	b.PutInt32(t.MonthCount)
	return nil
}

// Decode implements bin.Decoder.
func (t *TelegramPaymentPurposePremiumGiveaway) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode telegramPaymentPurposePremiumGiveaway#bfd4a227 to nil")
	}
	if err := b.ConsumeID(TelegramPaymentPurposePremiumGiveawayTypeID); err != nil {
		return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TelegramPaymentPurposePremiumGiveaway) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode telegramPaymentPurposePremiumGiveaway#bfd4a227 to nil")
	}
	{
		if err := t.Parameters.Decode(b); err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field parameters: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field currency: %w", err)
		}
		t.Currency = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field amount: %w", err)
		}
		t.Amount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field winner_count: %w", err)
		}
		t.WinnerCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field month_count: %w", err)
		}
		t.MonthCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TelegramPaymentPurposePremiumGiveaway) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode telegramPaymentPurposePremiumGiveaway#bfd4a227 as nil")
	}
	b.ObjStart()
	b.PutID("telegramPaymentPurposePremiumGiveaway")
	b.Comma()
	b.FieldStart("parameters")
	if err := t.Parameters.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode telegramPaymentPurposePremiumGiveaway#bfd4a227: field parameters: %w", err)
	}
	b.Comma()
	b.FieldStart("currency")
	b.PutString(t.Currency)
	b.Comma()
	b.FieldStart("amount")
	b.PutInt53(t.Amount)
	b.Comma()
	b.FieldStart("winner_count")
	b.PutInt32(t.WinnerCount)
	b.Comma()
	b.FieldStart("month_count")
	b.PutInt32(t.MonthCount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TelegramPaymentPurposePremiumGiveaway) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode telegramPaymentPurposePremiumGiveaway#bfd4a227 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("telegramPaymentPurposePremiumGiveaway"); err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: %w", err)
			}
		case "parameters":
			if err := t.Parameters.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field parameters: %w", err)
			}
		case "currency":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field currency: %w", err)
			}
			t.Currency = value
		case "amount":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field amount: %w", err)
			}
			t.Amount = value
		case "winner_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field winner_count: %w", err)
			}
			t.WinnerCount = value
		case "month_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposePremiumGiveaway#bfd4a227: field month_count: %w", err)
			}
			t.MonthCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetParameters returns value of Parameters field.
func (t *TelegramPaymentPurposePremiumGiveaway) GetParameters() (value PremiumGiveawayParameters) {
	if t == nil {
		return
	}
	return t.Parameters
}

// GetCurrency returns value of Currency field.
func (t *TelegramPaymentPurposePremiumGiveaway) GetCurrency() (value string) {
	if t == nil {
		return
	}
	return t.Currency
}

// GetAmount returns value of Amount field.
func (t *TelegramPaymentPurposePremiumGiveaway) GetAmount() (value int64) {
	if t == nil {
		return
	}
	return t.Amount
}

// GetWinnerCount returns value of WinnerCount field.
func (t *TelegramPaymentPurposePremiumGiveaway) GetWinnerCount() (value int32) {
	if t == nil {
		return
	}
	return t.WinnerCount
}

// GetMonthCount returns value of MonthCount field.
func (t *TelegramPaymentPurposePremiumGiveaway) GetMonthCount() (value int32) {
	if t == nil {
		return
	}
	return t.MonthCount
}

// TelegramPaymentPurposeStars represents TL type `telegramPaymentPurposeStars#e273ee52`.
type TelegramPaymentPurposeStars struct {
	// ISO 4217 currency code of the payment currency
	Currency string
	// Paid amount, in the smallest units of the currency
	Amount int64
	// Number of bought stars
	StarCount int64
}

// TelegramPaymentPurposeStarsTypeID is TL type id of TelegramPaymentPurposeStars.
const TelegramPaymentPurposeStarsTypeID = 0xe273ee52

// construct implements constructor of TelegramPaymentPurposeClass.
func (t TelegramPaymentPurposeStars) construct() TelegramPaymentPurposeClass { return &t }

// Ensuring interfaces in compile-time for TelegramPaymentPurposeStars.
var (
	_ bin.Encoder     = &TelegramPaymentPurposeStars{}
	_ bin.Decoder     = &TelegramPaymentPurposeStars{}
	_ bin.BareEncoder = &TelegramPaymentPurposeStars{}
	_ bin.BareDecoder = &TelegramPaymentPurposeStars{}

	_ TelegramPaymentPurposeClass = &TelegramPaymentPurposeStars{}
)

func (t *TelegramPaymentPurposeStars) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Currency == "") {
		return false
	}
	if !(t.Amount == 0) {
		return false
	}
	if !(t.StarCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TelegramPaymentPurposeStars) String() string {
	if t == nil {
		return "TelegramPaymentPurposeStars(nil)"
	}
	type Alias TelegramPaymentPurposeStars
	return fmt.Sprintf("TelegramPaymentPurposeStars%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TelegramPaymentPurposeStars) TypeID() uint32 {
	return TelegramPaymentPurposeStarsTypeID
}

// TypeName returns name of type in TL schema.
func (*TelegramPaymentPurposeStars) TypeName() string {
	return "telegramPaymentPurposeStars"
}

// TypeInfo returns info about TL type.
func (t *TelegramPaymentPurposeStars) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "telegramPaymentPurposeStars",
		ID:   TelegramPaymentPurposeStarsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Currency",
			SchemaName: "currency",
		},
		{
			Name:       "Amount",
			SchemaName: "amount",
		},
		{
			Name:       "StarCount",
			SchemaName: "star_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TelegramPaymentPurposeStars) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode telegramPaymentPurposeStars#e273ee52 as nil")
	}
	b.PutID(TelegramPaymentPurposeStarsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TelegramPaymentPurposeStars) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode telegramPaymentPurposeStars#e273ee52 as nil")
	}
	b.PutString(t.Currency)
	b.PutInt53(t.Amount)
	b.PutInt53(t.StarCount)
	return nil
}

// Decode implements bin.Decoder.
func (t *TelegramPaymentPurposeStars) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode telegramPaymentPurposeStars#e273ee52 to nil")
	}
	if err := b.ConsumeID(TelegramPaymentPurposeStarsTypeID); err != nil {
		return fmt.Errorf("unable to decode telegramPaymentPurposeStars#e273ee52: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TelegramPaymentPurposeStars) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode telegramPaymentPurposeStars#e273ee52 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposeStars#e273ee52: field currency: %w", err)
		}
		t.Currency = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposeStars#e273ee52: field amount: %w", err)
		}
		t.Amount = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode telegramPaymentPurposeStars#e273ee52: field star_count: %w", err)
		}
		t.StarCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TelegramPaymentPurposeStars) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode telegramPaymentPurposeStars#e273ee52 as nil")
	}
	b.ObjStart()
	b.PutID("telegramPaymentPurposeStars")
	b.Comma()
	b.FieldStart("currency")
	b.PutString(t.Currency)
	b.Comma()
	b.FieldStart("amount")
	b.PutInt53(t.Amount)
	b.Comma()
	b.FieldStart("star_count")
	b.PutInt53(t.StarCount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TelegramPaymentPurposeStars) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode telegramPaymentPurposeStars#e273ee52 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("telegramPaymentPurposeStars"); err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposeStars#e273ee52: %w", err)
			}
		case "currency":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposeStars#e273ee52: field currency: %w", err)
			}
			t.Currency = value
		case "amount":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposeStars#e273ee52: field amount: %w", err)
			}
			t.Amount = value
		case "star_count":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode telegramPaymentPurposeStars#e273ee52: field star_count: %w", err)
			}
			t.StarCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCurrency returns value of Currency field.
func (t *TelegramPaymentPurposeStars) GetCurrency() (value string) {
	if t == nil {
		return
	}
	return t.Currency
}

// GetAmount returns value of Amount field.
func (t *TelegramPaymentPurposeStars) GetAmount() (value int64) {
	if t == nil {
		return
	}
	return t.Amount
}

// GetStarCount returns value of StarCount field.
func (t *TelegramPaymentPurposeStars) GetStarCount() (value int64) {
	if t == nil {
		return
	}
	return t.StarCount
}

// TelegramPaymentPurposeClassName is schema name of TelegramPaymentPurposeClass.
const TelegramPaymentPurposeClassName = "TelegramPaymentPurpose"

// TelegramPaymentPurposeClass represents TelegramPaymentPurpose generic type.
//
// Example:
//
//	g, err := tdapi.DecodeTelegramPaymentPurpose(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.TelegramPaymentPurposePremiumGiftCodes: // telegramPaymentPurposePremiumGiftCodes#91009cc
//	case *tdapi.TelegramPaymentPurposePremiumGiveaway: // telegramPaymentPurposePremiumGiveaway#bfd4a227
//	case *tdapi.TelegramPaymentPurposeStars: // telegramPaymentPurposeStars#e273ee52
//	default: panic(v)
//	}
type TelegramPaymentPurposeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() TelegramPaymentPurposeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error

	// ISO 4217 currency code of the payment currency
	GetCurrency() (value string)
	// Paid amount, in the smallest units of the currency
	GetAmount() (value int64)
}

// DecodeTelegramPaymentPurpose implements binary de-serialization for TelegramPaymentPurposeClass.
func DecodeTelegramPaymentPurpose(buf *bin.Buffer) (TelegramPaymentPurposeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case TelegramPaymentPurposePremiumGiftCodesTypeID:
		// Decoding telegramPaymentPurposePremiumGiftCodes#91009cc.
		v := TelegramPaymentPurposePremiumGiftCodes{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TelegramPaymentPurposeClass: %w", err)
		}
		return &v, nil
	case TelegramPaymentPurposePremiumGiveawayTypeID:
		// Decoding telegramPaymentPurposePremiumGiveaway#bfd4a227.
		v := TelegramPaymentPurposePremiumGiveaway{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TelegramPaymentPurposeClass: %w", err)
		}
		return &v, nil
	case TelegramPaymentPurposeStarsTypeID:
		// Decoding telegramPaymentPurposeStars#e273ee52.
		v := TelegramPaymentPurposeStars{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TelegramPaymentPurposeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode TelegramPaymentPurposeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONTelegramPaymentPurpose implements binary de-serialization for TelegramPaymentPurposeClass.
func DecodeTDLibJSONTelegramPaymentPurpose(buf tdjson.Decoder) (TelegramPaymentPurposeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "telegramPaymentPurposePremiumGiftCodes":
		// Decoding telegramPaymentPurposePremiumGiftCodes#91009cc.
		v := TelegramPaymentPurposePremiumGiftCodes{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TelegramPaymentPurposeClass: %w", err)
		}
		return &v, nil
	case "telegramPaymentPurposePremiumGiveaway":
		// Decoding telegramPaymentPurposePremiumGiveaway#bfd4a227.
		v := TelegramPaymentPurposePremiumGiveaway{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TelegramPaymentPurposeClass: %w", err)
		}
		return &v, nil
	case "telegramPaymentPurposeStars":
		// Decoding telegramPaymentPurposeStars#e273ee52.
		v := TelegramPaymentPurposeStars{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TelegramPaymentPurposeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode TelegramPaymentPurposeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// TelegramPaymentPurpose boxes the TelegramPaymentPurposeClass providing a helper.
type TelegramPaymentPurposeBox struct {
	TelegramPaymentPurpose TelegramPaymentPurposeClass
}

// Decode implements bin.Decoder for TelegramPaymentPurposeBox.
func (b *TelegramPaymentPurposeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode TelegramPaymentPurposeBox to nil")
	}
	v, err := DecodeTelegramPaymentPurpose(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TelegramPaymentPurpose = v
	return nil
}

// Encode implements bin.Encode for TelegramPaymentPurposeBox.
func (b *TelegramPaymentPurposeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TelegramPaymentPurpose == nil {
		return fmt.Errorf("unable to encode TelegramPaymentPurposeClass as nil")
	}
	return b.TelegramPaymentPurpose.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for TelegramPaymentPurposeBox.
func (b *TelegramPaymentPurposeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode TelegramPaymentPurposeBox to nil")
	}
	v, err := DecodeTDLibJSONTelegramPaymentPurpose(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TelegramPaymentPurpose = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for TelegramPaymentPurposeBox.
func (b *TelegramPaymentPurposeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.TelegramPaymentPurpose == nil {
		return fmt.Errorf("unable to encode TelegramPaymentPurposeClass as nil")
	}
	return b.TelegramPaymentPurpose.EncodeTDLibJSON(buf)
}
