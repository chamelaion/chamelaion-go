// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chamelaion

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/chamelaion/chamelaion-go/internal/apiform"
	"github.com/chamelaion/chamelaion-go/internal/apijson"
	"github.com/chamelaion/chamelaion-go/internal/requestconfig"
	"github.com/chamelaion/chamelaion-go/option"
	"github.com/chamelaion/chamelaion-go/packages/param"
	"github.com/chamelaion/chamelaion-go/packages/respjson"
)

// VideoTranslateService contains endpoints for asynchronous video translation.
type VideoTranslateService struct {
	options  []option.RequestOption
	Requests VideoTranslateRequestService
	Voices   VideoTranslateVoiceService
}

func NewVideoTranslateService(opts ...option.RequestOption) (r VideoTranslateService) {
	r.options = opts
	r.Requests = NewVideoTranslateRequestService(opts...)
	r.Voices = NewVideoTranslateVoiceService(opts...)
	return
}

func (r *VideoTranslateService) Generate(ctx context.Context, body VideoTranslateGenerateParams, opts ...option.RequestOption) (res *VideoTranslateGenerate, err error) {
	opts = slices.Concat(r.options, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, "v1/video-translate/generate", body, &res, opts...)
	return res, err
}

func (r *VideoTranslateService) GenerateWithMedia(ctx context.Context, body VideoTranslateGenerateWithMediaParams, opts ...option.RequestOption) (res *VideoTranslateGenerate, err error) {
	opts = slices.Concat(r.options, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, "v1/video-translate/generate-with-media", body, &res, opts...)
	return res, err
}

type VideoTranslateGenerate struct {
	RequestID string                       `json:"request_id" api:"required" format:"uuid"`
	Status    VideoTranslateGenerateStatus `json:"status" api:"required"`
	JSON      struct {
		RequestID   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

func (r VideoTranslateGenerate) RawJSON() string { return r.JSON.raw }
func (r *VideoTranslateGenerate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTranslateGenerateStatus string

const VideoTranslateGenerateStatusSuccess VideoTranslateGenerateStatus = "success"

type VideoTranslateGenerateParams struct {
	Inputs                 []VideoTranslateGenerateParamsInput `json:"inputs,omitzero" api:"required"`
	TargetLanguages        []string                            `json:"target_languages,omitzero" api:"required"`
	Config                 map[string]any                      `json:"config,omitzero"`
	DisableBackgroundAudio param.Opt[bool]                     `json:"disable_background_audio,omitzero"`
	DisableLipSync         param.Opt[bool]                     `json:"disable_lip_sync,omitzero"`
	Formality              param.Opt[bool]                     `json:"formality,omitzero"`
	ReferenceID            param.Opt[string]                   `json:"reference_id,omitzero"`
	SourceLanguage         param.Opt[string]                   `json:"source_language,omitzero"`
	WebhookURL             param.Opt[string]                   `json:"webhook_url,omitzero"`
	paramObj
}

func (r VideoTranslateGenerateParams) MarshalJSON() ([]byte, error) {
	type shadow VideoTranslateGenerateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type VideoTranslateGenerateParamsInput struct {
	Type string `json:"type,omitzero" api:"required"`
	URL  string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r VideoTranslateGenerateParamsInput) MarshalJSON() ([]byte, error) {
	type shadow VideoTranslateGenerateParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() { apijson.RegisterFieldValidator[VideoTranslateGenerateParamsInput]("type", "video") }

type VideoTranslateGenerateWithMediaParams struct {
	Video                  io.Reader            `json:"video,omitzero" api:"required" format:"binary"`
	TargetLanguages        []string             `json:"target_languages,omitzero" api:"required"`
	Config                 param.Opt[string]    `json:"config,omitzero"`
	DisableBackgroundAudio param.Opt[bool]      `json:"disable_background_audio,omitzero"`
	DisableLipSync         param.Opt[bool]      `json:"disable_lip_sync,omitzero"`
	Formality              param.Opt[bool]      `json:"formality,omitzero"`
	ReferenceID            param.Opt[string]    `json:"reference_id,omitzero"`
	SourceLanguage         param.Opt[string]    `json:"source_language,omitzero"`
	SourceSRT              io.Reader            `json:"source_srt,omitzero" format:"binary"`
	TargetSRTs             map[string]io.Reader `json:"target_srt,omitzero"`
	WebhookURL             param.Opt[string]    `json:"webhook_url,omitzero"`
	paramObj
}

func (r VideoTranslateGenerateWithMediaParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	if err = writer.Close(); err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
