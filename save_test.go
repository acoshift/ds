package ds

import (
	"context"
	"testing"
)

func TestSave(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping save.")
	}
	ctx := context.Background()
	client, err := initClient()
	if err != nil {
		t.Fatal(err)
	}
	x := &ExampleModel{Name: "Test1", Value: 1}
	err = client.Save(ctx, "ExampleModel", x)
	if err != nil {
		t.Error(err)
	}
	if x.Key() == nil {
		t.Errorf("expetect key to be assigned")
	}
	if x.ID() == 0 {
		t.Errorf("expected id to be assigned")
	}
	err = client.DeleteByID(ctx, "ExampleModel", x.ID())
	if err != nil {
		t.Error(err)
	}
}
