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

// StatsMessageStats represents TL type `stats.messageStats#8999f295`.
// Message statistics
//
// See https://core.telegram.org/constructor/stats.messageStats for reference.
type StatsMessageStats struct {
	// Message view graph
	ViewsGraph StatsGraphClass
}

// StatsMessageStatsTypeID is TL type id of StatsMessageStats.
const StatsMessageStatsTypeID = 0x8999f295

// Ensuring interfaces in compile-time for StatsMessageStats.
var (
	_ bin.Encoder     = &StatsMessageStats{}
	_ bin.Decoder     = &StatsMessageStats{}
	_ bin.BareEncoder = &StatsMessageStats{}
	_ bin.BareDecoder = &StatsMessageStats{}
)

func (m *StatsMessageStats) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ViewsGraph == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *StatsMessageStats) String() string {
	if m == nil {
		return "StatsMessageStats(nil)"
	}
	type Alias StatsMessageStats
	return fmt.Sprintf("StatsMessageStats%+v", Alias(*m))
}

// FillFrom fills StatsMessageStats from given interface.
func (m *StatsMessageStats) FillFrom(from interface {
	GetViewsGraph() (value StatsGraphClass)
}) {
	m.ViewsGraph = from.GetViewsGraph()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsMessageStats) TypeID() uint32 {
	return StatsMessageStatsTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsMessageStats) TypeName() string {
	return "stats.messageStats"
}

// TypeInfo returns info about TL type.
func (m *StatsMessageStats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.messageStats",
		ID:   StatsMessageStatsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ViewsGraph",
			SchemaName: "views_graph",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *StatsMessageStats) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode stats.messageStats#8999f295 as nil")
	}
	b.PutID(StatsMessageStatsTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *StatsMessageStats) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode stats.messageStats#8999f295 as nil")
	}
	if m.ViewsGraph == nil {
		return fmt.Errorf("unable to encode stats.messageStats#8999f295: field views_graph is nil")
	}
	if err := m.ViewsGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.messageStats#8999f295: field views_graph: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *StatsMessageStats) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode stats.messageStats#8999f295 to nil")
	}
	if err := b.ConsumeID(StatsMessageStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.messageStats#8999f295: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *StatsMessageStats) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode stats.messageStats#8999f295 to nil")
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.messageStats#8999f295: field views_graph: %w", err)
		}
		m.ViewsGraph = value
	}
	return nil
}

// GetViewsGraph returns value of ViewsGraph field.
func (m *StatsMessageStats) GetViewsGraph() (value StatsGraphClass) {
	return m.ViewsGraph
}
