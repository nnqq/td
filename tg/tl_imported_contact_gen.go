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

// ImportedContact represents TL type `importedContact#c13e3c50`.
// Successfully imported contact.
//
// See https://core.telegram.org/constructor/importedContact for reference.
type ImportedContact struct {
	// User identifier
	UserID int64
	// The contact's client identifier (passed to one of the InputContact¹ constructors)
	//
	// Links:
	//  1) https://core.telegram.org/type/InputContact
	ClientID int64
}

// ImportedContactTypeID is TL type id of ImportedContact.
const ImportedContactTypeID = 0xc13e3c50

// Ensuring interfaces in compile-time for ImportedContact.
var (
	_ bin.Encoder     = &ImportedContact{}
	_ bin.Decoder     = &ImportedContact{}
	_ bin.BareEncoder = &ImportedContact{}
	_ bin.BareDecoder = &ImportedContact{}
)

func (i *ImportedContact) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.UserID == 0) {
		return false
	}
	if !(i.ClientID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *ImportedContact) String() string {
	if i == nil {
		return "ImportedContact(nil)"
	}
	type Alias ImportedContact
	return fmt.Sprintf("ImportedContact%+v", Alias(*i))
}

// FillFrom fills ImportedContact from given interface.
func (i *ImportedContact) FillFrom(from interface {
	GetUserID() (value int64)
	GetClientID() (value int64)
}) {
	i.UserID = from.GetUserID()
	i.ClientID = from.GetClientID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ImportedContact) TypeID() uint32 {
	return ImportedContactTypeID
}

// TypeName returns name of type in TL schema.
func (*ImportedContact) TypeName() string {
	return "importedContact"
}

// TypeInfo returns info about TL type.
func (i *ImportedContact) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "importedContact",
		ID:   ImportedContactTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "ClientID",
			SchemaName: "client_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *ImportedContact) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode importedContact#c13e3c50 as nil")
	}
	b.PutID(ImportedContactTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *ImportedContact) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode importedContact#c13e3c50 as nil")
	}
	b.PutLong(i.UserID)
	b.PutLong(i.ClientID)
	return nil
}

// Decode implements bin.Decoder.
func (i *ImportedContact) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode importedContact#c13e3c50 to nil")
	}
	if err := b.ConsumeID(ImportedContactTypeID); err != nil {
		return fmt.Errorf("unable to decode importedContact#c13e3c50: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *ImportedContact) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode importedContact#c13e3c50 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode importedContact#c13e3c50: field user_id: %w", err)
		}
		i.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode importedContact#c13e3c50: field client_id: %w", err)
		}
		i.ClientID = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (i *ImportedContact) GetUserID() (value int64) {
	return i.UserID
}

// GetClientID returns value of ClientID field.
func (i *ImportedContact) GetClientID() (value int64) {
	return i.ClientID
}
