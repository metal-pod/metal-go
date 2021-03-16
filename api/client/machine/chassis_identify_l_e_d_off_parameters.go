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

// NewChassisIdentifyLEDOffParams creates a new ChassisIdentifyLEDOffParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChassisIdentifyLEDOffParams() *ChassisIdentifyLEDOffParams {
	return &ChassisIdentifyLEDOffParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChassisIdentifyLEDOffParamsWithTimeout creates a new ChassisIdentifyLEDOffParams object
// with the ability to set a timeout on a request.
func NewChassisIdentifyLEDOffParamsWithTimeout(timeout time.Duration) *ChassisIdentifyLEDOffParams {
	return &ChassisIdentifyLEDOffParams{
		timeout: timeout,
	}
}

// NewChassisIdentifyLEDOffParamsWithContext creates a new ChassisIdentifyLEDOffParams object
// with the ability to set a context for a request.
func NewChassisIdentifyLEDOffParamsWithContext(ctx context.Context) *ChassisIdentifyLEDOffParams {
	return &ChassisIdentifyLEDOffParams{
		Context: ctx,
	}
}

// NewChassisIdentifyLEDOffParamsWithHTTPClient creates a new ChassisIdentifyLEDOffParams object
// with the ability to set a custom HTTPClient for a request.
func NewChassisIdentifyLEDOffParamsWithHTTPClient(client *http.Client) *ChassisIdentifyLEDOffParams {
	return &ChassisIdentifyLEDOffParams{
		HTTPClient: client,
	}
}

/* ChassisIdentifyLEDOffParams contains all the parameters to send to the API endpoint
   for the chassis identify l e d off operation.

   Typically these are written to a http.Request.
*/
type ChassisIdentifyLEDOffParams struct {

	// Body.
	Body models.V1EmptyBody

	/* Description.

	   reason why the chassis identify LED has been turned off
	*/
	Description *string

	/* ID.

	   identifier of the machine
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chassis identify l e d off params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChassisIdentifyLEDOffParams) WithDefaults() *ChassisIdentifyLEDOffParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chassis identify l e d off params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChassisIdentifyLEDOffParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) WithTimeout(timeout time.Duration) *ChassisIdentifyLEDOffParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) WithContext(ctx context.Context) *ChassisIdentifyLEDOffParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) WithHTTPClient(client *http.Client) *ChassisIdentifyLEDOffParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) WithBody(body models.V1EmptyBody) *ChassisIdentifyLEDOffParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) SetBody(body models.V1EmptyBody) {
	o.Body = body
}

// WithDescription adds the description to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) WithDescription(description *string) *ChassisIdentifyLEDOffParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) SetDescription(description *string) {
	o.Description = description
}

// WithID adds the id to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) WithID(id string) *ChassisIdentifyLEDOffParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the chassis identify l e d off params
func (o *ChassisIdentifyLEDOffParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChassisIdentifyLEDOffParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
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
