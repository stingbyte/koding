package compute_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIComputeProviderUpdateParams creates a new PostRemoteAPIComputeProviderUpdateParams object
// with the default values initialized.
func NewPostRemoteAPIComputeProviderUpdateParams() *PostRemoteAPIComputeProviderUpdateParams {
	var ()
	return &PostRemoteAPIComputeProviderUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIComputeProviderUpdateParamsWithTimeout creates a new PostRemoteAPIComputeProviderUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIComputeProviderUpdateParamsWithTimeout(timeout time.Duration) *PostRemoteAPIComputeProviderUpdateParams {
	var ()
	return &PostRemoteAPIComputeProviderUpdateParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIComputeProviderUpdateParamsWithContext creates a new PostRemoteAPIComputeProviderUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIComputeProviderUpdateParamsWithContext(ctx context.Context) *PostRemoteAPIComputeProviderUpdateParams {
	var ()
	return &PostRemoteAPIComputeProviderUpdateParams{

		Context: ctx,
	}
}

/*PostRemoteAPIComputeProviderUpdateParams contains all the parameters to send to the API endpoint
for the post remote API compute provider update operation typically these are written to a http.Request
*/
type PostRemoteAPIComputeProviderUpdateParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API compute provider update params
func (o *PostRemoteAPIComputeProviderUpdateParams) WithTimeout(timeout time.Duration) *PostRemoteAPIComputeProviderUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API compute provider update params
func (o *PostRemoteAPIComputeProviderUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API compute provider update params
func (o *PostRemoteAPIComputeProviderUpdateParams) WithContext(ctx context.Context) *PostRemoteAPIComputeProviderUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API compute provider update params
func (o *PostRemoteAPIComputeProviderUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API compute provider update params
func (o *PostRemoteAPIComputeProviderUpdateParams) WithBody(body models.DefaultSelector) *PostRemoteAPIComputeProviderUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API compute provider update params
func (o *PostRemoteAPIComputeProviderUpdateParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIComputeProviderUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
