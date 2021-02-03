// Code generated by go-swagger; DO NOT EDIT.

package machine

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

// NewMachineDiskParams creates a new MachineDiskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMachineDiskParams() *MachineDiskParams {
	return &MachineDiskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMachineDiskParamsWithTimeout creates a new MachineDiskParams object
// with the ability to set a timeout on a request.
func NewMachineDiskParamsWithTimeout(timeout time.Duration) *MachineDiskParams {
	return &MachineDiskParams{
		timeout: timeout,
	}
}

// NewMachineDiskParamsWithContext creates a new MachineDiskParams object
// with the ability to set a context for a request.
func NewMachineDiskParamsWithContext(ctx context.Context) *MachineDiskParams {
	return &MachineDiskParams{
		Context: ctx,
	}
}

// NewMachineDiskParamsWithHTTPClient creates a new MachineDiskParams object
// with the ability to set a custom HTTPClient for a request.
func NewMachineDiskParamsWithHTTPClient(client *http.Client) *MachineDiskParams {
	return &MachineDiskParams{
		HTTPClient: client,
	}
}

/* MachineDiskParams contains all the parameters to send to the API endpoint
   for the machine disk operation.

   Typically these are written to a http.Request.
*/
type MachineDiskParams struct {

	// Body.
	Body models.V1EmptyBody

	/* ID.

	   identifier of the machine
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the machine disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachineDiskParams) WithDefaults() *MachineDiskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the machine disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachineDiskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the machine disk params
func (o *MachineDiskParams) WithTimeout(timeout time.Duration) *MachineDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machine disk params
func (o *MachineDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machine disk params
func (o *MachineDiskParams) WithContext(ctx context.Context) *MachineDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machine disk params
func (o *MachineDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machine disk params
func (o *MachineDiskParams) WithHTTPClient(client *http.Client) *MachineDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machine disk params
func (o *MachineDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the machine disk params
func (o *MachineDiskParams) WithBody(body models.V1EmptyBody) *MachineDiskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the machine disk params
func (o *MachineDiskParams) SetBody(body models.V1EmptyBody) {
	o.Body = body
}

// WithID adds the id to the machine disk params
func (o *MachineDiskParams) WithID(id string) *MachineDiskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the machine disk params
func (o *MachineDiskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MachineDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
