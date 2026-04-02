// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chamelaion_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/chamelaion/chamelaion-go"
	"github.com/chamelaion/chamelaion-go/internal/testutil"
	"github.com/chamelaion/chamelaion-go/option"
)

func TestLipsyncGenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Lipsync.Generate(context.TODO(), chamelaion.LipsyncGenerateParams{
		Inputs: []chamelaion.LipsyncGenerateParamsInput{{
			Type: "video",
			URL:  "https://example.com/source/video.mp4",
		}, {
			Type: "audio",
			URL:  "https://example.com/source/audio.wav",
		}},
		DisableActiveSpeakerDetection: chamelaion.Bool(false),
		ReferenceID:                   chamelaion.String("demo-001"),
	})
	if err != nil {
		var apierr *chamelaion.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLipsyncGenerateWithMediaWithOptionalParams(t *testing.T) {
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
	_, err := client.Lipsync.GenerateWithMedia(context.TODO(), chamelaion.LipsyncGenerateWithMediaParams{
		Audio:                         io.Reader(bytes.NewBuffer([]byte("Example data"))),
		Video:                         io.Reader(bytes.NewBuffer([]byte("Example data"))),
		DisableActiveSpeakerDetection: chamelaion.Bool(true),
		Model:                         chamelaion.LipsyncGenerateWithMediaParamsModelLipsync2,
		ReferenceID:                   chamelaion.String("reference_id"),
	})
	if err != nil {
		var apierr *chamelaion.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
