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

// PhoneExportGroupCallInviteRequest represents TL type `phone.exportGroupCallInvite#e6aa647f`.
//
// See https://core.telegram.org/method/phone.exportGroupCallInvite for reference.
type PhoneExportGroupCallInviteRequest struct {
	// Flags field of PhoneExportGroupCallInviteRequest.
	Flags bin.Fields
	// CanSelfUnmute field of PhoneExportGroupCallInviteRequest.
	CanSelfUnmute bool
	// Call field of PhoneExportGroupCallInviteRequest.
	Call InputGroupCall
}

// PhoneExportGroupCallInviteRequestTypeID is TL type id of PhoneExportGroupCallInviteRequest.
const PhoneExportGroupCallInviteRequestTypeID = 0xe6aa647f

// Ensuring interfaces in compile-time for PhoneExportGroupCallInviteRequest.
var (
	_ bin.Encoder     = &PhoneExportGroupCallInviteRequest{}
	_ bin.Decoder     = &PhoneExportGroupCallInviteRequest{}
	_ bin.BareEncoder = &PhoneExportGroupCallInviteRequest{}
	_ bin.BareDecoder = &PhoneExportGroupCallInviteRequest{}
)

func (e *PhoneExportGroupCallInviteRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.CanSelfUnmute == false) {
		return false
	}
	if !(e.Call.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *PhoneExportGroupCallInviteRequest) String() string {
	if e == nil {
		return "PhoneExportGroupCallInviteRequest(nil)"
	}
	type Alias PhoneExportGroupCallInviteRequest
	return fmt.Sprintf("PhoneExportGroupCallInviteRequest%+v", Alias(*e))
}

// FillFrom fills PhoneExportGroupCallInviteRequest from given interface.
func (e *PhoneExportGroupCallInviteRequest) FillFrom(from interface {
	GetCanSelfUnmute() (value bool)
	GetCall() (value InputGroupCall)
}) {
	e.CanSelfUnmute = from.GetCanSelfUnmute()
	e.Call = from.GetCall()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneExportGroupCallInviteRequest) TypeID() uint32 {
	return PhoneExportGroupCallInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneExportGroupCallInviteRequest) TypeName() string {
	return "phone.exportGroupCallInvite"
}

// TypeInfo returns info about TL type.
func (e *PhoneExportGroupCallInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.exportGroupCallInvite",
		ID:   PhoneExportGroupCallInviteRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CanSelfUnmute",
			SchemaName: "can_self_unmute",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "Call",
			SchemaName: "call",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *PhoneExportGroupCallInviteRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode phone.exportGroupCallInvite#e6aa647f as nil")
	}
	b.PutID(PhoneExportGroupCallInviteRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *PhoneExportGroupCallInviteRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode phone.exportGroupCallInvite#e6aa647f as nil")
	}
	if !(e.CanSelfUnmute == false) {
		e.Flags.Set(0)
	}
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.exportGroupCallInvite#e6aa647f: field flags: %w", err)
	}
	if err := e.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.exportGroupCallInvite#e6aa647f: field call: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *PhoneExportGroupCallInviteRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode phone.exportGroupCallInvite#e6aa647f to nil")
	}
	if err := b.ConsumeID(PhoneExportGroupCallInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.exportGroupCallInvite#e6aa647f: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *PhoneExportGroupCallInviteRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode phone.exportGroupCallInvite#e6aa647f to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.exportGroupCallInvite#e6aa647f: field flags: %w", err)
		}
	}
	e.CanSelfUnmute = e.Flags.Has(0)
	{
		if err := e.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.exportGroupCallInvite#e6aa647f: field call: %w", err)
		}
	}
	return nil
}

// SetCanSelfUnmute sets value of CanSelfUnmute conditional field.
func (e *PhoneExportGroupCallInviteRequest) SetCanSelfUnmute(value bool) {
	if value {
		e.Flags.Set(0)
		e.CanSelfUnmute = true
	} else {
		e.Flags.Unset(0)
		e.CanSelfUnmute = false
	}
}

// GetCanSelfUnmute returns value of CanSelfUnmute conditional field.
func (e *PhoneExportGroupCallInviteRequest) GetCanSelfUnmute() (value bool) {
	return e.Flags.Has(0)
}

// GetCall returns value of Call field.
func (e *PhoneExportGroupCallInviteRequest) GetCall() (value InputGroupCall) {
	return e.Call
}

// PhoneExportGroupCallInvite invokes method phone.exportGroupCallInvite#e6aa647f returning error if any.
//
// See https://core.telegram.org/method/phone.exportGroupCallInvite for reference.
func (c *Client) PhoneExportGroupCallInvite(ctx context.Context, request *PhoneExportGroupCallInviteRequest) (*PhoneExportedGroupCallInvite, error) {
	var result PhoneExportedGroupCallInvite

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
