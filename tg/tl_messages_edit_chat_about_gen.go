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

// MessagesEditChatAboutRequest represents TL type `messages.editChatAbout#def60797`.
// Edit the description of a group/supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/messages.editChatAbout for reference.
type MessagesEditChatAboutRequest struct {
	// The group/supergroup/channel¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Peer InputPeerClass
	// The new description
	About string
}

// MessagesEditChatAboutRequestTypeID is TL type id of MessagesEditChatAboutRequest.
const MessagesEditChatAboutRequestTypeID = 0xdef60797

// Ensuring interfaces in compile-time for MessagesEditChatAboutRequest.
var (
	_ bin.Encoder     = &MessagesEditChatAboutRequest{}
	_ bin.Decoder     = &MessagesEditChatAboutRequest{}
	_ bin.BareEncoder = &MessagesEditChatAboutRequest{}
	_ bin.BareDecoder = &MessagesEditChatAboutRequest{}
)

func (e *MessagesEditChatAboutRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Peer == nil) {
		return false
	}
	if !(e.About == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEditChatAboutRequest) String() string {
	if e == nil {
		return "MessagesEditChatAboutRequest(nil)"
	}
	type Alias MessagesEditChatAboutRequest
	return fmt.Sprintf("MessagesEditChatAboutRequest%+v", Alias(*e))
}

// FillFrom fills MessagesEditChatAboutRequest from given interface.
func (e *MessagesEditChatAboutRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetAbout() (value string)
}) {
	e.Peer = from.GetPeer()
	e.About = from.GetAbout()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesEditChatAboutRequest) TypeID() uint32 {
	return MessagesEditChatAboutRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesEditChatAboutRequest) TypeName() string {
	return "messages.editChatAbout"
}

// TypeInfo returns info about TL type.
func (e *MessagesEditChatAboutRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.editChatAbout",
		ID:   MessagesEditChatAboutRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "About",
			SchemaName: "about",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *MessagesEditChatAboutRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatAbout#def60797 as nil")
	}
	b.PutID(MessagesEditChatAboutRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesEditChatAboutRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatAbout#def60797 as nil")
	}
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.editChatAbout#def60797: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editChatAbout#def60797: field peer: %w", err)
	}
	b.PutString(e.About)
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEditChatAboutRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatAbout#def60797 to nil")
	}
	if err := b.ConsumeID(MessagesEditChatAboutRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editChatAbout#def60797: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesEditChatAboutRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatAbout#def60797 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAbout#def60797: field peer: %w", err)
		}
		e.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAbout#def60797: field about: %w", err)
		}
		e.About = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (e *MessagesEditChatAboutRequest) GetPeer() (value InputPeerClass) {
	return e.Peer
}

// GetAbout returns value of About field.
func (e *MessagesEditChatAboutRequest) GetAbout() (value string) {
	return e.About
}

// MessagesEditChatAbout invokes method messages.editChatAbout#def60797 returning error if any.
// Edit the description of a group/supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ABOUT_NOT_MODIFIED: About text has not changed
//  400 CHAT_ABOUT_TOO_LONG: Chat about too long
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.editChatAbout for reference.
// Can be used by bots.
func (c *Client) MessagesEditChatAbout(ctx context.Context, request *MessagesEditChatAboutRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
