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

// StarTransactionSourceTelegram represents TL type `starTransactionSourceTelegram#8be6d54b`.
type StarTransactionSourceTelegram struct {
}

// StarTransactionSourceTelegramTypeID is TL type id of StarTransactionSourceTelegram.
const StarTransactionSourceTelegramTypeID = 0x8be6d54b

// construct implements constructor of StarTransactionSourceClass.
func (s StarTransactionSourceTelegram) construct() StarTransactionSourceClass { return &s }

// Ensuring interfaces in compile-time for StarTransactionSourceTelegram.
var (
	_ bin.Encoder     = &StarTransactionSourceTelegram{}
	_ bin.Decoder     = &StarTransactionSourceTelegram{}
	_ bin.BareEncoder = &StarTransactionSourceTelegram{}
	_ bin.BareDecoder = &StarTransactionSourceTelegram{}

	_ StarTransactionSourceClass = &StarTransactionSourceTelegram{}
)

func (s *StarTransactionSourceTelegram) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarTransactionSourceTelegram) String() string {
	if s == nil {
		return "StarTransactionSourceTelegram(nil)"
	}
	type Alias StarTransactionSourceTelegram
	return fmt.Sprintf("StarTransactionSourceTelegram%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarTransactionSourceTelegram) TypeID() uint32 {
	return StarTransactionSourceTelegramTypeID
}

// TypeName returns name of type in TL schema.
func (*StarTransactionSourceTelegram) TypeName() string {
	return "starTransactionSourceTelegram"
}

// TypeInfo returns info about TL type.
func (s *StarTransactionSourceTelegram) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starTransactionSourceTelegram",
		ID:   StarTransactionSourceTelegramTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarTransactionSourceTelegram) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceTelegram#8be6d54b as nil")
	}
	b.PutID(StarTransactionSourceTelegramTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarTransactionSourceTelegram) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceTelegram#8be6d54b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarTransactionSourceTelegram) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceTelegram#8be6d54b to nil")
	}
	if err := b.ConsumeID(StarTransactionSourceTelegramTypeID); err != nil {
		return fmt.Errorf("unable to decode starTransactionSourceTelegram#8be6d54b: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarTransactionSourceTelegram) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceTelegram#8be6d54b to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StarTransactionSourceTelegram) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceTelegram#8be6d54b as nil")
	}
	b.ObjStart()
	b.PutID("starTransactionSourceTelegram")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StarTransactionSourceTelegram) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceTelegram#8be6d54b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("starTransactionSourceTelegram"); err != nil {
				return fmt.Errorf("unable to decode starTransactionSourceTelegram#8be6d54b: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StarTransactionSourceAppStore represents TL type `starTransactionSourceAppStore#4b0d7348`.
type StarTransactionSourceAppStore struct {
}

// StarTransactionSourceAppStoreTypeID is TL type id of StarTransactionSourceAppStore.
const StarTransactionSourceAppStoreTypeID = 0x4b0d7348

// construct implements constructor of StarTransactionSourceClass.
func (s StarTransactionSourceAppStore) construct() StarTransactionSourceClass { return &s }

// Ensuring interfaces in compile-time for StarTransactionSourceAppStore.
var (
	_ bin.Encoder     = &StarTransactionSourceAppStore{}
	_ bin.Decoder     = &StarTransactionSourceAppStore{}
	_ bin.BareEncoder = &StarTransactionSourceAppStore{}
	_ bin.BareDecoder = &StarTransactionSourceAppStore{}

	_ StarTransactionSourceClass = &StarTransactionSourceAppStore{}
)

func (s *StarTransactionSourceAppStore) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarTransactionSourceAppStore) String() string {
	if s == nil {
		return "StarTransactionSourceAppStore(nil)"
	}
	type Alias StarTransactionSourceAppStore
	return fmt.Sprintf("StarTransactionSourceAppStore%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarTransactionSourceAppStore) TypeID() uint32 {
	return StarTransactionSourceAppStoreTypeID
}

// TypeName returns name of type in TL schema.
func (*StarTransactionSourceAppStore) TypeName() string {
	return "starTransactionSourceAppStore"
}

// TypeInfo returns info about TL type.
func (s *StarTransactionSourceAppStore) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starTransactionSourceAppStore",
		ID:   StarTransactionSourceAppStoreTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarTransactionSourceAppStore) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceAppStore#4b0d7348 as nil")
	}
	b.PutID(StarTransactionSourceAppStoreTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarTransactionSourceAppStore) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceAppStore#4b0d7348 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarTransactionSourceAppStore) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceAppStore#4b0d7348 to nil")
	}
	if err := b.ConsumeID(StarTransactionSourceAppStoreTypeID); err != nil {
		return fmt.Errorf("unable to decode starTransactionSourceAppStore#4b0d7348: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarTransactionSourceAppStore) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceAppStore#4b0d7348 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StarTransactionSourceAppStore) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceAppStore#4b0d7348 as nil")
	}
	b.ObjStart()
	b.PutID("starTransactionSourceAppStore")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StarTransactionSourceAppStore) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceAppStore#4b0d7348 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("starTransactionSourceAppStore"); err != nil {
				return fmt.Errorf("unable to decode starTransactionSourceAppStore#4b0d7348: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StarTransactionSourceGooglePlay represents TL type `starTransactionSourceGooglePlay#1faf9fac`.
type StarTransactionSourceGooglePlay struct {
}

// StarTransactionSourceGooglePlayTypeID is TL type id of StarTransactionSourceGooglePlay.
const StarTransactionSourceGooglePlayTypeID = 0x1faf9fac

// construct implements constructor of StarTransactionSourceClass.
func (s StarTransactionSourceGooglePlay) construct() StarTransactionSourceClass { return &s }

// Ensuring interfaces in compile-time for StarTransactionSourceGooglePlay.
var (
	_ bin.Encoder     = &StarTransactionSourceGooglePlay{}
	_ bin.Decoder     = &StarTransactionSourceGooglePlay{}
	_ bin.BareEncoder = &StarTransactionSourceGooglePlay{}
	_ bin.BareDecoder = &StarTransactionSourceGooglePlay{}

	_ StarTransactionSourceClass = &StarTransactionSourceGooglePlay{}
)

func (s *StarTransactionSourceGooglePlay) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarTransactionSourceGooglePlay) String() string {
	if s == nil {
		return "StarTransactionSourceGooglePlay(nil)"
	}
	type Alias StarTransactionSourceGooglePlay
	return fmt.Sprintf("StarTransactionSourceGooglePlay%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarTransactionSourceGooglePlay) TypeID() uint32 {
	return StarTransactionSourceGooglePlayTypeID
}

// TypeName returns name of type in TL schema.
func (*StarTransactionSourceGooglePlay) TypeName() string {
	return "starTransactionSourceGooglePlay"
}

// TypeInfo returns info about TL type.
func (s *StarTransactionSourceGooglePlay) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starTransactionSourceGooglePlay",
		ID:   StarTransactionSourceGooglePlayTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarTransactionSourceGooglePlay) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceGooglePlay#1faf9fac as nil")
	}
	b.PutID(StarTransactionSourceGooglePlayTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarTransactionSourceGooglePlay) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceGooglePlay#1faf9fac as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarTransactionSourceGooglePlay) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceGooglePlay#1faf9fac to nil")
	}
	if err := b.ConsumeID(StarTransactionSourceGooglePlayTypeID); err != nil {
		return fmt.Errorf("unable to decode starTransactionSourceGooglePlay#1faf9fac: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarTransactionSourceGooglePlay) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceGooglePlay#1faf9fac to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StarTransactionSourceGooglePlay) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceGooglePlay#1faf9fac as nil")
	}
	b.ObjStart()
	b.PutID("starTransactionSourceGooglePlay")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StarTransactionSourceGooglePlay) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceGooglePlay#1faf9fac to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("starTransactionSourceGooglePlay"); err != nil {
				return fmt.Errorf("unable to decode starTransactionSourceGooglePlay#1faf9fac: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StarTransactionSourceFragment represents TL type `starTransactionSourceFragment#8531efbd`.
type StarTransactionSourceFragment struct {
}

// StarTransactionSourceFragmentTypeID is TL type id of StarTransactionSourceFragment.
const StarTransactionSourceFragmentTypeID = 0x8531efbd

// construct implements constructor of StarTransactionSourceClass.
func (s StarTransactionSourceFragment) construct() StarTransactionSourceClass { return &s }

// Ensuring interfaces in compile-time for StarTransactionSourceFragment.
var (
	_ bin.Encoder     = &StarTransactionSourceFragment{}
	_ bin.Decoder     = &StarTransactionSourceFragment{}
	_ bin.BareEncoder = &StarTransactionSourceFragment{}
	_ bin.BareDecoder = &StarTransactionSourceFragment{}

	_ StarTransactionSourceClass = &StarTransactionSourceFragment{}
)

func (s *StarTransactionSourceFragment) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarTransactionSourceFragment) String() string {
	if s == nil {
		return "StarTransactionSourceFragment(nil)"
	}
	type Alias StarTransactionSourceFragment
	return fmt.Sprintf("StarTransactionSourceFragment%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarTransactionSourceFragment) TypeID() uint32 {
	return StarTransactionSourceFragmentTypeID
}

// TypeName returns name of type in TL schema.
func (*StarTransactionSourceFragment) TypeName() string {
	return "starTransactionSourceFragment"
}

// TypeInfo returns info about TL type.
func (s *StarTransactionSourceFragment) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starTransactionSourceFragment",
		ID:   StarTransactionSourceFragmentTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarTransactionSourceFragment) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceFragment#8531efbd as nil")
	}
	b.PutID(StarTransactionSourceFragmentTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarTransactionSourceFragment) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceFragment#8531efbd as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarTransactionSourceFragment) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceFragment#8531efbd to nil")
	}
	if err := b.ConsumeID(StarTransactionSourceFragmentTypeID); err != nil {
		return fmt.Errorf("unable to decode starTransactionSourceFragment#8531efbd: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarTransactionSourceFragment) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceFragment#8531efbd to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StarTransactionSourceFragment) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceFragment#8531efbd as nil")
	}
	b.ObjStart()
	b.PutID("starTransactionSourceFragment")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StarTransactionSourceFragment) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceFragment#8531efbd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("starTransactionSourceFragment"); err != nil {
				return fmt.Errorf("unable to decode starTransactionSourceFragment#8531efbd: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StarTransactionSourceUser represents TL type `starTransactionSourceUser#dd07c483`.
type StarTransactionSourceUser struct {
	// Identifier of the user
	UserID int64
	// Information about the bought product; may be null if none
	ProductInfo ProductInfo
}

// StarTransactionSourceUserTypeID is TL type id of StarTransactionSourceUser.
const StarTransactionSourceUserTypeID = 0xdd07c483

// construct implements constructor of StarTransactionSourceClass.
func (s StarTransactionSourceUser) construct() StarTransactionSourceClass { return &s }

// Ensuring interfaces in compile-time for StarTransactionSourceUser.
var (
	_ bin.Encoder     = &StarTransactionSourceUser{}
	_ bin.Decoder     = &StarTransactionSourceUser{}
	_ bin.BareEncoder = &StarTransactionSourceUser{}
	_ bin.BareDecoder = &StarTransactionSourceUser{}

	_ StarTransactionSourceClass = &StarTransactionSourceUser{}
)

func (s *StarTransactionSourceUser) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.ProductInfo.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarTransactionSourceUser) String() string {
	if s == nil {
		return "StarTransactionSourceUser(nil)"
	}
	type Alias StarTransactionSourceUser
	return fmt.Sprintf("StarTransactionSourceUser%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarTransactionSourceUser) TypeID() uint32 {
	return StarTransactionSourceUserTypeID
}

// TypeName returns name of type in TL schema.
func (*StarTransactionSourceUser) TypeName() string {
	return "starTransactionSourceUser"
}

// TypeInfo returns info about TL type.
func (s *StarTransactionSourceUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starTransactionSourceUser",
		ID:   StarTransactionSourceUserTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "ProductInfo",
			SchemaName: "product_info",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarTransactionSourceUser) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceUser#dd07c483 as nil")
	}
	b.PutID(StarTransactionSourceUserTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarTransactionSourceUser) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceUser#dd07c483 as nil")
	}
	b.PutInt53(s.UserID)
	if err := s.ProductInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starTransactionSourceUser#dd07c483: field product_info: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarTransactionSourceUser) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceUser#dd07c483 to nil")
	}
	if err := b.ConsumeID(StarTransactionSourceUserTypeID); err != nil {
		return fmt.Errorf("unable to decode starTransactionSourceUser#dd07c483: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarTransactionSourceUser) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceUser#dd07c483 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode starTransactionSourceUser#dd07c483: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		if err := s.ProductInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode starTransactionSourceUser#dd07c483: field product_info: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StarTransactionSourceUser) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceUser#dd07c483 as nil")
	}
	b.ObjStart()
	b.PutID("starTransactionSourceUser")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(s.UserID)
	b.Comma()
	b.FieldStart("product_info")
	if err := s.ProductInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode starTransactionSourceUser#dd07c483: field product_info: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StarTransactionSourceUser) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceUser#dd07c483 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("starTransactionSourceUser"); err != nil {
				return fmt.Errorf("unable to decode starTransactionSourceUser#dd07c483: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode starTransactionSourceUser#dd07c483: field user_id: %w", err)
			}
			s.UserID = value
		case "product_info":
			if err := s.ProductInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode starTransactionSourceUser#dd07c483: field product_info: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (s *StarTransactionSourceUser) GetUserID() (value int64) {
	if s == nil {
		return
	}
	return s.UserID
}

