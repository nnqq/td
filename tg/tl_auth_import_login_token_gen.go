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

// AuthImportLoginTokenRequest represents TL type `auth.importLoginToken#95ac5ce4`.
// Login using a redirected login token, generated in case of DC mismatch during QR code
// login¹.
// For more info, see login via QR code¹.
//
// Links:
//  1) https://core.telegram.org/api/qr-login
//  2) https://core.telegram.org/api/qr-login
//
// See https://core.telegram.org/method/auth.importLoginToken for reference.
type AuthImportLoginTokenRequest struct {
	// Login token
	Token []byte
}

// AuthImportLoginTokenRequestTypeID is TL type id of AuthImportLoginTokenRequest.
const AuthImportLoginTokenRequestTypeID = 0x95ac5ce4

// Ensuring interfaces in compile-time for AuthImportLoginTokenRequest.
var (
	_ bin.Encoder     = &AuthImportLoginTokenRequest{}
	_ bin.Decoder     = &AuthImportLoginTokenRequest{}
	_ bin.BareEncoder = &AuthImportLoginTokenRequest{}
	_ bin.BareDecoder = &AuthImportLoginTokenRequest{}
)

func (i *AuthImportLoginTokenRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Token == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AuthImportLoginTokenRequest) String() string {
	if i == nil {
		return "AuthImportLoginTokenRequest(nil)"
	}
	type Alias AuthImportLoginTokenRequest
	return fmt.Sprintf("AuthImportLoginTokenRequest%+v", Alias(*i))
}

// FillFrom fills AuthImportLoginTokenRequest from given interface.
func (i *AuthImportLoginTokenRequest) FillFrom(from interface {
	GetToken() (value []byte)
}) {
	i.Token = from.GetToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthImportLoginTokenRequest) TypeID() uint32 {
	return AuthImportLoginTokenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthImportLoginTokenRequest) TypeName() string {
	return "auth.importLoginToken"
}

// TypeInfo returns info about TL type.
func (i *AuthImportLoginTokenRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.importLoginToken",
		ID:   AuthImportLoginTokenRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Token",
			SchemaName: "token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *AuthImportLoginTokenRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importLoginToken#95ac5ce4 as nil")
	}
	b.PutID(AuthImportLoginTokenRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *AuthImportLoginTokenRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importLoginToken#95ac5ce4 as nil")
	}
	b.PutBytes(i.Token)
	return nil
}

// Decode implements bin.Decoder.
func (i *AuthImportLoginTokenRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importLoginToken#95ac5ce4 to nil")
	}
	if err := b.ConsumeID(AuthImportLoginTokenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.importLoginToken#95ac5ce4: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *AuthImportLoginTokenRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importLoginToken#95ac5ce4 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importLoginToken#95ac5ce4: field token: %w", err)
		}
		i.Token = value
	}
	return nil
}

// GetToken returns value of Token field.
func (i *AuthImportLoginTokenRequest) GetToken() (value []byte) {
	return i.Token
}

// AuthImportLoginToken invokes method auth.importLoginToken#95ac5ce4 returning error if any.
// Login using a redirected login token, generated in case of DC mismatch during QR code
// login¹.
// For more info, see login via QR code¹.
//
// Links:
//  1) https://core.telegram.org/api/qr-login
//  2) https://core.telegram.org/api/qr-login
//
// Possible errors:
//  400 AUTH_TOKEN_EXPIRED: The authorization token has expired
//
// See https://core.telegram.org/method/auth.importLoginToken for reference.
func (c *Client) AuthImportLoginToken(ctx context.Context, token []byte) (AuthLoginTokenClass, error) {
	var result AuthLoginTokenBox

	request := &AuthImportLoginTokenRequest{
		Token: token,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.LoginToken, nil
}
