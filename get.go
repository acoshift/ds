package ds

import (
	"context"
	"reflect"

	"cloud.google.com/go/datastore"
)

// SetKey sets key to model
func SetKey(key *datastore.Key, dst interface{}) {
	if x, ok := dst.(KeySetter); ok {
		x.SetKey(key)
	}
}

// SetKeys sets keys to models
func SetKeys(keys []*datastore.Key, dst interface{}) {
	xs := reflect.ValueOf(dst).Elem()
	for i := 0; i < xs.Len(); i++ {
		if x, ok := xs.Index(i).Interface().(KeySetter); ok {
			x.SetKey(keys[i])
		}
	}
}

// SetCommitKey sets commit pending key to model
func SetCommitKey(commit *datastore.Commit, pendingKey *datastore.PendingKey, dst interface{}) {
	if x, ok := dst.(KeySetter); ok {
		x.SetKey(commit.Key(pendingKey))
	}
}

// SetCommitKeys sets commit pending keys to models
func SetCommitKeys(commit *datastore.Commit, pendingKeys []*datastore.PendingKey, dst interface{}) {
	xs := reflect.ValueOf(dst)
	if xs.Kind() == reflect.Ptr {
		xs = xs.Elem()
	}
	for i := 0; i < xs.Len(); i++ {
		if x, ok := xs.Index(i).Interface().(KeySetter); ok {
			x.SetKey(commit.Key(pendingKeys[i]))
		}
	}
}

// GetByKey retrieves model from datastore by key
func (client *Client) GetByKey(ctx context.Context, key *datastore.Key, dst interface{}) error {
	err := client.Get(ctx, key, dst)
	if err != nil {
		return err
	}
	SetKey(key, dst)
	return nil
}

// GetByKeys retrieves models from datastore by keys
func (client *Client) GetByKeys(ctx context.Context, keys []*datastore.Key, dst interface{}) error {
	err := client.GetMulti(ctx, keys, dst)
	if err != nil {
		return err
	}
	SetKeys(keys, dst)
	return nil
}

// GetByID retrieves model from datastore by id
func (client *Client) GetByID(ctx context.Context, id string, dst interface{}) error {
	return client.GetByKey(ctx, datastore.IDKey(dst.(KindGetter).Kind(), parseID(id), nil), dst)
}

// GetByIDs retrieves models from datastore by ids
func (client *Client) GetByIDs(ctx context.Context, ids []string, kind KindGetter, dst interface{}) error {
	keys := buildIDKeys(kind.Kind(), ids)
	return client.GetByKeys(ctx, keys, dst)
}

// GetByName retrieves model from datastore by name
func (client *Client) GetByName(ctx context.Context, name string, dst interface{}) error {
	return client.GetByKey(ctx, datastore.NameKey(dst.(KindGetter).Kind(), name, nil), dst)
}

// GetByNames retrieves models from datastore by names
func (client *Client) GetByNames(ctx context.Context, names []string, kind KindGetter, dst interface{}) error {
	keys := buildNameKeys(kind.Kind(), names)
	return client.GetByKeys(ctx, keys, dst)
}

// GetByQuery retrieves model from datastore by datastore query
func (client *Client) GetByQuery(ctx context.Context, q *datastore.Query, dst interface{}) error {
	keys, err := client.GetAll(ctx, q, dst)
	if err != nil {
		return err
	}
	SetKeys(keys, dst)
	return nil
}
