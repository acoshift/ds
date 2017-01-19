package ds

import (
	"context"
	"reflect"

	"cloud.google.com/go/datastore"
)

func beforeSave(src interface{}) {
	x := src.(KeyGetSetter)

	// stamp model
	if x, ok := src.(Stampable); ok {
		x.Stamp()
	}

	// create new key
	if x.Key() == nil {
		if k, ok := src.(KindGetter); ok {
			x.SetKey(datastore.IncompleteKey(k.Kind(), nil))
		}
	}
}

// Save saves model to datastore
func (client *Client) Save(ctx context.Context, src interface{}) error {
	beforeSave(src)

	x := src.(KeyGetSetter)
	key, err := client.Put(ctx, x.Key(), x)
	if err != nil {
		return err
	}
	x.SetKey(key)
	return nil
}

// SaveMulti saves models to datastore
func (client *Client) SaveMulti(ctx context.Context, src interface{}) error {
	xs := reflect.ValueOf(src)
	keys := make([]*datastore.Key, xs.Len())
	for i := range keys {
		x := xs.Index(i).Interface()
		beforeSave(x)
		keys[i] = x.(KeyGetter).Key()
	}
	keys, err := client.PutMulti(ctx, keys, src)
	if err != nil {
		return err
	}
	SetKeys(keys, src)
	return nil
}
