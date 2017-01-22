package ds

import (
	"context"

	"cloud.google.com/go/datastore"
)

// DeleteByID deletes data from datastore by IDKey
func (client *Client) DeleteByID(ctx context.Context, kind string, id int64) error {
	if id == 0 {
		return datastore.ErrInvalidKey
	}
	return client.Delete(ctx, datastore.IDKey(kind, id, nil))
}

// DeleteByStringID deletes data from datastore by IDKey
func (client *Client) DeleteByStringID(ctx context.Context, kind string, id string) error {
	tid := parseID(id)
	if tid == 0 {
		return datastore.ErrInvalidKey
	}
	return client.Delete(ctx, datastore.IDKey(kind, tid, nil))
}

// DeleteByIDs deletes data from datastore by IDKeys
func (client *Client) DeleteByIDs(ctx context.Context, kind string, ids []int64) error {
	keys := make([]*datastore.Key, len(ids))
	for i, id := range ids {
		if id == 0 {
			return datastore.ErrInvalidKey
		}
		keys[i] = datastore.IDKey(kind, id, nil)
	}
	return client.DeleteMulti(ctx, keys)
}

// DeleteByStringIDs deletes data from datastore by IDKeys
func (client *Client) DeleteByStringIDs(ctx context.Context, kind string, ids []string) error {
	k := kind
	keys := make([]*datastore.Key, len(ids))
	for i, id := range ids {
		tid := parseID(id)
		if tid == 0 {
			return datastore.ErrInvalidKey
		}
		keys[i] = datastore.IDKey(k, tid, nil)
	}
	return client.DeleteMulti(ctx, keys)
}

// DeleteByName deletes data from datastore by NameKey
func (client *Client) DeleteByName(ctx context.Context, kind string, name string) error {
	if name == "" {
		return datastore.ErrInvalidKey
	}
	return client.Delete(ctx, datastore.NameKey(kind, name, nil))
}

// DeleteByNames deletes data from datastore by NameKeys
func (client *Client) DeleteByNames(ctx context.Context, kind string, names []string) error {
	keys := make([]*datastore.Key, len(names))
	for i, name := range names {
		if name == "" {
			return datastore.ErrInvalidKey
		}
		keys[i] = datastore.NameKey(kind, name, nil)
	}
	return client.DeleteMulti(ctx, keys)
}

// DeleteModel deletes data by get key from model
func (client *Client) DeleteModel(ctx context.Context, src interface{}) error {
	key := ExtractKey(src)
	if key == nil {
		return datastore.ErrInvalidKey
	}
	return client.Delete(ctx, key)
}

// DeleteModels deletes data by get keys from models
func (client *Client) DeleteModels(ctx context.Context, src interface{}) error {
	keys := ExtractKeys(src)
	if len(keys) == 0 {
		return datastore.ErrInvalidKey
	}
	return client.DeleteMulti(ctx, keys)
}
