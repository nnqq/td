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

// PhoneConnection represents TL type `phoneConnection#9d4c17c0`.
// Identifies an endpoint that can be used to connect to the other user in a phone call
//
// See https://core.telegram.org/constructor/phoneConnection for reference.
type PhoneConnection struct {
	// Endpoint ID
	ID int64
	// IP address of endpoint
	IP string
	// IPv6 address of endpoint
	Ipv6 string
	// Port ID
	Port int
	// Our peer tag
	PeerTag []byte
}

// PhoneConnectionTypeID is TL type id of PhoneConnection.
const PhoneConnectionTypeID = 0x9d4c17c0

// construct implements constructor of PhoneConnectionClass.
func (p PhoneConnection) construct() PhoneConnectionClass { return &p }

// Ensuring interfaces in compile-time for PhoneConnection.
var (
	_ bin.Encoder     = &PhoneConnection{}
	_ bin.Decoder     = &PhoneConnection{}
	_ bin.BareEncoder = &PhoneConnection{}
	_ bin.BareDecoder = &PhoneConnection{}

	_ PhoneConnectionClass = &PhoneConnection{}
)

func (p *PhoneConnection) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.IP == "") {
		return false
	}
	if !(p.Ipv6 == "") {
		return false
	}
	if !(p.Port == 0) {
		return false
	}
	if !(p.PeerTag == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneConnection) String() string {
	if p == nil {
		return "PhoneConnection(nil)"
	}
	type Alias PhoneConnection
	return fmt.Sprintf("PhoneConnection%+v", Alias(*p))
}

// FillFrom fills PhoneConnection from given interface.
func (p *PhoneConnection) FillFrom(from interface {
	GetID() (value int64)
	GetIP() (value string)
	GetIpv6() (value string)
	GetPort() (value int)
	GetPeerTag() (value []byte)
}) {
	p.ID = from.GetID()
	p.IP = from.GetIP()
	p.Ipv6 = from.GetIpv6()
	p.Port = from.GetPort()
	p.PeerTag = from.GetPeerTag()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneConnection) TypeID() uint32 {
	return PhoneConnectionTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneConnection) TypeName() string {
	return "phoneConnection"
}

// TypeInfo returns info about TL type.
func (p *PhoneConnection) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phoneConnection",
		ID:   PhoneConnectionTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "IP",
			SchemaName: "ip",
		},
		{
			Name:       "Ipv6",
			SchemaName: "ipv6",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "PeerTag",
			SchemaName: "peer_tag",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhoneConnection) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneConnection#9d4c17c0 as nil")
	}
	b.PutID(PhoneConnectionTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhoneConnection) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneConnection#9d4c17c0 as nil")
	}
	b.PutLong(p.ID)
	b.PutString(p.IP)
	b.PutString(p.Ipv6)
	b.PutInt(p.Port)
	b.PutBytes(p.PeerTag)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneConnection) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneConnection#9d4c17c0 to nil")
	}
	if err := b.ConsumeID(PhoneConnectionTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhoneConnection) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneConnection#9d4c17c0 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field ip: %w", err)
		}
		p.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field ipv6: %w", err)
		}
		p.Ipv6 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field port: %w", err)
		}
		p.Port = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field peer_tag: %w", err)
		}
		p.PeerTag = value
	}
	return nil
}

// GetID returns value of ID field.
func (p *PhoneConnection) GetID() (value int64) {
	return p.ID
}

// GetIP returns value of IP field.
func (p *PhoneConnection) GetIP() (value string) {
	return p.IP
}

// GetIpv6 returns value of Ipv6 field.
func (p *PhoneConnection) GetIpv6() (value string) {
	return p.Ipv6
}

// GetPort returns value of Port field.
func (p *PhoneConnection) GetPort() (value int) {
	return p.Port
}

// GetPeerTag returns value of PeerTag field.
func (p *PhoneConnection) GetPeerTag() (value []byte) {
	return p.PeerTag
}

