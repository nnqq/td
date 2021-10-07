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

// InputPhotoClassArray is adapter for slice of InputPhotoClass.
type InputPhotoClassArray []InputPhotoClass

// Sort sorts slice of InputPhotoClass.
func (s InputPhotoClassArray) Sort(less func(a, b InputPhotoClass) bool) InputPhotoClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPhotoClass.
func (s InputPhotoClassArray) SortStable(less func(a, b InputPhotoClass) bool) InputPhotoClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPhotoClass.
func (s InputPhotoClassArray) Retain(keep func(x InputPhotoClass) bool) InputPhotoClassArray {
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
func (s InputPhotoClassArray) First() (v InputPhotoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPhotoClassArray) Last() (v InputPhotoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPhotoClassArray) PopFirst() (v InputPhotoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPhotoClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPhotoClassArray) Pop() (v InputPhotoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputPhoto returns copy with only InputPhoto constructors.
func (s InputPhotoClassArray) AsInputPhoto() (to InputPhotoArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPhoto)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillNotEmptyMap fills only NotEmpty constructors to given map.
func (s InputPhotoClassArray) FillNotEmptyMap(to map[int64]*InputPhoto) {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// NotEmptyToMap collects only NotEmpty constructors to map.
func (s InputPhotoClassArray) NotEmptyToMap() map[int64]*InputPhoto {
	r := make(map[int64]*InputPhoto, len(s))
	s.FillNotEmptyMap(r)
	return r
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s InputPhotoClassArray) AppendOnlyNotEmpty(to []*InputPhoto) []*InputPhoto {
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
func (s InputPhotoClassArray) AsNotEmpty() (to []*InputPhoto) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s InputPhotoClassArray) FirstAsNotEmpty() (v *InputPhoto, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s InputPhotoClassArray) LastAsNotEmpty() (v *InputPhoto, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *InputPhotoClassArray) PopFirstAsNotEmpty() (v *InputPhoto, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *InputPhotoClassArray) PopAsNotEmpty() (v *InputPhoto, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// InputPhotoArray is adapter for slice of InputPhoto.
type InputPhotoArray []InputPhoto

// Sort sorts slice of InputPhoto.
func (s InputPhotoArray) Sort(less func(a, b InputPhoto) bool) InputPhotoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPhoto.
func (s InputPhotoArray) SortStable(less func(a, b InputPhoto) bool) InputPhotoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPhoto.
func (s InputPhotoArray) Retain(keep func(x InputPhoto) bool) InputPhotoArray {
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
func (s InputPhotoArray) First() (v InputPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPhotoArray) Last() (v InputPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPhotoArray) PopFirst() (v InputPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPhoto
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPhotoArray) Pop() (v InputPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputPhoto by ID.
func (s InputPhotoArray) SortByID() InputPhotoArray {
	return s.Sort(func(a, b InputPhoto) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputPhoto by ID.
func (s InputPhotoArray) SortStableByID() InputPhotoArray {
	return s.SortStable(func(a, b InputPhoto) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputPhotoArray) FillMap(to map[int64]InputPhoto) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputPhotoArray) ToMap() map[int64]InputPhoto {
	r := make(map[int64]InputPhoto, len(s))
	s.FillMap(r)
	return r
}
