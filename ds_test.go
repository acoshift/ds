package ds_test

import (
	"context"
	"encoding/base64"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
	. "github.com/acoshift/ds"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type ExampleModel struct {
	Model
	StampModel
	Name  string
	Value int
}

func initClient() (*Client, error) {
	// load service account from env
	serviceAccountStr := os.Getenv("service_account")
	opts := []option.ClientOption{}
	if serviceAccountStr != "" {
		serviceAccount, err := base64.StdEncoding.DecodeString(serviceAccountStr)
		if err != nil {
			return nil, err
		}
		cfg, err := google.JWTConfigFromJSON(serviceAccount, datastore.ScopeDatastore)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(cfg.TokenSource(context.Background())))
	}
	projectID := os.Getenv("project_id")
	if projectID == "" {
		projectID = "acoshift-test"
	}
	return NewClient(context.Background(), projectID, opts...)
}

func TestSave(t *testing.T) {
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