// PhoneConnectionWebrtc represents TL type `phoneConnectionWebrtc#635fe375`.
// WebRTC connection parameters
//
// See https://core.telegram.org/constructor/phoneConnectionWebrtc for reference.
type PhoneConnectionWebrtc struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a TURN endpoint
	Turn bool
	// Whether this is a STUN endpoint
	Stun bool
	// Endpoint ID
	ID int64
	// IP address
	IP string
	// IPv6 address
	Ipv6 string
	// Port
	Port int
	// Username
	Username string
	// Password
	Password string
}

// PhoneConnectionWebrtcTypeID is TL type id of PhoneConnectionWebrtc.
const PhoneConnectionWebrtcTypeID = 0x635fe375

// construct implements constructor of PhoneConnectionClass.
func (p PhoneConnectionWebrtc) construct() PhoneConnectionClass { return &p }

// Ensuring interfaces in compile-time for PhoneConnectionWebrtc.
var (
	_ bin.Encoder     = &PhoneConnectionWebrtc{}
	_ bin.Decoder     = &PhoneConnectionWebrtc{}
	_ bin.BareEncoder = &PhoneConnectionWebrtc{}
	_ bin.BareDecoder = &PhoneConnectionWebrtc{}

	_ PhoneConnectionClass = &PhoneConnectionWebrtc{}
)

func (p *PhoneConnectionWebrtc) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Turn == false) {
		return false
	}
	if !(p.Stun == false) {
		return false
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.IP == "") {
		return false
	}
	if !(p.Ipv6 == "") {
		return false
	}
	if !(p.Port == 0) {
		return false
	}
	if !(p.Username == "") {
		return false
	}
	if !(p.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneConnectionWebrtc) String() string {
	if p == nil {
		return "PhoneConnectionWebrtc(nil)"
	}
	type Alias PhoneConnectionWebrtc
	return fmt.Sprintf("PhoneConnectionWebrtc%+v", Alias(*p))
}

// FillFrom fills PhoneConnectionWebrtc from given interface.
func (p *PhoneConnectionWebrtc) FillFrom(from interface {
	GetTurn() (value bool)
	GetStun() (value bool)
	GetID() (value int64)
	GetIP() (value string)
	GetIpv6() (value string)
	GetPort() (value int)
	GetUsername() (value string)
	GetPassword() (value string)
}) {
	p.Turn = from.GetTurn()
	p.Stun = from.GetStun()
	p.ID = from.GetID()
	p.IP = from.GetIP()
	p.Ipv6 = from.GetIpv6()
	p.Port = from.GetPort()
	p.Username = from.GetUsername()
	p.Password = from.GetPassword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneConnectionWebrtc) TypeID() uint32 {
	return PhoneConnectionWebrtcTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneConnectionWebrtc) TypeName() string {
	return "phoneConnectionWebrtc"
}

// TypeInfo returns info about TL type.
func (p *PhoneConnectionWebrtc) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phoneConnectionWebrtc",
		ID:   PhoneConnectionWebrtcTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Turn",
			SchemaName: "turn",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "Stun",
			SchemaName: "stun",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "IP",
			SchemaName: "ip",
		},
		{
			Name:       "Ipv6",
			SchemaName: "ipv6",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhoneConnectionWebrtc) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneConnectionWebrtc#635fe375 as nil")
	}
	b.PutID(PhoneConnectionWebrtcTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhoneConnectionWebrtc) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneConnectionWebrtc#635fe375 as nil")
	}
	if !(p.Turn == false) {
		p.Flags.Set(0)
	}
	if !(p.Stun == false) {
		p.Flags.Set(1)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneConnectionWebrtc#635fe375: field flags: %w", err)
	}
	b.PutLong(p.ID)
	b.PutString(p.IP)
	b.PutString(p.Ipv6)
	b.PutInt(p.Port)
	b.PutString(p.Username)
	b.PutString(p.Password)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneConnectionWebrtc) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneConnectionWebrtc#635fe375 to nil")
	}
	if err := b.ConsumeID(PhoneConnectionWebrtcTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhoneConnectionWebrtc) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneConnectionWebrtc#635fe375 to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field flags: %w", err)
		}
	}
	p.Turn = p.Flags.Has(0)
	p.Stun = p.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field ip: %w", err)
		}
		p.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field ipv6: %w", err)
		}
		p.Ipv6 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field port: %w", err)
		}
		p.Port = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field username: %w", err)
		}
		p.Username = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field password: %w", err)
		}
		p.Password = value
	}
	return nil
}

