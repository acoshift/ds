package ds

import (
	"context"

	"cloud.google.com/go/datastore"
)

// PutModel puts a model to datastore
func (c *Client) PutModel(ctx context.Context, src interface{}) error {
	x := src.(KeyGetSetter)
	_, err := c.c.Put(ctx, x.GetKey(), x)
	return err
}

// PutModels puts models to datastore
func (c *Client) PutModels(ctx context.Context, src interface{}) error {
	xs := valueOf(src)
	keys := make([]*datastore.Key, xs.Len())
	for i := range keys {
		x := xs.Index(i).Interface()
		keys[i] = x.(KeyGetter).GetKey()
	}
	keys, err := c.c.PutMulti(ctx, keys, src)
	SetKeys(keys, src)
	return err
}
