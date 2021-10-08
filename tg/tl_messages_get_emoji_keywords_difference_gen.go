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

// MessagesGetEmojiKeywordsDifferenceRequest represents TL type `messages.getEmojiKeywordsDifference#1508b6af`.
// Get changed emoji keywords
//
// See https://core.telegram.org/method/messages.getEmojiKeywordsDifference for reference.
type MessagesGetEmojiKeywordsDifferenceRequest struct {
	// Language code
	LangCode string
	// Previous emoji keyword localization version
	FromVersion int
}

// MessagesGetEmojiKeywordsDifferenceRequestTypeID is TL type id of MessagesGetEmojiKeywordsDifferenceRequest.
const MessagesGetEmojiKeywordsDifferenceRequestTypeID = 0x1508b6af

// Ensuring interfaces in compile-time for MessagesGetEmojiKeywordsDifferenceRequest.
var (
	_ bin.Encoder     = &MessagesGetEmojiKeywordsDifferenceRequest{}
	_ bin.Decoder     = &MessagesGetEmojiKeywordsDifferenceRequest{}
	_ bin.BareEncoder = &MessagesGetEmojiKeywordsDifferenceRequest{}
	_ bin.BareDecoder = &MessagesGetEmojiKeywordsDifferenceRequest{}
)

func (g *MessagesGetEmojiKeywordsDifferenceRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangCode == "") {
		return false
	}
	if !(g.FromVersion == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) String() string {
	if g == nil {
		return "MessagesGetEmojiKeywordsDifferenceRequest(nil)"
	}
	type Alias MessagesGetEmojiKeywordsDifferenceRequest
	return fmt.Sprintf("MessagesGetEmojiKeywordsDifferenceRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetEmojiKeywordsDifferenceRequest from given interface.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) FillFrom(from interface {
	GetLangCode() (value string)
	GetFromVersion() (value int)
}) {
	g.LangCode = from.GetLangCode()
	g.FromVersion = from.GetFromVersion()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetEmojiKeywordsDifferenceRequest) TypeID() uint32 {
	return MessagesGetEmojiKeywordsDifferenceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetEmojiKeywordsDifferenceRequest) TypeName() string {
	return "messages.getEmojiKeywordsDifference"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getEmojiKeywordsDifference",
		ID:   MessagesGetEmojiKeywordsDifferenceRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "FromVersion",
			SchemaName: "from_version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywordsDifference#1508b6af as nil")
	}
	b.PutID(MessagesGetEmojiKeywordsDifferenceRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywordsDifference#1508b6af as nil")
	}
	b.PutString(g.LangCode)
	b.PutInt(g.FromVersion)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywordsDifference#1508b6af to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiKeywordsDifferenceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiKeywordsDifference#1508b6af: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywordsDifference#1508b6af to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywordsDifference#1508b6af: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywordsDifference#1508b6af: field from_version: %w", err)
		}
		g.FromVersion = value
	}
	return nil
}

// GetLangCode returns value of LangCode field.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) GetLangCode() (value string) {
	return g.LangCode
}

// GetFromVersion returns value of FromVersion field.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) GetFromVersion() (value int) {
	return g.FromVersion
}

// MessagesGetEmojiKeywordsDifference invokes method messages.getEmojiKeywordsDifference#1508b6af returning error if any.
// Get changed emoji keywords
//
// See https://core.telegram.org/method/messages.getEmojiKeywordsDifference for reference.
func (c *Client) MessagesGetEmojiKeywordsDifference(ctx context.Context, request *MessagesGetEmojiKeywordsDifferenceRequest) (*EmojiKeywordsDifference, error) {
	var result EmojiKeywordsDifference

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
