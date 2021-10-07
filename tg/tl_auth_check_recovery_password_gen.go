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

// AuthCheckRecoveryPasswordRequest represents TL type `auth.checkRecoveryPassword#d36bf79`.
//
// See https://core.telegram.org/method/auth.checkRecoveryPassword for reference.
type AuthCheckRecoveryPasswordRequest struct {
	// Code field of AuthCheckRecoveryPasswordRequest.
	Code string
}

// AuthCheckRecoveryPasswordRequestTypeID is TL type id of AuthCheckRecoveryPasswordRequest.
const AuthCheckRecoveryPasswordRequestTypeID = 0xd36bf79

// Ensuring interfaces in compile-time for AuthCheckRecoveryPasswordRequest.
var (
	_ bin.Encoder     = &AuthCheckRecoveryPasswordRequest{}
	_ bin.Decoder     = &AuthCheckRecoveryPasswordRequest{}
	_ bin.BareEncoder = &AuthCheckRecoveryPasswordRequest{}
	_ bin.BareDecoder = &AuthCheckRecoveryPasswordRequest{}
)

func (c *AuthCheckRecoveryPasswordRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Code == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AuthCheckRecoveryPasswordRequest) String() string {
	if c == nil {
		return "AuthCheckRecoveryPasswordRequest(nil)"
	}
	type Alias AuthCheckRecoveryPasswordRequest
	return fmt.Sprintf("AuthCheckRecoveryPasswordRequest%+v", Alias(*c))
}

// FillFrom fills AuthCheckRecoveryPasswordRequest from given interface.
func (c *AuthCheckRecoveryPasswordRequest) FillFrom(from interface {
	GetCode() (value string)
}) {
	c.Code = from.GetCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthCheckRecoveryPasswordRequest) TypeID() uint32 {
	return AuthCheckRecoveryPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthCheckRecoveryPasswordRequest) TypeName() string {
	return "auth.checkRecoveryPassword"
}

// TypeInfo returns info about TL type.
func (c *AuthCheckRecoveryPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.checkRecoveryPassword",
		ID:   AuthCheckRecoveryPasswordRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *AuthCheckRecoveryPasswordRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.checkRecoveryPassword#d36bf79 as nil")
	}
	b.PutID(AuthCheckRecoveryPasswordRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AuthCheckRecoveryPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.checkRecoveryPassword#d36bf79 as nil")
	}
	b.PutString(c.Code)
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCheckRecoveryPasswordRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.checkRecoveryPassword#d36bf79 to nil")
	}
	if err := b.ConsumeID(AuthCheckRecoveryPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.checkRecoveryPassword#d36bf79: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AuthCheckRecoveryPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.checkRecoveryPassword#d36bf79 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.checkRecoveryPassword#d36bf79: field code: %w", err)
		}
		c.Code = value
	}
	return nil
}

// GetCode returns value of Code field.
func (c *AuthCheckRecoveryPasswordRequest) GetCode() (value string) {
	return c.Code
}

// AuthCheckRecoveryPassword invokes method auth.checkRecoveryPassword#d36bf79 returning error if any.
//
// See https://core.telegram.org/method/auth.checkRecoveryPassword for reference.
func (c *Client) AuthCheckRecoveryPassword(ctx context.Context, code string) (bool, error) {
	var result BoolBox

	request := &AuthCheckRecoveryPasswordRequest{
		Code: code,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
