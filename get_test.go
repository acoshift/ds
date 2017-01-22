package ds

import (
	"context"
	"testing"
)

func checkSkipGet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping get")
	}
}

func TestGetByKey(t *testing.T) {
	checkSkipGet(t)
	client, err := initClient()
	if err != nil {
		t.Fatal(err)
	}

	keys := prepareData(client)
	defer removeData(client)

	ctx := context.Background()
	var x ExampleModel

	err = client.GetByKey(ctx, keys[0], &x)
	if err != nil {
		t.Error(err)
	}
	if !x.Key().Equal(keys[0]) {
		t.Errorf("key not equals")
	}

	xs := make([]*ExampleModel, len(keys))
	err = client.GetByKeys(ctx, keys, xs)
	if err != nil {
		t.Error(err)
	}
	if len(keys) != len(xs) {
		t.Errorf("keys and result len not equals")
	}
	for i := range keys {
		if !keys[i].Equal(xs[i].Key()) {
			t.Errorf("key not equals")
		}
	}

	var xs2 []*ExampleModel
	err = client.GetByKeys(ctx, keys, &xs2)
	if err != nil {
		t.Error(err)
	}
	if len(keys) != len(xs2) {
		t.Errorf("keys and result len not equals")
	}
	for i := range keys {
		if !keys[i].Equal(xs2[i].Key()) {
			t.Errorf("key not equals")
		}
	}
}
