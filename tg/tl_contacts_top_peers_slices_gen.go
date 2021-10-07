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

// ContactsTopPeersClassArray is adapter for slice of ContactsTopPeersClass.
type ContactsTopPeersClassArray []ContactsTopPeersClass

// Sort sorts slice of ContactsTopPeersClass.
func (s ContactsTopPeersClassArray) Sort(less func(a, b ContactsTopPeersClass) bool) ContactsTopPeersClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ContactsTopPeersClass.
func (s ContactsTopPeersClassArray) SortStable(less func(a, b ContactsTopPeersClass) bool) ContactsTopPeersClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ContactsTopPeersClass.
func (s ContactsTopPeersClassArray) Retain(keep func(x ContactsTopPeersClass) bool) ContactsTopPeersClassArray {
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
func (s ContactsTopPeersClassArray) First() (v ContactsTopPeersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ContactsTopPeersClassArray) Last() (v ContactsTopPeersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ContactsTopPeersClassArray) PopFirst() (v ContactsTopPeersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ContactsTopPeersClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ContactsTopPeersClassArray) Pop() (v ContactsTopPeersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsContactsTopPeers returns copy with only ContactsTopPeers constructors.
func (s ContactsTopPeersClassArray) AsContactsTopPeers() (to ContactsTopPeersArray) {
	for _, elem := range s {
		value, ok := elem.(*ContactsTopPeers)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ContactsTopPeersArray is adapter for slice of ContactsTopPeers.
type ContactsTopPeersArray []ContactsTopPeers

// Sort sorts slice of ContactsTopPeers.
func (s ContactsTopPeersArray) Sort(less func(a, b ContactsTopPeers) bool) ContactsTopPeersArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ContactsTopPeers.
func (s ContactsTopPeersArray) SortStable(less func(a, b ContactsTopPeers) bool) ContactsTopPeersArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ContactsTopPeers.
func (s ContactsTopPeersArray) Retain(keep func(x ContactsTopPeers) bool) ContactsTopPeersArray {
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
func (s ContactsTopPeersArray) First() (v ContactsTopPeers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ContactsTopPeersArray) Last() (v ContactsTopPeers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ContactsTopPeersArray) PopFirst() (v ContactsTopPeers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ContactsTopPeers
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ContactsTopPeersArray) Pop() (v ContactsTopPeers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
