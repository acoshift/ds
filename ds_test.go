package ds

import (
	"context"
	"encoding/base64"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type ExampleModel struct {
	Model
	StampModel
	Name  string
	Value int
}

type ExampleNotModel struct {
	Name string
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

func TestInvalidNewClient(t *testing.T) {
	client, err := NewClient(context.Background(), "invalid-project-id", option.WithServiceAccountFile("invalid-file"))
	if err == nil {
		t.Errorf("expected error not nil")
	}
	if client != nil {
		t.Errorf("expected client to be nil")
	}
}
