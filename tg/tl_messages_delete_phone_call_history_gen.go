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

// MessagesDeletePhoneCallHistoryRequest represents TL type `messages.deletePhoneCallHistory#f9cbe409`.
//
// See https://core.telegram.org/method/messages.deletePhoneCallHistory for reference.
type MessagesDeletePhoneCallHistoryRequest struct {
	// Flags field of MessagesDeletePhoneCallHistoryRequest.
	Flags bin.Fields
	// Revoke field of MessagesDeletePhoneCallHistoryRequest.
	Revoke bool
}

// MessagesDeletePhoneCallHistoryRequestTypeID is TL type id of MessagesDeletePhoneCallHistoryRequest.
const MessagesDeletePhoneCallHistoryRequestTypeID = 0xf9cbe409

// Ensuring interfaces in compile-time for MessagesDeletePhoneCallHistoryRequest.
var (
	_ bin.Encoder     = &MessagesDeletePhoneCallHistoryRequest{}
	_ bin.Decoder     = &MessagesDeletePhoneCallHistoryRequest{}
	_ bin.BareEncoder = &MessagesDeletePhoneCallHistoryRequest{}
	_ bin.BareDecoder = &MessagesDeletePhoneCallHistoryRequest{}
)

func (d *MessagesDeletePhoneCallHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Revoke == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeletePhoneCallHistoryRequest) String() string {
	if d == nil {
		return "MessagesDeletePhoneCallHistoryRequest(nil)"
	}
	type Alias MessagesDeletePhoneCallHistoryRequest
	return fmt.Sprintf("MessagesDeletePhoneCallHistoryRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeletePhoneCallHistoryRequest from given interface.
func (d *MessagesDeletePhoneCallHistoryRequest) FillFrom(from interface {
	GetRevoke() (value bool)
}) {
	d.Revoke = from.GetRevoke()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeletePhoneCallHistoryRequest) TypeID() uint32 {
	return MessagesDeletePhoneCallHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeletePhoneCallHistoryRequest) TypeName() string {
	return "messages.deletePhoneCallHistory"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeletePhoneCallHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deletePhoneCallHistory",
		ID:   MessagesDeletePhoneCallHistoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Revoke",
			SchemaName: "revoke",
			Null:       !d.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDeletePhoneCallHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deletePhoneCallHistory#f9cbe409 as nil")
	}
	b.PutID(MessagesDeletePhoneCallHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeletePhoneCallHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deletePhoneCallHistory#f9cbe409 as nil")
	}
	if !(d.Revoke == false) {
		d.Flags.Set(0)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deletePhoneCallHistory#f9cbe409: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDeletePhoneCallHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deletePhoneCallHistory#f9cbe409 to nil")
	}
	if err := b.ConsumeID(MessagesDeletePhoneCallHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deletePhoneCallHistory#f9cbe409: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeletePhoneCallHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deletePhoneCallHistory#f9cbe409 to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.deletePhoneCallHistory#f9cbe409: field flags: %w", err)
		}
	}
	d.Revoke = d.Flags.Has(0)
	return nil
}

// SetRevoke sets value of Revoke conditional field.
func (d *MessagesDeletePhoneCallHistoryRequest) SetRevoke(value bool) {
	if value {
		d.Flags.Set(0)
		d.Revoke = true
	} else {
		d.Flags.Unset(0)
		d.Revoke = false
	}
}

// GetRevoke returns value of Revoke conditional field.
func (d *MessagesDeletePhoneCallHistoryRequest) GetRevoke() (value bool) {
	return d.Flags.Has(0)
}

// MessagesDeletePhoneCallHistory invokes method messages.deletePhoneCallHistory#f9cbe409 returning error if any.
//
// See https://core.telegram.org/method/messages.deletePhoneCallHistory for reference.
func (c *Client) MessagesDeletePhoneCallHistory(ctx context.Context, request *MessagesDeletePhoneCallHistoryRequest) (*MessagesAffectedFoundMessages, error) {
	var result MessagesAffectedFoundMessages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
