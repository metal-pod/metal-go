// Code generated by go-swagger; DO NOT EDIT.

package firmware

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
)

// NewUploadFirmwareParams creates a new UploadFirmwareParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadFirmwareParams() *UploadFirmwareParams {
	return &UploadFirmwareParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadFirmwareParamsWithTimeout creates a new UploadFirmwareParams object
// with the ability to set a timeout on a request.
func NewUploadFirmwareParamsWithTimeout(timeout time.Duration) *UploadFirmwareParams {
	return &UploadFirmwareParams{
		timeout: timeout,
	}
}

// NewUploadFirmwareParamsWithContext creates a new UploadFirmwareParams object
// with the ability to set a context for a request.
func NewUploadFirmwareParamsWithContext(ctx context.Context) *UploadFirmwareParams {
	return &UploadFirmwareParams{
		Context: ctx,
	}
}

// NewUploadFirmwareParamsWithHTTPClient creates a new UploadFirmwareParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadFirmwareParamsWithHTTPClient(client *http.Client) *UploadFirmwareParams {
	return &UploadFirmwareParams{
		HTTPClient: client,
	}
}

/* UploadFirmwareParams contains all the parameters to send to the API endpoint
   for the upload firmware operation.

   Typically these are written to a http.Request.
*/
type UploadFirmwareParams struct {

	/* Board.

	   the board
	*/
	Board string

	/* File.

	   the firmware file
	*/
	File runtime.NamedReadCloser

	/* Kind.

	   the firmware kind [bios|bmc]
	*/
	Kind string

	/* Revision.

	   the firmware revision
	*/
	Revision string

	/* Vendor.

	   the vendor
	*/
	Vendor string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload firmware params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadFirmwareParams) WithDefaults() *UploadFirmwareParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload firmware params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadFirmwareParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload firmware params
func (o *UploadFirmwareParams) WithTimeout(timeout time.Duration) *UploadFirmwareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload firmware params
func (o *UploadFirmwareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload firmware params
func (o *UploadFirmwareParams) WithContext(ctx context.Context) *UploadFirmwareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload firmware params
func (o *UploadFirmwareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload firmware params
func (o *UploadFirmwareParams) WithHTTPClient(client *http.Client) *UploadFirmwareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload firmware params
func (o *UploadFirmwareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBoard adds the board to the upload firmware params
func (o *UploadFirmwareParams) WithBoard(board string) *UploadFirmwareParams {
	o.SetBoard(board)
	return o
}

// SetBoard adds the board to the upload firmware params
func (o *UploadFirmwareParams) SetBoard(board string) {
	o.Board = board
}

// WithFile adds the file to the upload firmware params
func (o *UploadFirmwareParams) WithFile(file runtime.NamedReadCloser) *UploadFirmwareParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the upload firmware params
func (o *UploadFirmwareParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithKind adds the kind to the upload firmware params
func (o *UploadFirmwareParams) WithKind(kind string) *UploadFirmwareParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the upload firmware params
func (o *UploadFirmwareParams) SetKind(kind string) {
	o.Kind = kind
}

// WithRevision adds the revision to the upload firmware params
func (o *UploadFirmwareParams) WithRevision(revision string) *UploadFirmwareParams {
	o.SetRevision(revision)
	return o
}

// SetRevision adds the revision to the upload firmware params
func (o *UploadFirmwareParams) SetRevision(revision string) {
	o.Revision = revision
}

// WithVendor adds the vendor to the upload firmware params
func (o *UploadFirmwareParams) WithVendor(vendor string) *UploadFirmwareParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the upload firmware params
func (o *UploadFirmwareParams) SetVendor(vendor string) {
	o.Vendor = vendor
}

// WriteToRequest writes these params to a swagger request
func (o *UploadFirmwareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param board
	if err := r.SetPathParam("board", o.Board); err != nil {
		return err
	}

	if o.File != nil {

		if o.File != nil {
			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}
		}
	}

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	// path param revision
	if err := r.SetPathParam("revision", o.Revision); err != nil {
		return err
	}

	// path param vendor
	if err := r.SetPathParam("vendor", o.Vendor); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
