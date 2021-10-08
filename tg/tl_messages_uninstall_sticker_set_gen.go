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

// MessagesUninstallStickerSetRequest represents TL type `messages.uninstallStickerSet#f96e55de`.
// Uninstall a stickerset
//
// See https://core.telegram.org/method/messages.uninstallStickerSet for reference.
type MessagesUninstallStickerSetRequest struct {
	// The stickerset to uninstall
	Stickerset InputStickerSetClass
}

// MessagesUninstallStickerSetRequestTypeID is TL type id of MessagesUninstallStickerSetRequest.
const MessagesUninstallStickerSetRequestTypeID = 0xf96e55de

// Ensuring interfaces in compile-time for MessagesUninstallStickerSetRequest.
var (
	_ bin.Encoder     = &MessagesUninstallStickerSetRequest{}
	_ bin.Decoder     = &MessagesUninstallStickerSetRequest{}
	_ bin.BareEncoder = &MessagesUninstallStickerSetRequest{}
	_ bin.BareDecoder = &MessagesUninstallStickerSetRequest{}
)

func (u *MessagesUninstallStickerSetRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Stickerset == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *MessagesUninstallStickerSetRequest) String() string {
	if u == nil {
		return "MessagesUninstallStickerSetRequest(nil)"
	}
	type Alias MessagesUninstallStickerSetRequest
	return fmt.Sprintf("MessagesUninstallStickerSetRequest%+v", Alias(*u))
}

// FillFrom fills MessagesUninstallStickerSetRequest from given interface.
func (u *MessagesUninstallStickerSetRequest) FillFrom(from interface {
	GetStickerset() (value InputStickerSetClass)
}) {
	u.Stickerset = from.GetStickerset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesUninstallStickerSetRequest) TypeID() uint32 {
	return MessagesUninstallStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesUninstallStickerSetRequest) TypeName() string {
	return "messages.uninstallStickerSet"
}

// TypeInfo returns info about TL type.
func (u *MessagesUninstallStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.uninstallStickerSet",
		ID:   MessagesUninstallStickerSetRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *MessagesUninstallStickerSetRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.uninstallStickerSet#f96e55de as nil")
	}
	b.PutID(MessagesUninstallStickerSetRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *MessagesUninstallStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.uninstallStickerSet#f96e55de as nil")
	}
	if u.Stickerset == nil {
		return fmt.Errorf("unable to encode messages.uninstallStickerSet#f96e55de: field stickerset is nil")
	}
	if err := u.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uninstallStickerSet#f96e55de: field stickerset: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *MessagesUninstallStickerSetRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.uninstallStickerSet#f96e55de to nil")
	}
	if err := b.ConsumeID(MessagesUninstallStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.uninstallStickerSet#f96e55de: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *MessagesUninstallStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.uninstallStickerSet#f96e55de to nil")
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.uninstallStickerSet#f96e55de: field stickerset: %w", err)
		}
		u.Stickerset = value
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (u *MessagesUninstallStickerSetRequest) GetStickerset() (value InputStickerSetClass) {
	return u.Stickerset
}

// MessagesUninstallStickerSet invokes method messages.uninstallStickerSet#f96e55de returning error if any.
// Uninstall a stickerset
//
// Possible errors:
//  400 STICKERSET_INVALID: The provided sticker set is invalid
//
// See https://core.telegram.org/method/messages.uninstallStickerSet for reference.
func (c *Client) MessagesUninstallStickerSet(ctx context.Context, stickerset InputStickerSetClass) (bool, error) {
	var result BoolBox

	request := &MessagesUninstallStickerSetRequest{
		Stickerset: stickerset,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
