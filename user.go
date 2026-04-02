// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chamelaion

import (
	"context"
	"net/http"
	"slices"

	"github.com/chamelaion/chamelaion-go/internal/apijson"
	"github.com/chamelaion/chamelaion-go/internal/requestconfig"
	"github.com/chamelaion/chamelaion-go/option"
	"github.com/chamelaion/chamelaion-go/packages/respjson"
)

// Endpoints for API token identity.
//
// UserService contains methods and other services that help with interacting with
// the chamelaion API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.options = opts
	return
}

// Returns the identity associated with the authenticated API token.
func (r *UserService) GetCurrent(ctx context.Context, opts ...option.RequestOption) (res *UserGetCurrentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UserGetCurrentResponse struct {
	// Display name of the API token.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetCurrentResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetCurrentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
