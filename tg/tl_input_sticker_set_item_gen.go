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

// InputStickerSetItem represents TL type `inputStickerSetItem#ffa0a496`.
// Sticker in a stickerset
//
// See https://core.telegram.org/constructor/inputStickerSetItem for reference.
type InputStickerSetItem struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The sticker
	Document InputDocumentClass
	// Associated emoji
	Emoji string
	// Coordinates for mask sticker
	//
	// Use SetMaskCoords and GetMaskCoords helpers.
	MaskCoords MaskCoords
}

// InputStickerSetItemTypeID is TL type id of InputStickerSetItem.
const InputStickerSetItemTypeID = 0xffa0a496

// Ensuring interfaces in compile-time for InputStickerSetItem.
var (
	_ bin.Encoder     = &InputStickerSetItem{}
	_ bin.Decoder     = &InputStickerSetItem{}
	_ bin.BareEncoder = &InputStickerSetItem{}
	_ bin.BareDecoder = &InputStickerSetItem{}
)

func (i *InputStickerSetItem) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Document == nil) {
		return false
	}
	if !(i.Emoji == "") {
		return false
	}
	if !(i.MaskCoords.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetItem) String() string {
	if i == nil {
		return "InputStickerSetItem(nil)"
	}
	type Alias InputStickerSetItem
	return fmt.Sprintf("InputStickerSetItem%+v", Alias(*i))
}

// FillFrom fills InputStickerSetItem from given interface.
func (i *InputStickerSetItem) FillFrom(from interface {
	GetDocument() (value InputDocumentClass)
	GetEmoji() (value string)
	GetMaskCoords() (value MaskCoords, ok bool)
}) {
	i.Document = from.GetDocument()
	i.Emoji = from.GetEmoji()
	if val, ok := from.GetMaskCoords(); ok {
		i.MaskCoords = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickerSetItem) TypeID() uint32 {
	return InputStickerSetItemTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickerSetItem) TypeName() string {
	return "inputStickerSetItem"
}

// TypeInfo returns info about TL type.
func (i *InputStickerSetItem) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickerSetItem",
		ID:   InputStickerSetItemTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Document",
			SchemaName: "document",
		},
		{
			Name:       "Emoji",
			SchemaName: "emoji",
		},
		{
			Name:       "MaskCoords",
			SchemaName: "mask_coords",
			Null:       !i.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStickerSetItem) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetItem#ffa0a496 as nil")
	}
	b.PutID(InputStickerSetItemTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickerSetItem) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetItem#ffa0a496 as nil")
	}
	if !(i.MaskCoords.Zero()) {
		i.Flags.Set(0)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerSetItem#ffa0a496: field flags: %w", err)
	}
	if i.Document == nil {
		return fmt.Errorf("unable to encode inputStickerSetItem#ffa0a496: field document is nil")
	}
	if err := i.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerSetItem#ffa0a496: field document: %w", err)
	}
	b.PutString(i.Emoji)
	if i.Flags.Has(0) {
		if err := i.MaskCoords.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputStickerSetItem#ffa0a496: field mask_coords: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerSetItem) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetItem#ffa0a496 to nil")
	}
	if err := b.ConsumeID(InputStickerSetItemTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickerSetItem) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetItem#ffa0a496 to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: field document: %w", err)
		}
		i.Document = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: field emoji: %w", err)
		}
		i.Emoji = value
	}
	if i.Flags.Has(0) {
		if err := i.MaskCoords.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: field mask_coords: %w", err)
		}
	}
	return nil
}

// GetDocument returns value of Document field.
func (i *InputStickerSetItem) GetDocument() (value InputDocumentClass) {
	return i.Document
}

// GetEmoji returns value of Emoji field.
func (i *InputStickerSetItem) GetEmoji() (value string) {
	return i.Emoji
}

// SetMaskCoords sets value of MaskCoords conditional field.
func (i *InputStickerSetItem) SetMaskCoords(value MaskCoords) {
	i.Flags.Set(0)
	i.MaskCoords = value
}

// GetMaskCoords returns value of MaskCoords conditional field and
// boolean which is true if field was set.
func (i *InputStickerSetItem) GetMaskCoords() (value MaskCoords, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.MaskCoords, true
}

// GetDocumentAsNotEmpty returns mapped value of Document field.
func (i *InputStickerSetItem) GetDocumentAsNotEmpty() (*InputDocument, bool) {
	return i.Document.AsNotEmpty()
}
