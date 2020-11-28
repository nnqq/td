// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// SendRequest represents TL type `send#f74488a`.
type SendRequest struct {
	// Msg field of SendRequest.
	Msg SMS
}

// SendRequestTypeID is TL type id of SendRequest.
const SendRequestTypeID = 0xf74488a

// Encode implements bin.Encoder.
func (s *SendRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode send#f74488a as nil")
	}
	b.PutID(SendRequestTypeID)
	if err := s.Msg.Encode(b); err != nil {
		return fmt.Errorf("unable to encode send#f74488a: field msg: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SendRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode send#f74488a to nil")
	}
	if err := b.ConsumeID(SendRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode send#f74488a: %w", err)
	}
	{
		if err := s.Msg.Decode(b); err != nil {
			return fmt.Errorf("unable to decode send#f74488a: field msg: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for SendRequest.
var (
	_ bin.Encoder = &SendRequest{}
	_ bin.Decoder = &SendRequest{}
)
