package ds

import (
	"reflect"

	"cloud.google.com/go/datastore"
)

func valueOf(src interface{}) reflect.Value {
	xs := reflect.ValueOf(src)
	if xs.Kind() == reflect.Ptr {
		xs = xs.Elem()
	}
	return xs
}

func prepareQuery(kind string, qs []Query) *datastore.Query {
	q := datastore.NewQuery(kind)
	for _, setter := range qs {
		q = setter(q)
	}
	return q
}
