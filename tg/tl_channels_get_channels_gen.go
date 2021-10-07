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

// ChannelsGetChannelsRequest represents TL type `channels.getChannels#a7f6bbb`.
// Get info about channels/supergroups¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getChannels for reference.
type ChannelsGetChannelsRequest struct {
	// IDs of channels/supergroups to get info about
	ID []InputChannelClass
}

// ChannelsGetChannelsRequestTypeID is TL type id of ChannelsGetChannelsRequest.
const ChannelsGetChannelsRequestTypeID = 0xa7f6bbb

// Ensuring interfaces in compile-time for ChannelsGetChannelsRequest.
var (
	_ bin.Encoder     = &ChannelsGetChannelsRequest{}
	_ bin.Decoder     = &ChannelsGetChannelsRequest{}
	_ bin.BareEncoder = &ChannelsGetChannelsRequest{}
	_ bin.BareDecoder = &ChannelsGetChannelsRequest{}
)

func (g *ChannelsGetChannelsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetChannelsRequest) String() string {
	if g == nil {
		return "ChannelsGetChannelsRequest(nil)"
	}
	type Alias ChannelsGetChannelsRequest
	return fmt.Sprintf("ChannelsGetChannelsRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetChannelsRequest from given interface.
func (g *ChannelsGetChannelsRequest) FillFrom(from interface {
	GetID() (value []InputChannelClass)
}) {
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetChannelsRequest) TypeID() uint32 {
	return ChannelsGetChannelsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetChannelsRequest) TypeName() string {
	return "channels.getChannels"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetChannelsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getChannels",
		ID:   ChannelsGetChannelsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChannelsGetChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getChannels#a7f6bbb as nil")
	}
	b.PutID(ChannelsGetChannelsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetChannelsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getChannels#a7f6bbb as nil")
	}
	b.PutVectorHeader(len(g.ID))
	for idx, v := range g.ID {
		if v == nil {
			return fmt.Errorf("unable to encode channels.getChannels#a7f6bbb: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.getChannels#a7f6bbb: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getChannels#a7f6bbb to nil")
	}
	if err := b.ConsumeID(ChannelsGetChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getChannels#a7f6bbb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetChannelsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getChannels#a7f6bbb to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getChannels#a7f6bbb: field id: %w", err)
		}

		if headerLen > 0 {
			g.ID = make([]InputChannelClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputChannel(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.getChannels#a7f6bbb: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (g *ChannelsGetChannelsRequest) GetID() (value []InputChannelClass) {
	return g.ID
}

// MapID returns field ID wrapped in InputChannelClassArray helper.
func (g *ChannelsGetChannelsRequest) MapID() (value InputChannelClassArray) {
	return InputChannelClassArray(g.ID)
}

// ChannelsGetChannels invokes method channels.getChannels#a7f6bbb returning error if any.
// Get info about channels/supergroups¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MSG_ID_INVALID: Invalid message ID provided
//
// See https://core.telegram.org/method/channels.getChannels for reference.
// Can be used by bots.
func (c *Client) ChannelsGetChannels(ctx context.Context, id []InputChannelClass) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	request := &ChannelsGetChannelsRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
