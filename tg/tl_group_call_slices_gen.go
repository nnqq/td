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

// GroupCallClassArray is adapter for slice of GroupCallClass.
type GroupCallClassArray []GroupCallClass

// Sort sorts slice of GroupCallClass.
func (s GroupCallClassArray) Sort(less func(a, b GroupCallClass) bool) GroupCallClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of GroupCallClass.
func (s GroupCallClassArray) SortStable(less func(a, b GroupCallClass) bool) GroupCallClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of GroupCallClass.
func (s GroupCallClassArray) Retain(keep func(x GroupCallClass) bool) GroupCallClassArray {
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
func (s GroupCallClassArray) First() (v GroupCallClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s GroupCallClassArray) Last() (v GroupCallClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *GroupCallClassArray) PopFirst() (v GroupCallClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero GroupCallClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *GroupCallClassArray) Pop() (v GroupCallClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of GroupCallClass by ID.
func (s GroupCallClassArray) SortByID() GroupCallClassArray {
	return s.Sort(func(a, b GroupCallClass) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of GroupCallClass by ID.
func (s GroupCallClassArray) SortStableByID() GroupCallClassArray {
	return s.SortStable(func(a, b GroupCallClass) bool {
		return a.GetID() < b.GetID()
	})
}

// FillGroupCallDiscardedMap fills only GroupCallDiscarded constructors to given map.
func (s GroupCallClassArray) FillGroupCallDiscardedMap(to map[int64]*GroupCallDiscarded) {
	for _, elem := range s {
		value, ok := elem.(*GroupCallDiscarded)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// GroupCallDiscardedToMap collects only GroupCallDiscarded constructors to map.
func (s GroupCallClassArray) GroupCallDiscardedToMap() map[int64]*GroupCallDiscarded {
	r := make(map[int64]*GroupCallDiscarded, len(s))
	s.FillGroupCallDiscardedMap(r)
	return r
}

// AsGroupCallDiscarded returns copy with only GroupCallDiscarded constructors.
func (s GroupCallClassArray) AsGroupCallDiscarded() (to GroupCallDiscardedArray) {
	for _, elem := range s {
		value, ok := elem.(*GroupCallDiscarded)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillGroupCallMap fills only GroupCall constructors to given map.
func (s GroupCallClassArray) FillGroupCallMap(to map[int64]*GroupCall) {
	for _, elem := range s {
		value, ok := elem.(*GroupCall)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// GroupCallToMap collects only GroupCall constructors to map.
func (s GroupCallClassArray) GroupCallToMap() map[int64]*GroupCall {
	r := make(map[int64]*GroupCall, len(s))
	s.FillGroupCallMap(r)
	return r
}

// AsGroupCall returns copy with only GroupCall constructors.
func (s GroupCallClassArray) AsGroupCall() (to GroupCallArray) {
	for _, elem := range s {
		value, ok := elem.(*GroupCall)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// GroupCallDiscardedArray is adapter for slice of GroupCallDiscarded.
type GroupCallDiscardedArray []GroupCallDiscarded

// Sort sorts slice of GroupCallDiscarded.
func (s GroupCallDiscardedArray) Sort(less func(a, b GroupCallDiscarded) bool) GroupCallDiscardedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of GroupCallDiscarded.
func (s GroupCallDiscardedArray) SortStable(less func(a, b GroupCallDiscarded) bool) GroupCallDiscardedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of GroupCallDiscarded.
func (s GroupCallDiscardedArray) Retain(keep func(x GroupCallDiscarded) bool) GroupCallDiscardedArray {
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
func (s GroupCallDiscardedArray) First() (v GroupCallDiscarded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s GroupCallDiscardedArray) Last() (v GroupCallDiscarded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *GroupCallDiscardedArray) PopFirst() (v GroupCallDiscarded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero GroupCallDiscarded
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *GroupCallDiscardedArray) Pop() (v GroupCallDiscarded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of GroupCallDiscarded by ID.
func (s GroupCallDiscardedArray) SortByID() GroupCallDiscardedArray {
	return s.Sort(func(a, b GroupCallDiscarded) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of GroupCallDiscarded by ID.
func (s GroupCallDiscardedArray) SortStableByID() GroupCallDiscardedArray {
	return s.SortStable(func(a, b GroupCallDiscarded) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s GroupCallDiscardedArray) FillMap(to map[int64]GroupCallDiscarded) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s GroupCallDiscardedArray) ToMap() map[int64]GroupCallDiscarded {
	r := make(map[int64]GroupCallDiscarded, len(s))
	s.FillMap(r)
	return r
}

// GroupCallArray is adapter for slice of GroupCall.
type GroupCallArray []GroupCall

// Sort sorts slice of GroupCall.
func (s GroupCallArray) Sort(less func(a, b GroupCall) bool) GroupCallArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of GroupCall.
func (s GroupCallArray) SortStable(less func(a, b GroupCall) bool) GroupCallArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of GroupCall.
func (s GroupCallArray) Retain(keep func(x GroupCall) bool) GroupCallArray {
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
func (s GroupCallArray) First() (v GroupCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s GroupCallArray) Last() (v GroupCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *GroupCallArray) PopFirst() (v GroupCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero GroupCall
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *GroupCallArray) Pop() (v GroupCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of GroupCall by ID.
func (s GroupCallArray) SortByID() GroupCallArray {
	return s.Sort(func(a, b GroupCall) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of GroupCall by ID.
func (s GroupCallArray) SortStableByID() GroupCallArray {
	return s.SortStable(func(a, b GroupCall) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s GroupCallArray) FillMap(to map[int64]GroupCall) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s GroupCallArray) ToMap() map[int64]GroupCall {
	r := make(map[int64]GroupCall, len(s))
	s.FillMap(r)
	return r
}
