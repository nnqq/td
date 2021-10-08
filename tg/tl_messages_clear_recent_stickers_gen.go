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

// MessagesClearRecentStickersRequest represents TL type `messages.clearRecentStickers#8999602d`.
// Clear recent stickers
//
// See https://core.telegram.org/method/messages.clearRecentStickers for reference.
type MessagesClearRecentStickersRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag to clear the list of stickers recently attached to photo or video files
	Attached bool
}

// MessagesClearRecentStickersRequestTypeID is TL type id of MessagesClearRecentStickersRequest.
const MessagesClearRecentStickersRequestTypeID = 0x8999602d

// Ensuring interfaces in compile-time for MessagesClearRecentStickersRequest.
var (
	_ bin.Encoder     = &MessagesClearRecentStickersRequest{}
	_ bin.Decoder     = &MessagesClearRecentStickersRequest{}
	_ bin.BareEncoder = &MessagesClearRecentStickersRequest{}
	_ bin.BareDecoder = &MessagesClearRecentStickersRequest{}
)

func (c *MessagesClearRecentStickersRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Attached == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesClearRecentStickersRequest) String() string {
	if c == nil {
		return "MessagesClearRecentStickersRequest(nil)"
	}
	type Alias MessagesClearRecentStickersRequest
	return fmt.Sprintf("MessagesClearRecentStickersRequest%+v", Alias(*c))
}

// FillFrom fills MessagesClearRecentStickersRequest from given interface.
func (c *MessagesClearRecentStickersRequest) FillFrom(from interface {
	GetAttached() (value bool)
}) {
	c.Attached = from.GetAttached()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesClearRecentStickersRequest) TypeID() uint32 {
	return MessagesClearRecentStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesClearRecentStickersRequest) TypeName() string {
	return "messages.clearRecentStickers"
}

// TypeInfo returns info about TL type.
func (c *MessagesClearRecentStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.clearRecentStickers",
		ID:   MessagesClearRecentStickersRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Attached",
			SchemaName: "attached",
			Null:       !c.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesClearRecentStickersRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.clearRecentStickers#8999602d as nil")
	}
	b.PutID(MessagesClearRecentStickersRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesClearRecentStickersRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.clearRecentStickers#8999602d as nil")
	}
	if !(c.Attached == false) {
		c.Flags.Set(0)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.clearRecentStickers#8999602d: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesClearRecentStickersRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.clearRecentStickers#8999602d to nil")
	}
	if err := b.ConsumeID(MessagesClearRecentStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.clearRecentStickers#8999602d: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesClearRecentStickersRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.clearRecentStickers#8999602d to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.clearRecentStickers#8999602d: field flags: %w", err)
		}
	}
	c.Attached = c.Flags.Has(0)
	return nil
}

// SetAttached sets value of Attached conditional field.
func (c *MessagesClearRecentStickersRequest) SetAttached(value bool) {
	if value {
		c.Flags.Set(0)
		c.Attached = true
	} else {
		c.Flags.Unset(0)
		c.Attached = false
	}
}

// GetAttached returns value of Attached conditional field.
func (c *MessagesClearRecentStickersRequest) GetAttached() (value bool) {
	return c.Flags.Has(0)
}

// MessagesClearRecentStickers invokes method messages.clearRecentStickers#8999602d returning error if any.
// Clear recent stickers
//
// See https://core.telegram.org/method/messages.clearRecentStickers for reference.
func (c *Client) MessagesClearRecentStickers(ctx context.Context, request *MessagesClearRecentStickersRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
