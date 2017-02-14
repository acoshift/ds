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
	if x.GetKey() == nil && kind != "" {
		x.NewKey(kind)
	}
}

// SaveModel saves model to datastore
// kind is optional, if key already set
// if key was not set in model, will call NewKey with given kind
// if key is incomplete key and AllocateIncompleteID is true, will call allocate ids before put model to datastore
func (client *Client) SaveModel(ctx context.Context, kind string, src interface{}) error {
	beforeSave(kind, src)

	x := src.(KeyGetSetter)
	if key := x.GetKey(); key.Incomplete() && client.AllocateIncompleteID {
		keys, err := client.AllocateIDs(ctx, []*datastore.Key{key})
		if err != nil {
			return err
		}
		x.SetKey(keys[0])
	}
	key, err := client.Put(ctx, x.GetKey(), x)
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
	for i := 0; i < xs.Len(); i++ {
		x := xs.Index(i).Interface()
		beforeSave(kind, x)
	}
	err := client.PutModels(ctx, src)
	if err != nil {
		return err
	}
	return nil
}

// AllocateModel calls AllocateIDModel and SaveModel
func (client *Client) AllocateModel(ctx context.Context, kind string, src interface{}) error {
	err := client.AllocateIDModel(ctx, kind, src)
	if err != nil {
		return err
	}
	err = client.SaveModel(ctx, kind, src)
	if err != nil {
		return err
	}
	return nil
}

// AllocateModels calls AllocateIDModels and SaveModels
func (client *Client) AllocateModels(ctx context.Context, kind string, src interface{}) error {
	err := client.AllocateIDModels(ctx, kind, src)
	if err != nil {
		return err
	}
	err = client.SaveModels(ctx, kind, src)
	if err != nil {
		return err
	}
	return nil
}
