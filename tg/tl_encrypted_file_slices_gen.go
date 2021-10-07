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

// EncryptedFileClassArray is adapter for slice of EncryptedFileClass.
type EncryptedFileClassArray []EncryptedFileClass

// Sort sorts slice of EncryptedFileClass.
func (s EncryptedFileClassArray) Sort(less func(a, b EncryptedFileClass) bool) EncryptedFileClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EncryptedFileClass.
func (s EncryptedFileClassArray) SortStable(less func(a, b EncryptedFileClass) bool) EncryptedFileClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EncryptedFileClass.
func (s EncryptedFileClassArray) Retain(keep func(x EncryptedFileClass) bool) EncryptedFileClassArray {
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
func (s EncryptedFileClassArray) First() (v EncryptedFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EncryptedFileClassArray) Last() (v EncryptedFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EncryptedFileClassArray) PopFirst() (v EncryptedFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EncryptedFileClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EncryptedFileClassArray) Pop() (v EncryptedFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsEncryptedFile returns copy with only EncryptedFile constructors.
func (s EncryptedFileClassArray) AsEncryptedFile() (to EncryptedFileArray) {
	for _, elem := range s {
		value, ok := elem.(*EncryptedFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillNotEmptyMap fills only NotEmpty constructors to given map.
func (s EncryptedFileClassArray) FillNotEmptyMap(to map[int64]*EncryptedFile) {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// NotEmptyToMap collects only NotEmpty constructors to map.
func (s EncryptedFileClassArray) NotEmptyToMap() map[int64]*EncryptedFile {
	r := make(map[int64]*EncryptedFile, len(s))
	s.FillNotEmptyMap(r)
	return r
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s EncryptedFileClassArray) AppendOnlyNotEmpty(to []*EncryptedFile) []*EncryptedFile {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s EncryptedFileClassArray) AsNotEmpty() (to []*EncryptedFile) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s EncryptedFileClassArray) FirstAsNotEmpty() (v *EncryptedFile, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s EncryptedFileClassArray) LastAsNotEmpty() (v *EncryptedFile, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *EncryptedFileClassArray) PopFirstAsNotEmpty() (v *EncryptedFile, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *EncryptedFileClassArray) PopAsNotEmpty() (v *EncryptedFile, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// EncryptedFileArray is adapter for slice of EncryptedFile.
type EncryptedFileArray []EncryptedFile

// Sort sorts slice of EncryptedFile.
func (s EncryptedFileArray) Sort(less func(a, b EncryptedFile) bool) EncryptedFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EncryptedFile.
func (s EncryptedFileArray) SortStable(less func(a, b EncryptedFile) bool) EncryptedFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EncryptedFile.
func (s EncryptedFileArray) Retain(keep func(x EncryptedFile) bool) EncryptedFileArray {
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
func (s EncryptedFileArray) First() (v EncryptedFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EncryptedFileArray) Last() (v EncryptedFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EncryptedFileArray) PopFirst() (v EncryptedFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EncryptedFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EncryptedFileArray) Pop() (v EncryptedFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of EncryptedFile by ID.
func (s EncryptedFileArray) SortByID() EncryptedFileArray {
	return s.Sort(func(a, b EncryptedFile) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of EncryptedFile by ID.
func (s EncryptedFileArray) SortStableByID() EncryptedFileArray {
	return s.SortStable(func(a, b EncryptedFile) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s EncryptedFileArray) FillMap(to map[int64]EncryptedFile) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s EncryptedFileArray) ToMap() map[int64]EncryptedFile {
	r := make(map[int64]EncryptedFile, len(s))
	s.FillMap(r)
	return r
}
