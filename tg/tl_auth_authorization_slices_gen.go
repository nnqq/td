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

// AuthAuthorizationClassArray is adapter for slice of AuthAuthorizationClass.
type AuthAuthorizationClassArray []AuthAuthorizationClass

// Sort sorts slice of AuthAuthorizationClass.
func (s AuthAuthorizationClassArray) Sort(less func(a, b AuthAuthorizationClass) bool) AuthAuthorizationClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthAuthorizationClass.
func (s AuthAuthorizationClassArray) SortStable(less func(a, b AuthAuthorizationClass) bool) AuthAuthorizationClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthAuthorizationClass.
func (s AuthAuthorizationClassArray) Retain(keep func(x AuthAuthorizationClass) bool) AuthAuthorizationClassArray {
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
func (s AuthAuthorizationClassArray) First() (v AuthAuthorizationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthAuthorizationClassArray) Last() (v AuthAuthorizationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthAuthorizationClassArray) PopFirst() (v AuthAuthorizationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthAuthorizationClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthAuthorizationClassArray) Pop() (v AuthAuthorizationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAuthAuthorization returns copy with only AuthAuthorization constructors.
func (s AuthAuthorizationClassArray) AsAuthAuthorization() (to AuthAuthorizationArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthAuthorization)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAuthAuthorizationSignUpRequired returns copy with only AuthAuthorizationSignUpRequired constructors.
func (s AuthAuthorizationClassArray) AsAuthAuthorizationSignUpRequired() (to AuthAuthorizationSignUpRequiredArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthAuthorizationSignUpRequired)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AuthAuthorizationArray is adapter for slice of AuthAuthorization.
type AuthAuthorizationArray []AuthAuthorization

// Sort sorts slice of AuthAuthorization.
func (s AuthAuthorizationArray) Sort(less func(a, b AuthAuthorization) bool) AuthAuthorizationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthAuthorization.
func (s AuthAuthorizationArray) SortStable(less func(a, b AuthAuthorization) bool) AuthAuthorizationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthAuthorization.
func (s AuthAuthorizationArray) Retain(keep func(x AuthAuthorization) bool) AuthAuthorizationArray {
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
func (s AuthAuthorizationArray) First() (v AuthAuthorization, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthAuthorizationArray) Last() (v AuthAuthorization, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthAuthorizationArray) PopFirst() (v AuthAuthorization, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthAuthorization
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthAuthorizationArray) Pop() (v AuthAuthorization, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AuthAuthorizationSignUpRequiredArray is adapter for slice of AuthAuthorizationSignUpRequired.
type AuthAuthorizationSignUpRequiredArray []AuthAuthorizationSignUpRequired

// Sort sorts slice of AuthAuthorizationSignUpRequired.
func (s AuthAuthorizationSignUpRequiredArray) Sort(less func(a, b AuthAuthorizationSignUpRequired) bool) AuthAuthorizationSignUpRequiredArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthAuthorizationSignUpRequired.
func (s AuthAuthorizationSignUpRequiredArray) SortStable(less func(a, b AuthAuthorizationSignUpRequired) bool) AuthAuthorizationSignUpRequiredArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthAuthorizationSignUpRequired.
func (s AuthAuthorizationSignUpRequiredArray) Retain(keep func(x AuthAuthorizationSignUpRequired) bool) AuthAuthorizationSignUpRequiredArray {
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
func (s AuthAuthorizationSignUpRequiredArray) First() (v AuthAuthorizationSignUpRequired, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthAuthorizationSignUpRequiredArray) Last() (v AuthAuthorizationSignUpRequired, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthAuthorizationSignUpRequiredArray) PopFirst() (v AuthAuthorizationSignUpRequired, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthAuthorizationSignUpRequired
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthAuthorizationSignUpRequiredArray) Pop() (v AuthAuthorizationSignUpRequired, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
