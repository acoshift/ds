package ds

import (
	"context"

	"cloud.google.com/go/datastore"
)

// GetByKey retrieves model from datastore by key
func (client *Client) GetByKey(ctx context.Context, key *datastore.Key, dst interface{}) error {
	err := client.Get(ctx, key, dst)
	SetKey(key, dst)
	if err != nil {
		return err
	}
	return nil
}

// GetByKeys retrieves models from datastore by keys
func (client *Client) GetByKeys(ctx context.Context, keys []*datastore.Key, dst interface{}) error {
	err := client.GetMulti(ctx, keys, dst)
	SetKeys(keys, dst)
	if err != nil {
		return err
	}
	return nil
}

// GetByModel retrieves model from datastore by key from model
func (client *Client) GetByModel(ctx context.Context, dst interface{}) error {
	key := ExtractKey(dst)
	return client.GetByKey(ctx, key, dst)
}

// GetByModels retrieves models from datastore by keys from models
func (client *Client) GetByModels(ctx context.Context, dst interface{}) error {
	keys := ExtractKeys(dst)
	return client.GetByKeys(ctx, keys, dst)
}

// GetByID retrieves model from datastore by id
func (client *Client) GetByID(ctx context.Context, kind string, id int64, dst interface{}) error {
	return client.GetByKey(ctx, datastore.IDKey(kind, id, nil), dst)
}

// GetByIDs retrieves models from datastore by ids
func (client *Client) GetByIDs(ctx context.Context, kind string, ids []int64, dst interface{}) error {
	keys := BuildIDKeys(kind, ids)
	return client.GetByKeys(ctx, keys, dst)
}

// GetByStringID retrieves model from datastore by string id
func (client *Client) GetByStringID(ctx context.Context, kind string, id string, dst interface{}) error {
	tid := parseID(id)
	if tid == 0 {
		return datastore.ErrInvalidKey
	}
	return client.GetByKey(ctx, datastore.IDKey(kind, tid, nil), dst)
}

// GetByStringIDs retrieves models from datastore by string ids
func (client *Client) GetByStringIDs(ctx context.Context, kind string, ids []string, dst interface{}) error {
	keys := BuildStringIDKeys(kind, ids)
	return client.GetByKeys(ctx, keys, dst)
}

// GetByName retrieves model from datastore by name
func (client *Client) GetByName(ctx context.Context, kind string, name string, dst interface{}) error {
	return client.GetByKey(ctx, datastore.NameKey(kind, name, nil), dst)
}

// GetByNames retrieves models from datastore by names
func (client *Client) GetByNames(ctx context.Context, kind string, names []string, dst interface{}) error {
	keys := BuildNameKeys(kind, names)
	return client.GetByKeys(ctx, keys, dst)
}

// GetByQuery retrieves model from datastore by datastore query
func (client *Client) GetByQuery(ctx context.Context, q *datastore.Query, dst interface{}) error {
	keys, err := client.GetAll(ctx, q, dst)
	SetKeys(keys, dst)
	if err != nil {
		return err
	}
	return nil
}
