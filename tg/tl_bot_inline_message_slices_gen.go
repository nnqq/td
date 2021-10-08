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

// BotInlineMessageClassArray is adapter for slice of BotInlineMessageClass.
type BotInlineMessageClassArray []BotInlineMessageClass

// Sort sorts slice of BotInlineMessageClass.
func (s BotInlineMessageClassArray) Sort(less func(a, b BotInlineMessageClass) bool) BotInlineMessageClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineMessageClass.
func (s BotInlineMessageClassArray) SortStable(less func(a, b BotInlineMessageClass) bool) BotInlineMessageClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineMessageClass.
func (s BotInlineMessageClassArray) Retain(keep func(x BotInlineMessageClass) bool) BotInlineMessageClassArray {
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
func (s BotInlineMessageClassArray) First() (v BotInlineMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineMessageClassArray) Last() (v BotInlineMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineMessageClassArray) PopFirst() (v BotInlineMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineMessageClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineMessageClassArray) Pop() (v BotInlineMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsBotInlineMessageMediaAuto returns copy with only BotInlineMessageMediaAuto constructors.
func (s BotInlineMessageClassArray) AsBotInlineMessageMediaAuto() (to BotInlineMessageMediaAutoArray) {
	for _, elem := range s {
		value, ok := elem.(*BotInlineMessageMediaAuto)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsBotInlineMessageText returns copy with only BotInlineMessageText constructors.
func (s BotInlineMessageClassArray) AsBotInlineMessageText() (to BotInlineMessageTextArray) {
	for _, elem := range s {
		value, ok := elem.(*BotInlineMessageText)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsBotInlineMessageMediaGeo returns copy with only BotInlineMessageMediaGeo constructors.
func (s BotInlineMessageClassArray) AsBotInlineMessageMediaGeo() (to BotInlineMessageMediaGeoArray) {
	for _, elem := range s {
		value, ok := elem.(*BotInlineMessageMediaGeo)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsBotInlineMessageMediaVenue returns copy with only BotInlineMessageMediaVenue constructors.
func (s BotInlineMessageClassArray) AsBotInlineMessageMediaVenue() (to BotInlineMessageMediaVenueArray) {
	for _, elem := range s {
		value, ok := elem.(*BotInlineMessageMediaVenue)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsBotInlineMessageMediaContact returns copy with only BotInlineMessageMediaContact constructors.
func (s BotInlineMessageClassArray) AsBotInlineMessageMediaContact() (to BotInlineMessageMediaContactArray) {
	for _, elem := range s {
		value, ok := elem.(*BotInlineMessageMediaContact)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsBotInlineMessageMediaInvoice returns copy with only BotInlineMessageMediaInvoice constructors.
func (s BotInlineMessageClassArray) AsBotInlineMessageMediaInvoice() (to BotInlineMessageMediaInvoiceArray) {
	for _, elem := range s {
		value, ok := elem.(*BotInlineMessageMediaInvoice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// BotInlineMessageMediaAutoArray is adapter for slice of BotInlineMessageMediaAuto.
type BotInlineMessageMediaAutoArray []BotInlineMessageMediaAuto

// Sort sorts slice of BotInlineMessageMediaAuto.
func (s BotInlineMessageMediaAutoArray) Sort(less func(a, b BotInlineMessageMediaAuto) bool) BotInlineMessageMediaAutoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineMessageMediaAuto.
func (s BotInlineMessageMediaAutoArray) SortStable(less func(a, b BotInlineMessageMediaAuto) bool) BotInlineMessageMediaAutoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineMessageMediaAuto.
func (s BotInlineMessageMediaAutoArray) Retain(keep func(x BotInlineMessageMediaAuto) bool) BotInlineMessageMediaAutoArray {
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
func (s BotInlineMessageMediaAutoArray) First() (v BotInlineMessageMediaAuto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineMessageMediaAutoArray) Last() (v BotInlineMessageMediaAuto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaAutoArray) PopFirst() (v BotInlineMessageMediaAuto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineMessageMediaAuto
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaAutoArray) Pop() (v BotInlineMessageMediaAuto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// BotInlineMessageTextArray is adapter for slice of BotInlineMessageText.
type BotInlineMessageTextArray []BotInlineMessageText

// Sort sorts slice of BotInlineMessageText.
func (s BotInlineMessageTextArray) Sort(less func(a, b BotInlineMessageText) bool) BotInlineMessageTextArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineMessageText.
func (s BotInlineMessageTextArray) SortStable(less func(a, b BotInlineMessageText) bool) BotInlineMessageTextArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineMessageText.
func (s BotInlineMessageTextArray) Retain(keep func(x BotInlineMessageText) bool) BotInlineMessageTextArray {
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
func (s BotInlineMessageTextArray) First() (v BotInlineMessageText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineMessageTextArray) Last() (v BotInlineMessageText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineMessageTextArray) PopFirst() (v BotInlineMessageText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineMessageText
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineMessageTextArray) Pop() (v BotInlineMessageText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// BotInlineMessageMediaGeoArray is adapter for slice of BotInlineMessageMediaGeo.
type BotInlineMessageMediaGeoArray []BotInlineMessageMediaGeo

// Sort sorts slice of BotInlineMessageMediaGeo.
func (s BotInlineMessageMediaGeoArray) Sort(less func(a, b BotInlineMessageMediaGeo) bool) BotInlineMessageMediaGeoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineMessageMediaGeo.
func (s BotInlineMessageMediaGeoArray) SortStable(less func(a, b BotInlineMessageMediaGeo) bool) BotInlineMessageMediaGeoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineMessageMediaGeo.
func (s BotInlineMessageMediaGeoArray) Retain(keep func(x BotInlineMessageMediaGeo) bool) BotInlineMessageMediaGeoArray {
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
func (s BotInlineMessageMediaGeoArray) First() (v BotInlineMessageMediaGeo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineMessageMediaGeoArray) Last() (v BotInlineMessageMediaGeo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaGeoArray) PopFirst() (v BotInlineMessageMediaGeo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineMessageMediaGeo
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaGeoArray) Pop() (v BotInlineMessageMediaGeo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// BotInlineMessageMediaVenueArray is adapter for slice of BotInlineMessageMediaVenue.
type BotInlineMessageMediaVenueArray []BotInlineMessageMediaVenue

// Sort sorts slice of BotInlineMessageMediaVenue.
func (s BotInlineMessageMediaVenueArray) Sort(less func(a, b BotInlineMessageMediaVenue) bool) BotInlineMessageMediaVenueArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineMessageMediaVenue.
func (s BotInlineMessageMediaVenueArray) SortStable(less func(a, b BotInlineMessageMediaVenue) bool) BotInlineMessageMediaVenueArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineMessageMediaVenue.
func (s BotInlineMessageMediaVenueArray) Retain(keep func(x BotInlineMessageMediaVenue) bool) BotInlineMessageMediaVenueArray {
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
func (s BotInlineMessageMediaVenueArray) First() (v BotInlineMessageMediaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineMessageMediaVenueArray) Last() (v BotInlineMessageMediaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaVenueArray) PopFirst() (v BotInlineMessageMediaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineMessageMediaVenue
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaVenueArray) Pop() (v BotInlineMessageMediaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// BotInlineMessageMediaContactArray is adapter for slice of BotInlineMessageMediaContact.
type BotInlineMessageMediaContactArray []BotInlineMessageMediaContact

// Sort sorts slice of BotInlineMessageMediaContact.
func (s BotInlineMessageMediaContactArray) Sort(less func(a, b BotInlineMessageMediaContact) bool) BotInlineMessageMediaContactArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineMessageMediaContact.
func (s BotInlineMessageMediaContactArray) SortStable(less func(a, b BotInlineMessageMediaContact) bool) BotInlineMessageMediaContactArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineMessageMediaContact.
func (s BotInlineMessageMediaContactArray) Retain(keep func(x BotInlineMessageMediaContact) bool) BotInlineMessageMediaContactArray {
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
func (s BotInlineMessageMediaContactArray) First() (v BotInlineMessageMediaContact, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineMessageMediaContactArray) Last() (v BotInlineMessageMediaContact, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaContactArray) PopFirst() (v BotInlineMessageMediaContact, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineMessageMediaContact
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaContactArray) Pop() (v BotInlineMessageMediaContact, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// BotInlineMessageMediaInvoiceArray is adapter for slice of BotInlineMessageMediaInvoice.
type BotInlineMessageMediaInvoiceArray []BotInlineMessageMediaInvoice

// Sort sorts slice of BotInlineMessageMediaInvoice.
func (s BotInlineMessageMediaInvoiceArray) Sort(less func(a, b BotInlineMessageMediaInvoice) bool) BotInlineMessageMediaInvoiceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineMessageMediaInvoice.
func (s BotInlineMessageMediaInvoiceArray) SortStable(less func(a, b BotInlineMessageMediaInvoice) bool) BotInlineMessageMediaInvoiceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineMessageMediaInvoice.
func (s BotInlineMessageMediaInvoiceArray) Retain(keep func(x BotInlineMessageMediaInvoice) bool) BotInlineMessageMediaInvoiceArray {
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
func (s BotInlineMessageMediaInvoiceArray) First() (v BotInlineMessageMediaInvoice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineMessageMediaInvoiceArray) Last() (v BotInlineMessageMediaInvoice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaInvoiceArray) PopFirst() (v BotInlineMessageMediaInvoice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineMessageMediaInvoice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineMessageMediaInvoiceArray) Pop() (v BotInlineMessageMediaInvoice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
