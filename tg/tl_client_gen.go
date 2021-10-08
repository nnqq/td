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

// Invoker can invoke raw MTProto rpc calls.
type Invoker interface {
	Invoke(ctx context.Context, input bin.Encoder, output bin.Decoder) error
}

// Client implement methods for calling functions from TL schema via Invoker.
type Client struct {
	rpc Invoker
}

// Invoker returns Invoker used by this client.
func (c *Client) Invoker() Invoker {
	return c.rpc
}

// NewClient creates new Client.
func NewClient(invoker Invoker) *Client {
	return &Client{
		rpc: invoker,
	}
}
