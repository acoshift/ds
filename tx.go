package ds

import (
	"context"
	"reflect"

	"cloud.google.com/go/datastore"
)

// Tx is the datastore transaction wrapper
type Tx struct {
	*datastore.Transaction
}

// RunInTx is the RunInTransaction wrapper
func (client *Client) RunInTx(ctx context.Context, f func(tx *Tx) error, opts ...datastore.TransactionOption) (*datastore.Commit, error) {
	return client.RunInTransaction(ctx, func(t *datastore.Transaction) error {
		return f(&Tx{t})
	})
}

// GetByKey retrieves model from datastore by key
func (tx *Tx) GetByKey(key *datastore.Key, dst interface{}) error {
	err := tx.Get(key, dst)
	if err != nil {
		return err
	}
	return nil
}

// GetByKeys retrieves models from datastore by keys
func (tx *Tx) GetByKeys(keys []*datastore.Key, dst interface{}) error {
	// prepare slice if dst is pointer to 0 len slice
	if rf := reflect.ValueOf(dst); rf.Kind() == reflect.Ptr {
		rs := rf.Elem()
		if rs.Kind() == reflect.Slice && rs.Len() == 0 {
			l := len(keys)
			rs.Set(reflect.MakeSlice(rs.Type(), l, l))
		}
		dst = rs.Interface()
	}

	err := tx.GetMulti(keys, dst)
	if err != nil {
		return err
	}
	return nil
}

// GetByModel retrieves model from datastore by key from model
func (tx *Tx) GetByModel(dst interface{}) error {
	key := ExtractKey(dst)
	return tx.GetByKey(key, dst)
}

// GetByModels retrieves models from datastore by keys from models
func (tx *Tx) GetByModels(dst interface{}) error {
	keys := ExtractKeys(dst)
	return tx.GetByKeys(keys, dst)
}

// GetByID retrieves model from datastore by id
func (tx *Tx) GetByID(kind string, id int64, dst interface{}) error {
	return tx.GetByKey(datastore.IDKey(kind, id, nil), dst)
}

// GetByIDs retrieves models from datastore by ids
func (tx *Tx) GetByIDs(kind string, ids []int64, dst interface{}) error {
	keys := BuildIDKeys(kind, ids)
	return tx.GetByKeys(keys, dst)
}

// GetByStringID retrieves model from datastore by string id
func (tx *Tx) GetByStringID(kind string, id string, dst interface{}) error {
	tid := parseID(id)
	if tid == 0 {
		return datastore.ErrInvalidKey
	}
	return tx.GetByKey(datastore.IDKey(kind, tid, nil), dst)
}

// GetByStringIDs retrieves models from datastore by string ids
func (tx *Tx) GetByStringIDs(kind string, ids []string, dst interface{}) error {
	keys := BuildStringIDKeys(kind, ids)
	return tx.GetByKeys(keys, dst)
}

// GetByName retrieves model from datastore by name
func (tx *Tx) GetByName(kind string, name string, dst interface{}) error {
	return tx.GetByKey(datastore.NameKey(kind, name, nil), dst)
}

// GetByNames retrieves models from datastore by names
func (tx *Tx) GetByNames(kind string, names []string, dst interface{}) error {
	keys := BuildNameKeys(kind, names)
	return tx.GetByKeys(keys, dst)
}

// PutModel puts a model to datastore
func (tx *Tx) PutModel(src interface{}) error {
	x := src.(KeyGetSetter)
	_, err := tx.Put(x.GetKey(), x)
	return err
}

// PutModels puts models to datastore
func (tx *Tx) PutModels(src interface{}) ([]*datastore.PendingKey, error) {
	xs := valueOf(src)
	keys := make([]*datastore.Key, xs.Len())
	for i := range keys {
		x := xs.Index(i).Interface()
		keys[i] = x.(KeyGetter).GetKey()
	}
	// TODO: store pending key inside model ?
	return tx.PutMulti(keys, src)
}

// SaveModel saves model to datastore
func (tx *Tx) SaveModel(kind string, src interface{}) (*datastore.PendingKey, error) {
	beforeSave(kind, src)

	x := src.(KeyGetSetter)
	key, err := tx.Put(x.GetKey(), x)
	// x.SetKey(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// SaveModels saves models to datastore
func (tx *Tx) SaveModels(kind string, src interface{}) ([]*datastore.PendingKey, error) {
	xs := valueOf(src)
	for i := 0; i < xs.Len(); i++ {
		x := xs.Index(i).Interface()
		beforeSave(kind, x)
	}
	keys, err := tx.PutModels(src)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
