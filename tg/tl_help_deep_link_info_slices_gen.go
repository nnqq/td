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

// HelpDeepLinkInfoClassArray is adapter for slice of HelpDeepLinkInfoClass.
type HelpDeepLinkInfoClassArray []HelpDeepLinkInfoClass

// Sort sorts slice of HelpDeepLinkInfoClass.
func (s HelpDeepLinkInfoClassArray) Sort(less func(a, b HelpDeepLinkInfoClass) bool) HelpDeepLinkInfoClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpDeepLinkInfoClass.
func (s HelpDeepLinkInfoClassArray) SortStable(less func(a, b HelpDeepLinkInfoClass) bool) HelpDeepLinkInfoClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpDeepLinkInfoClass.
func (s HelpDeepLinkInfoClassArray) Retain(keep func(x HelpDeepLinkInfoClass) bool) HelpDeepLinkInfoClassArray {
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
func (s HelpDeepLinkInfoClassArray) First() (v HelpDeepLinkInfoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpDeepLinkInfoClassArray) Last() (v HelpDeepLinkInfoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpDeepLinkInfoClassArray) PopFirst() (v HelpDeepLinkInfoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpDeepLinkInfoClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpDeepLinkInfoClassArray) Pop() (v HelpDeepLinkInfoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpDeepLinkInfo returns copy with only HelpDeepLinkInfo constructors.
func (s HelpDeepLinkInfoClassArray) AsHelpDeepLinkInfo() (to HelpDeepLinkInfoArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpDeepLinkInfo)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s HelpDeepLinkInfoClassArray) AppendOnlyNotEmpty(to []*HelpDeepLinkInfo) []*HelpDeepLinkInfo {
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
func (s HelpDeepLinkInfoClassArray) AsNotEmpty() (to []*HelpDeepLinkInfo) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s HelpDeepLinkInfoClassArray) FirstAsNotEmpty() (v *HelpDeepLinkInfo, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s HelpDeepLinkInfoClassArray) LastAsNotEmpty() (v *HelpDeepLinkInfo, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *HelpDeepLinkInfoClassArray) PopFirstAsNotEmpty() (v *HelpDeepLinkInfo, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *HelpDeepLinkInfoClassArray) PopAsNotEmpty() (v *HelpDeepLinkInfo, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// HelpDeepLinkInfoArray is adapter for slice of HelpDeepLinkInfo.
type HelpDeepLinkInfoArray []HelpDeepLinkInfo

// Sort sorts slice of HelpDeepLinkInfo.
func (s HelpDeepLinkInfoArray) Sort(less func(a, b HelpDeepLinkInfo) bool) HelpDeepLinkInfoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpDeepLinkInfo.
func (s HelpDeepLinkInfoArray) SortStable(less func(a, b HelpDeepLinkInfo) bool) HelpDeepLinkInfoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpDeepLinkInfo.
func (s HelpDeepLinkInfoArray) Retain(keep func(x HelpDeepLinkInfo) bool) HelpDeepLinkInfoArray {
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
func (s HelpDeepLinkInfoArray) First() (v HelpDeepLinkInfo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpDeepLinkInfoArray) Last() (v HelpDeepLinkInfo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpDeepLinkInfoArray) PopFirst() (v HelpDeepLinkInfo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpDeepLinkInfo
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpDeepLinkInfoArray) Pop() (v HelpDeepLinkInfo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
