package ds

import (
	"context"

	"cloud.google.com/go/datastore"
)

// PutModel puts a model to datastore
func (client *Client) PutModel(ctx context.Context, src interface{}) error {
	x := src.(KeyGetSetter)
	_, err := client.Put(ctx, x.GetKey(), x)
	return err
}

// PutModels puts models to datastore
func (client *Client) PutModels(ctx context.Context, src interface{}) error {
	xs := valueOf(src)
	keys := make([]*datastore.Key, xs.Len())
	for i := range keys {
		x := xs.Index(i).Interface()
		keys[i] = x.(KeyGetter).GetKey()
	}
	_, err := client.PutMulti(ctx, keys, src)
	return err
}
