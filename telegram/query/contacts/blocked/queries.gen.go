// Code generated by itergen, DO NOT EDIT.

package blocked

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/nnqq/td/tg"
)

// No-op definition for keeping imports.
var _ = context.Background()

// Request is a parameter for Query.
type Request struct {
	Offset int
	Limit  int
}

// Query is an abstraction for blocked request.
// NB: iterator mutates returned data (sorts, at least).
type Query interface {
	Query(ctx context.Context, req Request) (tg.ContactsBlockedClass, error)
}

// QueryFunc is a function adapter for Query.
type QueryFunc func(ctx context.Context, req Request) (tg.ContactsBlockedClass, error)

// Query implements Query interface.
func (q QueryFunc) Query(ctx context.Context, req Request) (tg.ContactsBlockedClass, error) {
	return q(ctx, req)
}

// QueryBuilder is a helper to create message queries.
type QueryBuilder struct {
	raw *tg.Client
}

// NewQueryBuilder creates new QueryBuilder.
func NewQueryBuilder(raw *tg.Client) *QueryBuilder {
	return &QueryBuilder{raw: raw}
}

// GetBlockedQueryBuilder is query builder of ContactsGetBlocked.
type GetBlockedQueryBuilder struct {
	raw       *tg.Client
	req       tg.ContactsGetBlockedRequest
	batchSize int
	offset    int
}

// GetBlocked creates query builder of ContactsGetBlocked.
func (q *QueryBuilder) GetBlocked() *GetBlockedQueryBuilder {
	b := &GetBlockedQueryBuilder{
		raw:       q.raw,
		batchSize: 1,
		req:       tg.ContactsGetBlockedRequest{},
	}

	return b
}

// BatchSize sets buffer of message loaded from one request.
// Be carefully, when set this limit, because Telegram does not return error if limit is too big,
// so results can be incorrect.
func (b *GetBlockedQueryBuilder) BatchSize(batchSize int) *GetBlockedQueryBuilder {
	b.batchSize = batchSize
	return b
}

// Query implements Query interface.
func (b *GetBlockedQueryBuilder) Query(ctx context.Context, req Request) (tg.ContactsBlockedClass, error) {
	r := &tg.ContactsGetBlockedRequest{
		Limit: req.Limit,
	}

	r.Offset = req.Offset
	return b.raw.ContactsGetBlocked(ctx, r)
}

// Iter returns iterator using built query.
func (b *GetBlockedQueryBuilder) Iter() *Iterator {
	iter := NewIterator(b, b.batchSize)
	iter = iter.Offset(b.offset)
	return iter
}

// ForEach calls given callback on each iterator element.
func (b *GetBlockedQueryBuilder) ForEach(ctx context.Context, cb func(context.Context, Elem) error) error {
	iter := b.Iter()
	for iter.Next(ctx) {
		if err := cb(ctx, iter.Value()); err != nil {
			return err
		}
	}
	return iter.Err()
}

// Count fetches remote state to get number of elements.
func (b *GetBlockedQueryBuilder) Count(ctx context.Context) (int, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return 0, xerrors.Errorf("get total: %w", err)
	}
	return c, nil
}

// Collect creates iterator and collects all elements to slice.
func (b *GetBlockedQueryBuilder) Collect(ctx context.Context) ([]Elem, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return nil, xerrors.Errorf("get total: %w", err)
	}

	r := make([]Elem, 0, c)
	for iter.Next(ctx) {
		r = append(r, iter.Value())
	}

	return r, iter.Err()
}
