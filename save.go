package ds

import (
	"context"

	"cloud.google.com/go/datastore"
)

func beforeSave(kind string, src interface{}) {
	x := src.(KeyGetSetter)

	// stamp model
	if x, ok := src.(Stampable); ok {
		x.Stamp()
	}

	// create new key
	if x.Key() == nil && kind != "" {
		x.NewKey(kind)
	}
}

// SaveModel saves model to datastore
// kind is optional, if key already set
// if key was not set in model, will call NewKey with given kind
func (client *Client) SaveModel(ctx context.Context, kind string, src interface{}) error {
	beforeSave(kind, src)

	x := src.(KeyGetSetter)
	key, err := client.Put(ctx, x.Key(), x)
	x.SetKey(key)
	if err != nil {
		return err
	}
	return nil
}

// SaveModels saves models to datastore
// see more in SaveModel
func (client *Client) SaveModels(ctx context.Context, kind string, src interface{}) error {
	xs := valueOf(src)
	keys := make([]*datastore.Key, xs.Len())
	for i := range keys {
		x := xs.Index(i).Interface()
		beforeSave(kind, x)
		keys[i] = x.(KeyGetter).Key()
	}
	keys, err := client.PutMulti(ctx, keys, src)
	SetKeys(keys, src)
	if err != nil {
		return err
	}
	return nil
}
