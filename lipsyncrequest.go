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

	"github.com/stainless-sdks/chamelaion-go/internal/apijson"
	"github.com/stainless-sdks/chamelaion-go/internal/apiquery"
	"github.com/stainless-sdks/chamelaion-go/internal/requestconfig"
	"github.com/stainless-sdks/chamelaion-go/option"
	"github.com/stainless-sdks/chamelaion-go/packages/param"
	"github.com/stainless-sdks/chamelaion-go/packages/respjson"
)

// LipsyncRequestService contains methods and other services that help with
// interacting with the chamelaion API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLipsyncRequestService] method instead.
type LipsyncRequestService struct {
	options []option.RequestOption
}

// NewLipsyncRequestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLipsyncRequestService(opts ...option.RequestOption) (r LipsyncRequestService) {
	r = LipsyncRequestService{}
	r.options = opts
	return
}

// Returns a single lipsync request by ID or reference ID.
func (r *LipsyncRequestService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LipsyncRequest, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/lipsync/requests/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of lipsync requests for the authenticated account.
func (r *LipsyncRequestService) List(ctx context.Context, query LipsyncRequestListParams, opts ...option.RequestOption) (res *LipsyncRequestListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/lipsync/requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type LipsyncRequest struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	CreatedAt    time.Time `json:"created_at" api:"required" format:"date-time"`
	Status       string    `json:"status" api:"required"`
	ErrorMessage string    `json:"error_message"`
	FinishedAt   time.Time `json:"finished_at" format:"date-time"`
	// Time-limited signed download URL for the generated output file.
	OutputURL   string    `json:"output_url" format:"uri"`
	ReferenceID string    `json:"reference_id"`
	StartedAt   time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Status       respjson.Field
		ErrorMessage respjson.Field
		FinishedAt   respjson.Field
		OutputURL    respjson.Field
		ReferenceID  respjson.Field
		StartedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LipsyncRequest) RawJSON() string { return r.JSON.raw }
func (r *LipsyncRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LipsyncRequestListResponse struct {
	Data       []LipsyncRequest                     `json:"data" api:"required"`
	Pagination LipsyncRequestListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LipsyncRequestListResponse) RawJSON() string { return r.JSON.raw }
func (r *LipsyncRequestListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LipsyncRequestListResponsePagination struct {
	Limit  int64 `json:"limit" api:"required"`
	Offset int64 `json:"offset" api:"required"`
	Total  int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LipsyncRequestListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *LipsyncRequestListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LipsyncRequestListParams struct {
	// Maximum number of results to return (default 20)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of results to skip (default 0)
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by reference ID (exact match)
	ReferenceID param.Opt[string] `query:"reference_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LipsyncRequestListParams]'s query parameters as
// `url.Values`.
func (r LipsyncRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
