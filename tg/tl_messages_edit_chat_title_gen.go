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

// MessagesEditChatTitleRequest represents TL type `messages.editChatTitle#73783ffd`.
// Chanages chat name and sends a service message on it.
//
// See https://core.telegram.org/method/messages.editChatTitle for reference.
type MessagesEditChatTitleRequest struct {
	// Chat ID
	ChatID int64
	// New chat name, different from the old one
	Title string
}

// MessagesEditChatTitleRequestTypeID is TL type id of MessagesEditChatTitleRequest.
const MessagesEditChatTitleRequestTypeID = 0x73783ffd

// Ensuring interfaces in compile-time for MessagesEditChatTitleRequest.
var (
	_ bin.Encoder     = &MessagesEditChatTitleRequest{}
	_ bin.Decoder     = &MessagesEditChatTitleRequest{}
	_ bin.BareEncoder = &MessagesEditChatTitleRequest{}
	_ bin.BareDecoder = &MessagesEditChatTitleRequest{}
)

func (e *MessagesEditChatTitleRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEditChatTitleRequest) String() string {
	if e == nil {
		return "MessagesEditChatTitleRequest(nil)"
	}
	type Alias MessagesEditChatTitleRequest
	return fmt.Sprintf("MessagesEditChatTitleRequest%+v", Alias(*e))
}

// FillFrom fills MessagesEditChatTitleRequest from given interface.
func (e *MessagesEditChatTitleRequest) FillFrom(from interface {
	GetChatID() (value int64)
	GetTitle() (value string)
}) {
	e.ChatID = from.GetChatID()
	e.Title = from.GetTitle()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesEditChatTitleRequest) TypeID() uint32 {
	return MessagesEditChatTitleRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesEditChatTitleRequest) TypeName() string {
	return "messages.editChatTitle"
}

// TypeInfo returns info about TL type.
func (e *MessagesEditChatTitleRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.editChatTitle",
		ID:   MessagesEditChatTitleRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *MessagesEditChatTitleRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatTitle#73783ffd as nil")
	}
	b.PutID(MessagesEditChatTitleRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesEditChatTitleRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatTitle#73783ffd as nil")
	}
	b.PutLong(e.ChatID)
	b.PutString(e.Title)
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEditChatTitleRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatTitle#73783ffd to nil")
	}
	if err := b.ConsumeID(MessagesEditChatTitleRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editChatTitle#73783ffd: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesEditChatTitleRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatTitle#73783ffd to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatTitle#73783ffd: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatTitle#73783ffd: field title: %w", err)
		}
		e.Title = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (e *MessagesEditChatTitleRequest) GetChatID() (value int64) {
	return e.ChatID
}

// GetTitle returns value of Title field.
func (e *MessagesEditChatTitleRequest) GetTitle() (value string) {
	return e.Title
}

// MessagesEditChatTitle invokes method messages.editChatTitle#73783ffd returning error if any.
// Chanages chat name and sends a service message on it.
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  400 CHAT_TITLE_EMPTY: No chat title provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.editChatTitle for reference.
// Can be used by bots.
func (c *Client) MessagesEditChatTitle(ctx context.Context, request *MessagesEditChatTitleRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
