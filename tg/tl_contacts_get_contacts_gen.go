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

// ContactsGetContactsRequest represents TL type `contacts.getContacts#5dd69e12`.
// Returns the current user's contact list.
//
// See https://core.telegram.org/method/contacts.getContacts for reference.
type ContactsGetContactsRequest struct {
	// If there already is a full contact list on the client, a hash¹ of a the list of
	// contact IDs in ascending order may be passed in this parameter. If the contact set was
	// not changed, (contacts.contactsNotModified)² will be returned.
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	//  2) https://core.telegram.org/constructor/contacts.contactsNotModified
	Hash int64
}

// ContactsGetContactsRequestTypeID is TL type id of ContactsGetContactsRequest.
const ContactsGetContactsRequestTypeID = 0x5dd69e12

// Ensuring interfaces in compile-time for ContactsGetContactsRequest.
var (
	_ bin.Encoder     = &ContactsGetContactsRequest{}
	_ bin.Decoder     = &ContactsGetContactsRequest{}
	_ bin.BareEncoder = &ContactsGetContactsRequest{}
	_ bin.BareDecoder = &ContactsGetContactsRequest{}
)

func (g *ContactsGetContactsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetContactsRequest) String() string {
	if g == nil {
		return "ContactsGetContactsRequest(nil)"
	}
	type Alias ContactsGetContactsRequest
	return fmt.Sprintf("ContactsGetContactsRequest%+v", Alias(*g))
}

// FillFrom fills ContactsGetContactsRequest from given interface.
func (g *ContactsGetContactsRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsGetContactsRequest) TypeID() uint32 {
	return ContactsGetContactsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsGetContactsRequest) TypeName() string {
	return "contacts.getContacts"
}

// TypeInfo returns info about TL type.
func (g *ContactsGetContactsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.getContacts",
		ID:   ContactsGetContactsRequestTypeID,
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
func (g *ContactsGetContactsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getContacts#5dd69e12 as nil")
	}
	b.PutID(ContactsGetContactsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ContactsGetContactsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getContacts#5dd69e12 as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetContactsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getContacts#5dd69e12 to nil")
	}
	if err := b.ConsumeID(ContactsGetContactsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getContacts#5dd69e12: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ContactsGetContactsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getContacts#5dd69e12 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getContacts#5dd69e12: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *ContactsGetContactsRequest) GetHash() (value int64) {
	return g.Hash
}

// ContactsGetContacts invokes method contacts.getContacts#5dd69e12 returning error if any.
// Returns the current user's contact list.
//
// See https://core.telegram.org/method/contacts.getContacts for reference.
func (c *Client) ContactsGetContacts(ctx context.Context, hash int64) (ContactsContactsClass, error) {
	var result ContactsContactsBox

	request := &ContactsGetContactsRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Contacts, nil
}
