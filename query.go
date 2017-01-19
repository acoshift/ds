package ds

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
)

// Query is the function for set datastore query
type Query func(q *datastore.Query) *datastore.Query

// Query run Get All
// dst is *[]*Model
func (client *Client) Query(ctx context.Context, kind KindGetter, dst interface{}, qs ...Query) error {
	q := datastore.NewQuery(kind.Kind())
	for _, setter := range qs {
		q = setter(q)
	}

	keys, err := client.GetAll(ctx, q, dst)
	if err != nil {
		return err
	}
	setKeys(keys, dst)
	return nil
}

// QueryFirst run Get to get the first result
func (client *Client) QueryFirst(ctx context.Context, dst interface{}, qs ...Query) error {
	q := datastore.NewQuery(dst.(KindGetter).Kind())
	for _, setter := range qs {
		q = setter(q)
	}

	key, err := client.Run(ctx, q).Next(dst)
	if err != nil {
		return err
	}

	if x, ok := dst.(KeySetter); ok {
		x.SetKey(key)
	}
	return nil
}

// QueryKeys queries only key
func (client *Client) QueryKeys(ctx context.Context, kind KindGetter, qs ...Query) ([]*datastore.Key, error) {
	q := datastore.NewQuery(kind.Kind())
	for _, setter := range qs {
		q = setter(q)
	}
	q = q.KeysOnly()

	keys, err := client.GetAll(ctx, q, nil)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

// Query Helper functions

// Filter func
func Filter(filterStr string, value interface{}) Query {
	return func(q *datastore.Query) *datastore.Query {
		return q.Filter(filterStr, value)
	}
}

// CreateAfter quries is model created after (or equals) given time
func CreateAfter(t time.Time, equals bool) Query {
	p := "CreatedAt >"
	if equals {
		p += "="
	}
	return Filter(p, t)
}

// UpdateBefore queries is model updated before (or equals) given time
func UpdateBefore(t time.Time, equals bool) Query {
	p := "UpdatedAt <"
	if equals {
		p += "="
	}
	return Filter(p, t)
}

// UpdateAfter queries is model updated after (or equals) given time
func UpdateAfter(t time.Time, equals bool) Query {
	p := "UpdatedAt >"
	if equals {
		p += "="
	}
	return Filter(p, t)
}

// Offset adds offset to query
func Offset(offset int) Query {
	return func(q *datastore.Query) *datastore.Query {
		return q.Offset(offset)
	}
}

// Limit adds limit to query
func Limit(limit int) Query {
	return func(q *datastore.Query) *datastore.Query {
		return q.Limit(limit)
	}
}

// Namespace adds namespace to query
func Namespace(ns string) Query {
	return func(q *datastore.Query) *datastore.Query {
		return q.Namespace(ns)
	}
}

// Order adds order to query
func Order(fieldName string) Query {
	return func(q *datastore.Query) *datastore.Query {
		return q.Order(fieldName)
	}
}

// Project adds order to query
func Project(fieldNames ...string) Query {
	return func(q *datastore.Query) *datastore.Query {
		return q.Project(fieldNames...)
	}
}
