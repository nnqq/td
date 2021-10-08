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

// DocumentClassArray is adapter for slice of DocumentClass.
type DocumentClassArray []DocumentClass

// Sort sorts slice of DocumentClass.
func (s DocumentClassArray) Sort(less func(a, b DocumentClass) bool) DocumentClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DocumentClass.
func (s DocumentClassArray) SortStable(less func(a, b DocumentClass) bool) DocumentClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DocumentClass.
func (s DocumentClassArray) Retain(keep func(x DocumentClass) bool) DocumentClassArray {
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
func (s DocumentClassArray) First() (v DocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentClassArray) Last() (v DocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentClassArray) PopFirst() (v DocumentClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DocumentClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentClassArray) Pop() (v DocumentClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of DocumentClass by ID.
func (s DocumentClassArray) SortByID() DocumentClassArray {
	return s.Sort(func(a, b DocumentClass) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of DocumentClass by ID.
func (s DocumentClassArray) SortStableByID() DocumentClassArray {
	return s.SortStable(func(a, b DocumentClass) bool {
		return a.GetID() < b.GetID()
	})
}

// FillDocumentEmptyMap fills only DocumentEmpty constructors to given map.
func (s DocumentClassArray) FillDocumentEmptyMap(to map[int64]*DocumentEmpty) {
	for _, elem := range s {
		value, ok := elem.(*DocumentEmpty)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// DocumentEmptyToMap collects only DocumentEmpty constructors to map.
func (s DocumentClassArray) DocumentEmptyToMap() map[int64]*DocumentEmpty {
	r := make(map[int64]*DocumentEmpty, len(s))
	s.FillDocumentEmptyMap(r)
	return r
}

// AsDocumentEmpty returns copy with only DocumentEmpty constructors.
func (s DocumentClassArray) AsDocumentEmpty() (to DocumentEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*DocumentEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillDocumentMap fills only Document constructors to given map.
func (s DocumentClassArray) FillDocumentMap(to map[int64]*Document) {
	for _, elem := range s {
		value, ok := elem.(*Document)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// DocumentToMap collects only Document constructors to map.
func (s DocumentClassArray) DocumentToMap() map[int64]*Document {
	r := make(map[int64]*Document, len(s))
	s.FillDocumentMap(r)
	return r
}

// AsDocument returns copy with only Document constructors.
func (s DocumentClassArray) AsDocument() (to DocumentArray) {
	for _, elem := range s {
		value, ok := elem.(*Document)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillNotEmptyMap fills only NotEmpty constructors to given map.
func (s DocumentClassArray) FillNotEmptyMap(to map[int64]*Document) {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// NotEmptyToMap collects only NotEmpty constructors to map.
func (s DocumentClassArray) NotEmptyToMap() map[int64]*Document {
	r := make(map[int64]*Document, len(s))
	s.FillNotEmptyMap(r)
	return r
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s DocumentClassArray) AppendOnlyNotEmpty(to []*Document) []*Document {
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
func (s DocumentClassArray) AsNotEmpty() (to []*Document) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s DocumentClassArray) FirstAsNotEmpty() (v *Document, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s DocumentClassArray) LastAsNotEmpty() (v *Document, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *DocumentClassArray) PopFirstAsNotEmpty() (v *Document, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *DocumentClassArray) PopAsNotEmpty() (v *Document, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// DocumentEmptyArray is adapter for slice of DocumentEmpty.
type DocumentEmptyArray []DocumentEmpty

// Sort sorts slice of DocumentEmpty.
func (s DocumentEmptyArray) Sort(less func(a, b DocumentEmpty) bool) DocumentEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DocumentEmpty.
func (s DocumentEmptyArray) SortStable(less func(a, b DocumentEmpty) bool) DocumentEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DocumentEmpty.
func (s DocumentEmptyArray) Retain(keep func(x DocumentEmpty) bool) DocumentEmptyArray {
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
func (s DocumentEmptyArray) First() (v DocumentEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentEmptyArray) Last() (v DocumentEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentEmptyArray) PopFirst() (v DocumentEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DocumentEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentEmptyArray) Pop() (v DocumentEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of DocumentEmpty by ID.
func (s DocumentEmptyArray) SortByID() DocumentEmptyArray {
	return s.Sort(func(a, b DocumentEmpty) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of DocumentEmpty by ID.
func (s DocumentEmptyArray) SortStableByID() DocumentEmptyArray {
	return s.SortStable(func(a, b DocumentEmpty) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s DocumentEmptyArray) FillMap(to map[int64]DocumentEmpty) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s DocumentEmptyArray) ToMap() map[int64]DocumentEmpty {
	r := make(map[int64]DocumentEmpty, len(s))
	s.FillMap(r)
	return r
}

// DocumentArray is adapter for slice of Document.
type DocumentArray []Document

// Sort sorts slice of Document.
func (s DocumentArray) Sort(less func(a, b Document) bool) DocumentArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of Document.
func (s DocumentArray) SortStable(less func(a, b Document) bool) DocumentArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of Document.
func (s DocumentArray) Retain(keep func(x Document) bool) DocumentArray {
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
func (s DocumentArray) First() (v Document, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentArray) Last() (v Document, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentArray) PopFirst() (v Document, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero Document
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentArray) Pop() (v Document, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of Document by ID.
func (s DocumentArray) SortByID() DocumentArray {
	return s.Sort(func(a, b Document) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of Document by ID.
func (s DocumentArray) SortStableByID() DocumentArray {
	return s.SortStable(func(a, b Document) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of Document by Date.
func (s DocumentArray) SortByDate() DocumentArray {
	return s.Sort(func(a, b Document) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of Document by Date.
func (s DocumentArray) SortStableByDate() DocumentArray {
	return s.SortStable(func(a, b Document) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s DocumentArray) FillMap(to map[int64]Document) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s DocumentArray) ToMap() map[int64]Document {
	r := make(map[int64]Document, len(s))
	s.FillMap(r)
	return r
}
