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

// MessagesReadMentionsRequest represents TL type `messages.readMentions#f0189d3`.
// Mark mentions as read
//
// See https://core.telegram.org/method/messages.readMentions for reference.
type MessagesReadMentionsRequest struct {
	// Dialog
	Peer InputPeerClass
}

// MessagesReadMentionsRequestTypeID is TL type id of MessagesReadMentionsRequest.
const MessagesReadMentionsRequestTypeID = 0xf0189d3

// Ensuring interfaces in compile-time for MessagesReadMentionsRequest.
var (
	_ bin.Encoder     = &MessagesReadMentionsRequest{}
	_ bin.Decoder     = &MessagesReadMentionsRequest{}
	_ bin.BareEncoder = &MessagesReadMentionsRequest{}
	_ bin.BareDecoder = &MessagesReadMentionsRequest{}
)

func (r *MessagesReadMentionsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReadMentionsRequest) String() string {
	if r == nil {
		return "MessagesReadMentionsRequest(nil)"
	}
	type Alias MessagesReadMentionsRequest
	return fmt.Sprintf("MessagesReadMentionsRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReadMentionsRequest from given interface.
func (r *MessagesReadMentionsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	r.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReadMentionsRequest) TypeID() uint32 {
	return MessagesReadMentionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReadMentionsRequest) TypeName() string {
	return "messages.readMentions"
}

// TypeInfo returns info about TL type.
func (r *MessagesReadMentionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.readMentions",
		ID:   MessagesReadMentionsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReadMentionsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readMentions#f0189d3 as nil")
	}
	b.PutID(MessagesReadMentionsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReadMentionsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readMentions#f0189d3 as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.readMentions#f0189d3: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.readMentions#f0189d3: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReadMentionsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readMentions#f0189d3 to nil")
	}
	if err := b.ConsumeID(MessagesReadMentionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readMentions#f0189d3: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReadMentionsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readMentions#f0189d3 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.readMentions#f0189d3: field peer: %w", err)
		}
		r.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesReadMentionsRequest) GetPeer() (value InputPeerClass) {
	return r.Peer
}

// MessagesReadMentions invokes method messages.readMentions#f0189d3 returning error if any.
// Mark mentions as read
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.readMentions for reference.
func (c *Client) MessagesReadMentions(ctx context.Context, peer InputPeerClass) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	request := &MessagesReadMentionsRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
