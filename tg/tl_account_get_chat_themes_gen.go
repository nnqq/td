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

// AccountGetChatThemesRequest represents TL type `account.getChatThemes#d6d71d7b`.
//
// See https://core.telegram.org/method/account.getChatThemes for reference.
type AccountGetChatThemesRequest struct {
	// Hash field of AccountGetChatThemesRequest.
	Hash int
}

// AccountGetChatThemesRequestTypeID is TL type id of AccountGetChatThemesRequest.
const AccountGetChatThemesRequestTypeID = 0xd6d71d7b

// Ensuring interfaces in compile-time for AccountGetChatThemesRequest.
var (
	_ bin.Encoder     = &AccountGetChatThemesRequest{}
	_ bin.Decoder     = &AccountGetChatThemesRequest{}
	_ bin.BareEncoder = &AccountGetChatThemesRequest{}
	_ bin.BareDecoder = &AccountGetChatThemesRequest{}
)

func (g *AccountGetChatThemesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetChatThemesRequest) String() string {
	if g == nil {
		return "AccountGetChatThemesRequest(nil)"
	}
	type Alias AccountGetChatThemesRequest
	return fmt.Sprintf("AccountGetChatThemesRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetChatThemesRequest from given interface.
func (g *AccountGetChatThemesRequest) FillFrom(from interface {
	GetHash() (value int)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetChatThemesRequest) TypeID() uint32 {
	return AccountGetChatThemesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetChatThemesRequest) TypeName() string {
	return "account.getChatThemes"
}

// TypeInfo returns info about TL type.
func (g *AccountGetChatThemesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getChatThemes",
		ID:   AccountGetChatThemesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetChatThemesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getChatThemes#d6d71d7b as nil")
	}
	b.PutID(AccountGetChatThemesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetChatThemesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getChatThemes#d6d71d7b as nil")
	}
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetChatThemesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getChatThemes#d6d71d7b to nil")
	}
	if err := b.ConsumeID(AccountGetChatThemesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getChatThemes#d6d71d7b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetChatThemesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getChatThemes#d6d71d7b to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.getChatThemes#d6d71d7b: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *AccountGetChatThemesRequest) GetHash() (value int) {
	return g.Hash
}

// AccountGetChatThemes invokes method account.getChatThemes#d6d71d7b returning error if any.
//
// See https://core.telegram.org/method/account.getChatThemes for reference.
func (c *Client) AccountGetChatThemes(ctx context.Context, hash int) (AccountChatThemesClass, error) {
	var result AccountChatThemesBox

	request := &AccountGetChatThemesRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChatThemes, nil
}
