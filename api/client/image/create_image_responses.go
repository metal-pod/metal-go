// Code generated by go-swagger; DO NOT EDIT.

package image

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

// CreateImageReader is a Reader for the CreateImage structure.
type CreateImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateImageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateImageCreated creates a CreateImageCreated with default headers values
func NewCreateImageCreated() *CreateImageCreated {
	return &CreateImageCreated{}
}

/* CreateImageCreated describes a response with status code 201, with default header values.

Created
*/
type CreateImageCreated struct {
	Payload *models.V1ImageResponse
}

func (o *CreateImageCreated) Error() string {
	return fmt.Sprintf("[PUT /v1/image][%d] createImageCreated  %+v", 201, o.Payload)
}
func (o *CreateImageCreated) GetPayload() *models.V1ImageResponse {
	return o.Payload
}

func (o *CreateImageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ImageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImageConflict creates a CreateImageConflict with default headers values
func NewCreateImageConflict() *CreateImageConflict {
	return &CreateImageConflict{}
}

/* CreateImageConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateImageConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

func (o *CreateImageConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/image][%d] createImageConflict  %+v", 409, o.Payload)
}
func (o *CreateImageConflict) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreateImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImageDefault creates a CreateImageDefault with default headers values
func NewCreateImageDefault(code int) *CreateImageDefault {
	return &CreateImageDefault{
		_statusCode: code,
	}
}

/* CreateImageDefault describes a response with status code -1, with default header values.

Error
*/
type CreateImageDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the create image default response
func (o *CreateImageDefault) Code() int {
	return o._statusCode
}

func (o *CreateImageDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/image][%d] createImage default  %+v", o._statusCode, o.Payload)
}
func (o *CreateImageDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreateImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
