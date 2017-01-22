package ds_test

import (
	"context"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/datastore"

	"google.golang.org/api/option"

	. "github.com/acoshift/ds"
	"golang.org/x/oauth2/google"
)

type ExampleModel struct {
	Model
	StampModel
	Name  string
	Value int
}

func initClient() (*Client, error) {
	// load service account from env
	serviceAccount := os.Getenv("service_account")
	log.Println(serviceAccount)
	opts := []option.ClientOption{}
	if serviceAccount != "" {
		cfg, err := google.JWTConfigFromJSON([]byte(serviceAccount), datastore.ScopeDatastore)
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
	if x.GetID() == 0 {
		t.Errorf("expected id to be assigned")
	}
	err = client.DeleteByID(ctx, "ExampleModel", x.GetID())
	if err != nil {
		t.Error(err)
	}
}
