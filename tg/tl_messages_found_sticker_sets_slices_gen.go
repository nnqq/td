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

// MessagesFoundStickerSetsClassArray is adapter for slice of MessagesFoundStickerSetsClass.
type MessagesFoundStickerSetsClassArray []MessagesFoundStickerSetsClass

// Sort sorts slice of MessagesFoundStickerSetsClass.
func (s MessagesFoundStickerSetsClassArray) Sort(less func(a, b MessagesFoundStickerSetsClass) bool) MessagesFoundStickerSetsClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFoundStickerSetsClass.
func (s MessagesFoundStickerSetsClassArray) SortStable(less func(a, b MessagesFoundStickerSetsClass) bool) MessagesFoundStickerSetsClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFoundStickerSetsClass.
func (s MessagesFoundStickerSetsClassArray) Retain(keep func(x MessagesFoundStickerSetsClass) bool) MessagesFoundStickerSetsClassArray {
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
func (s MessagesFoundStickerSetsClassArray) First() (v MessagesFoundStickerSetsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFoundStickerSetsClassArray) Last() (v MessagesFoundStickerSetsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFoundStickerSetsClassArray) PopFirst() (v MessagesFoundStickerSetsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFoundStickerSetsClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFoundStickerSetsClassArray) Pop() (v MessagesFoundStickerSetsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesFoundStickerSets returns copy with only MessagesFoundStickerSets constructors.
func (s MessagesFoundStickerSetsClassArray) AsMessagesFoundStickerSets() (to MessagesFoundStickerSetsArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesFoundStickerSets)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesFoundStickerSetsClassArray) AppendOnlyModified(to []*MessagesFoundStickerSets) []*MessagesFoundStickerSets {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s MessagesFoundStickerSetsClassArray) AsModified() (to []*MessagesFoundStickerSets) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesFoundStickerSetsClassArray) FirstAsModified() (v *MessagesFoundStickerSets, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesFoundStickerSetsClassArray) LastAsModified() (v *MessagesFoundStickerSets, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesFoundStickerSetsClassArray) PopFirstAsModified() (v *MessagesFoundStickerSets, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesFoundStickerSetsClassArray) PopAsModified() (v *MessagesFoundStickerSets, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesFoundStickerSetsArray is adapter for slice of MessagesFoundStickerSets.
type MessagesFoundStickerSetsArray []MessagesFoundStickerSets

// Sort sorts slice of MessagesFoundStickerSets.
func (s MessagesFoundStickerSetsArray) Sort(less func(a, b MessagesFoundStickerSets) bool) MessagesFoundStickerSetsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFoundStickerSets.
func (s MessagesFoundStickerSetsArray) SortStable(less func(a, b MessagesFoundStickerSets) bool) MessagesFoundStickerSetsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFoundStickerSets.
func (s MessagesFoundStickerSetsArray) Retain(keep func(x MessagesFoundStickerSets) bool) MessagesFoundStickerSetsArray {
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
func (s MessagesFoundStickerSetsArray) First() (v MessagesFoundStickerSets, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFoundStickerSetsArray) Last() (v MessagesFoundStickerSets, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFoundStickerSetsArray) PopFirst() (v MessagesFoundStickerSets, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFoundStickerSets
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFoundStickerSetsArray) Pop() (v MessagesFoundStickerSets, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
