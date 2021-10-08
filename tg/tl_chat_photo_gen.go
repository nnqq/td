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

// ChatPhotoEmpty represents TL type `chatPhotoEmpty#37c1011c`.
// Group photo is not set.
//
// See https://core.telegram.org/constructor/chatPhotoEmpty for reference.
type ChatPhotoEmpty struct {
}

// ChatPhotoEmptyTypeID is TL type id of ChatPhotoEmpty.
const ChatPhotoEmptyTypeID = 0x37c1011c

// construct implements constructor of ChatPhotoClass.
func (c ChatPhotoEmpty) construct() ChatPhotoClass { return &c }

// Ensuring interfaces in compile-time for ChatPhotoEmpty.
var (
	_ bin.Encoder     = &ChatPhotoEmpty{}
	_ bin.Decoder     = &ChatPhotoEmpty{}
	_ bin.BareEncoder = &ChatPhotoEmpty{}
	_ bin.BareDecoder = &ChatPhotoEmpty{}

	_ ChatPhotoClass = &ChatPhotoEmpty{}
)

func (c *ChatPhotoEmpty) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatPhotoEmpty) String() string {
	if c == nil {
		return "ChatPhotoEmpty(nil)"
	}
	type Alias ChatPhotoEmpty
	return fmt.Sprintf("ChatPhotoEmpty%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatPhotoEmpty) TypeID() uint32 {
	return ChatPhotoEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatPhotoEmpty) TypeName() string {
	return "chatPhotoEmpty"
}

// TypeInfo returns info about TL type.
func (c *ChatPhotoEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatPhotoEmpty",
		ID:   ChatPhotoEmptyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatPhotoEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPhotoEmpty#37c1011c as nil")
	}
	b.PutID(ChatPhotoEmptyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatPhotoEmpty) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPhotoEmpty#37c1011c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatPhotoEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPhotoEmpty#37c1011c to nil")
	}
	if err := b.ConsumeID(ChatPhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatPhotoEmpty#37c1011c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatPhotoEmpty) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPhotoEmpty#37c1011c to nil")
	}
	return nil
}

// ChatPhoto represents TL type `chatPhoto#1c6e1c11`.
// Group profile photo.
//
// See https://core.telegram.org/constructor/chatPhoto for reference.
type ChatPhoto struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the user has an animated profile picture
	HasVideo bool
	// PhotoID field of ChatPhoto.
	PhotoID int64
	// StrippedThumb field of ChatPhoto.
	//
	// Use SetStrippedThumb and GetStrippedThumb helpers.
	StrippedThumb []byte
	// DC where this photo is stored
	DCID int
}

// ChatPhotoTypeID is TL type id of ChatPhoto.
const ChatPhotoTypeID = 0x1c6e1c11

// construct implements constructor of ChatPhotoClass.
func (c ChatPhoto) construct() ChatPhotoClass { return &c }

// Ensuring interfaces in compile-time for ChatPhoto.
var (
	_ bin.Encoder     = &ChatPhoto{}
	_ bin.Decoder     = &ChatPhoto{}
	_ bin.BareEncoder = &ChatPhoto{}
	_ bin.BareDecoder = &ChatPhoto{}

	_ ChatPhotoClass = &ChatPhoto{}
)

func (c *ChatPhoto) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.HasVideo == false) {
		return false
	}
	if !(c.PhotoID == 0) {
		return false
	}
	if !(c.StrippedThumb == nil) {
		return false
	}
	if !(c.DCID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatPhoto) String() string {
	if c == nil {
		return "ChatPhoto(nil)"
	}
	type Alias ChatPhoto
	return fmt.Sprintf("ChatPhoto%+v", Alias(*c))
}

// FillFrom fills ChatPhoto from given interface.
func (c *ChatPhoto) FillFrom(from interface {
	GetHasVideo() (value bool)
	GetPhotoID() (value int64)
	GetStrippedThumb() (value []byte, ok bool)
	GetDCID() (value int)
}) {
	c.HasVideo = from.GetHasVideo()
	c.PhotoID = from.GetPhotoID()
	if val, ok := from.GetStrippedThumb(); ok {
		c.StrippedThumb = val
	}

	c.DCID = from.GetDCID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatPhoto) TypeID() uint32 {
	return ChatPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatPhoto) TypeName() string {
	return "chatPhoto"
}

