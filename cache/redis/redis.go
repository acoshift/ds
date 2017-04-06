package redis

import (
	"bytes"
	"encoding/gob"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/acoshift/ds"
	"github.com/garyburd/redigo/redis"
)

// Cache implement Cache interface
type Cache struct {
	Pool   *redis.Pool
	Prefix string
	TTL    time.Duration
}

func encode(v interface{}) ([]byte, error) {
	w := &bytes.Buffer{}
	err := gob.NewEncoder(w).Encode(v)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func decode(b []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(b)).Decode(v)
}

// Get gets data
func (cache *Cache) Get(key *datastore.Key, dst interface{}) error {
	db := cache.Pool.Get()
	defer db.Close()
	b, err := redis.Bytes(db.Do("GET", cache.Prefix+key.String()))
	if err != nil {
		return err
	}
	if len(b) == 0 {
		return ds.ErrCacheNotFound
	}
	return decode(b, dst)
}

// GetMulti gets multi data
func (cache *Cache) GetMulti(keys []*datastore.Key, dst []interface{}) error {
	return nil
}

// Set sets data
func (cache *Cache) Set(key *datastore.Key, src interface{}) error {
	db := cache.Pool.Get()
	defer db.Close()
	b, err := encode(src)
	if err != nil {
		return err
	}
	if cache.TTL > 0 {
		_, err = db.Do("SETEX", cache.Prefix+key.String(), int(cache.TTL/time.Second), b)
		return err
	}
	_, err = db.Do("SET", cache.Prefix+key.String(), b)
	return err
}

// SetMulti sets data
func (cache *Cache) SetMulti(keys []*datastore.Key, src interface{}) error {
	return nil
}

// Del dels data
func (cache *Cache) Del(key *datastore.Key) error {
	db := cache.Pool.Get()
	defer db.Close()
	_, err := db.Do("DEL", cache.Prefix+key.String())
	return err
}

// DelMulti dels multi data
func (cache *Cache) DelMulti(keys []*datastore.Key) error {
	db := cache.Pool.Get()
	defer db.Close()
	for _, key := range keys {
		db.Send("DEL", cache.Prefix+key.String())
	}
	return db.Flush()
}
