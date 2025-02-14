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

// BotInfo represents TL type `botInfo#1b74b335`.
// Info about bots (available bot commands, etc)
//
// See https://core.telegram.org/constructor/botInfo for reference.
type BotInfo struct {
	// ID of the bot
	UserID int64
	// Description of the bot
	Description string
	// Bot commands that can be used in the chat
	Commands []BotCommand
}

// BotInfoTypeID is TL type id of BotInfo.
const BotInfoTypeID = 0x1b74b335

// Ensuring interfaces in compile-time for BotInfo.
var (
	_ bin.Encoder     = &BotInfo{}
	_ bin.Decoder     = &BotInfo{}
	_ bin.BareEncoder = &BotInfo{}
	_ bin.BareDecoder = &BotInfo{}
)

func (b *BotInfo) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.UserID == 0) {
		return false
	}
	if !(b.Description == "") {
		return false
	}
	if !(b.Commands == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotInfo) String() string {
	if b == nil {
		return "BotInfo(nil)"
	}
	type Alias BotInfo
	return fmt.Sprintf("BotInfo%+v", Alias(*b))
}

// FillFrom fills BotInfo from given interface.
func (b *BotInfo) FillFrom(from interface {
	GetUserID() (value int64)
	GetDescription() (value string)
	GetCommands() (value []BotCommand)
}) {
	b.UserID = from.GetUserID()
	b.Description = from.GetDescription()
	b.Commands = from.GetCommands()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotInfo) TypeID() uint32 {
	return BotInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*BotInfo) TypeName() string {
	return "botInfo"
}

// TypeInfo returns info about TL type.
func (b *BotInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botInfo",
		ID:   BotInfoTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "Commands",
			SchemaName: "commands",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotInfo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInfo#1b74b335 as nil")
	}
	buf.PutID(BotInfoTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotInfo) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInfo#1b74b335 as nil")
	}
	buf.PutLong(b.UserID)
	buf.PutString(b.Description)
	buf.PutVectorHeader(len(b.Commands))
	for idx, v := range b.Commands {
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInfo#1b74b335: field commands element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotInfo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInfo#1b74b335 to nil")
	}
	if err := buf.ConsumeID(BotInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode botInfo#1b74b335: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotInfo) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInfo#1b74b335 to nil")
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#1b74b335: field user_id: %w", err)
		}
		b.UserID = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#1b74b335: field description: %w", err)
		}
		b.Description = value
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#1b74b335: field commands: %w", err)
		}

		if headerLen > 0 {
			b.Commands = make([]BotCommand, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotCommand
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode botInfo#1b74b335: field commands: %w", err)
			}
			b.Commands = append(b.Commands, value)
		}
	}
	return nil
}

// GetUserID returns value of UserID field.
func (b *BotInfo) GetUserID() (value int64) {
	return b.UserID
}

// GetDescription returns value of Description field.
func (b *BotInfo) GetDescription() (value string) {
	return b.Description
}

// GetCommands returns value of Commands field.
func (b *BotInfo) GetCommands() (value []BotCommand) {
	return b.Commands
}
