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

// MsgsStateInfo represents TL type `msgs_state_info#4deb57d`.
type MsgsStateInfo struct {
	// ReqMsgID field of MsgsStateInfo.
	ReqMsgID int64
	// Info field of MsgsStateInfo.
	Info []byte
}

// MsgsStateInfoTypeID is TL type id of MsgsStateInfo.
const MsgsStateInfoTypeID = 0x4deb57d

// Ensuring interfaces in compile-time for MsgsStateInfo.
var (
	_ bin.Encoder     = &MsgsStateInfo{}
	_ bin.Decoder     = &MsgsStateInfo{}
	_ bin.BareEncoder = &MsgsStateInfo{}
	_ bin.BareDecoder = &MsgsStateInfo{}
)

func (m *MsgsStateInfo) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ReqMsgID == 0) {
		return false
	}
	if !(m.Info == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgsStateInfo) String() string {
	if m == nil {
		return "MsgsStateInfo(nil)"
	}
	type Alias MsgsStateInfo
	return fmt.Sprintf("MsgsStateInfo%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgsStateInfo) TypeID() uint32 {
	return MsgsStateInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgsStateInfo) TypeName() string {
	return "msgs_state_info"
}

// TypeInfo returns info about TL type.
func (m *MsgsStateInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msgs_state_info",
		ID:   MsgsStateInfoTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReqMsgID",
			SchemaName: "req_msg_id",
		},
		{
			Name:       "Info",
			SchemaName: "info",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MsgsStateInfo) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_state_info#4deb57d as nil")
	}
	b.PutID(MsgsStateInfoTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MsgsStateInfo) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_state_info#4deb57d as nil")
	}
	b.PutLong(m.ReqMsgID)
	b.PutBytes(m.Info)
	return nil
}

// Decode implements bin.Decoder.
func (m *MsgsStateInfo) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_state_info#4deb57d to nil")
	}
	if err := b.ConsumeID(MsgsStateInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode msgs_state_info#4deb57d: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MsgsStateInfo) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_state_info#4deb57d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_state_info#4deb57d: field req_msg_id: %w", err)
		}
		m.ReqMsgID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_state_info#4deb57d: field info: %w", err)
		}
		m.Info = value
	}
	return nil
}

// GetReqMsgID returns value of ReqMsgID field.
func (m *MsgsStateInfo) GetReqMsgID() (value int64) {
	return m.ReqMsgID
}

// GetInfo returns value of Info field.
func (m *MsgsStateInfo) GetInfo() (value []byte) {
	return m.Info
}
