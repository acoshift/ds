package ds

import "cloud.google.com/go/datastore"
import "strconv"

// Model is the base model
// every model should have this Model
type Model struct {
	key *datastore.Key
	ID  string `datastore:"-"`
}

// Key returns key from model
func (x *Model) Key() *datastore.Key {
	return x.key
}

// SetKey sets model key to given key
func (x *Model) SetKey(key *datastore.Key) {
	x.key = key
	if key == nil {
		x.ID = ""
		return
	}
	if key.Name != "" {
		x.ID = key.Name
		return
	}
	if key.ID != 0 {
		x.ID = strconv.FormatInt(key.ID, 10)
	}
}

// KeyGetter interface
type KeyGetter interface {
	Key() *datastore.Key
}

// KeySetter interface
type KeySetter interface {
	SetKey(*datastore.Key)
}

// KeyGetSetter interface
type KeyGetSetter interface {
	KeyGetter
	KeySetter
}

// KindGetter interface
type KindGetter interface {
	Kind() string
}
