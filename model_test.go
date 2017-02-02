package ds

import (
	"testing"

	"fmt"

	"cloud.google.com/go/datastore"
)

func TestModel(t *testing.T) {
	var x *Model
	if x.GetKey() != nil {
		t.Fatalf("expected key of nil to be nil")
	}

	// Should not panic
	x.SetKey(nil)
	x.SetKey(datastore.IDKey("Test", int64(10), nil))
	x.SetID("Test", 10)
	x.NewKey("Test")
	if x.ID() != 0 {
		t.Fatalf("expected id to be 0")
	}

	x = &Model{}
	x.NewKey("Test")
	if x.GetKey() == nil {
		t.Fatalf("expected key not nil")
	}

	x.SetKey(nil)
	if x.GetKey() != nil {
		t.Fatalf("expected key to be nil")
	}
	if x.ID() != 0 {
		t.Fatalf("expected id to be 0")
	}

	x.SetID("Test", 10)
	if x.GetKey() == nil {
		t.Fatalf("expected key not nil")
	}
	if x.ID() == 0 {
		t.Fatalf("expected id not 0")
	}
}

func TestStringIDModel(t *testing.T) {
	var x *StringIDModel
	if x.GetKey() != nil {
		t.Fatalf("expected key of nil to be nil")
	}

	// Should not panic
	x.SetKey(nil)
	x.SetKey(datastore.IDKey("Test", int64(10), nil))
	x.SetID("Test", 10)
	x.SetStringID("Test", "aaa")
	x.SetNameID("Test", "bbb")
	x.NewKey("Test")
	if x.ID() != "" {
		t.Fatalf("expected id to be empty")
	}

	x = &StringIDModel{}
	x.NewKey("Test")
	if x.GetKey() == nil {
		t.Fatalf("expected key not nil")
	}

	x.SetKey(nil)
	if x.GetKey() != nil {
		t.Fatalf("expected key to be nil")
	}
	if x.ID() != "" {
		t.Fatalf("expected id to be empty")
	}

	x.SetID("Test", 10)
	if x.GetKey() == nil {
		t.Fatalf("expected key not nil")
	}
	if x.ID() == "" {
		t.Fatalf("expected id not empty")
	}

	x.SetStringID("Test", "10")
	if x.GetKey() == nil {
		t.Fatalf("expected key not nil")
	}
	if x.ID() != "10" {
		t.Fatalf("expected id to be %s; got %s", "10", x.ID())
	}

	x.SetNameID("Test", "aaa")
	if x.GetKey() == nil {
		t.Fatalf("expected key not nil")
	}
	if x.ID() != "aaa" {
		t.Fatalf("expected id to be %s; got %s", "aaa", x.ID())
	}
}

func TestSetKey(t *testing.T) {
	x := &Model{}
	key := datastore.IDKey("Test", 1, nil)
	SetKey(nil, nil)
	SetKey(nil, x)
	SetKey(key, nil)
	SetKey(key, x)
	if x.GetKey() == nil {
		t.Fatalf("expected key not nil")
	}
	y := Model{}
	// Set to unpointer should not have side-effect
	SetKey(key, y)
	if y.GetKey() != nil {
		t.Fatalf("expected key to be nil")
	}
}

func TestSetKeys(t *testing.T) {
	xs := []interface{}{
		&Model{},
		Model{},
		nil,
		ExampleModel{},
		&ExampleModel{},
		ExampleNotModel{},
		&ExampleNotModel{},
	}
	keys := make([]*datastore.Key, len(xs))
	for i := range xs {
		keys[i] = datastore.IDKey("Test", int64(i), nil)
	}
	SetKeys(nil, nil)
	SetKeys(keys, nil)
	SetKeys(nil, xs)
	SetKeys(keys, xs)
	SetKeys(keys, &xs)
	for i := range xs {
		if x, ok := xs[i].(KeyGetter); ok {
			if x.GetKey() == nil {
				t.Fatalf("expected key not nil")
			}
			if x.GetKey() != keys[i] {
				t.Fatalf("wrong key")
			}
		}
	}
}

func TestSetIDs(t *testing.T) {
	xs := []interface{}{
		&Model{},
		Model{},
		nil,
		ExampleModel{},
		&ExampleModel{},
		ExampleNotModel{},
		&ExampleNotModel{},
	}
	ids := make([]int64, len(xs))
	for i := range xs {
		ids[i] = int64(i + 1)
	}
	validate := func() {
		for i := range xs {
			if x, ok := xs[i].(KeyGetter); ok {
				if x.GetKey().ID != ids[i] {
					t.Fatalf("expected id to be %d; got %d", ids[i], x.GetKey().ID)
				}
			}
		}
	}
	SetIDs("Test", ids, xs)
	validate()

	xs = []interface{}{
		&Model{},
		Model{},
		nil,
		ExampleModel{},
		&ExampleModel{},
		ExampleNotModel{},
		&ExampleNotModel{},
	}
	SetIDs("Test", ids, &xs)
	validate()
}

func TestSetStringIDs(t *testing.T) {
	xs := []interface{}{
		&Model{},
		Model{},
		nil,
		ExampleModel{},
		&ExampleModel{},
		ExampleNotModel{},
		&ExampleNotModel{},
	}
	ids := make([]string, len(xs))
	for i := range xs {
		ids[i] = fmt.Sprintf("%d", i+1)
	}
	validate := func() {
		for i := range xs {
			if x, ok := xs[i].(KeyGetter); ok {
				if x.GetKey().ID != parseID(ids[i]) {
					t.Fatalf("expected id to be %s; got %d", ids[i], x.GetKey().ID)
				}
			}
		}
	}
	SetStringIDs("Test", ids, xs)
	validate()

	SetStringIDs("Test", ids, &xs)
	validate()
}

func TestSetNameIDs(t *testing.T) {
	xs := []interface{}{
		&Model{},
		Model{},
		nil,
		ExampleModel{},
		&ExampleModel{},
		ExampleNotModel{},
		&ExampleNotModel{},
	}
	ids := make([]string, len(xs))
	for i := range xs {
		ids[i] = fmt.Sprintf("test%d", i+1)
	}
	validate := func() {
		for i := range xs {
			if x, ok := xs[i].(KeyGetter); ok {
				if x.GetKey().ID != parseID(ids[i]) {
					t.Fatalf("expected id to be %s; got %d", ids[i], x.GetKey().ID)
				}
			}
		}
	}
	SetNameIDs("Test", ids, xs)
	validate()

	SetNameIDs("Test", ids, &xs)
	validate()
}
