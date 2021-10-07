// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// InputStickerSetShortName represents TL type `inputStickerSetShortName#861cc8a0`.
// Stickerset by short name, from tg://addstickers?set=short_name
//
// See https://core.telegram.org/constructor/inputStickerSetShortName for reference.
type InputStickerSetShortName struct {
	// From tg://addstickers?set=short_name
	ShortName string
}

// InputStickerSetShortNameTypeID is TL type id of InputStickerSetShortName.
const InputStickerSetShortNameTypeID = 0x861cc8a0

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetShortName) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetShortName.
var (
	_ bin.Encoder     = &InputStickerSetShortName{}
	_ bin.Decoder     = &InputStickerSetShortName{}
	_ bin.BareEncoder = &InputStickerSetShortName{}
	_ bin.BareDecoder = &InputStickerSetShortName{}

	_ InputStickerSetClass = &InputStickerSetShortName{}
)

func (i *InputStickerSetShortName) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ShortName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetShortName) String() string {
	if i == nil {
		return "InputStickerSetShortName(nil)"
	}
	type Alias InputStickerSetShortName
	return fmt.Sprintf("InputStickerSetShortName%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickerSetShortName) TypeID() uint32 {
	return InputStickerSetShortNameTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickerSetShortName) TypeName() string {
	return "inputStickerSetShortName"
}

// TypeInfo returns info about TL type.
func (i *InputStickerSetShortName) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickerSetShortName",
		ID:   InputStickerSetShortNameTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStickerSetShortName) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetShortName#861cc8a0 as nil")
	}
	b.PutID(InputStickerSetShortNameTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickerSetShortName) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetShortName#861cc8a0 as nil")
	}
	b.PutString(i.ShortName)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerSetShortName) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetShortName#861cc8a0 to nil")
	}
	if err := b.ConsumeID(InputStickerSetShortNameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetShortName#861cc8a0: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickerSetShortName) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetShortName#861cc8a0 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetShortName#861cc8a0: field short_name: %w", err)
		}
		i.ShortName = value
	}
	return nil
}

// GetShortName returns value of ShortName field.
func (i *InputStickerSetShortName) GetShortName() (value string) {
	return i.ShortName
}

// InputStickerSetEmpty represents TL type `inputStickerSetEmpty#ffb62b95`.
// Empty constructor
//
// See https://core.telegram.org/constructor/inputStickerSetEmpty for reference.
type InputStickerSetEmpty struct {
}

// InputStickerSetEmptyTypeID is TL type id of InputStickerSetEmpty.
const InputStickerSetEmptyTypeID = 0xffb62b95

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetEmpty) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetEmpty.
var (
	_ bin.Encoder     = &InputStickerSetEmpty{}
	_ bin.Decoder     = &InputStickerSetEmpty{}
	_ bin.BareEncoder = &InputStickerSetEmpty{}
	_ bin.BareDecoder = &InputStickerSetEmpty{}

	_ InputStickerSetClass = &InputStickerSetEmpty{}
)

func (i *InputStickerSetEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetEmpty) String() string {
	if i == nil {
		return "InputStickerSetEmpty(nil)"
	}
	type Alias InputStickerSetEmpty
	return fmt.Sprintf("InputStickerSetEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickerSetEmpty) TypeID() uint32 {
	return InputStickerSetEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickerSetEmpty) TypeName() string {
	return "inputStickerSetEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputStickerSetEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickerSetEmpty",
		ID:   InputStickerSetEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStickerSetEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetEmpty#ffb62b95 as nil")
	}
	b.PutID(InputStickerSetEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickerSetEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetEmpty#ffb62b95 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerSetEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetEmpty#ffb62b95 to nil")
	}
	if err := b.ConsumeID(InputStickerSetEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetEmpty#ffb62b95: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickerSetEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetEmpty#ffb62b95 to nil")
	}
	return nil
}

// InputStickerSetClass represents InputStickerSet generic type.
//
// See https://core.telegram.org/type/InputStickerSet for reference.
//
// Example:
//  g, err := e2e.DecodeInputStickerSet(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *e2e.InputStickerSetShortName: // inputStickerSetShortName#861cc8a0
//  case *e2e.InputStickerSetEmpty: // inputStickerSetEmpty#ffb62b95
//  default: panic(v)
//  }
type InputStickerSetClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputStickerSetClass

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
}

// DecodeInputStickerSet implements binary de-serialization for InputStickerSetClass.
func DecodeInputStickerSet(buf *bin.Buffer) (InputStickerSetClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStickerSetShortNameTypeID:
		// Decoding inputStickerSetShortName#861cc8a0.
		v := InputStickerSetShortName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", err)
		}
		return &v, nil
	case InputStickerSetEmptyTypeID:
		// Decoding inputStickerSetEmpty#ffb62b95.
		v := InputStickerSetEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputStickerSet boxes the InputStickerSetClass providing a helper.
type InputStickerSetBox struct {
	InputStickerSet InputStickerSetClass
}

// Decode implements bin.Decoder for InputStickerSetBox.
func (b *InputStickerSetBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStickerSetBox to nil")
	}
	v, err := DecodeInputStickerSet(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStickerSet = v
	return nil
}

// Encode implements bin.Encode for InputStickerSetBox.
func (b *InputStickerSetBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputStickerSet == nil {
		return fmt.Errorf("unable to encode InputStickerSetClass as nil")
	}
	return b.InputStickerSet.Encode(buf)
}
