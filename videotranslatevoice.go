// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chamelaion

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/chamelaion/chamelaion-go/internal/apiquery"
	"github.com/chamelaion/chamelaion-go/internal/requestconfig"
	"github.com/chamelaion/chamelaion-go/option"
	"github.com/chamelaion/chamelaion-go/packages/param"
)

type VideoTranslateVoiceService struct{ options []option.RequestOption }

func NewVideoTranslateVoiceService(opts ...option.RequestOption) (r VideoTranslateVoiceService) {
	r.options = opts
	return
}

func (r *VideoTranslateVoiceService) List(ctx context.Context, query VideoTranslateVoiceListParams, opts ...option.RequestOption) (res *VideoTranslateVoiceListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "v1/video-translate/voices", query, &res, opts...)
	return res, err
}

type VideoTranslateVoice struct {
	VoiceID     string `json:"voice_id" api:"required"`
	Name        string `json:"name" api:"required"`
	DefaultName string `json:"default_name" api:"required"`
	Description string `json:"description" api:"required"`
	Language    string `json:"language" api:"required"`
	Application string `json:"application" api:"required"`
	Gender      string `json:"gender" api:"required"`
	IsFavorite  bool   `json:"is_favorite" api:"required"`
}

type VideoTranslateVoiceListResponse struct {
	Voices     []VideoTranslateVoice                     `json:"voices" api:"required"`
	Pagination VideoTranslateVoiceListResponsePagination `json:"pagination" api:"required"`
}

type VideoTranslateVoiceListResponsePagination struct {
	Limit  int64 `json:"limit" api:"required"`
	Offset int64 `json:"offset" api:"required"`
	Total  int64 `json:"total" api:"required"`
}

type VideoTranslateVoiceListParams struct {
	Application param.Opt[string] `query:"application,omitzero" json:"-"`
	Favorite    param.Opt[bool]   `query:"favorite,omitzero" json:"-"`
	Gender      param.Opt[string] `query:"gender,omitzero" json:"-"`
	Language    param.Opt[string] `query:"language,omitzero" json:"-"`
	Limit       param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Offset      param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	Search      param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

func (r VideoTranslateVoiceListParams) URLQuery() (url.Values, error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat: apiquery.ArrayQueryFormatComma, NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
