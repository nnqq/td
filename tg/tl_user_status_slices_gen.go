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

// UserStatusClassArray is adapter for slice of UserStatusClass.
type UserStatusClassArray []UserStatusClass

// Sort sorts slice of UserStatusClass.
func (s UserStatusClassArray) Sort(less func(a, b UserStatusClass) bool) UserStatusClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UserStatusClass.
func (s UserStatusClassArray) SortStable(less func(a, b UserStatusClass) bool) UserStatusClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UserStatusClass.
func (s UserStatusClassArray) Retain(keep func(x UserStatusClass) bool) UserStatusClassArray {
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
func (s UserStatusClassArray) First() (v UserStatusClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UserStatusClassArray) Last() (v UserStatusClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UserStatusClassArray) PopFirst() (v UserStatusClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UserStatusClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UserStatusClassArray) Pop() (v UserStatusClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUserStatusOnline returns copy with only UserStatusOnline constructors.
func (s UserStatusClassArray) AsUserStatusOnline() (to UserStatusOnlineArray) {
	for _, elem := range s {
		value, ok := elem.(*UserStatusOnline)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUserStatusOffline returns copy with only UserStatusOffline constructors.
func (s UserStatusClassArray) AsUserStatusOffline() (to UserStatusOfflineArray) {
	for _, elem := range s {
		value, ok := elem.(*UserStatusOffline)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// UserStatusOnlineArray is adapter for slice of UserStatusOnline.
type UserStatusOnlineArray []UserStatusOnline

// Sort sorts slice of UserStatusOnline.
func (s UserStatusOnlineArray) Sort(less func(a, b UserStatusOnline) bool) UserStatusOnlineArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UserStatusOnline.
func (s UserStatusOnlineArray) SortStable(less func(a, b UserStatusOnline) bool) UserStatusOnlineArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UserStatusOnline.
func (s UserStatusOnlineArray) Retain(keep func(x UserStatusOnline) bool) UserStatusOnlineArray {
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
func (s UserStatusOnlineArray) First() (v UserStatusOnline, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UserStatusOnlineArray) Last() (v UserStatusOnline, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UserStatusOnlineArray) PopFirst() (v UserStatusOnline, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UserStatusOnline
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UserStatusOnlineArray) Pop() (v UserStatusOnline, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UserStatusOfflineArray is adapter for slice of UserStatusOffline.
type UserStatusOfflineArray []UserStatusOffline

// Sort sorts slice of UserStatusOffline.
func (s UserStatusOfflineArray) Sort(less func(a, b UserStatusOffline) bool) UserStatusOfflineArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UserStatusOffline.
func (s UserStatusOfflineArray) SortStable(less func(a, b UserStatusOffline) bool) UserStatusOfflineArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UserStatusOffline.
func (s UserStatusOfflineArray) Retain(keep func(x UserStatusOffline) bool) UserStatusOfflineArray {
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
func (s UserStatusOfflineArray) First() (v UserStatusOffline, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UserStatusOfflineArray) Last() (v UserStatusOffline, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UserStatusOfflineArray) PopFirst() (v UserStatusOffline, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UserStatusOffline
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UserStatusOfflineArray) Pop() (v UserStatusOffline, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
