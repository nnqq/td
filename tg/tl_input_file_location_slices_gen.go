//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// InputFileLocationClassArray is adapter for slice of InputFileLocationClass.
type InputFileLocationClassArray []InputFileLocationClass

// Sort sorts slice of InputFileLocationClass.
func (s InputFileLocationClassArray) Sort(less func(a, b InputFileLocationClass) bool) InputFileLocationClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputFileLocationClass.
func (s InputFileLocationClassArray) SortStable(less func(a, b InputFileLocationClass) bool) InputFileLocationClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputFileLocationClass.
func (s InputFileLocationClassArray) Retain(keep func(x InputFileLocationClass) bool) InputFileLocationClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputFileLocationClassArray) First() (v InputFileLocationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputFileLocationClassArray) Last() (v InputFileLocationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputFileLocationClassArray) PopFirst() (v InputFileLocationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputFileLocationClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputFileLocationClassArray) Pop() (v InputFileLocationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputPeerPhotoFileLocationLegacy returns copy with only InputPeerPhotoFileLocationLegacy constructors.
func (s InputFileLocationClassArray) AsInputPeerPhotoFileLocationLegacy() (to InputPeerPhotoFileLocationLegacyArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPeerPhotoFileLocationLegacy)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputStickerSetThumbLegacy returns copy with only InputStickerSetThumbLegacy constructors.
func (s InputFileLocationClassArray) AsInputStickerSetThumbLegacy() (to InputStickerSetThumbLegacyArray) {
	for _, elem := range s {
		value, ok := elem.(*InputStickerSetThumbLegacy)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputFileLocation returns copy with only InputFileLocation constructors.
func (s InputFileLocationClassArray) AsInputFileLocation() (to InputFileLocationArray) {
	for _, elem := range s {
		value, ok := elem.(*InputFileLocation)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputEncryptedFileLocation returns copy with only InputEncryptedFileLocation constructors.
func (s InputFileLocationClassArray) AsInputEncryptedFileLocation() (to InputEncryptedFileLocationArray) {
	for _, elem := range s {
		value, ok := elem.(*InputEncryptedFileLocation)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputDocumentFileLocation returns copy with only InputDocumentFileLocation constructors.
func (s InputFileLocationClassArray) AsInputDocumentFileLocation() (to InputDocumentFileLocationArray) {
	for _, elem := range s {
		value, ok := elem.(*InputDocumentFileLocation)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputSecureFileLocation returns copy with only InputSecureFileLocation constructors.
func (s InputFileLocationClassArray) AsInputSecureFileLocation() (to InputSecureFileLocationArray) {
	for _, elem := range s {
		value, ok := elem.(*InputSecureFileLocation)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPhotoFileLocation returns copy with only InputPhotoFileLocation constructors.
func (s InputFileLocationClassArray) AsInputPhotoFileLocation() (to InputPhotoFileLocationArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPhotoFileLocation)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPhotoLegacyFileLocation returns copy with only InputPhotoLegacyFileLocation constructors.
func (s InputFileLocationClassArray) AsInputPhotoLegacyFileLocation() (to InputPhotoLegacyFileLocationArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPhotoLegacyFileLocation)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPeerPhotoFileLocation returns copy with only InputPeerPhotoFileLocation constructors.
func (s InputFileLocationClassArray) AsInputPeerPhotoFileLocation() (to InputPeerPhotoFileLocationArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPeerPhotoFileLocation)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputStickerSetThumb returns copy with only InputStickerSetThumb constructors.
func (s InputFileLocationClassArray) AsInputStickerSetThumb() (to InputStickerSetThumbArray) {
	for _, elem := range s {
		value, ok := elem.(*InputStickerSetThumb)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputGroupCallStream returns copy with only InputGroupCallStream constructors.
func (s InputFileLocationClassArray) AsInputGroupCallStream() (to InputGroupCallStreamArray) {
	for _, elem := range s {
		value, ok := elem.(*InputGroupCallStream)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputPeerPhotoFileLocationLegacyArray is adapter for slice of InputPeerPhotoFileLocationLegacy.
type InputPeerPhotoFileLocationLegacyArray []InputPeerPhotoFileLocationLegacy

// Sort sorts slice of InputPeerPhotoFileLocationLegacy.
func (s InputPeerPhotoFileLocationLegacyArray) Sort(less func(a, b InputPeerPhotoFileLocationLegacy) bool) InputPeerPhotoFileLocationLegacyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPeerPhotoFileLocationLegacy.
func (s InputPeerPhotoFileLocationLegacyArray) SortStable(less func(a, b InputPeerPhotoFileLocationLegacy) bool) InputPeerPhotoFileLocationLegacyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPeerPhotoFileLocationLegacy.
func (s InputPeerPhotoFileLocationLegacyArray) Retain(keep func(x InputPeerPhotoFileLocationLegacy) bool) InputPeerPhotoFileLocationLegacyArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputPeerPhotoFileLocationLegacyArray) First() (v InputPeerPhotoFileLocationLegacy, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPeerPhotoFileLocationLegacyArray) Last() (v InputPeerPhotoFileLocationLegacy, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPeerPhotoFileLocationLegacyArray) PopFirst() (v InputPeerPhotoFileLocationLegacy, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPeerPhotoFileLocationLegacy
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPeerPhotoFileLocationLegacyArray) Pop() (v InputPeerPhotoFileLocationLegacy, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputStickerSetThumbLegacyArray is adapter for slice of InputStickerSetThumbLegacy.
type InputStickerSetThumbLegacyArray []InputStickerSetThumbLegacy

// Sort sorts slice of InputStickerSetThumbLegacy.
func (s InputStickerSetThumbLegacyArray) Sort(less func(a, b InputStickerSetThumbLegacy) bool) InputStickerSetThumbLegacyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputStickerSetThumbLegacy.
func (s InputStickerSetThumbLegacyArray) SortStable(less func(a, b InputStickerSetThumbLegacy) bool) InputStickerSetThumbLegacyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputStickerSetThumbLegacy.
func (s InputStickerSetThumbLegacyArray) Retain(keep func(x InputStickerSetThumbLegacy) bool) InputStickerSetThumbLegacyArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputStickerSetThumbLegacyArray) First() (v InputStickerSetThumbLegacy, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputStickerSetThumbLegacyArray) Last() (v InputStickerSetThumbLegacy, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputStickerSetThumbLegacyArray) PopFirst() (v InputStickerSetThumbLegacy, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputStickerSetThumbLegacy
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputStickerSetThumbLegacyArray) Pop() (v InputStickerSetThumbLegacy, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputFileLocationArray is adapter for slice of InputFileLocation.
type InputFileLocationArray []InputFileLocation

// Sort sorts slice of InputFileLocation.
func (s InputFileLocationArray) Sort(less func(a, b InputFileLocation) bool) InputFileLocationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputFileLocation.
func (s InputFileLocationArray) SortStable(less func(a, b InputFileLocation) bool) InputFileLocationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputFileLocation.
func (s InputFileLocationArray) Retain(keep func(x InputFileLocation) bool) InputFileLocationArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputFileLocationArray) First() (v InputFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputFileLocationArray) Last() (v InputFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputFileLocationArray) PopFirst() (v InputFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputFileLocation
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputFileLocationArray) Pop() (v InputFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputEncryptedFileLocationArray is adapter for slice of InputEncryptedFileLocation.
type InputEncryptedFileLocationArray []InputEncryptedFileLocation

// Sort sorts slice of InputEncryptedFileLocation.
func (s InputEncryptedFileLocationArray) Sort(less func(a, b InputEncryptedFileLocation) bool) InputEncryptedFileLocationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputEncryptedFileLocation.
func (s InputEncryptedFileLocationArray) SortStable(less func(a, b InputEncryptedFileLocation) bool) InputEncryptedFileLocationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputEncryptedFileLocation.
func (s InputEncryptedFileLocationArray) Retain(keep func(x InputEncryptedFileLocation) bool) InputEncryptedFileLocationArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputEncryptedFileLocationArray) First() (v InputEncryptedFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputEncryptedFileLocationArray) Last() (v InputEncryptedFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputEncryptedFileLocationArray) PopFirst() (v InputEncryptedFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputEncryptedFileLocation
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputEncryptedFileLocationArray) Pop() (v InputEncryptedFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputEncryptedFileLocation by ID.
func (s InputEncryptedFileLocationArray) SortByID() InputEncryptedFileLocationArray {
	return s.Sort(func(a, b InputEncryptedFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputEncryptedFileLocation by ID.
func (s InputEncryptedFileLocationArray) SortStableByID() InputEncryptedFileLocationArray {
	return s.SortStable(func(a, b InputEncryptedFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputEncryptedFileLocationArray) FillMap(to map[int64]InputEncryptedFileLocation) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputEncryptedFileLocationArray) ToMap() map[int64]InputEncryptedFileLocation {
	r := make(map[int64]InputEncryptedFileLocation, len(s))
	s.FillMap(r)
	return r
}

// InputDocumentFileLocationArray is adapter for slice of InputDocumentFileLocation.
type InputDocumentFileLocationArray []InputDocumentFileLocation

// Sort sorts slice of InputDocumentFileLocation.
func (s InputDocumentFileLocationArray) Sort(less func(a, b InputDocumentFileLocation) bool) InputDocumentFileLocationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputDocumentFileLocation.
func (s InputDocumentFileLocationArray) SortStable(less func(a, b InputDocumentFileLocation) bool) InputDocumentFileLocationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputDocumentFileLocation.
func (s InputDocumentFileLocationArray) Retain(keep func(x InputDocumentFileLocation) bool) InputDocumentFileLocationArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputDocumentFileLocationArray) First() (v InputDocumentFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputDocumentFileLocationArray) Last() (v InputDocumentFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputDocumentFileLocationArray) PopFirst() (v InputDocumentFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputDocumentFileLocation
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputDocumentFileLocationArray) Pop() (v InputDocumentFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputDocumentFileLocation by ID.
func (s InputDocumentFileLocationArray) SortByID() InputDocumentFileLocationArray {
	return s.Sort(func(a, b InputDocumentFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputDocumentFileLocation by ID.
func (s InputDocumentFileLocationArray) SortStableByID() InputDocumentFileLocationArray {
	return s.SortStable(func(a, b InputDocumentFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputDocumentFileLocationArray) FillMap(to map[int64]InputDocumentFileLocation) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputDocumentFileLocationArray) ToMap() map[int64]InputDocumentFileLocation {
	r := make(map[int64]InputDocumentFileLocation, len(s))
	s.FillMap(r)
	return r
}

// InputSecureFileLocationArray is adapter for slice of InputSecureFileLocation.
type InputSecureFileLocationArray []InputSecureFileLocation

// Sort sorts slice of InputSecureFileLocation.
func (s InputSecureFileLocationArray) Sort(less func(a, b InputSecureFileLocation) bool) InputSecureFileLocationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputSecureFileLocation.
func (s InputSecureFileLocationArray) SortStable(less func(a, b InputSecureFileLocation) bool) InputSecureFileLocationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputSecureFileLocation.
func (s InputSecureFileLocationArray) Retain(keep func(x InputSecureFileLocation) bool) InputSecureFileLocationArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputSecureFileLocationArray) First() (v InputSecureFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputSecureFileLocationArray) Last() (v InputSecureFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputSecureFileLocationArray) PopFirst() (v InputSecureFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputSecureFileLocation
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputSecureFileLocationArray) Pop() (v InputSecureFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputSecureFileLocation by ID.
func (s InputSecureFileLocationArray) SortByID() InputSecureFileLocationArray {
	return s.Sort(func(a, b InputSecureFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputSecureFileLocation by ID.
func (s InputSecureFileLocationArray) SortStableByID() InputSecureFileLocationArray {
	return s.SortStable(func(a, b InputSecureFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputSecureFileLocationArray) FillMap(to map[int64]InputSecureFileLocation) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputSecureFileLocationArray) ToMap() map[int64]InputSecureFileLocation {
	r := make(map[int64]InputSecureFileLocation, len(s))
	s.FillMap(r)
	return r
}

// InputPhotoFileLocationArray is adapter for slice of InputPhotoFileLocation.
type InputPhotoFileLocationArray []InputPhotoFileLocation

// Sort sorts slice of InputPhotoFileLocation.
func (s InputPhotoFileLocationArray) Sort(less func(a, b InputPhotoFileLocation) bool) InputPhotoFileLocationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPhotoFileLocation.
func (s InputPhotoFileLocationArray) SortStable(less func(a, b InputPhotoFileLocation) bool) InputPhotoFileLocationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPhotoFileLocation.
func (s InputPhotoFileLocationArray) Retain(keep func(x InputPhotoFileLocation) bool) InputPhotoFileLocationArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputPhotoFileLocationArray) First() (v InputPhotoFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPhotoFileLocationArray) Last() (v InputPhotoFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPhotoFileLocationArray) PopFirst() (v InputPhotoFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPhotoFileLocation
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPhotoFileLocationArray) Pop() (v InputPhotoFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputPhotoFileLocation by ID.
func (s InputPhotoFileLocationArray) SortByID() InputPhotoFileLocationArray {
	return s.Sort(func(a, b InputPhotoFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputPhotoFileLocation by ID.
func (s InputPhotoFileLocationArray) SortStableByID() InputPhotoFileLocationArray {
	return s.SortStable(func(a, b InputPhotoFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputPhotoFileLocationArray) FillMap(to map[int64]InputPhotoFileLocation) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputPhotoFileLocationArray) ToMap() map[int64]InputPhotoFileLocation {
	r := make(map[int64]InputPhotoFileLocation, len(s))
	s.FillMap(r)
	return r
}

// InputPhotoLegacyFileLocationArray is adapter for slice of InputPhotoLegacyFileLocation.
type InputPhotoLegacyFileLocationArray []InputPhotoLegacyFileLocation

// Sort sorts slice of InputPhotoLegacyFileLocation.
func (s InputPhotoLegacyFileLocationArray) Sort(less func(a, b InputPhotoLegacyFileLocation) bool) InputPhotoLegacyFileLocationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPhotoLegacyFileLocation.
func (s InputPhotoLegacyFileLocationArray) SortStable(less func(a, b InputPhotoLegacyFileLocation) bool) InputPhotoLegacyFileLocationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPhotoLegacyFileLocation.
func (s InputPhotoLegacyFileLocationArray) Retain(keep func(x InputPhotoLegacyFileLocation) bool) InputPhotoLegacyFileLocationArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputPhotoLegacyFileLocationArray) First() (v InputPhotoLegacyFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPhotoLegacyFileLocationArray) Last() (v InputPhotoLegacyFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPhotoLegacyFileLocationArray) PopFirst() (v InputPhotoLegacyFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPhotoLegacyFileLocation
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPhotoLegacyFileLocationArray) Pop() (v InputPhotoLegacyFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputPhotoLegacyFileLocation by ID.
func (s InputPhotoLegacyFileLocationArray) SortByID() InputPhotoLegacyFileLocationArray {
	return s.Sort(func(a, b InputPhotoLegacyFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputPhotoLegacyFileLocation by ID.
func (s InputPhotoLegacyFileLocationArray) SortStableByID() InputPhotoLegacyFileLocationArray {
	return s.SortStable(func(a, b InputPhotoLegacyFileLocation) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputPhotoLegacyFileLocationArray) FillMap(to map[int64]InputPhotoLegacyFileLocation) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputPhotoLegacyFileLocationArray) ToMap() map[int64]InputPhotoLegacyFileLocation {
	r := make(map[int64]InputPhotoLegacyFileLocation, len(s))
	s.FillMap(r)
	return r
}

// InputPeerPhotoFileLocationArray is adapter for slice of InputPeerPhotoFileLocation.
type InputPeerPhotoFileLocationArray []InputPeerPhotoFileLocation

// Sort sorts slice of InputPeerPhotoFileLocation.
func (s InputPeerPhotoFileLocationArray) Sort(less func(a, b InputPeerPhotoFileLocation) bool) InputPeerPhotoFileLocationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPeerPhotoFileLocation.
func (s InputPeerPhotoFileLocationArray) SortStable(less func(a, b InputPeerPhotoFileLocation) bool) InputPeerPhotoFileLocationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPeerPhotoFileLocation.
func (s InputPeerPhotoFileLocationArray) Retain(keep func(x InputPeerPhotoFileLocation) bool) InputPeerPhotoFileLocationArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputPeerPhotoFileLocationArray) First() (v InputPeerPhotoFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPeerPhotoFileLocationArray) Last() (v InputPeerPhotoFileLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPeerPhotoFileLocationArray) PopFirst() (v InputPeerPhotoFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPeerPhotoFileLocation
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPeerPhotoFileLocationArray) Pop() (v InputPeerPhotoFileLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputStickerSetThumbArray is adapter for slice of InputStickerSetThumb.
type InputStickerSetThumbArray []InputStickerSetThumb

// Sort sorts slice of InputStickerSetThumb.
func (s InputStickerSetThumbArray) Sort(less func(a, b InputStickerSetThumb) bool) InputStickerSetThumbArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputStickerSetThumb.
func (s InputStickerSetThumbArray) SortStable(less func(a, b InputStickerSetThumb) bool) InputStickerSetThumbArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputStickerSetThumb.
func (s InputStickerSetThumbArray) Retain(keep func(x InputStickerSetThumb) bool) InputStickerSetThumbArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputStickerSetThumbArray) First() (v InputStickerSetThumb, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputStickerSetThumbArray) Last() (v InputStickerSetThumb, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputStickerSetThumbArray) PopFirst() (v InputStickerSetThumb, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputStickerSetThumb
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputStickerSetThumbArray) Pop() (v InputStickerSetThumb, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputGroupCallStreamArray is adapter for slice of InputGroupCallStream.
type InputGroupCallStreamArray []InputGroupCallStream

// Sort sorts slice of InputGroupCallStream.
func (s InputGroupCallStreamArray) Sort(less func(a, b InputGroupCallStream) bool) InputGroupCallStreamArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputGroupCallStream.
func (s InputGroupCallStreamArray) SortStable(less func(a, b InputGroupCallStream) bool) InputGroupCallStreamArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputGroupCallStream.
func (s InputGroupCallStreamArray) Retain(keep func(x InputGroupCallStream) bool) InputGroupCallStreamArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputGroupCallStreamArray) First() (v InputGroupCallStream, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputGroupCallStreamArray) Last() (v InputGroupCallStream, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputGroupCallStreamArray) PopFirst() (v InputGroupCallStream, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputGroupCallStream
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputGroupCallStreamArray) Pop() (v InputGroupCallStream, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
