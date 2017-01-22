package ds

import (
	"context"
	"testing"
)

func TestSave(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping save")
	}
	ctx := context.Background()
	client, err := initClient()
	if err != nil {
		t.Fatal(err)
	}
	x := &ExampleModel{Name: "Test1", Value: 1}
	err = client.SaveModel(ctx, "ExampleModel", x)
	if err != nil {
		t.Error(err)
	}
	if x.Key() == nil {
		t.Errorf("expetect key to be assigned")
	}
	if x.CreatedAt.IsZero() || x.UpdatedAt.IsZero() {
		t.Errorf("expetect stamp model to be assigned")
	}
	if x.ID() == 0 {
		t.Errorf("expected id to be assigned")
	}
	err = client.DeleteModel(ctx, x)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveMulti(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping save multi")
	}
	ctx := context.Background()
	client, err := initClient()
	if err != nil {
		t.Fatal(err)
	}
	xs := []*ExampleModel{
		{Name: "Test1", Value: 1},
		{Name: "Test2", Value: 2},
	}
	err = client.SaveModels(ctx, "Test", xs)
	if err != nil {
		t.Error(err)
	}
	for _, x := range xs {
		if x.Key() == nil {
			t.Errorf("expetect key to be assigned")
		}
		if x.CreatedAt.IsZero() || x.UpdatedAt.IsZero() {
			t.Errorf("expetect stamp model to be assigned")
		}
	}
	err = client.DeleteModels(ctx, xs)
	if err != nil {
		t.Error(err)
	}
}
