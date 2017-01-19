package ds

import (
	"strconv"

	"cloud.google.com/go/datastore"
)

func parseID(id string) int64 {
	r, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0
	}
	return r
}

func buildIDKeys(kind string, ids []string) []*datastore.Key {
	keys := make([]*datastore.Key, len(ids))
	for i, id := range ids {
		keys[i] = datastore.IDKey(kind, parseID(id), nil)
	}
	return keys
}

func buildNameKeys(kind string, names []string) []*datastore.Key {
	keys := make([]*datastore.Key, len(names))
	for i, name := range names {
		keys[i] = datastore.NameKey(kind, name, nil)
	}
	return keys
}
