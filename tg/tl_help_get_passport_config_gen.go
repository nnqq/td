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

// HelpGetPassportConfigRequest represents TL type `help.getPassportConfig#c661ad08`.
// Get passport¹ configuration
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/help.getPassportConfig for reference.
type HelpGetPassportConfigRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// HelpGetPassportConfigRequestTypeID is TL type id of HelpGetPassportConfigRequest.
const HelpGetPassportConfigRequestTypeID = 0xc661ad08

// Ensuring interfaces in compile-time for HelpGetPassportConfigRequest.
var (
	_ bin.Encoder     = &HelpGetPassportConfigRequest{}
	_ bin.Decoder     = &HelpGetPassportConfigRequest{}
	_ bin.BareEncoder = &HelpGetPassportConfigRequest{}
	_ bin.BareDecoder = &HelpGetPassportConfigRequest{}
)

func (g *HelpGetPassportConfigRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetPassportConfigRequest) String() string {
	if g == nil {
		return "HelpGetPassportConfigRequest(nil)"
	}
	type Alias HelpGetPassportConfigRequest
	return fmt.Sprintf("HelpGetPassportConfigRequest%+v", Alias(*g))
}

// FillFrom fills HelpGetPassportConfigRequest from given interface.
func (g *HelpGetPassportConfigRequest) FillFrom(from interface {
	GetHash() (value int)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetPassportConfigRequest) TypeID() uint32 {
	return HelpGetPassportConfigRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetPassportConfigRequest) TypeName() string {
	return "help.getPassportConfig"
}

// TypeInfo returns info about TL type.
func (g *HelpGetPassportConfigRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getPassportConfig",
		ID:   HelpGetPassportConfigRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetPassportConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getPassportConfig#c661ad08 as nil")
	}
	b.PutID(HelpGetPassportConfigRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetPassportConfigRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getPassportConfig#c661ad08 as nil")
	}
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetPassportConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getPassportConfig#c661ad08 to nil")
	}
	if err := b.ConsumeID(HelpGetPassportConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getPassportConfig#c661ad08: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetPassportConfigRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getPassportConfig#c661ad08 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.getPassportConfig#c661ad08: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *HelpGetPassportConfigRequest) GetHash() (value int) {
	return g.Hash
}

// HelpGetPassportConfig invokes method help.getPassportConfig#c661ad08 returning error if any.
// Get passport¹ configuration
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/help.getPassportConfig for reference.
func (c *Client) HelpGetPassportConfig(ctx context.Context, hash int) (HelpPassportConfigClass, error) {
	var result HelpPassportConfigBox

	request := &HelpGetPassportConfigRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PassportConfig, nil
}
