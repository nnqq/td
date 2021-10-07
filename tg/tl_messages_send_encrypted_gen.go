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

// MessagesSendEncryptedRequest represents TL type `messages.sendEncrypted#44fa7a15`.
// Sends a text message to a secret chat.
//
// See https://core.telegram.org/method/messages.sendEncrypted for reference.
type MessagesSendEncryptedRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Send encrypted message without a notification
	Silent bool
	// Secret chat ID
	Peer InputEncryptedChat
	// Unique client message ID, necessary to avoid message resending
	RandomID int64
	// TL-serialization of DecryptedMessage¹ type, encrypted with a key that was created
	// during chat initialization
	//
	// Links:
	//  1) https://core.telegram.org/type/DecryptedMessage
	Data []byte
}

// MessagesSendEncryptedRequestTypeID is TL type id of MessagesSendEncryptedRequest.
const MessagesSendEncryptedRequestTypeID = 0x44fa7a15

// Ensuring interfaces in compile-time for MessagesSendEncryptedRequest.
var (
	_ bin.Encoder     = &MessagesSendEncryptedRequest{}
	_ bin.Decoder     = &MessagesSendEncryptedRequest{}
	_ bin.BareEncoder = &MessagesSendEncryptedRequest{}
	_ bin.BareDecoder = &MessagesSendEncryptedRequest{}
)

func (s *MessagesSendEncryptedRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendEncryptedRequest) String() string {
	if s == nil {
		return "MessagesSendEncryptedRequest(nil)"
	}
	type Alias MessagesSendEncryptedRequest
	return fmt.Sprintf("MessagesSendEncryptedRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendEncryptedRequest from given interface.
func (s *MessagesSendEncryptedRequest) FillFrom(from interface {
	GetSilent() (value bool)
	GetPeer() (value InputEncryptedChat)
	GetRandomID() (value int64)
	GetData() (value []byte)
}) {
	s.Silent = from.GetSilent()
	s.Peer = from.GetPeer()
	s.RandomID = from.GetRandomID()
	s.Data = from.GetData()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendEncryptedRequest) TypeID() uint32 {
	return MessagesSendEncryptedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendEncryptedRequest) TypeName() string {
	return "messages.sendEncrypted"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendEncryptedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendEncrypted",
		ID:   MessagesSendEncryptedRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSendEncryptedRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendEncrypted#44fa7a15 as nil")
	}
	b.PutID(MessagesSendEncryptedRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendEncryptedRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendEncrypted#44fa7a15 as nil")
	}
	if !(s.Silent == false) {
		s.Flags.Set(0)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncrypted#44fa7a15: field flags: %w", err)
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncrypted#44fa7a15: field peer: %w", err)
	}
	b.PutLong(s.RandomID)
	b.PutBytes(s.Data)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendEncryptedRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendEncrypted#44fa7a15 to nil")
	}
	if err := b.ConsumeID(MessagesSendEncryptedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendEncryptedRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendEncrypted#44fa7a15 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(0)
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: field peer: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncrypted#44fa7a15: field data: %w", err)
		}
		s.Data = value
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendEncryptedRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(0)
		s.Silent = true
	} else {
		s.Flags.Unset(0)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendEncryptedRequest) GetSilent() (value bool) {
	return s.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendEncryptedRequest) GetPeer() (value InputEncryptedChat) {
	return s.Peer
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendEncryptedRequest) GetRandomID() (value int64) {
	return s.RandomID
}

// GetData returns value of Data field.
func (s *MessagesSendEncryptedRequest) GetData() (value []byte) {
	return s.Data
}

// MessagesSendEncrypted invokes method messages.sendEncrypted#44fa7a15 returning error if any.
// Sends a text message to a secret chat.
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 DATA_INVALID: Encrypted data invalid
//  400 ENCRYPTION_DECLINED: The secret chat was declined
//  400 MSG_WAIT_FAILED: A waiting call returned an error
//
// See https://core.telegram.org/method/messages.sendEncrypted for reference.
func (c *Client) MessagesSendEncrypted(ctx context.Context, request *MessagesSendEncryptedRequest) (MessagesSentEncryptedMessageClass, error) {
	var result MessagesSentEncryptedMessageBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SentEncryptedMessage, nil
}
