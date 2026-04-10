// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chamelaion_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/chamelaion/chamelaion-go"
	"github.com/chamelaion/chamelaion-go/internal/testutil"
	"github.com/chamelaion/chamelaion-go/option"
)

func TestLipsyncRequestGet(t *testing.T) {
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
	_, err := client.Lipsync.Requests.Get(context.TODO(), "6f82a2d8-a6d4-4e8a-a0fa-e8b09823a2d8")
	if err != nil {
		var apierr *chamelaion.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLipsyncRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.Lipsync.Requests.List(context.TODO(), chamelaion.LipsyncRequestListParams{
		Limit:       chamelaion.Int(20),
		Offset:      chamelaion.Int(0),
		ReferenceID: chamelaion.String("batch-2026-04"),
	})
	if err != nil {
		var apierr *chamelaion.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
