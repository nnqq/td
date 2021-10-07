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

// AccountSendVerifyEmailCodeRequest represents TL type `account.sendVerifyEmailCode#7011509f`.
// Send the verification email code for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.sendVerifyEmailCode for reference.
type AccountSendVerifyEmailCodeRequest struct {
	// The email where to send the code
	Email string
}

// AccountSendVerifyEmailCodeRequestTypeID is TL type id of AccountSendVerifyEmailCodeRequest.
const AccountSendVerifyEmailCodeRequestTypeID = 0x7011509f

// Ensuring interfaces in compile-time for AccountSendVerifyEmailCodeRequest.
var (
	_ bin.Encoder     = &AccountSendVerifyEmailCodeRequest{}
	_ bin.Decoder     = &AccountSendVerifyEmailCodeRequest{}
	_ bin.BareEncoder = &AccountSendVerifyEmailCodeRequest{}
	_ bin.BareDecoder = &AccountSendVerifyEmailCodeRequest{}
)

func (s *AccountSendVerifyEmailCodeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Email == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSendVerifyEmailCodeRequest) String() string {
	if s == nil {
		return "AccountSendVerifyEmailCodeRequest(nil)"
	}
	type Alias AccountSendVerifyEmailCodeRequest
	return fmt.Sprintf("AccountSendVerifyEmailCodeRequest%+v", Alias(*s))
}

// FillFrom fills AccountSendVerifyEmailCodeRequest from given interface.
func (s *AccountSendVerifyEmailCodeRequest) FillFrom(from interface {
	GetEmail() (value string)
}) {
	s.Email = from.GetEmail()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSendVerifyEmailCodeRequest) TypeID() uint32 {
	return AccountSendVerifyEmailCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSendVerifyEmailCodeRequest) TypeName() string {
	return "account.sendVerifyEmailCode"
}

// TypeInfo returns info about TL type.
func (s *AccountSendVerifyEmailCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.sendVerifyEmailCode",
		ID:   AccountSendVerifyEmailCodeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Email",
			SchemaName: "email",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSendVerifyEmailCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendVerifyEmailCode#7011509f as nil")
	}
	b.PutID(AccountSendVerifyEmailCodeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSendVerifyEmailCodeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendVerifyEmailCode#7011509f as nil")
	}
	b.PutString(s.Email)
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSendVerifyEmailCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendVerifyEmailCode#7011509f to nil")
	}
	if err := b.ConsumeID(AccountSendVerifyEmailCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sendVerifyEmailCode#7011509f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSendVerifyEmailCodeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendVerifyEmailCode#7011509f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sendVerifyEmailCode#7011509f: field email: %w", err)
		}
		s.Email = value
	}
	return nil
}

// GetEmail returns value of Email field.
func (s *AccountSendVerifyEmailCodeRequest) GetEmail() (value string) {
	return s.Email
}

// AccountSendVerifyEmailCode invokes method account.sendVerifyEmailCode#7011509f returning error if any.
// Send the verification email code for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.sendVerifyEmailCode for reference.
func (c *Client) AccountSendVerifyEmailCode(ctx context.Context, email string) (*AccountSentEmailCode, error) {
	var result AccountSentEmailCode

	request := &AccountSendVerifyEmailCodeRequest{
		Email: email,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
