package ds

import (
	"context"

	"cloud.google.com/go/datastore"
)

// DeleteByKey deletes data from datastore by Key
func (c *Client) DeleteByKey(ctx context.Context, key *datastore.Key) error {
	return c.c.Delete(ctx, key)
}

// DeleteByKeys deletes data from datastore by Keys
func (c *Client) DeleteByKeys(ctx context.Context, keys []*datastore.Key) error {
	return c.c.DeleteMulti(ctx, keys)
}

// DeleteByID deletes data from datastore by IDKey
func (c *Client) DeleteByID(ctx context.Context, kind string, id int64) error {
	if id == 0 {
		return datastore.ErrInvalidKey
	}
	return c.c.Delete(ctx, datastore.IDKey(kind, id, nil))
}

// DeleteByStringID deletes data from datastore by IDKey
func (c *Client) DeleteByStringID(ctx context.Context, kind string, id string) error {
	tid := parseID(id)
	if tid == 0 {
		return datastore.ErrInvalidKey
	}
	return c.c.Delete(ctx, datastore.IDKey(kind, tid, nil))
}

// DeleteByIDs deletes data from datastore by IDKeys
func (c *Client) DeleteByIDs(ctx context.Context, kind string, ids []int64) error {
	keys := make([]*datastore.Key, len(ids))
	for i, id := range ids {
		if id == 0 {
			return datastore.ErrInvalidKey
		}
		keys[i] = datastore.IDKey(kind, id, nil)
	}
	return c.c.DeleteMulti(ctx, keys)
}

// DeleteByStringIDs deletes data from datastore by IDKeys
func (c *Client) DeleteByStringIDs(ctx context.Context, kind string, ids []string) error {
	k := kind
	keys := make([]*datastore.Key, len(ids))
	for i, id := range ids {
		tid := parseID(id)
		if tid == 0 {
			return datastore.ErrInvalidKey
		}
		keys[i] = datastore.IDKey(k, tid, nil)
	}
	return c.c.DeleteMulti(ctx, keys)
}

// DeleteByName deletes data from datastore by NameKey
func (c *Client) DeleteByName(ctx context.Context, kind string, name string) error {
	if len(name) == 0 {
		return datastore.ErrInvalidKey
	}
	return c.c.Delete(ctx, datastore.NameKey(kind, name, nil))
}

// DeleteByNames deletes data from datastore by NameKeys
func (c *Client) DeleteByNames(ctx context.Context, kind string, names []string) error {
	keys := make([]*datastore.Key, len(names))
	for i, name := range names {
		if len(name) == 0 {
			return datastore.ErrInvalidKey
		}
		keys[i] = datastore.NameKey(kind, name, nil)
	}
	return c.c.DeleteMulti(ctx, keys)
}

// DeleteModel deletes data by get key from model
func (c *Client) DeleteModel(ctx context.Context, src interface{}) error {
	key := ExtractKey(src)
	if key == nil {
		return datastore.ErrInvalidKey
	}
	return c.c.Delete(ctx, key)
}

// DeleteModels deletes data by get keys from models
func (c *Client) DeleteModels(ctx context.Context, src interface{}) error {
	keys := ExtractKeys(src)
	if len(keys) == 0 {
		return datastore.ErrInvalidKey
	}
	return c.c.DeleteMulti(ctx, keys)
}
