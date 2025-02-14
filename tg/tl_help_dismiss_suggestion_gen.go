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

// HelpDismissSuggestionRequest represents TL type `help.dismissSuggestion#f50dbaa1`.
// Dismiss a suggestion
//
// See https://core.telegram.org/method/help.dismissSuggestion for reference.
type HelpDismissSuggestionRequest struct {
	// Peer field of HelpDismissSuggestionRequest.
	Peer InputPeerClass
	// Suggestion
	Suggestion string
}

// HelpDismissSuggestionRequestTypeID is TL type id of HelpDismissSuggestionRequest.
const HelpDismissSuggestionRequestTypeID = 0xf50dbaa1

// Ensuring interfaces in compile-time for HelpDismissSuggestionRequest.
var (
	_ bin.Encoder     = &HelpDismissSuggestionRequest{}
	_ bin.Decoder     = &HelpDismissSuggestionRequest{}
	_ bin.BareEncoder = &HelpDismissSuggestionRequest{}
	_ bin.BareDecoder = &HelpDismissSuggestionRequest{}
)

func (d *HelpDismissSuggestionRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Peer == nil) {
		return false
	}
	if !(d.Suggestion == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *HelpDismissSuggestionRequest) String() string {
	if d == nil {
		return "HelpDismissSuggestionRequest(nil)"
	}
	type Alias HelpDismissSuggestionRequest
	return fmt.Sprintf("HelpDismissSuggestionRequest%+v", Alias(*d))
}

// FillFrom fills HelpDismissSuggestionRequest from given interface.
func (d *HelpDismissSuggestionRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetSuggestion() (value string)
}) {
	d.Peer = from.GetPeer()
	d.Suggestion = from.GetSuggestion()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpDismissSuggestionRequest) TypeID() uint32 {
	return HelpDismissSuggestionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpDismissSuggestionRequest) TypeName() string {
	return "help.dismissSuggestion"
}

// TypeInfo returns info about TL type.
func (d *HelpDismissSuggestionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.dismissSuggestion",
		ID:   HelpDismissSuggestionRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Suggestion",
			SchemaName: "suggestion",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *HelpDismissSuggestionRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode help.dismissSuggestion#f50dbaa1 as nil")
	}
	b.PutID(HelpDismissSuggestionRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *HelpDismissSuggestionRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode help.dismissSuggestion#f50dbaa1 as nil")
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode help.dismissSuggestion#f50dbaa1: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.dismissSuggestion#f50dbaa1: field peer: %w", err)
	}
	b.PutString(d.Suggestion)
	return nil
}

// Decode implements bin.Decoder.
func (d *HelpDismissSuggestionRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode help.dismissSuggestion#f50dbaa1 to nil")
	}
	if err := b.ConsumeID(HelpDismissSuggestionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.dismissSuggestion#f50dbaa1: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *HelpDismissSuggestionRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode help.dismissSuggestion#f50dbaa1 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.dismissSuggestion#f50dbaa1: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.dismissSuggestion#f50dbaa1: field suggestion: %w", err)
		}
		d.Suggestion = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (d *HelpDismissSuggestionRequest) GetPeer() (value InputPeerClass) {
	return d.Peer
}

// GetSuggestion returns value of Suggestion field.
func (d *HelpDismissSuggestionRequest) GetSuggestion() (value string) {
	return d.Suggestion
}

// HelpDismissSuggestion invokes method help.dismissSuggestion#f50dbaa1 returning error if any.
// Dismiss a suggestion
//
// See https://core.telegram.org/method/help.dismissSuggestion for reference.
// Can be used by bots.
func (c *Client) HelpDismissSuggestion(ctx context.Context, request *HelpDismissSuggestionRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
