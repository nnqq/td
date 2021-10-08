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

// InputPaymentCredentialsSaved represents TL type `inputPaymentCredentialsSaved#c10eb2cf`.
// Saved payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentialsSaved for reference.
type InputPaymentCredentialsSaved struct {
	// Credential ID
	ID string
	// Temporary password
	TmpPassword []byte
}

// InputPaymentCredentialsSavedTypeID is TL type id of InputPaymentCredentialsSaved.
const InputPaymentCredentialsSavedTypeID = 0xc10eb2cf

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsSaved) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsSaved.
var (
	_ bin.Encoder     = &InputPaymentCredentialsSaved{}
	_ bin.Decoder     = &InputPaymentCredentialsSaved{}
	_ bin.BareEncoder = &InputPaymentCredentialsSaved{}
	_ bin.BareDecoder = &InputPaymentCredentialsSaved{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsSaved{}
)

func (i *InputPaymentCredentialsSaved) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == "") {
		return false
	}
	if !(i.TmpPassword == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentialsSaved) String() string {
	if i == nil {
		return "InputPaymentCredentialsSaved(nil)"
	}
	type Alias InputPaymentCredentialsSaved
	return fmt.Sprintf("InputPaymentCredentialsSaved%+v", Alias(*i))
}

// FillFrom fills InputPaymentCredentialsSaved from given interface.
func (i *InputPaymentCredentialsSaved) FillFrom(from interface {
	GetID() (value string)
	GetTmpPassword() (value []byte)
}) {
	i.ID = from.GetID()
	i.TmpPassword = from.GetTmpPassword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPaymentCredentialsSaved) TypeID() uint32 {
	return InputPaymentCredentialsSavedTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPaymentCredentialsSaved) TypeName() string {
	return "inputPaymentCredentialsSaved"
}

// TypeInfo returns info about TL type.
func (i *InputPaymentCredentialsSaved) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPaymentCredentialsSaved",
		ID:   InputPaymentCredentialsSavedTypeID,
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
			Name:       "TmpPassword",
			SchemaName: "tmp_password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsSaved) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsSaved#c10eb2cf as nil")
	}
	b.PutID(InputPaymentCredentialsSavedTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPaymentCredentialsSaved) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsSaved#c10eb2cf as nil")
	}
	b.PutString(i.ID)
	b.PutBytes(i.TmpPassword)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsSaved) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsSaved#c10eb2cf to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsSavedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPaymentCredentialsSaved) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsSaved#c10eb2cf to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: field tmp_password: %w", err)
		}
		i.TmpPassword = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputPaymentCredentialsSaved) GetID() (value string) {
	return i.ID
}

// GetTmpPassword returns value of TmpPassword field.
func (i *InputPaymentCredentialsSaved) GetTmpPassword() (value []byte) {
	return i.TmpPassword
}

// InputPaymentCredentials represents TL type `inputPaymentCredentials#3417d728`.
// Payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentials for reference.
type InputPaymentCredentials struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Save payment credential for future use
	Save bool
	// Payment credentials
	Data DataJSON
}

// InputPaymentCredentialsTypeID is TL type id of InputPaymentCredentials.
const InputPaymentCredentialsTypeID = 0x3417d728

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentials) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentials.
var (
	_ bin.Encoder     = &InputPaymentCredentials{}
	_ bin.Decoder     = &InputPaymentCredentials{}
	_ bin.BareEncoder = &InputPaymentCredentials{}
	_ bin.BareDecoder = &InputPaymentCredentials{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentials{}
)

func (i *InputPaymentCredentials) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Save == false) {
		return false
	}
	if !(i.Data.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentials) String() string {
	if i == nil {
		return "InputPaymentCredentials(nil)"
	}
	type Alias InputPaymentCredentials
	return fmt.Sprintf("InputPaymentCredentials%+v", Alias(*i))
}

// FillFrom fills InputPaymentCredentials from given interface.
func (i *InputPaymentCredentials) FillFrom(from interface {
	GetSave() (value bool)
	GetData() (value DataJSON)
}) {
	i.Save = from.GetSave()
	i.Data = from.GetData()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPaymentCredentials) TypeID() uint32 {
	return InputPaymentCredentialsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPaymentCredentials) TypeName() string {
	return "inputPaymentCredentials"
}

