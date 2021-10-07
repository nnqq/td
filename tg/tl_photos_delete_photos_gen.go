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

// PhotosDeletePhotosRequest represents TL type `photos.deletePhotos#87cf7f2f`.
// Deletes profile photos.
//
// See https://core.telegram.org/method/photos.deletePhotos for reference.
type PhotosDeletePhotosRequest struct {
	// Input photos to delete
	ID []InputPhotoClass
}

// PhotosDeletePhotosRequestTypeID is TL type id of PhotosDeletePhotosRequest.
const PhotosDeletePhotosRequestTypeID = 0x87cf7f2f

// Ensuring interfaces in compile-time for PhotosDeletePhotosRequest.
var (
	_ bin.Encoder     = &PhotosDeletePhotosRequest{}
	_ bin.Decoder     = &PhotosDeletePhotosRequest{}
	_ bin.BareEncoder = &PhotosDeletePhotosRequest{}
	_ bin.BareDecoder = &PhotosDeletePhotosRequest{}
)

func (d *PhotosDeletePhotosRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *PhotosDeletePhotosRequest) String() string {
	if d == nil {
		return "PhotosDeletePhotosRequest(nil)"
	}
	type Alias PhotosDeletePhotosRequest
	return fmt.Sprintf("PhotosDeletePhotosRequest%+v", Alias(*d))
}

// FillFrom fills PhotosDeletePhotosRequest from given interface.
func (d *PhotosDeletePhotosRequest) FillFrom(from interface {
	GetID() (value []InputPhotoClass)
}) {
	d.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotosDeletePhotosRequest) TypeID() uint32 {
	return PhotosDeletePhotosRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotosDeletePhotosRequest) TypeName() string {
	return "photos.deletePhotos"
}

// TypeInfo returns info about TL type.
func (d *PhotosDeletePhotosRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photos.deletePhotos",
		ID:   PhotosDeletePhotosRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *PhotosDeletePhotosRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode photos.deletePhotos#87cf7f2f as nil")
	}
	b.PutID(PhotosDeletePhotosRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *PhotosDeletePhotosRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode photos.deletePhotos#87cf7f2f as nil")
	}
	b.PutVectorHeader(len(d.ID))
	for idx, v := range d.ID {
		if v == nil {
			return fmt.Errorf("unable to encode photos.deletePhotos#87cf7f2f: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.deletePhotos#87cf7f2f: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *PhotosDeletePhotosRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode photos.deletePhotos#87cf7f2f to nil")
	}
	if err := b.ConsumeID(PhotosDeletePhotosRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.deletePhotos#87cf7f2f: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *PhotosDeletePhotosRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode photos.deletePhotos#87cf7f2f to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.deletePhotos#87cf7f2f: field id: %w", err)
		}

		if headerLen > 0 {
			d.ID = make([]InputPhotoClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPhoto(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.deletePhotos#87cf7f2f: field id: %w", err)
			}
			d.ID = append(d.ID, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (d *PhotosDeletePhotosRequest) GetID() (value []InputPhotoClass) {
	return d.ID
}

// MapID returns field ID wrapped in InputPhotoClassArray helper.
func (d *PhotosDeletePhotosRequest) MapID() (value InputPhotoClassArray) {
	return InputPhotoClassArray(d.ID)
}

// PhotosDeletePhotos invokes method photos.deletePhotos#87cf7f2f returning error if any.
// Deletes profile photos.
//
// See https://core.telegram.org/method/photos.deletePhotos for reference.
func (c *Client) PhotosDeletePhotos(ctx context.Context, id []InputPhotoClass) ([]int64, error) {
	var result LongVector

	request := &PhotosDeletePhotosRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []int64(result.Elems), nil
}
