// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-stack/metal-go/api/models"
)

// UpdateSwitchReader is a Reader for the UpdateSwitch structure.
type UpdateSwitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSwitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSwitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewUpdateSwitchConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateSwitchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSwitchOK creates a UpdateSwitchOK with default headers values
func NewUpdateSwitchOK() *UpdateSwitchOK {
	return &UpdateSwitchOK{}
}

/*UpdateSwitchOK handles this case with default header values.

OK
*/
type UpdateSwitchOK struct {
	Payload *models.V1SwitchResponse
}

func (o *UpdateSwitchOK) Error() string {
	return fmt.Sprintf("[POST /v1/switch][%d] updateSwitchOK  %+v", 200, o.Payload)
}

func (o *UpdateSwitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SwitchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSwitchConflict creates a UpdateSwitchConflict with default headers values
func NewUpdateSwitchConflict() *UpdateSwitchConflict {
	return &UpdateSwitchConflict{}
}

/*UpdateSwitchConflict handles this case with default header values.

Conflict
*/
type UpdateSwitchConflict struct {
	Payload *models.HttperrorsHTTPErrorResponse
}

func (o *UpdateSwitchConflict) Error() string {
	return fmt.Sprintf("[POST /v1/switch][%d] updateSwitchConflict  %+v", 409, o.Payload)
}

func (o *UpdateSwitchConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSwitchDefault creates a UpdateSwitchDefault with default headers values
func NewUpdateSwitchDefault(code int) *UpdateSwitchDefault {
	return &UpdateSwitchDefault{
		_statusCode: code,
	}
}

/*UpdateSwitchDefault handles this case with default header values.

Error
*/
type UpdateSwitchDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the update switch default response
func (o *UpdateSwitchDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSwitchDefault) Error() string {
	return fmt.Sprintf("[POST /v1/switch][%d] updateSwitch default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSwitchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}