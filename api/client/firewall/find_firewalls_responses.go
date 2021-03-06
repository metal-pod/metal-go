// Code generated by go-swagger; DO NOT EDIT.

package firewall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// FindFirewallsReader is a Reader for the FindFirewalls structure.
type FindFirewallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindFirewallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindFirewallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindFirewallsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindFirewallsOK creates a FindFirewallsOK with default headers values
func NewFindFirewallsOK() *FindFirewallsOK {
	return &FindFirewallsOK{}
}

/* FindFirewallsOK describes a response with status code 200, with default header values.

OK
*/
type FindFirewallsOK struct {
	Payload []*models.V1FirewallResponse
}

func (o *FindFirewallsOK) Error() string {
	return fmt.Sprintf("[POST /v1/firewall/find][%d] findFirewallsOK  %+v", 200, o.Payload)
}
func (o *FindFirewallsOK) GetPayload() []*models.V1FirewallResponse {
	return o.Payload
}

func (o *FindFirewallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindFirewallsDefault creates a FindFirewallsDefault with default headers values
func NewFindFirewallsDefault(code int) *FindFirewallsDefault {
	return &FindFirewallsDefault{
		_statusCode: code,
	}
}

/* FindFirewallsDefault describes a response with status code -1, with default header values.

Error
*/
type FindFirewallsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find firewalls default response
func (o *FindFirewallsDefault) Code() int {
	return o._statusCode
}

func (o *FindFirewallsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/firewall/find][%d] findFirewalls default  %+v", o._statusCode, o.Payload)
}
func (o *FindFirewallsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindFirewallsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
