package ds

import (
	"testing"
)

func TestQuery(t *testing.T) {
	skipShort(t, "Query")
	client, err := initClient()
	if err != nil {
		t.Fatal(err)
	}
	keys := prepareData(client)
	defer removeData(client)

	var xs []*ExampleModel
	err = client.Query(ctx, keys[0].Kind, &xs)
	if err != nil {
		t.Error(err)
	}

	var x ExampleModel
	err = client.QueryFirst(ctx, keys[0].Kind, &x)
	if err != nil {
		t.Error(err)
	}

	ks, err := client.QueryKeys(ctx, keys[0].Kind)
	if err != nil {
		t.Error(err)
	}
	if len(keys) != len(ks) {
		t.Error("expected query by keys have same length as prepare data")
	}
}
