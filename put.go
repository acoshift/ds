package ds

import (
	"context"

	"cloud.google.com/go/datastore"
)

// PutModel puts a model to datastore
func (client *Client) PutModel(ctx context.Context, src interface{}) error {
	key := src.(KeyGetSetter).GetKey()
	key, err := client.Put(ctx, key, src)
	SetKey(key, src)
	if client.Cache != nil {
		client.Cache.Del(key)
	}
	if err != nil {
		return err
	}
	return nil
}

// PutModels puts models to datastore
func (client *Client) PutModels(ctx context.Context, src interface{}) error {
	xs := valueOf(src)
	keys := make([]*datastore.Key, xs.Len())
	for i := range keys {
		x := xs.Index(i).Interface()
		keys[i] = x.(KeyGetter).GetKey()
	}
	keys, err := client.PutMulti(ctx, keys, src)
	SetKeys(keys, src)
	if client.Cache != nil {
		client.Cache.DelMulti(keys)
	}
	if err != nil {
		return err
	}
	return nil
}
