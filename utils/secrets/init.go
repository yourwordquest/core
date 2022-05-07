package secrets

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var project string = os.Getenv("PROJECT_ID")
var client *secretmanager.Client

var testEnv = os.Getenv("TEST_ENVIRONMENT") == "true"

func init() {
	var err error

	client, err = secretmanager.NewClient(context.Background())
	if err != nil {
		// Ignore secrets setup in tests
		if testEnv {
			return
		}
		panic(err)
	}
}

func GetSecret(name string) (string, error) {
	if testEnv {
		value, ok := os.LookupEnv(name)
		if !ok {
			return value, errors.New("secret value not set")
		}
		return value, nil
	}
	if client == nil {
		return "", errors.New("secrets client not initialized")
	}
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", project, name),
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	resp, err := client.AccessSecretVersion(ctx, req)

	if err != nil {
		return "", err
	}

	payload := resp.GetPayload()
	return string(payload.GetData()), nil
}
