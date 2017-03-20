package ds

import (
	"context"
	"reflect"

	"cloud.google.com/go/datastore"
)

// GetByKey retrieves model from datastore by key
func (c *Client) GetByKey(ctx context.Context, key *datastore.Key, dst interface{}) error {
	err := c.c.Get(ctx, key, dst)
	SetKey(key, dst)
	if err != nil {
		return err
	}
	return nil
}

// GetByKeys retrieves models from datastore by keys
func (c *Client) GetByKeys(ctx context.Context, keys []*datastore.Key, dst interface{}) error {
	// prepare slice if dst is pointer to 0 len slice
	if rf := reflect.ValueOf(dst); rf.Kind() == reflect.Ptr {
		rs := rf.Elem()
		if rs.Kind() == reflect.Slice && rs.Len() == 0 {
			l := len(keys)
			rs.Set(reflect.MakeSlice(rs.Type(), l, l))
		}
		dst = rs.Interface()
	}

	err := c.c.GetMulti(ctx, keys, dst)
	SetKeys(keys, dst)
	if err != nil {
		return err
	}
	return nil
}

// GetByModel retrieves model from datastore by key from model
func (c *Client) GetByModel(ctx context.Context, dst interface{}) error {
	key := ExtractKey(dst)
	return c.GetByKey(ctx, key, dst)
}

// GetByModels retrieves models from datastore by keys from models
func (c *Client) GetByModels(ctx context.Context, dst interface{}) error {
	keys := ExtractKeys(dst)
	return c.GetByKeys(ctx, keys, dst)
}

// GetByID retrieves model from datastore by id
func (c *Client) GetByID(ctx context.Context, kind string, id int64, dst interface{}) error {
	return c.GetByKey(ctx, datastore.IDKey(kind, id, nil), dst)
}

// GetByIDs retrieves models from datastore by ids
func (c *Client) GetByIDs(ctx context.Context, kind string, ids []int64, dst interface{}) error {
	keys := BuildIDKeys(kind, ids)
	return c.GetByKeys(ctx, keys, dst)
}

// GetByStringID retrieves model from datastore by string id
func (c *Client) GetByStringID(ctx context.Context, kind string, id string, dst interface{}) error {
	tid := parseID(id)
	if tid == 0 {
		return datastore.ErrInvalidKey
	}
	return c.GetByKey(ctx, datastore.IDKey(kind, tid, nil), dst)
}

// GetByStringIDs retrieves models from datastore by string ids
func (c *Client) GetByStringIDs(ctx context.Context, kind string, ids []string, dst interface{}) error {
	keys := BuildStringIDKeys(kind, ids)
	return c.GetByKeys(ctx, keys, dst)
}

// GetByName retrieves model from datastore by name
func (c *Client) GetByName(ctx context.Context, kind string, name string, dst interface{}) error {
	return c.GetByKey(ctx, datastore.NameKey(kind, name, nil), dst)
}

// GetByNames retrieves models from datastore by names
func (c *Client) GetByNames(ctx context.Context, kind string, names []string, dst interface{}) error {
	keys := BuildNameKeys(kind, names)
	return c.GetByKeys(ctx, keys, dst)
}

// GetByQuery retrieves model from datastore by datastore query
func (c *Client) GetByQuery(ctx context.Context, q *datastore.Query, dst interface{}) error {
	_, err := c.c.GetAll(ctx, q, dst)
	if err != nil {
		return err
	}
	return nil
}
