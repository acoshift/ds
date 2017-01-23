package ds

import (
	"testing"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

func TestError(t *testing.T) {
	if !NotFound(iterator.Done) {
		t.Errorf("expected iterator.Done to be not found error")
	}
	if !NotFound(datastore.ErrNoSuchEntity) {
		t.Errorf("expected datastore.ErrNoSuchEntity to be not found error")
	}
	if !NotFound(datastore.ErrInvalidKey) {
		t.Errorf("expected datastore.ErrInvalidKey to be not found error")
	}

	if IgnoreNotFound(iterator.Done) != nil {
		t.Errorf("expected not found error to be ignored")
	}
	if IgnoreNotFound(datastore.ErrInvalidEntityType) == nil {
		t.Errorf("expected other error than not found error not ignored")
	}

	if !FieldMismatch(&datastore.ErrFieldMismatch{}) {
		t.Errorf("expected datastore.ErrFieldMismatch is field mismatch error")
	}

	if IgnoreFieldMismatch(&datastore.ErrFieldMismatch{}) != nil {
		t.Errorf("expected field mismatch error to be ignored")
	}
	if IgnoreFieldMismatch(datastore.ErrInvalidKey) == nil {
		t.Errorf("expected other error than field mismatch error not ignored")
	}

	errs := datastore.MultiError{&datastore.ErrFieldMismatch{}, &datastore.ErrFieldMismatch{}}
	if IgnoreFieldMismatch(errs) != nil {
		t.Errorf("expected field mismatch errors to be ignored")
	}
	if IgnoreNotFound(errs) == nil {
		t.Errorf("expected errors not ignored")
	}
}