// GetProductInfo returns value of ProductInfo field.
func (s *StarTransactionSourceUser) GetProductInfo() (value ProductInfo) {
	if s == nil {
		return
	}
	return s.ProductInfo
}

// StarTransactionSourceUnsupported represents TL type `starTransactionSourceUnsupported#1c536b37`.
type StarTransactionSourceUnsupported struct {
}

// StarTransactionSourceUnsupportedTypeID is TL type id of StarTransactionSourceUnsupported.
const StarTransactionSourceUnsupportedTypeID = 0x1c536b37

// construct implements constructor of StarTransactionSourceClass.
func (s StarTransactionSourceUnsupported) construct() StarTransactionSourceClass { return &s }

// Ensuring interfaces in compile-time for StarTransactionSourceUnsupported.
var (
	_ bin.Encoder     = &StarTransactionSourceUnsupported{}
	_ bin.Decoder     = &StarTransactionSourceUnsupported{}
	_ bin.BareEncoder = &StarTransactionSourceUnsupported{}
	_ bin.BareDecoder = &StarTransactionSourceUnsupported{}

	_ StarTransactionSourceClass = &StarTransactionSourceUnsupported{}
)

func (s *StarTransactionSourceUnsupported) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarTransactionSourceUnsupported) String() string {
	if s == nil {
		return "StarTransactionSourceUnsupported(nil)"
	}
	type Alias StarTransactionSourceUnsupported
	return fmt.Sprintf("StarTransactionSourceUnsupported%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarTransactionSourceUnsupported) TypeID() uint32 {
	return StarTransactionSourceUnsupportedTypeID
}

// TypeName returns name of type in TL schema.
func (*StarTransactionSourceUnsupported) TypeName() string {
	return "starTransactionSourceUnsupported"
}

// TypeInfo returns info about TL type.
func (s *StarTransactionSourceUnsupported) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starTransactionSourceUnsupported",
		ID:   StarTransactionSourceUnsupportedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarTransactionSourceUnsupported) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceUnsupported#1c536b37 as nil")
	}
	b.PutID(StarTransactionSourceUnsupportedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarTransactionSourceUnsupported) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceUnsupported#1c536b37 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarTransactionSourceUnsupported) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceUnsupported#1c536b37 to nil")
	}
	if err := b.ConsumeID(StarTransactionSourceUnsupportedTypeID); err != nil {
		return fmt.Errorf("unable to decode starTransactionSourceUnsupported#1c536b37: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarTransactionSourceUnsupported) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceUnsupported#1c536b37 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StarTransactionSourceUnsupported) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransactionSourceUnsupported#1c536b37 as nil")
	}
	b.ObjStart()
	b.PutID("starTransactionSourceUnsupported")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StarTransactionSourceUnsupported) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransactionSourceUnsupported#1c536b37 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("starTransactionSourceUnsupported"); err != nil {
				return fmt.Errorf("unable to decode starTransactionSourceUnsupported#1c536b37: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StarTransactionSourceClassName is schema name of StarTransactionSourceClass.
const StarTransactionSourceClassName = "StarTransactionSource"

// StarTransactionSourceClass represents StarTransactionSource generic type.
//
// Example:
//
//	g, err := tdapi.DecodeStarTransactionSource(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.StarTransactionSourceTelegram: // starTransactionSourceTelegram#8be6d54b
//	case *tdapi.StarTransactionSourceAppStore: // starTransactionSourceAppStore#4b0d7348
//	case *tdapi.StarTransactionSourceGooglePlay: // starTransactionSourceGooglePlay#1faf9fac
//	case *tdapi.StarTransactionSourceFragment: // starTransactionSourceFragment#8531efbd
//	case *tdapi.StarTransactionSourceUser: // starTransactionSourceUser#dd07c483
//	case *tdapi.StarTransactionSourceUnsupported: // starTransactionSourceUnsupported#1c536b37
//	default: panic(v)
//	}
type StarTransactionSourceClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StarTransactionSourceClass

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
}

