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

// AccountSetAccountTTLRequest represents TL type `account.setAccountTTL#2442485e`.
// Set account self-destruction period
//
// See https://core.telegram.org/method/account.setAccountTTL for reference.
type AccountSetAccountTTLRequest struct {
	// Time to live in days
	TTL AccountDaysTTL
}

// AccountSetAccountTTLRequestTypeID is TL type id of AccountSetAccountTTLRequest.
const AccountSetAccountTTLRequestTypeID = 0x2442485e

// Ensuring interfaces in compile-time for AccountSetAccountTTLRequest.
var (
	_ bin.Encoder     = &AccountSetAccountTTLRequest{}
	_ bin.Decoder     = &AccountSetAccountTTLRequest{}
	_ bin.BareEncoder = &AccountSetAccountTTLRequest{}
	_ bin.BareDecoder = &AccountSetAccountTTLRequest{}
)

func (s *AccountSetAccountTTLRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.TTL.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSetAccountTTLRequest) String() string {
	if s == nil {
		return "AccountSetAccountTTLRequest(nil)"
	}
	type Alias AccountSetAccountTTLRequest
	return fmt.Sprintf("AccountSetAccountTTLRequest%+v", Alias(*s))
}

// FillFrom fills AccountSetAccountTTLRequest from given interface.
func (s *AccountSetAccountTTLRequest) FillFrom(from interface {
	GetTTL() (value AccountDaysTTL)
}) {
	s.TTL = from.GetTTL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSetAccountTTLRequest) TypeID() uint32 {
	return AccountSetAccountTTLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSetAccountTTLRequest) TypeName() string {
	return "account.setAccountTTL"
}

// TypeInfo returns info about TL type.
func (s *AccountSetAccountTTLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.setAccountTTL",
		ID:   AccountSetAccountTTLRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TTL",
			SchemaName: "ttl",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSetAccountTTLRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setAccountTTL#2442485e as nil")
	}
	b.PutID(AccountSetAccountTTLRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSetAccountTTLRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setAccountTTL#2442485e as nil")
	}
	if err := s.TTL.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.setAccountTTL#2442485e: field ttl: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSetAccountTTLRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setAccountTTL#2442485e to nil")
	}
	if err := b.ConsumeID(AccountSetAccountTTLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setAccountTTL#2442485e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSetAccountTTLRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setAccountTTL#2442485e to nil")
	}
	{
		if err := s.TTL.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.setAccountTTL#2442485e: field ttl: %w", err)
		}
	}
	return nil
}

// GetTTL returns value of TTL field.
func (s *AccountSetAccountTTLRequest) GetTTL() (value AccountDaysTTL) {
	return s.TTL
}

// AccountSetAccountTTL invokes method account.setAccountTTL#2442485e returning error if any.
// Set account self-destruction period
//
// Possible errors:
//  400 TTL_DAYS_INVALID: The provided TTL is invalid
//
// See https://core.telegram.org/method/account.setAccountTTL for reference.
func (c *Client) AccountSetAccountTTL(ctx context.Context, ttl AccountDaysTTL) (bool, error) {
	var result BoolBox

	request := &AccountSetAccountTTLRequest{
		TTL: ttl,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
