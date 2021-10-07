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

// PasswordKdfAlgoClassArray is adapter for slice of PasswordKdfAlgoClass.
type PasswordKdfAlgoClassArray []PasswordKdfAlgoClass

// Sort sorts slice of PasswordKdfAlgoClass.
func (s PasswordKdfAlgoClassArray) Sort(less func(a, b PasswordKdfAlgoClass) bool) PasswordKdfAlgoClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PasswordKdfAlgoClass.
func (s PasswordKdfAlgoClassArray) SortStable(less func(a, b PasswordKdfAlgoClass) bool) PasswordKdfAlgoClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PasswordKdfAlgoClass.
func (s PasswordKdfAlgoClassArray) Retain(keep func(x PasswordKdfAlgoClass) bool) PasswordKdfAlgoClassArray {
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
func (s PasswordKdfAlgoClassArray) First() (v PasswordKdfAlgoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PasswordKdfAlgoClassArray) Last() (v PasswordKdfAlgoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PasswordKdfAlgoClassArray) PopFirst() (v PasswordKdfAlgoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PasswordKdfAlgoClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PasswordKdfAlgoClassArray) Pop() (v PasswordKdfAlgoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow returns copy with only PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow constructors.
func (s PasswordKdfAlgoClassArray) AsPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow() (to PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray) {
	for _, elem := range s {
		value, ok := elem.(*PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray is adapter for slice of PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow.
type PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray []PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow

// Sort sorts slice of PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow.
func (s PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray) Sort(less func(a, b PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) bool) PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow.
func (s PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray) SortStable(less func(a, b PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) bool) PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow.
func (s PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray) Retain(keep func(x PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) bool) PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray {
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
func (s PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray) First() (v PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray) Last() (v PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray) PopFirst() (v PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowArray) Pop() (v PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
