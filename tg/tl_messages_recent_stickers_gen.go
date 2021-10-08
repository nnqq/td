// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/nnqq/td/bin"
	"github.com/nnqq/td/tdp"
	"github.com/nnqq/td/tgerr"
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
)

// MessagesRecentStickersNotModified represents TL type `messages.recentStickersNotModified#b17f890`.
// No new recent sticker was found
//
// See https://core.telegram.org/constructor/messages.recentStickersNotModified for reference.
type MessagesRecentStickersNotModified struct {
}

// MessagesRecentStickersNotModifiedTypeID is TL type id of MessagesRecentStickersNotModified.
const MessagesRecentStickersNotModifiedTypeID = 0xb17f890

// construct implements constructor of MessagesRecentStickersClass.
func (r MessagesRecentStickersNotModified) construct() MessagesRecentStickersClass { return &r }

// Ensuring interfaces in compile-time for MessagesRecentStickersNotModified.
var (
	_ bin.Encoder     = &MessagesRecentStickersNotModified{}
	_ bin.Decoder     = &MessagesRecentStickersNotModified{}
	_ bin.BareEncoder = &MessagesRecentStickersNotModified{}
	_ bin.BareDecoder = &MessagesRecentStickersNotModified{}

	_ MessagesRecentStickersClass = &MessagesRecentStickersNotModified{}
)

func (r *MessagesRecentStickersNotModified) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRecentStickersNotModified) String() string {
	if r == nil {
		return "MessagesRecentStickersNotModified(nil)"
	}
	type Alias MessagesRecentStickersNotModified
	return fmt.Sprintf("MessagesRecentStickersNotModified%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesRecentStickersNotModified) TypeID() uint32 {
	return MessagesRecentStickersNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesRecentStickersNotModified) TypeName() string {
	return "messages.recentStickersNotModified"
}

// TypeInfo returns info about TL type.
func (r *MessagesRecentStickersNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.recentStickersNotModified",
		ID:   MessagesRecentStickersNotModifiedTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesRecentStickersNotModified) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.recentStickersNotModified#b17f890 as nil")
	}
	b.PutID(MessagesRecentStickersNotModifiedTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesRecentStickersNotModified) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.recentStickersNotModified#b17f890 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRecentStickersNotModified) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.recentStickersNotModified#b17f890 to nil")
	}
	if err := b.ConsumeID(MessagesRecentStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.recentStickersNotModified#b17f890: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesRecentStickersNotModified) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.recentStickersNotModified#b17f890 to nil")
	}
	return nil
}

// MessagesRecentStickers represents TL type `messages.recentStickers#88d37c56`.
// Recently used stickers
//
// See https://core.telegram.org/constructor/messages.recentStickers for reference.
type MessagesRecentStickers struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
	// Emojis associated to stickers
	Packs []StickerPack
	// Recent stickers
	Stickers []DocumentClass
	// When was each sticker last used
	Dates []int
}

// MessagesRecentStickersTypeID is TL type id of MessagesRecentStickers.
const MessagesRecentStickersTypeID = 0x88d37c56

// construct implements constructor of MessagesRecentStickersClass.
func (r MessagesRecentStickers) construct() MessagesRecentStickersClass { return &r }

// Ensuring interfaces in compile-time for MessagesRecentStickers.
var (
	_ bin.Encoder     = &MessagesRecentStickers{}
	_ bin.Decoder     = &MessagesRecentStickers{}
	_ bin.BareEncoder = &MessagesRecentStickers{}
	_ bin.BareDecoder = &MessagesRecentStickers{}

	_ MessagesRecentStickersClass = &MessagesRecentStickers{}
)

func (r *MessagesRecentStickers) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Hash == 0) {
		return false
	}
	if !(r.Packs == nil) {
		return false
	}
	if !(r.Stickers == nil) {
		return false
	}
	if !(r.Dates == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRecentStickers) String() string {
	if r == nil {
		return "MessagesRecentStickers(nil)"
	}
	type Alias MessagesRecentStickers
	return fmt.Sprintf("MessagesRecentStickers%+v", Alias(*r))
}

// FillFrom fills MessagesRecentStickers from given interface.
func (r *MessagesRecentStickers) FillFrom(from interface {
	GetHash() (value int64)
	GetPacks() (value []StickerPack)
	GetStickers() (value []DocumentClass)
	GetDates() (value []int)
}) {
	r.Hash = from.GetHash()
	r.Packs = from.GetPacks()
	r.Stickers = from.GetStickers()
	r.Dates = from.GetDates()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesRecentStickers) TypeID() uint32 {
	return MessagesRecentStickersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesRecentStickers) TypeName() string {
	return "messages.recentStickers"
}