// SetTurn sets value of Turn conditional field.
func (p *PhoneConnectionWebrtc) SetTurn(value bool) {
	if value {
		p.Flags.Set(0)
		p.Turn = true
	} else {
		p.Flags.Unset(0)
		p.Turn = false
	}
}

// GetTurn returns value of Turn conditional field.
func (p *PhoneConnectionWebrtc) GetTurn() (value bool) {
	return p.Flags.Has(0)
}

// SetStun sets value of Stun conditional field.
func (p *PhoneConnectionWebrtc) SetStun(value bool) {
	if value {
		p.Flags.Set(1)
		p.Stun = true
	} else {
		p.Flags.Unset(1)
		p.Stun = false
	}
}

// GetStun returns value of Stun conditional field.
func (p *PhoneConnectionWebrtc) GetStun() (value bool) {
	return p.Flags.Has(1)
}

// GetID returns value of ID field.
func (p *PhoneConnectionWebrtc) GetID() (value int64) {
	return p.ID
}

// GetIP returns value of IP field.
func (p *PhoneConnectionWebrtc) GetIP() (value string) {
	return p.IP
}

// GetIpv6 returns value of Ipv6 field.
func (p *PhoneConnectionWebrtc) GetIpv6() (value string) {
	return p.Ipv6
}

// GetPort returns value of Port field.
func (p *PhoneConnectionWebrtc) GetPort() (value int) {
	return p.Port
}

// GetUsername returns value of Username field.
func (p *PhoneConnectionWebrtc) GetUsername() (value string) {
	return p.Username
}

// GetPassword returns value of Password field.
func (p *PhoneConnectionWebrtc) GetPassword() (value string) {
	return p.Password
}

// PhoneConnectionClass represents PhoneConnection generic type.
//
// See https://core.telegram.org/type/PhoneConnection for reference.
//
// Example:
//  g, err := tg.DecodePhoneConnection(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.PhoneConnection: // phoneConnection#9d4c17c0
//  case *tg.PhoneConnectionWebrtc: // phoneConnectionWebrtc#635fe375
//  default: panic(v)
//  }
type PhoneConnectionClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PhoneConnectionClass

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

	// Endpoint ID
	GetID() (value int64)

	// IP address of endpoint
	GetIP() (value string)

	// IPv6 address of endpoint
	GetIpv6() (value string)

	// Port ID
	GetPort() (value int)
}

// DecodePhoneConnection implements binary de-serialization for PhoneConnectionClass.
func DecodePhoneConnection(buf *bin.Buffer) (PhoneConnectionClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhoneConnectionTypeID:
		// Decoding phoneConnection#9d4c17c0.
		v := PhoneConnection{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneConnectionClass: %w", err)
		}
		return &v, nil
	case PhoneConnectionWebrtcTypeID:
		// Decoding phoneConnectionWebrtc#635fe375.
		v := PhoneConnectionWebrtc{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneConnectionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhoneConnectionClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhoneConnection boxes the PhoneConnectionClass providing a helper.
type PhoneConnectionBox struct {
	PhoneConnection PhoneConnectionClass
}

// Decode implements bin.Decoder for PhoneConnectionBox.
func (b *PhoneConnectionBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhoneConnectionBox to nil")
	}
	v, err := DecodePhoneConnection(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PhoneConnection = v
	return nil
}

// Encode implements bin.Encode for PhoneConnectionBox.
func (b *PhoneConnectionBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PhoneConnection == nil {
		return fmt.Errorf("unable to encode PhoneConnectionClass as nil")
	}
	return b.PhoneConnection.Encode(buf)
}
