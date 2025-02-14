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

// PhoneToggleGroupCallSettingsRequest represents TL type `phone.toggleGroupCallSettings#74bbb43d`.
//
// See https://core.telegram.org/method/phone.toggleGroupCallSettings for reference.
type PhoneToggleGroupCallSettingsRequest struct {
	// Flags field of PhoneToggleGroupCallSettingsRequest.
	Flags bin.Fields
	// ResetInviteHash field of PhoneToggleGroupCallSettingsRequest.
	ResetInviteHash bool
	// Call field of PhoneToggleGroupCallSettingsRequest.
	Call InputGroupCall
	// JoinMuted field of PhoneToggleGroupCallSettingsRequest.
	//
	// Use SetJoinMuted and GetJoinMuted helpers.
	JoinMuted bool
}

// PhoneToggleGroupCallSettingsRequestTypeID is TL type id of PhoneToggleGroupCallSettingsRequest.
const PhoneToggleGroupCallSettingsRequestTypeID = 0x74bbb43d

// Ensuring interfaces in compile-time for PhoneToggleGroupCallSettingsRequest.
var (
	_ bin.Encoder     = &PhoneToggleGroupCallSettingsRequest{}
	_ bin.Decoder     = &PhoneToggleGroupCallSettingsRequest{}
	_ bin.BareEncoder = &PhoneToggleGroupCallSettingsRequest{}
	_ bin.BareDecoder = &PhoneToggleGroupCallSettingsRequest{}
)

func (t *PhoneToggleGroupCallSettingsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.ResetInviteHash == false) {
		return false
	}
	if !(t.Call.Zero()) {
		return false
	}
	if !(t.JoinMuted == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *PhoneToggleGroupCallSettingsRequest) String() string {
	if t == nil {
		return "PhoneToggleGroupCallSettingsRequest(nil)"
	}
	type Alias PhoneToggleGroupCallSettingsRequest
	return fmt.Sprintf("PhoneToggleGroupCallSettingsRequest%+v", Alias(*t))
}

// FillFrom fills PhoneToggleGroupCallSettingsRequest from given interface.
func (t *PhoneToggleGroupCallSettingsRequest) FillFrom(from interface {
	GetResetInviteHash() (value bool)
	GetCall() (value InputGroupCall)
	GetJoinMuted() (value bool, ok bool)
}) {
	t.ResetInviteHash = from.GetResetInviteHash()
	t.Call = from.GetCall()
	if val, ok := from.GetJoinMuted(); ok {
		t.JoinMuted = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneToggleGroupCallSettingsRequest) TypeID() uint32 {
	return PhoneToggleGroupCallSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneToggleGroupCallSettingsRequest) TypeName() string {
	return "phone.toggleGroupCallSettings"
}

// TypeInfo returns info about TL type.
func (t *PhoneToggleGroupCallSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.toggleGroupCallSettings",
		ID:   PhoneToggleGroupCallSettingsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ResetInviteHash",
			SchemaName: "reset_invite_hash",
			Null:       !t.Flags.Has(1),
		},
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "JoinMuted",
			SchemaName: "join_muted",
			Null:       !t.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *PhoneToggleGroupCallSettingsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode phone.toggleGroupCallSettings#74bbb43d as nil")
	}
	b.PutID(PhoneToggleGroupCallSettingsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *PhoneToggleGroupCallSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode phone.toggleGroupCallSettings#74bbb43d as nil")
	}
	if !(t.ResetInviteHash == false) {
		t.Flags.Set(1)
	}
	if !(t.JoinMuted == false) {
		t.Flags.Set(0)
	}
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.toggleGroupCallSettings#74bbb43d: field flags: %w", err)
	}
	if err := t.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.toggleGroupCallSettings#74bbb43d: field call: %w", err)
	}
	if t.Flags.Has(0) {
		b.PutBool(t.JoinMuted)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *PhoneToggleGroupCallSettingsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode phone.toggleGroupCallSettings#74bbb43d to nil")
	}
	if err := b.ConsumeID(PhoneToggleGroupCallSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.toggleGroupCallSettings#74bbb43d: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *PhoneToggleGroupCallSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode phone.toggleGroupCallSettings#74bbb43d to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallSettings#74bbb43d: field flags: %w", err)
		}
	}
	t.ResetInviteHash = t.Flags.Has(1)
	{
		if err := t.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallSettings#74bbb43d: field call: %w", err)
		}
	}
	if t.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallSettings#74bbb43d: field join_muted: %w", err)
		}
		t.JoinMuted = value
	}
	return nil
}

// SetResetInviteHash sets value of ResetInviteHash conditional field.
func (t *PhoneToggleGroupCallSettingsRequest) SetResetInviteHash(value bool) {
	if value {
		t.Flags.Set(1)
		t.ResetInviteHash = true
	} else {
		t.Flags.Unset(1)
		t.ResetInviteHash = false
	}
}

// GetResetInviteHash returns value of ResetInviteHash conditional field.
func (t *PhoneToggleGroupCallSettingsRequest) GetResetInviteHash() (value bool) {
	return t.Flags.Has(1)
}

// GetCall returns value of Call field.
func (t *PhoneToggleGroupCallSettingsRequest) GetCall() (value InputGroupCall) {
	return t.Call
}

// SetJoinMuted sets value of JoinMuted conditional field.
func (t *PhoneToggleGroupCallSettingsRequest) SetJoinMuted(value bool) {
	t.Flags.Set(0)
	t.JoinMuted = value
}

// GetJoinMuted returns value of JoinMuted conditional field and
// boolean which is true if field was set.
func (t *PhoneToggleGroupCallSettingsRequest) GetJoinMuted() (value bool, ok bool) {
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.JoinMuted, true
}

// PhoneToggleGroupCallSettings invokes method phone.toggleGroupCallSettings#74bbb43d returning error if any.
//
// See https://core.telegram.org/method/phone.toggleGroupCallSettings for reference.
func (c *Client) PhoneToggleGroupCallSettings(ctx context.Context, request *PhoneToggleGroupCallSettingsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
