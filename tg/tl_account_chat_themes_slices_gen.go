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

// AccountChatThemesClassArray is adapter for slice of AccountChatThemesClass.
type AccountChatThemesClassArray []AccountChatThemesClass

// Sort sorts slice of AccountChatThemesClass.
func (s AccountChatThemesClassArray) Sort(less func(a, b AccountChatThemesClass) bool) AccountChatThemesClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountChatThemesClass.
func (s AccountChatThemesClassArray) SortStable(less func(a, b AccountChatThemesClass) bool) AccountChatThemesClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountChatThemesClass.
func (s AccountChatThemesClassArray) Retain(keep func(x AccountChatThemesClass) bool) AccountChatThemesClassArray {
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
func (s AccountChatThemesClassArray) First() (v AccountChatThemesClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountChatThemesClassArray) Last() (v AccountChatThemesClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountChatThemesClassArray) PopFirst() (v AccountChatThemesClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountChatThemesClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountChatThemesClassArray) Pop() (v AccountChatThemesClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAccountChatThemes returns copy with only AccountChatThemes constructors.
func (s AccountChatThemesClassArray) AsAccountChatThemes() (to AccountChatThemesArray) {
	for _, elem := range s {
		value, ok := elem.(*AccountChatThemes)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s AccountChatThemesClassArray) AppendOnlyModified(to []*AccountChatThemes) []*AccountChatThemes {
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
func (s AccountChatThemesClassArray) AsModified() (to []*AccountChatThemes) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s AccountChatThemesClassArray) FirstAsModified() (v *AccountChatThemes, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s AccountChatThemesClassArray) LastAsModified() (v *AccountChatThemes, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *AccountChatThemesClassArray) PopFirstAsModified() (v *AccountChatThemes, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *AccountChatThemesClassArray) PopAsModified() (v *AccountChatThemes, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// AccountChatThemesArray is adapter for slice of AccountChatThemes.
type AccountChatThemesArray []AccountChatThemes

// Sort sorts slice of AccountChatThemes.
func (s AccountChatThemesArray) Sort(less func(a, b AccountChatThemes) bool) AccountChatThemesArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountChatThemes.
func (s AccountChatThemesArray) SortStable(less func(a, b AccountChatThemes) bool) AccountChatThemesArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountChatThemes.
func (s AccountChatThemesArray) Retain(keep func(x AccountChatThemes) bool) AccountChatThemesArray {
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
func (s AccountChatThemesArray) First() (v AccountChatThemes, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountChatThemesArray) Last() (v AccountChatThemes, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountChatThemesArray) PopFirst() (v AccountChatThemes, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountChatThemes
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountChatThemesArray) Pop() (v AccountChatThemes, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
