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

// RecentMeURLClassArray is adapter for slice of RecentMeURLClass.
type RecentMeURLClassArray []RecentMeURLClass

// Sort sorts slice of RecentMeURLClass.
func (s RecentMeURLClassArray) Sort(less func(a, b RecentMeURLClass) bool) RecentMeURLClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLClass.
func (s RecentMeURLClassArray) SortStable(less func(a, b RecentMeURLClass) bool) RecentMeURLClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLClass.
func (s RecentMeURLClassArray) Retain(keep func(x RecentMeURLClass) bool) RecentMeURLClassArray {
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
func (s RecentMeURLClassArray) First() (v RecentMeURLClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLClassArray) Last() (v RecentMeURLClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLClassArray) PopFirst() (v RecentMeURLClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLClassArray) Pop() (v RecentMeURLClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsRecentMeURLUnknown returns copy with only RecentMeURLUnknown constructors.
func (s RecentMeURLClassArray) AsRecentMeURLUnknown() (to RecentMeURLUnknownArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLUnknown)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeURLUser returns copy with only RecentMeURLUser constructors.
func (s RecentMeURLClassArray) AsRecentMeURLUser() (to RecentMeURLUserArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLUser)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeURLChat returns copy with only RecentMeURLChat constructors.
func (s RecentMeURLClassArray) AsRecentMeURLChat() (to RecentMeURLChatArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLChat)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeURLChatInvite returns copy with only RecentMeURLChatInvite constructors.
func (s RecentMeURLClassArray) AsRecentMeURLChatInvite() (to RecentMeURLChatInviteArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLChatInvite)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeURLStickerSet returns copy with only RecentMeURLStickerSet constructors.
func (s RecentMeURLClassArray) AsRecentMeURLStickerSet() (to RecentMeURLStickerSetArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLStickerSet)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// RecentMeURLUnknownArray is adapter for slice of RecentMeURLUnknown.
type RecentMeURLUnknownArray []RecentMeURLUnknown

// Sort sorts slice of RecentMeURLUnknown.
func (s RecentMeURLUnknownArray) Sort(less func(a, b RecentMeURLUnknown) bool) RecentMeURLUnknownArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLUnknown.
func (s RecentMeURLUnknownArray) SortStable(less func(a, b RecentMeURLUnknown) bool) RecentMeURLUnknownArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLUnknown.
func (s RecentMeURLUnknownArray) Retain(keep func(x RecentMeURLUnknown) bool) RecentMeURLUnknownArray {
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
func (s RecentMeURLUnknownArray) First() (v RecentMeURLUnknown, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLUnknownArray) Last() (v RecentMeURLUnknown, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLUnknownArray) PopFirst() (v RecentMeURLUnknown, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLUnknown
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLUnknownArray) Pop() (v RecentMeURLUnknown, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeURLUserArray is adapter for slice of RecentMeURLUser.
type RecentMeURLUserArray []RecentMeURLUser

// Sort sorts slice of RecentMeURLUser.
func (s RecentMeURLUserArray) Sort(less func(a, b RecentMeURLUser) bool) RecentMeURLUserArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLUser.
func (s RecentMeURLUserArray) SortStable(less func(a, b RecentMeURLUser) bool) RecentMeURLUserArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLUser.
func (s RecentMeURLUserArray) Retain(keep func(x RecentMeURLUser) bool) RecentMeURLUserArray {
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
func (s RecentMeURLUserArray) First() (v RecentMeURLUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLUserArray) Last() (v RecentMeURLUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLUserArray) PopFirst() (v RecentMeURLUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLUser
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLUserArray) Pop() (v RecentMeURLUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeURLChatArray is adapter for slice of RecentMeURLChat.
type RecentMeURLChatArray []RecentMeURLChat

// Sort sorts slice of RecentMeURLChat.
func (s RecentMeURLChatArray) Sort(less func(a, b RecentMeURLChat) bool) RecentMeURLChatArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLChat.
func (s RecentMeURLChatArray) SortStable(less func(a, b RecentMeURLChat) bool) RecentMeURLChatArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLChat.
func (s RecentMeURLChatArray) Retain(keep func(x RecentMeURLChat) bool) RecentMeURLChatArray {
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
func (s RecentMeURLChatArray) First() (v RecentMeURLChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLChatArray) Last() (v RecentMeURLChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLChatArray) PopFirst() (v RecentMeURLChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLChat
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLChatArray) Pop() (v RecentMeURLChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeURLChatInviteArray is adapter for slice of RecentMeURLChatInvite.
type RecentMeURLChatInviteArray []RecentMeURLChatInvite

// Sort sorts slice of RecentMeURLChatInvite.
func (s RecentMeURLChatInviteArray) Sort(less func(a, b RecentMeURLChatInvite) bool) RecentMeURLChatInviteArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLChatInvite.
func (s RecentMeURLChatInviteArray) SortStable(less func(a, b RecentMeURLChatInvite) bool) RecentMeURLChatInviteArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLChatInvite.
func (s RecentMeURLChatInviteArray) Retain(keep func(x RecentMeURLChatInvite) bool) RecentMeURLChatInviteArray {
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
func (s RecentMeURLChatInviteArray) First() (v RecentMeURLChatInvite, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLChatInviteArray) Last() (v RecentMeURLChatInvite, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLChatInviteArray) PopFirst() (v RecentMeURLChatInvite, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLChatInvite
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLChatInviteArray) Pop() (v RecentMeURLChatInvite, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeURLStickerSetArray is adapter for slice of RecentMeURLStickerSet.
type RecentMeURLStickerSetArray []RecentMeURLStickerSet

// Sort sorts slice of RecentMeURLStickerSet.
func (s RecentMeURLStickerSetArray) Sort(less func(a, b RecentMeURLStickerSet) bool) RecentMeURLStickerSetArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLStickerSet.
func (s RecentMeURLStickerSetArray) SortStable(less func(a, b RecentMeURLStickerSet) bool) RecentMeURLStickerSetArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLStickerSet.
func (s RecentMeURLStickerSetArray) Retain(keep func(x RecentMeURLStickerSet) bool) RecentMeURLStickerSetArray {
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
func (s RecentMeURLStickerSetArray) First() (v RecentMeURLStickerSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLStickerSetArray) Last() (v RecentMeURLStickerSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLStickerSetArray) PopFirst() (v RecentMeURLStickerSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLStickerSet
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLStickerSetArray) Pop() (v RecentMeURLStickerSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
