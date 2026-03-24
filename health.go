// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chamelaion

import (
	"context"
	"net/http"
	"slices"

	"github.com/chamelaion/chamelaion-go/internal/requestconfig"
	"github.com/chamelaion/chamelaion-go/option"
)

// HealthService contains methods and other services that help with interacting
// with the chamelaion API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.options = opts
	return
}

// Returns 200 if the service is running. No authentication required.
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}
