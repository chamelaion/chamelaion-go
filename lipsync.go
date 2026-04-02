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

// Endpoints for creating and retrieving lip sync requests.
//
// LipsyncService contains methods and other services that help with interacting
// with the chamelaion API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLipsyncService] method instead.
type LipsyncService struct {
	options []option.RequestOption
	// Endpoints for creating and retrieving lip sync requests.
	Requests LipsyncRequestService
}

// NewLipsyncService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLipsyncService(opts ...option.RequestOption) (r LipsyncService) {
	r = LipsyncService{}
	r.options = opts
	r.Requests = NewLipsyncRequestService(opts...)
	return
}

// Starts a lip sync job from exactly one remote video URL and one remote audio
// URL. The response includes a request identifier to poll later.
func (r *LipsyncService) Generate(ctx context.Context, body LipsyncGenerateParams, opts ...option.RequestOption) (res *LipsyncGenerate, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/lipsync/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Starts a lip sync job by uploading one video file and one audio file as
// multipart form-data.
func (r *LipsyncService) GenerateWithMedia(ctx context.Context, body LipsyncGenerateWithMediaParams, opts ...option.RequestOption) (res *LipsyncGenerate, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/lipsync/generate-with-media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type LipsyncGenerate struct {
	// Identifier of the created lip sync request.
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// Current state of the newly created request.
	//
	// Any of "success".
	Status LipsyncGenerateStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LipsyncGenerate) RawJSON() string { return r.JSON.raw }
func (r *LipsyncGenerate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of the newly created request.
type LipsyncGenerateStatus string

const (
	LipsyncGenerateStatusSuccess LipsyncGenerateStatus = "success"
)

type LipsyncGenerateParams struct {
	// Exactly one video input and one audio input (order does not matter).
	Inputs []LipsyncGenerateParamsInput `json:"inputs,omitzero" api:"required"`
	// Disable active speaker detection and use max-face lipsync preprocessing.
	DisableActiveSpeakerDetection param.Opt[bool] `json:"disable_active_speaker_detection,omitzero"`
	// Optional client-provided identifier for later retrieval.
	ReferenceID param.Opt[string] `json:"reference_id,omitzero"`
	paramObj
}

func (r LipsyncGenerateParams) MarshalJSON() (data []byte, err error) {
	type shadow LipsyncGenerateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LipsyncGenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, URL are required.
type LipsyncGenerateParamsInput struct {
	// Type of input media.
	//
	// Any of "video", "audio".
	Type string `json:"type,omitzero" api:"required"`
	// URL of the media resource.
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r LipsyncGenerateParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow LipsyncGenerateParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LipsyncGenerateParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LipsyncGenerateParamsInput](
		"type", "video", "audio",
	)
}

type LipsyncGenerateWithMediaParams struct {
	// Target audio file.
	Audio io.Reader `json:"audio,omitzero" api:"required" format:"binary"`
	// Source video file.
	Video io.Reader `json:"video,omitzero" api:"required" format:"binary"`
	// Disable active speaker detection and use max-face lipsync preprocessing.
	DisableActiveSpeakerDetection param.Opt[bool] `json:"disable_active_speaker_detection,omitzero"`
	// Optional client-provided identifier for later retrieval.
	ReferenceID param.Opt[string] `json:"reference_id,omitzero"`
	// Optional model selector.
	//
	// Any of "lipsync-2".
	Model LipsyncGenerateWithMediaParamsModel `json:"model,omitzero"`
	paramObj
}

func (r LipsyncGenerateWithMediaParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Optional model selector.
type LipsyncGenerateWithMediaParamsModel string

const (
	LipsyncGenerateWithMediaParamsModelLipsync2 LipsyncGenerateWithMediaParamsModel = "lipsync-2"
)
