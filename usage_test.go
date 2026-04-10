// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chamelaion_test

import (
	"context"
	"os"
	"testing"

	"github.com/chamelaion/chamelaion-go"
	"github.com/chamelaion/chamelaion-go/internal/testutil"
	"github.com/chamelaion/chamelaion-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chamelaion.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
		option.WithAPIKey("My API Key"),
	)
	response, err := client.Health.Check(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Status)
}
