package ds

import (
	"errors"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

// Errors
var (
	ErrInvalidID = errors.New("ds: invalid id")
)

// NotFound checks is error means not found
func NotFound(err error) bool {
	return err == iterator.Done || err == datastore.ErrNoSuchEntity
}

// FieldMismatch checks is error field mismatch
func FieldMismatch(err error) bool {
	_, ok := err.(*datastore.ErrFieldMismatch)
	return ok
}

// IgnoreFieldMismatch returns nil if err is field mismatch error(s)
func IgnoreFieldMismatch(err error) error {
	if FieldMismatch(err) {
		return nil
	}

	// check is multi errors
	if errs, ok := err.(datastore.MultiError); ok {
		es := make(datastore.MultiError, 0)
		for _, err := range errs {
			if !FieldMismatch(err) {
				es = append(es, err)
			}
		}
		if len(es) > 0 {
			return es
		}
		return nil
	}
	return err
}
