package ds

import (
	"context"

	"cloud.google.com/go/datastore"
)

// DeleteByID deletes data from datastore by IDKey
func (client *Client) DeleteByID(ctx context.Context, kind KindGetter, id string) error {
	tid := parseID(id)
	if tid == 0 {
		return ErrInvalidID
	}
	return client.Delete(ctx, datastore.IDKey(kind.Kind(), parseID(id), nil))
}

// DeleteByIDs deletes data from datastore by IDKeys
func (client *Client) DeleteByIDs(ctx context.Context, kind KindGetter, ids []string) error {
	k := kind.Kind()
	keys := make([]*datastore.Key, len(ids))
	for i, id := range ids {
		tid := parseID(id)
		if tid == 0 {
			return ErrInvalidID
		}
		keys[i] = datastore.IDKey(k, tid, nil)
	}
	return client.DeleteMulti(ctx, keys)
}

// DeleteByName deletes data from datastore by NameKey
func (client *Client) DeleteByName(ctx context.Context, kind KindGetter, name string) error {
	if name == "" {
		return ErrInvalidID
	}
	return client.Delete(ctx, datastore.NameKey(kind.Kind(), name, nil))
}

// DeleteByNames deletes data from datastore by NameKeys
func (client *Client) DeleteByNames(ctx context.Context, kind KindGetter, names []string) error {
	k := kind.Kind()
	keys := make([]*datastore.Key, len(names))
	for i, name := range names {
		if name == "" {
			return ErrInvalidID
		}
		keys[i] = datastore.NameKey(k, name, nil)
	}
	return client.DeleteMulti(ctx, keys)
}