// TypeInfo returns info about TL type.
func (i *InputPaymentCredentials) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPaymentCredentials",
		ID:   InputPaymentCredentialsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Save",
			SchemaName: "save",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentials) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentials#3417d728 as nil")
	}
	b.PutID(InputPaymentCredentialsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPaymentCredentials) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentials#3417d728 as nil")
	}
	if !(i.Save == false) {
		i.Flags.Set(0)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentials#3417d728: field flags: %w", err)
	}
	if err := i.Data.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentials#3417d728: field data: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentials) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentials#3417d728 to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPaymentCredentials) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentials#3417d728 to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: field flags: %w", err)
		}
	}
	i.Save = i.Flags.Has(0)
	{
		if err := i.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: field data: %w", err)
		}
	}
	return nil
}

// SetSave sets value of Save conditional field.
func (i *InputPaymentCredentials) SetSave(value bool) {
	if value {
		i.Flags.Set(0)
		i.Save = true
	} else {
		i.Flags.Unset(0)
		i.Save = false
	}
}

// GetSave returns value of Save conditional field.
func (i *InputPaymentCredentials) GetSave() (value bool) {
	return i.Flags.Has(0)
}

// GetData returns value of Data field.
func (i *InputPaymentCredentials) GetData() (value DataJSON) {
	return i.Data
}

// InputPaymentCredentialsApplePay represents TL type `inputPaymentCredentialsApplePay#aa1c39f`.
// Apple pay payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentialsApplePay for reference.
type InputPaymentCredentialsApplePay struct {
	// Payment data
	PaymentData DataJSON
}

// InputPaymentCredentialsApplePayTypeID is TL type id of InputPaymentCredentialsApplePay.
const InputPaymentCredentialsApplePayTypeID = 0xaa1c39f

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsApplePay) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsApplePay.
var (
	_ bin.Encoder     = &InputPaymentCredentialsApplePay{}
	_ bin.Decoder     = &InputPaymentCredentialsApplePay{}
	_ bin.BareEncoder = &InputPaymentCredentialsApplePay{}
	_ bin.BareDecoder = &InputPaymentCredentialsApplePay{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsApplePay{}
)

func (i *InputPaymentCredentialsApplePay) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.PaymentData.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentialsApplePay) String() string {
	if i == nil {
		return "InputPaymentCredentialsApplePay(nil)"
	}
	type Alias InputPaymentCredentialsApplePay
	return fmt.Sprintf("InputPaymentCredentialsApplePay%+v", Alias(*i))
}

// FillFrom fills InputPaymentCredentialsApplePay from given interface.
func (i *InputPaymentCredentialsApplePay) FillFrom(from interface {
	GetPaymentData() (value DataJSON)
}) {
	i.PaymentData = from.GetPaymentData()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPaymentCredentialsApplePay) TypeID() uint32 {
	return InputPaymentCredentialsApplePayTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPaymentCredentialsApplePay) TypeName() string {
	return "inputPaymentCredentialsApplePay"
}

// TypeInfo returns info about TL type.
func (i *InputPaymentCredentialsApplePay) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPaymentCredentialsApplePay",
		ID:   InputPaymentCredentialsApplePayTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PaymentData",
			SchemaName: "payment_data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsApplePay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsApplePay#aa1c39f as nil")
	}
	b.PutID(InputPaymentCredentialsApplePayTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPaymentCredentialsApplePay) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsApplePay#aa1c39f as nil")
	}
	if err := i.PaymentData.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentialsApplePay#aa1c39f: field payment_data: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsApplePay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsApplePay#aa1c39f to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsApplePayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsApplePay#aa1c39f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPaymentCredentialsApplePay) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsApplePay#aa1c39f to nil")
	}
	{
		if err := i.PaymentData.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsApplePay#aa1c39f: field payment_data: %w", err)
		}
	}
	return nil
}

// GetPaymentData returns value of PaymentData field.
func (i *InputPaymentCredentialsApplePay) GetPaymentData() (value DataJSON) {
	return i.PaymentData
}

// InputPaymentCredentialsGooglePay represents TL type `inputPaymentCredentialsGooglePay#8ac32801`.
//
// See https://core.telegram.org/constructor/inputPaymentCredentialsGooglePay for reference.
type InputPaymentCredentialsGooglePay struct {
	// PaymentToken field of InputPaymentCredentialsGooglePay.
	PaymentToken DataJSON
}

