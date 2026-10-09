// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chamelaion

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/chamelaion/chamelaion-go/internal/apijson"
	"github.com/chamelaion/chamelaion-go/internal/apiquery"
	"github.com/chamelaion/chamelaion-go/internal/requestconfig"
	"github.com/chamelaion/chamelaion-go/option"
	"github.com/chamelaion/chamelaion-go/packages/param"
	"github.com/chamelaion/chamelaion-go/packages/respjson"
)

type VideoTranslateRequestService struct{ options []option.RequestOption }

func NewVideoTranslateRequestService(opts ...option.RequestOption) (r VideoTranslateRequestService) {
	r.options = opts
	return
}

func (r *VideoTranslateRequestService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoTranslateRequest, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		return nil, errors.New("missing required id parameter")
	}
	path := fmt.Sprintf("v1/video-translate/requests/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *VideoTranslateRequestService) List(ctx context.Context, query VideoTranslateRequestListParams, opts ...option.RequestOption) (res *VideoTranslateRequestListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "v1/video-translate/requests", query, &res, opts...)
	return res, err
}

type VideoTranslateArtifactSet struct {
	Video string `json:"video" api:"required" format:"uri"`
	Audio string `json:"audio" api:"required" format:"uri"`
	SRT   string `json:"srt" api:"required" format:"uri"`
}

type VideoTranslateRequest struct {
	ID                     string                               `json:"id" api:"required" format:"uuid"`
	TargetLanguages        []string                             `json:"target_languages" api:"required"`
	DisableBackgroundAudio bool                                 `json:"disable_background_audio" api:"required"`
	DisableLipSync         bool                                 `json:"disable_lip_sync" api:"required"`
	Formality              bool                                 `json:"formality" api:"required"`
	Config                 map[string]any                       `json:"config" api:"required"`
	Status                 VideoTranslateRequestStatus          `json:"status" api:"required"`
	CreatedAt              time.Time                            `json:"created_at" api:"required" format:"date-time"`
	ErrorMessage           string                               `json:"error_message"`
	FinishedAt             time.Time                            `json:"finished_at" format:"date-time"`
	Outputs                map[string]VideoTranslateArtifactSet `json:"outputs"`
	ReferenceID            string                               `json:"reference_id"`
	SourceLanguage         string                               `json:"source_language"`
	StartedAt              time.Time                            `json:"started_at" format:"date-time"`
	WebhookURL             string                               `json:"webhook_url" format:"uri"`
	JSON                   struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

func (r VideoTranslateRequest) RawJSON() string { return r.JSON.raw }
func (r *VideoTranslateRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTranslateRequestStatus string

const (
	VideoTranslateRequestStatusQueued     VideoTranslateRequestStatus = "queued"
	VideoTranslateRequestStatusUploading  VideoTranslateRequestStatus = "uploading"
	VideoTranslateRequestStatusProcessing VideoTranslateRequestStatus = "processing"
	VideoTranslateRequestStatusCompleted  VideoTranslateRequestStatus = "completed"
	VideoTranslateRequestStatusFailed     VideoTranslateRequestStatus = "failed"
)

type VideoTranslateRequestListResponse struct {
	Data       []VideoTranslateRequest                     `json:"data" api:"required"`
	Pagination VideoTranslateRequestListResponsePagination `json:"pagination" api:"required"`
}

type VideoTranslateRequestListResponsePagination struct {
	Limit  int64 `json:"limit" api:"required"`
	Offset int64 `json:"offset" api:"required"`
	Total  int64 `json:"total" api:"required"`
}

type VideoTranslateRequestListParams struct {
	CreatedAfter   param.Opt[time.Time] `query:"created_after,omitzero" json:"-" format:"date-time"`
	CreatedBefore  param.Opt[time.Time] `query:"created_before,omitzero" json:"-" format:"date-time"`
	Limit          param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	Offset         param.Opt[int64]     `query:"offset,omitzero" json:"-"`
	ReferenceID    param.Opt[string]    `query:"reference_id,omitzero" json:"-"`
	SourceLanguage param.Opt[string]    `query:"source_language,omitzero" json:"-"`
	Status         param.Opt[string]    `query:"status,omitzero" json:"-"`
	TargetLanguage param.Opt[string]    `query:"target_language,omitzero" json:"-"`
	paramObj
}

func (r VideoTranslateRequestListParams) URLQuery() (url.Values, error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat: apiquery.ArrayQueryFormatComma, NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
