package ds

import (
	"context"

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
func (client *Client) QueryFirst(ctx context.Context, kind KindGetter, dst interface{}, qs ...Query) error {
	q := datastore.NewQuery(kind.Kind())
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