// InputPaymentCredentialsGooglePayTypeID is TL type id of InputPaymentCredentialsGooglePay.
const InputPaymentCredentialsGooglePayTypeID = 0x8ac32801

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsGooglePay) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsGooglePay.
var (
	_ bin.Encoder     = &InputPaymentCredentialsGooglePay{}
	_ bin.Decoder     = &InputPaymentCredentialsGooglePay{}
	_ bin.BareEncoder = &InputPaymentCredentialsGooglePay{}
	_ bin.BareDecoder = &InputPaymentCredentialsGooglePay{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsGooglePay{}
)

func (i *InputPaymentCredentialsGooglePay) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.PaymentToken.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentialsGooglePay) String() string {
	if i == nil {
		return "InputPaymentCredentialsGooglePay(nil)"
	}
	type Alias InputPaymentCredentialsGooglePay
	return fmt.Sprintf("InputPaymentCredentialsGooglePay%+v", Alias(*i))
}

// FillFrom fills InputPaymentCredentialsGooglePay from given interface.
func (i *InputPaymentCredentialsGooglePay) FillFrom(from interface {
	GetPaymentToken() (value DataJSON)
}) {
	i.PaymentToken = from.GetPaymentToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPaymentCredentialsGooglePay) TypeID() uint32 {
	return InputPaymentCredentialsGooglePayTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPaymentCredentialsGooglePay) TypeName() string {
	return "inputPaymentCredentialsGooglePay"
}

// TypeInfo returns info about TL type.
func (i *InputPaymentCredentialsGooglePay) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPaymentCredentialsGooglePay",
		ID:   InputPaymentCredentialsGooglePayTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PaymentToken",
			SchemaName: "payment_token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsGooglePay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsGooglePay#8ac32801 as nil")
	}
	b.PutID(InputPaymentCredentialsGooglePayTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPaymentCredentialsGooglePay) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsGooglePay#8ac32801 as nil")
	}
	if err := i.PaymentToken.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentialsGooglePay#8ac32801: field payment_token: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsGooglePay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsGooglePay#8ac32801 to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsGooglePayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsGooglePay#8ac32801: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPaymentCredentialsGooglePay) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsGooglePay#8ac32801 to nil")
	}
	{
		if err := i.PaymentToken.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsGooglePay#8ac32801: field payment_token: %w", err)
		}
	}
	return nil
}

// GetPaymentToken returns value of PaymentToken field.
func (i *InputPaymentCredentialsGooglePay) GetPaymentToken() (value DataJSON) {
	return i.PaymentToken
}

// InputPaymentCredentialsClass represents InputPaymentCredentials generic type.
//
// See https://core.telegram.org/type/InputPaymentCredentials for reference.
//
// Example:
//  g, err := tg.DecodeInputPaymentCredentials(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputPaymentCredentialsSaved: // inputPaymentCredentialsSaved#c10eb2cf
//  case *tg.InputPaymentCredentials: // inputPaymentCredentials#3417d728
//  case *tg.InputPaymentCredentialsApplePay: // inputPaymentCredentialsApplePay#aa1c39f
//  case *tg.InputPaymentCredentialsGooglePay: // inputPaymentCredentialsGooglePay#8ac32801
//  default: panic(v)
//  }
type InputPaymentCredentialsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputPaymentCredentialsClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputPaymentCredentials implements binary de-serialization for InputPaymentCredentialsClass.
func DecodeInputPaymentCredentials(buf *bin.Buffer) (InputPaymentCredentialsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPaymentCredentialsSavedTypeID:
		// Decoding inputPaymentCredentialsSaved#c10eb2cf.
		v := InputPaymentCredentialsSaved{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsTypeID:
		// Decoding inputPaymentCredentials#3417d728.
		v := InputPaymentCredentials{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsApplePayTypeID:
		// Decoding inputPaymentCredentialsApplePay#aa1c39f.
		v := InputPaymentCredentialsApplePay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsGooglePayTypeID:
		// Decoding inputPaymentCredentialsGooglePay#8ac32801.
		v := InputPaymentCredentialsGooglePay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPaymentCredentials boxes the InputPaymentCredentialsClass providing a helper.
type InputPaymentCredentialsBox struct {
	InputPaymentCredentials InputPaymentCredentialsClass
}

// Decode implements bin.Decoder for InputPaymentCredentialsBox.
func (b *InputPaymentCredentialsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPaymentCredentialsBox to nil")
	}
	v, err := DecodeInputPaymentCredentials(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPaymentCredentials = v
	return nil
}

// Encode implements bin.Encode for InputPaymentCredentialsBox.
func (b *InputPaymentCredentialsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPaymentCredentials == nil {
		return fmt.Errorf("unable to encode InputPaymentCredentialsClass as nil")
	}
	return b.InputPaymentCredentials.Encode(buf)
}
