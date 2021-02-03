// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
)

// NewAllocateNetworkParams creates a new AllocateNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllocateNetworkParams() *AllocateNetworkParams {
	return &AllocateNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllocateNetworkParamsWithTimeout creates a new AllocateNetworkParams object
// with the ability to set a timeout on a request.
func NewAllocateNetworkParamsWithTimeout(timeout time.Duration) *AllocateNetworkParams {
	return &AllocateNetworkParams{
		timeout: timeout,
	}
}

// NewAllocateNetworkParamsWithContext creates a new AllocateNetworkParams object
// with the ability to set a context for a request.
func NewAllocateNetworkParamsWithContext(ctx context.Context) *AllocateNetworkParams {
	return &AllocateNetworkParams{
		Context: ctx,
	}
}

// NewAllocateNetworkParamsWithHTTPClient creates a new AllocateNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllocateNetworkParamsWithHTTPClient(client *http.Client) *AllocateNetworkParams {
	return &AllocateNetworkParams{
		HTTPClient: client,
	}
}

/* AllocateNetworkParams contains all the parameters to send to the API endpoint
   for the allocate network operation.

   Typically these are written to a http.Request.
*/
type AllocateNetworkParams struct {

	// Body.
	Body *models.V1NetworkAllocateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the allocate network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocateNetworkParams) WithDefaults() *AllocateNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the allocate network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllocateNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the allocate network params
func (o *AllocateNetworkParams) WithTimeout(timeout time.Duration) *AllocateNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allocate network params
func (o *AllocateNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allocate network params
func (o *AllocateNetworkParams) WithContext(ctx context.Context) *AllocateNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allocate network params
func (o *AllocateNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allocate network params
func (o *AllocateNetworkParams) WithHTTPClient(client *http.Client) *AllocateNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allocate network params
func (o *AllocateNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the allocate network params
func (o *AllocateNetworkParams) WithBody(body *models.V1NetworkAllocateRequest) *AllocateNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the allocate network params
func (o *AllocateNetworkParams) SetBody(body *models.V1NetworkAllocateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AllocateNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
