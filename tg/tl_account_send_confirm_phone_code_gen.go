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

// AccountSendConfirmPhoneCodeRequest represents TL type `account.sendConfirmPhoneCode#1b3faa88`.
// Send confirmation code to cancel account deletion, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/account-deletion
//
// See https://core.telegram.org/method/account.sendConfirmPhoneCode for reference.
type AccountSendConfirmPhoneCodeRequest struct {
	// The hash from the service notification, for more info click here »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/account-deletion
	Hash string
	// Phone code settings
	Settings CodeSettings
}

// AccountSendConfirmPhoneCodeRequestTypeID is TL type id of AccountSendConfirmPhoneCodeRequest.
const AccountSendConfirmPhoneCodeRequestTypeID = 0x1b3faa88

// Ensuring interfaces in compile-time for AccountSendConfirmPhoneCodeRequest.
var (
	_ bin.Encoder     = &AccountSendConfirmPhoneCodeRequest{}
	_ bin.Decoder     = &AccountSendConfirmPhoneCodeRequest{}
	_ bin.BareEncoder = &AccountSendConfirmPhoneCodeRequest{}
	_ bin.BareDecoder = &AccountSendConfirmPhoneCodeRequest{}
)

func (s *AccountSendConfirmPhoneCodeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Hash == "") {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSendConfirmPhoneCodeRequest) String() string {
	if s == nil {
		return "AccountSendConfirmPhoneCodeRequest(nil)"
	}
	type Alias AccountSendConfirmPhoneCodeRequest
	return fmt.Sprintf("AccountSendConfirmPhoneCodeRequest%+v", Alias(*s))
}

// FillFrom fills AccountSendConfirmPhoneCodeRequest from given interface.
func (s *AccountSendConfirmPhoneCodeRequest) FillFrom(from interface {
	GetHash() (value string)
	GetSettings() (value CodeSettings)
}) {
	s.Hash = from.GetHash()
	s.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSendConfirmPhoneCodeRequest) TypeID() uint32 {
	return AccountSendConfirmPhoneCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSendConfirmPhoneCodeRequest) TypeName() string {
	return "account.sendConfirmPhoneCode"
}

// TypeInfo returns info about TL type.
func (s *AccountSendConfirmPhoneCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.sendConfirmPhoneCode",
		ID:   AccountSendConfirmPhoneCodeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSendConfirmPhoneCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendConfirmPhoneCode#1b3faa88 as nil")
	}
	b.PutID(AccountSendConfirmPhoneCodeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSendConfirmPhoneCodeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendConfirmPhoneCode#1b3faa88 as nil")
	}
	b.PutString(s.Hash)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.sendConfirmPhoneCode#1b3faa88: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSendConfirmPhoneCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendConfirmPhoneCode#1b3faa88 to nil")
	}
	if err := b.ConsumeID(AccountSendConfirmPhoneCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sendConfirmPhoneCode#1b3faa88: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSendConfirmPhoneCodeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendConfirmPhoneCode#1b3faa88 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sendConfirmPhoneCode#1b3faa88: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.sendConfirmPhoneCode#1b3faa88: field settings: %w", err)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (s *AccountSendConfirmPhoneCodeRequest) GetHash() (value string) {
	return s.Hash
}

// GetSettings returns value of Settings field.
func (s *AccountSendConfirmPhoneCodeRequest) GetSettings() (value CodeSettings) {
	return s.Settings
}

// AccountSendConfirmPhoneCode invokes method account.sendConfirmPhoneCode#1b3faa88 returning error if any.
// Send confirmation code to cancel account deletion, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/account-deletion
//
// Possible errors:
//  400 HASH_INVALID: The provided hash is invalid
//
// See https://core.telegram.org/method/account.sendConfirmPhoneCode for reference.
func (c *Client) AccountSendConfirmPhoneCode(ctx context.Context, request *AccountSendConfirmPhoneCodeRequest) (*AuthSentCode, error) {
	var result AuthSentCode

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
