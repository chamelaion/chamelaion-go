package chamelaion

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"strings"
	"testing"
)

func TestVideoTranslateGenerateJSON(t *testing.T) {
	body := VideoTranslateGenerateParams{
		Inputs:          []VideoTranslateGenerateParamsInput{{Type: "video", URL: "https://example.com/video.mp4"}},
		TargetLanguages: []string{"de"},
		WebhookURL:      String("https://hooks.example.com/video-translate"),
		Config:          map[string]any{"speaker_voices": map[string]string{"0": "voice-id"}},
	}
	encoded, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{`"target_languages":["de"]`, `"webhook_url":"https://hooks.example.com/video-translate"`, `"speaker_voices"`} {
		if !bytes.Contains(encoded, []byte(expected)) {
			t.Fatalf("body %s does not contain %s", encoded, expected)
		}
	}
}

func TestVideoTranslateMultipartUsesSRTPartNames(t *testing.T) {
	body := VideoTranslateGenerateWithMediaParams{
		Video:           File(strings.NewReader("video"), "input.mp4", "video/mp4"),
		TargetLanguages: []string{"de", "fr"},
		Config:          String(`{"speaker_voices":{"0":"voice-id"}}`),
		SourceSRT:       File(strings.NewReader("source"), "source.srt", "application/x-subrip"),
		TargetSRTs: map[string]io.Reader{
			"de": File(strings.NewReader("target"), "target.de.srt", "application/x-subrip"),
		},
	}
	encoded, contentType, err := body.MarshalMultipart()
	if err != nil {
		t.Fatal(err)
	}
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		t.Fatal(err)
	}
	reader := multipart.NewReader(bytes.NewReader(encoded), params["boundary"])
	fields := map[string]string{}
	for {
		part, nextErr := reader.NextPart()
		if nextErr == io.EOF {
			break
		}
		if nextErr != nil {
			t.Fatal(nextErr)
		}
		data, readErr := io.ReadAll(part)
		if readErr != nil {
			t.Fatal(readErr)
		}
		fields[part.FormName()] = string(data)
	}
	if fields["target_languages"] != "de,fr" {
		t.Fatalf("target_languages = %q", fields["target_languages"])
	}
	for _, name := range []string{"video", "source_srt", "target_srt.de", "config"} {
		if _, ok := fields[name]; !ok {
			t.Fatalf("missing multipart field %q: %#v", name, fields)
		}
	}
}
