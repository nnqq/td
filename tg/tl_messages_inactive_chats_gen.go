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

// MessagesInactiveChats represents TL type `messages.inactiveChats#a927fec5`.
// Inactive chat list
//
// See https://core.telegram.org/constructor/messages.inactiveChats for reference.
type MessagesInactiveChats struct {
	// When was the chat last active
	Dates []int
	// Chat list
	Chats []ChatClass
	// Users mentioned in the chat list
	Users []UserClass
}

// MessagesInactiveChatsTypeID is TL type id of MessagesInactiveChats.
const MessagesInactiveChatsTypeID = 0xa927fec5

// Ensuring interfaces in compile-time for MessagesInactiveChats.
var (
	_ bin.Encoder     = &MessagesInactiveChats{}
	_ bin.Decoder     = &MessagesInactiveChats{}
	_ bin.BareEncoder = &MessagesInactiveChats{}
	_ bin.BareDecoder = &MessagesInactiveChats{}
)

func (i *MessagesInactiveChats) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Dates == nil) {
		return false
	}
	if !(i.Chats == nil) {
		return false
	}
	if !(i.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *MessagesInactiveChats) String() string {
	if i == nil {
		return "MessagesInactiveChats(nil)"
	}
	type Alias MessagesInactiveChats
	return fmt.Sprintf("MessagesInactiveChats%+v", Alias(*i))
}

// FillFrom fills MessagesInactiveChats from given interface.
func (i *MessagesInactiveChats) FillFrom(from interface {
	GetDates() (value []int)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	i.Dates = from.GetDates()
	i.Chats = from.GetChats()
	i.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesInactiveChats) TypeID() uint32 {
	return MessagesInactiveChatsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesInactiveChats) TypeName() string {
	return "messages.inactiveChats"
}

// TypeInfo returns info about TL type.
func (i *MessagesInactiveChats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.inactiveChats",
		ID:   MessagesInactiveChatsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dates",
			SchemaName: "dates",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *MessagesInactiveChats) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.inactiveChats#a927fec5 as nil")
	}
	b.PutID(MessagesInactiveChatsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *MessagesInactiveChats) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.inactiveChats#a927fec5 as nil")
	}
	b.PutVectorHeader(len(i.Dates))
	for _, v := range i.Dates {
		b.PutInt(v)
	}
	b.PutVectorHeader(len(i.Chats))
	for idx, v := range i.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.inactiveChats#a927fec5: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.inactiveChats#a927fec5: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.inactiveChats#a927fec5: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.inactiveChats#a927fec5: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *MessagesInactiveChats) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.inactiveChats#a927fec5 to nil")
	}
	if err := b.ConsumeID(MessagesInactiveChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *MessagesInactiveChats) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.inactiveChats#a927fec5 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field dates: %w", err)
		}

		if headerLen > 0 {
			i.Dates = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field dates: %w", err)
			}
			i.Dates = append(i.Dates, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field chats: %w", err)
		}

		if headerLen > 0 {
			i.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field chats: %w", err)
			}
			i.Chats = append(i.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field users: %w", err)
		}

		if headerLen > 0 {
			i.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.inactiveChats#a927fec5: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// GetDates returns value of Dates field.
func (i *MessagesInactiveChats) GetDates() (value []int) {
	return i.Dates
}

// GetChats returns value of Chats field.
func (i *MessagesInactiveChats) GetChats() (value []ChatClass) {
	return i.Chats
}

// GetUsers returns value of Users field.
func (i *MessagesInactiveChats) GetUsers() (value []UserClass) {
	return i.Users
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (i *MessagesInactiveChats) MapChats() (value ChatClassArray) {
	return ChatClassArray(i.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (i *MessagesInactiveChats) MapUsers() (value UserClassArray) {
	return UserClassArray(i.Users)
}
