// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// NewSessionCreated represents TL type `new_session_created#9ec20908`.
type NewSessionCreated struct {
	// FirstMsgID field of NewSessionCreated.
	FirstMsgID int64
	// UniqueID field of NewSessionCreated.
	UniqueID int64
	// ServerSalt field of NewSessionCreated.
	ServerSalt int64
}

// NewSessionCreatedTypeID is TL type id of NewSessionCreated.
const NewSessionCreatedTypeID = 0x9ec20908

// Ensuring interfaces in compile-time for NewSessionCreated.
var (
	_ bin.Encoder     = &NewSessionCreated{}
	_ bin.Decoder     = &NewSessionCreated{}
	_ bin.BareEncoder = &NewSessionCreated{}
	_ bin.BareDecoder = &NewSessionCreated{}
)

func (n *NewSessionCreated) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.FirstMsgID == 0) {
		return false
	}
	if !(n.UniqueID == 0) {
		return false
	}
	if !(n.ServerSalt == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NewSessionCreated) String() string {
	if n == nil {
		return "NewSessionCreated(nil)"
	}
	type Alias NewSessionCreated
	return fmt.Sprintf("NewSessionCreated%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NewSessionCreated) TypeID() uint32 {
	return NewSessionCreatedTypeID
}

// TypeName returns name of type in TL schema.
func (*NewSessionCreated) TypeName() string {
	return "new_session_created"
}

// TypeInfo returns info about TL type.
func (n *NewSessionCreated) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "new_session_created",
		ID:   NewSessionCreatedTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FirstMsgID",
			SchemaName: "first_msg_id",
		},
		{
			Name:       "UniqueID",
			SchemaName: "unique_id",
		},
		{
			Name:       "ServerSalt",
			SchemaName: "server_salt",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NewSessionCreated) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode new_session_created#9ec20908 as nil")
	}
	b.PutID(NewSessionCreatedTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NewSessionCreated) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode new_session_created#9ec20908 as nil")
	}
	b.PutLong(n.FirstMsgID)
	b.PutLong(n.UniqueID)
	b.PutLong(n.ServerSalt)
	return nil
}

// Decode implements bin.Decoder.
func (n *NewSessionCreated) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode new_session_created#9ec20908 to nil")
	}
	if err := b.ConsumeID(NewSessionCreatedTypeID); err != nil {
		return fmt.Errorf("unable to decode new_session_created#9ec20908: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NewSessionCreated) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode new_session_created#9ec20908 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode new_session_created#9ec20908: field first_msg_id: %w", err)
		}
		n.FirstMsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode new_session_created#9ec20908: field unique_id: %w", err)
		}
		n.UniqueID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode new_session_created#9ec20908: field server_salt: %w", err)
		}
		n.ServerSalt = value
	}
	return nil
}

// GetFirstMsgID returns value of FirstMsgID field.
func (n *NewSessionCreated) GetFirstMsgID() (value int64) {
	return n.FirstMsgID
}

// GetUniqueID returns value of UniqueID field.
func (n *NewSessionCreated) GetUniqueID() (value int64) {
	return n.UniqueID
}

// GetServerSalt returns value of ServerSalt field.
func (n *NewSessionCreated) GetServerSalt() (value int64) {
	return n.ServerSalt
}