// TypeInfo returns info about TL type.
func (r *MessagesRecentStickers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.recentStickers",
		ID:   MessagesRecentStickersTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Packs",
			SchemaName: "packs",
		},
		{
			Name:       "Stickers",
			SchemaName: "stickers",
		},
		{
			Name:       "Dates",
			SchemaName: "dates",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesRecentStickers) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.recentStickers#88d37c56 as nil")
	}
	b.PutID(MessagesRecentStickersTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesRecentStickers) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.recentStickers#88d37c56 as nil")
	}
	b.PutLong(r.Hash)
	b.PutVectorHeader(len(r.Packs))
	for idx, v := range r.Packs {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.recentStickers#88d37c56: field packs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Stickers))
	for idx, v := range r.Stickers {
		if v == nil {
			return fmt.Errorf("unable to encode messages.recentStickers#88d37c56: field stickers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.recentStickers#88d37c56: field stickers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Dates))
	for _, v := range r.Dates {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRecentStickers) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.recentStickers#88d37c56 to nil")
	}
	if err := b.ConsumeID(MessagesRecentStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.recentStickers#88d37c56: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesRecentStickers) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.recentStickers#88d37c56 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.recentStickers#88d37c56: field hash: %w", err)
		}
		r.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.recentStickers#88d37c56: field packs: %w", err)
		}

		if headerLen > 0 {
			r.Packs = make([]StickerPack, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StickerPack
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.recentStickers#88d37c56: field packs: %w", err)
			}
			r.Packs = append(r.Packs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.recentStickers#88d37c56: field stickers: %w", err)
		}

		if headerLen > 0 {
			r.Stickers = make([]DocumentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.recentStickers#88d37c56: field stickers: %w", err)
			}
			r.Stickers = append(r.Stickers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.recentStickers#88d37c56: field dates: %w", err)
		}

		if headerLen > 0 {
			r.Dates = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.recentStickers#88d37c56: field dates: %w", err)
			}
			r.Dates = append(r.Dates, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (r *MessagesRecentStickers) GetHash() (value int64) {
	return r.Hash
}

// GetPacks returns value of Packs field.
func (r *MessagesRecentStickers) GetPacks() (value []StickerPack) {
	return r.Packs
}

// GetStickers returns value of Stickers field.
func (r *MessagesRecentStickers) GetStickers() (value []DocumentClass) {
	return r.Stickers
}

// GetDates returns value of Dates field.
func (r *MessagesRecentStickers) GetDates() (value []int) {
	return r.Dates
}

// MapStickers returns field Stickers wrapped in DocumentClassArray helper.
func (r *MessagesRecentStickers) MapStickers() (value DocumentClassArray) {
	return DocumentClassArray(r.Stickers)
}

// MessagesRecentStickersClass represents messages.RecentStickers generic type.
//
// See https://core.telegram.org/type/messages.RecentStickers for reference.
//
// Example:
//  g, err := tg.DecodeMessagesRecentStickers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesRecentStickersNotModified: // messages.recentStickersNotModified#b17f890
//  case *tg.MessagesRecentStickers: // messages.recentStickers#88d37c56
//  default: panic(v)
//  }
type MessagesRecentStickersClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesRecentStickersClass

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

	// AsModified tries to map MessagesRecentStickersClass to MessagesRecentStickers.
	AsModified() (*MessagesRecentStickers, bool)
}

// AsModified tries to map MessagesRecentStickersNotModified to MessagesRecentStickers.
func (r *MessagesRecentStickersNotModified) AsModified() (*MessagesRecentStickers, bool) {
	return nil, false
}

// AsModified tries to map MessagesRecentStickers to MessagesRecentStickers.
func (r *MessagesRecentStickers) AsModified() (*MessagesRecentStickers, bool) {
	return r, true
}

// DecodeMessagesRecentStickers implements binary de-serialization for MessagesRecentStickersClass.
func DecodeMessagesRecentStickers(buf *bin.Buffer) (MessagesRecentStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesRecentStickersNotModifiedTypeID:
		// Decoding messages.recentStickersNotModified#b17f890.
		v := MessagesRecentStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesRecentStickersClass: %w", err)
		}
		return &v, nil
	case MessagesRecentStickersTypeID:
		// Decoding messages.recentStickers#88d37c56.
		v := MessagesRecentStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesRecentStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesRecentStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesRecentStickers boxes the MessagesRecentStickersClass providing a helper.
type MessagesRecentStickersBox struct {
	RecentStickers MessagesRecentStickersClass
}

// Decode implements bin.Decoder for MessagesRecentStickersBox.
func (b *MessagesRecentStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesRecentStickersBox to nil")
	}
	v, err := DecodeMessagesRecentStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RecentStickers = v
	return nil
}

// Encode implements bin.Encode for MessagesRecentStickersBox.
func (b *MessagesRecentStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RecentStickers == nil {
		return fmt.Errorf("unable to encode MessagesRecentStickersClass as nil")
	}
	return b.RecentStickers.Encode(buf)
}
