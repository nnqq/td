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

// ReplyMarkupClassArray is adapter for slice of ReplyMarkupClass.
type ReplyMarkupClassArray []ReplyMarkupClass

// Sort sorts slice of ReplyMarkupClass.
func (s ReplyMarkupClassArray) Sort(less func(a, b ReplyMarkupClass) bool) ReplyMarkupClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyMarkupClass.
func (s ReplyMarkupClassArray) SortStable(less func(a, b ReplyMarkupClass) bool) ReplyMarkupClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyMarkupClass.
func (s ReplyMarkupClassArray) Retain(keep func(x ReplyMarkupClass) bool) ReplyMarkupClassArray {
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
func (s ReplyMarkupClassArray) First() (v ReplyMarkupClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyMarkupClassArray) Last() (v ReplyMarkupClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyMarkupClassArray) PopFirst() (v ReplyMarkupClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyMarkupClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyMarkupClassArray) Pop() (v ReplyMarkupClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsReplyKeyboardHide returns copy with only ReplyKeyboardHide constructors.
func (s ReplyMarkupClassArray) AsReplyKeyboardHide() (to ReplyKeyboardHideArray) {
	for _, elem := range s {
		value, ok := elem.(*ReplyKeyboardHide)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsReplyKeyboardForceReply returns copy with only ReplyKeyboardForceReply constructors.
func (s ReplyMarkupClassArray) AsReplyKeyboardForceReply() (to ReplyKeyboardForceReplyArray) {
	for _, elem := range s {
		value, ok := elem.(*ReplyKeyboardForceReply)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsReplyKeyboardMarkup returns copy with only ReplyKeyboardMarkup constructors.
func (s ReplyMarkupClassArray) AsReplyKeyboardMarkup() (to ReplyKeyboardMarkupArray) {
	for _, elem := range s {
		value, ok := elem.(*ReplyKeyboardMarkup)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsReplyInlineMarkup returns copy with only ReplyInlineMarkup constructors.
func (s ReplyMarkupClassArray) AsReplyInlineMarkup() (to ReplyInlineMarkupArray) {
	for _, elem := range s {
		value, ok := elem.(*ReplyInlineMarkup)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ReplyKeyboardHideArray is adapter for slice of ReplyKeyboardHide.
type ReplyKeyboardHideArray []ReplyKeyboardHide

// Sort sorts slice of ReplyKeyboardHide.
func (s ReplyKeyboardHideArray) Sort(less func(a, b ReplyKeyboardHide) bool) ReplyKeyboardHideArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyKeyboardHide.
func (s ReplyKeyboardHideArray) SortStable(less func(a, b ReplyKeyboardHide) bool) ReplyKeyboardHideArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyKeyboardHide.
func (s ReplyKeyboardHideArray) Retain(keep func(x ReplyKeyboardHide) bool) ReplyKeyboardHideArray {
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
func (s ReplyKeyboardHideArray) First() (v ReplyKeyboardHide, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyKeyboardHideArray) Last() (v ReplyKeyboardHide, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyKeyboardHideArray) PopFirst() (v ReplyKeyboardHide, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyKeyboardHide
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyKeyboardHideArray) Pop() (v ReplyKeyboardHide, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ReplyKeyboardForceReplyArray is adapter for slice of ReplyKeyboardForceReply.
type ReplyKeyboardForceReplyArray []ReplyKeyboardForceReply

// Sort sorts slice of ReplyKeyboardForceReply.
func (s ReplyKeyboardForceReplyArray) Sort(less func(a, b ReplyKeyboardForceReply) bool) ReplyKeyboardForceReplyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyKeyboardForceReply.
func (s ReplyKeyboardForceReplyArray) SortStable(less func(a, b ReplyKeyboardForceReply) bool) ReplyKeyboardForceReplyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyKeyboardForceReply.
func (s ReplyKeyboardForceReplyArray) Retain(keep func(x ReplyKeyboardForceReply) bool) ReplyKeyboardForceReplyArray {
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
func (s ReplyKeyboardForceReplyArray) First() (v ReplyKeyboardForceReply, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyKeyboardForceReplyArray) Last() (v ReplyKeyboardForceReply, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyKeyboardForceReplyArray) PopFirst() (v ReplyKeyboardForceReply, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyKeyboardForceReply
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyKeyboardForceReplyArray) Pop() (v ReplyKeyboardForceReply, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ReplyKeyboardMarkupArray is adapter for slice of ReplyKeyboardMarkup.
type ReplyKeyboardMarkupArray []ReplyKeyboardMarkup

// Sort sorts slice of ReplyKeyboardMarkup.
func (s ReplyKeyboardMarkupArray) Sort(less func(a, b ReplyKeyboardMarkup) bool) ReplyKeyboardMarkupArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyKeyboardMarkup.
func (s ReplyKeyboardMarkupArray) SortStable(less func(a, b ReplyKeyboardMarkup) bool) ReplyKeyboardMarkupArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyKeyboardMarkup.
func (s ReplyKeyboardMarkupArray) Retain(keep func(x ReplyKeyboardMarkup) bool) ReplyKeyboardMarkupArray {
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
func (s ReplyKeyboardMarkupArray) First() (v ReplyKeyboardMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyKeyboardMarkupArray) Last() (v ReplyKeyboardMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyKeyboardMarkupArray) PopFirst() (v ReplyKeyboardMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyKeyboardMarkup
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyKeyboardMarkupArray) Pop() (v ReplyKeyboardMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ReplyInlineMarkupArray is adapter for slice of ReplyInlineMarkup.
type ReplyInlineMarkupArray []ReplyInlineMarkup

// Sort sorts slice of ReplyInlineMarkup.
func (s ReplyInlineMarkupArray) Sort(less func(a, b ReplyInlineMarkup) bool) ReplyInlineMarkupArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyInlineMarkup.
func (s ReplyInlineMarkupArray) SortStable(less func(a, b ReplyInlineMarkup) bool) ReplyInlineMarkupArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyInlineMarkup.
func (s ReplyInlineMarkupArray) Retain(keep func(x ReplyInlineMarkup) bool) ReplyInlineMarkupArray {
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
func (s ReplyInlineMarkupArray) First() (v ReplyInlineMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyInlineMarkupArray) Last() (v ReplyInlineMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyInlineMarkupArray) PopFirst() (v ReplyInlineMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyInlineMarkup
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyInlineMarkupArray) Pop() (v ReplyInlineMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
