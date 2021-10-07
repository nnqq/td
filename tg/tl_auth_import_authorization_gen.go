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

// AuthImportAuthorizationRequest represents TL type `auth.importAuthorization#a57a7dad`.
// Logs in a user using a key transmitted from his native data-centre.
//
// See https://core.telegram.org/method/auth.importAuthorization for reference.
type AuthImportAuthorizationRequest struct {
	// User ID
	ID int64
	// Authorization key
	Bytes []byte
}

// AuthImportAuthorizationRequestTypeID is TL type id of AuthImportAuthorizationRequest.
const AuthImportAuthorizationRequestTypeID = 0xa57a7dad

// Ensuring interfaces in compile-time for AuthImportAuthorizationRequest.
var (
	_ bin.Encoder     = &AuthImportAuthorizationRequest{}
	_ bin.Decoder     = &AuthImportAuthorizationRequest{}
	_ bin.BareEncoder = &AuthImportAuthorizationRequest{}
	_ bin.BareDecoder = &AuthImportAuthorizationRequest{}
)

func (i *AuthImportAuthorizationRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AuthImportAuthorizationRequest) String() string {
	if i == nil {
		return "AuthImportAuthorizationRequest(nil)"
	}
	type Alias AuthImportAuthorizationRequest
	return fmt.Sprintf("AuthImportAuthorizationRequest%+v", Alias(*i))
}

// FillFrom fills AuthImportAuthorizationRequest from given interface.
func (i *AuthImportAuthorizationRequest) FillFrom(from interface {
	GetID() (value int64)
	GetBytes() (value []byte)
}) {
	i.ID = from.GetID()
	i.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthImportAuthorizationRequest) TypeID() uint32 {
	return AuthImportAuthorizationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthImportAuthorizationRequest) TypeName() string {
	return "auth.importAuthorization"
}

// TypeInfo returns info about TL type.
func (i *AuthImportAuthorizationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.importAuthorization",
		ID:   AuthImportAuthorizationRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *AuthImportAuthorizationRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importAuthorization#a57a7dad as nil")
	}
	b.PutID(AuthImportAuthorizationRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *AuthImportAuthorizationRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importAuthorization#a57a7dad as nil")
	}
	b.PutLong(i.ID)
	b.PutBytes(i.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (i *AuthImportAuthorizationRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importAuthorization#a57a7dad to nil")
	}
	if err := b.ConsumeID(AuthImportAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.importAuthorization#a57a7dad: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *AuthImportAuthorizationRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importAuthorization#a57a7dad to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importAuthorization#a57a7dad: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importAuthorization#a57a7dad: field bytes: %w", err)
		}
		i.Bytes = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *AuthImportAuthorizationRequest) GetID() (value int64) {
	return i.ID
}

// GetBytes returns value of Bytes field.
func (i *AuthImportAuthorizationRequest) GetBytes() (value []byte) {
	return i.Bytes
}

// AuthImportAuthorization invokes method auth.importAuthorization#a57a7dad returning error if any.
// Logs in a user using a key transmitted from his native data-centre.
//
// Possible errors:
//  400 AUTH_BYTES_INVALID: The provided authorization is invalid
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/auth.importAuthorization for reference.
// Can be used by bots.
func (c *Client) AuthImportAuthorization(ctx context.Context, request *AuthImportAuthorizationRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