// TypeInfo returns info about TL type.
func (c *ChatPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatPhoto",
		ID:   ChatPhotoTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasVideo",
			SchemaName: "has_video",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "PhotoID",
			SchemaName: "photo_id",
		},
		{
			Name:       "StrippedThumb",
			SchemaName: "stripped_thumb",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatPhoto) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPhoto#1c6e1c11 as nil")
	}
	b.PutID(ChatPhotoTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatPhoto) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPhoto#1c6e1c11 as nil")
	}
	if !(c.HasVideo == false) {
		c.Flags.Set(0)
	}
	if !(c.StrippedThumb == nil) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatPhoto#1c6e1c11: field flags: %w", err)
	}
	b.PutLong(c.PhotoID)
	if c.Flags.Has(1) {
		b.PutBytes(c.StrippedThumb)
	}
	b.PutInt(c.DCID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatPhoto) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPhoto#1c6e1c11 to nil")
	}
	if err := b.ConsumeID(ChatPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatPhoto#1c6e1c11: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatPhoto) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPhoto#1c6e1c11 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatPhoto#1c6e1c11: field flags: %w", err)
		}
	}
	c.HasVideo = c.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatPhoto#1c6e1c11: field photo_id: %w", err)
		}
		c.PhotoID = value
	}
	if c.Flags.Has(1) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode chatPhoto#1c6e1c11: field stripped_thumb: %w", err)
		}
		c.StrippedThumb = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatPhoto#1c6e1c11: field dc_id: %w", err)
		}
		c.DCID = value
	}
	return nil
}

// SetHasVideo sets value of HasVideo conditional field.
func (c *ChatPhoto) SetHasVideo(value bool) {
	if value {
		c.Flags.Set(0)
		c.HasVideo = true
	} else {
		c.Flags.Unset(0)
		c.HasVideo = false
	}
}

// GetHasVideo returns value of HasVideo conditional field.
func (c *ChatPhoto) GetHasVideo() (value bool) {
	return c.Flags.Has(0)
}

// GetPhotoID returns value of PhotoID field.
func (c *ChatPhoto) GetPhotoID() (value int64) {
	return c.PhotoID
}

// SetStrippedThumb sets value of StrippedThumb conditional field.
func (c *ChatPhoto) SetStrippedThumb(value []byte) {
	c.Flags.Set(1)
	c.StrippedThumb = value
}

// GetStrippedThumb returns value of StrippedThumb conditional field and
// boolean which is true if field was set.
func (c *ChatPhoto) GetStrippedThumb() (value []byte, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.StrippedThumb, true
}

// GetDCID returns value of DCID field.
func (c *ChatPhoto) GetDCID() (value int) {
	return c.DCID
}

// ChatPhotoClass represents ChatPhoto generic type.
//
// See https://core.telegram.org/type/ChatPhoto for reference.
//
// Example:
//  g, err := tg.DecodeChatPhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.ChatPhotoEmpty: // chatPhotoEmpty#37c1011c
//  case *tg.ChatPhoto: // chatPhoto#1c6e1c11
//  default: panic(v)
//  }
type ChatPhotoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatPhotoClass

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

	// AsNotEmpty tries to map ChatPhotoClass to ChatPhoto.
	AsNotEmpty() (*ChatPhoto, bool)
}

// AsNotEmpty tries to map ChatPhotoEmpty to ChatPhoto.
func (c *ChatPhotoEmpty) AsNotEmpty() (*ChatPhoto, bool) {
	return nil, false
}

// AsNotEmpty tries to map ChatPhoto to ChatPhoto.
func (c *ChatPhoto) AsNotEmpty() (*ChatPhoto, bool) {
	return c, true
}

// DecodeChatPhoto implements binary de-serialization for ChatPhotoClass.
func DecodeChatPhoto(buf *bin.Buffer) (ChatPhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatPhotoEmptyTypeID:
		// Decoding chatPhotoEmpty#37c1011c.
		v := ChatPhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatPhotoClass: %w", err)
		}
		return &v, nil
	case ChatPhotoTypeID:
		// Decoding chatPhoto#1c6e1c11.
		v := ChatPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatPhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatPhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatPhoto boxes the ChatPhotoClass providing a helper.
type ChatPhotoBox struct {
	ChatPhoto ChatPhotoClass
}

// Decode implements bin.Decoder for ChatPhotoBox.
func (b *ChatPhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatPhotoBox to nil")
	}
	v, err := DecodeChatPhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatPhoto = v
	return nil
}

// Encode implements bin.Encode for ChatPhotoBox.
func (b *ChatPhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatPhoto == nil {
		return fmt.Errorf("unable to encode ChatPhotoClass as nil")
	}
	return b.ChatPhoto.Encode(buf)
}
