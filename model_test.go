package ds

import (
	"testing"

	"cloud.google.com/go/datastore"
)

func TestModel(t *testing.T) {
	var x *Model
	if x.Key() != nil {
		t.Errorf("expected key of nil to be nil")
	}

	// Should not panic
	x.SetKey(nil)
	x.SetKey(datastore.IDKey("Test", int64(10), nil))
	x.SetID("Test", 10)
	x.NewKey("Test")
	if x.ID() != 0 {
		t.Errorf("expected id to be 0")
	}

	x = &Model{}
	x.NewKey("Test")
	if x.Key() == nil {
		t.Errorf("expected key not nil")
	}

	x.SetKey(nil)
	if x.Key() != nil {
		t.Errorf("expected key to be nil")
	}
	if x.ID() != 0 {
		t.Errorf("expected id to be 0")
	}

	x.SetID("Test", 10)
	if x.Key() == nil {
		t.Errorf("expected key not nil")
	}
	if x.ID() == 0 {
		t.Errorf("expected id not 0")
	}
}

func TestStringIDModel(t *testing.T) {
	var x *StringIDModel
	if x.Key() != nil {
		t.Errorf("expected key of nil to be nil")
	}

	// Should not panic
	x.SetKey(nil)
	x.SetKey(datastore.IDKey("Test", int64(10), nil))
	x.SetID("Test", 10)
	x.SetStringID("Test", "aaa")
	x.SetNameID("Test", "bbb")
	x.NewKey("Test")
	if x.ID() != "" {
		t.Errorf("expected id to be empty")
	}

	x = &StringIDModel{}
	x.NewKey("Test")
	if x.Key() == nil {
		t.Errorf("expected key not nil")
	}

	x.SetKey(nil)
	if x.Key() != nil {
		t.Errorf("expected key to be nil")
	}
	if x.ID() != "" {
		t.Errorf("expected id to be empty")
	}

	x.SetID("Test", 10)
	if x.Key() == nil {
		t.Errorf("expected key not nil")
	}
	if x.ID() == "" {
		t.Errorf("expected id not empty")
	}

	x.SetStringID("Test", "10")
	if x.Key() == nil {
		t.Errorf("expected key not nil")
	}
	if x.ID() != "10" {
		t.Errorf("expected id to be %s; got %s", "10", x.ID())
	}

	x.SetNameID("Test", "aaa")
	if x.Key() == nil {
		t.Errorf("expected key not nil")
	}
	if x.ID() != "aaa" {
		t.Errorf("expected id to be %s; got %s", "aaa", x.ID())
	}
}