// DecodeStarTransactionSource implements binary de-serialization for StarTransactionSourceClass.
func DecodeStarTransactionSource(buf *bin.Buffer) (StarTransactionSourceClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StarTransactionSourceTelegramTypeID:
		// Decoding starTransactionSourceTelegram#8be6d54b.
		v := StarTransactionSourceTelegram{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case StarTransactionSourceAppStoreTypeID:
		// Decoding starTransactionSourceAppStore#4b0d7348.
		v := StarTransactionSourceAppStore{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case StarTransactionSourceGooglePlayTypeID:
		// Decoding starTransactionSourceGooglePlay#1faf9fac.
		v := StarTransactionSourceGooglePlay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case StarTransactionSourceFragmentTypeID:
		// Decoding starTransactionSourceFragment#8531efbd.
		v := StarTransactionSourceFragment{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case StarTransactionSourceUserTypeID:
		// Decoding starTransactionSourceUser#dd07c483.
		v := StarTransactionSourceUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case StarTransactionSourceUnsupportedTypeID:
		// Decoding starTransactionSourceUnsupported#1c536b37.
		v := StarTransactionSourceUnsupported{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONStarTransactionSource implements binary de-serialization for StarTransactionSourceClass.
func DecodeTDLibJSONStarTransactionSource(buf tdjson.Decoder) (StarTransactionSourceClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "starTransactionSourceTelegram":
		// Decoding starTransactionSourceTelegram#8be6d54b.
		v := StarTransactionSourceTelegram{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case "starTransactionSourceAppStore":
		// Decoding starTransactionSourceAppStore#4b0d7348.
		v := StarTransactionSourceAppStore{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case "starTransactionSourceGooglePlay":
		// Decoding starTransactionSourceGooglePlay#1faf9fac.
		v := StarTransactionSourceGooglePlay{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case "starTransactionSourceFragment":
		// Decoding starTransactionSourceFragment#8531efbd.
		v := StarTransactionSourceFragment{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case "starTransactionSourceUser":
		// Decoding starTransactionSourceUser#dd07c483.
		v := StarTransactionSourceUser{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	case "starTransactionSourceUnsupported":
		// Decoding starTransactionSourceUnsupported#1c536b37.
		v := StarTransactionSourceUnsupported{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StarTransactionSourceClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// StarTransactionSource boxes the StarTransactionSourceClass providing a helper.
type StarTransactionSourceBox struct {
	StarTransactionSource StarTransactionSourceClass
}

// Decode implements bin.Decoder for StarTransactionSourceBox.
func (b *StarTransactionSourceBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StarTransactionSourceBox to nil")
	}
	v, err := DecodeStarTransactionSource(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StarTransactionSource = v
	return nil
}

// Encode implements bin.Encode for StarTransactionSourceBox.
func (b *StarTransactionSourceBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StarTransactionSource == nil {
		return fmt.Errorf("unable to encode StarTransactionSourceClass as nil")
	}
	return b.StarTransactionSource.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for StarTransactionSourceBox.
func (b *StarTransactionSourceBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode StarTransactionSourceBox to nil")
	}
	v, err := DecodeTDLibJSONStarTransactionSource(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StarTransactionSource = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for StarTransactionSourceBox.
func (b *StarTransactionSourceBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.StarTransactionSource == nil {
		return fmt.Errorf("unable to encode StarTransactionSourceClass as nil")
	}
	return b.StarTransactionSource.EncodeTDLibJSON(buf)
}
