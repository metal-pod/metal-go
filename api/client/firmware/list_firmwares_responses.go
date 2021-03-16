// Code generated by go-swagger; DO NOT EDIT.

package firmware

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

// ListFirmwaresReader is a Reader for the ListFirmwares structure.
type ListFirmwaresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFirmwaresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFirmwaresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListFirmwaresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListFirmwaresOK creates a ListFirmwaresOK with default headers values
func NewListFirmwaresOK() *ListFirmwaresOK {
	return &ListFirmwaresOK{}
}

/*ListFirmwaresOK handles this case with default header values.

OK
*/
type ListFirmwaresOK struct {
	Payload []*models.V1Firmwares
}

func (o *ListFirmwaresOK) Error() string {
	return fmt.Sprintf("[GET /v1/firmware][%d] listFirmwaresOK  %+v", 200, o.Payload)
}

func (o *ListFirmwaresOK) GetPayload() []*models.V1Firmwares {
	return o.Payload
}

func (o *ListFirmwaresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFirmwaresDefault creates a ListFirmwaresDefault with default headers values
func NewListFirmwaresDefault(code int) *ListFirmwaresDefault {
	return &ListFirmwaresDefault{
		_statusCode: code,
	}
}

/*ListFirmwaresDefault handles this case with default header values.

Error
*/
type ListFirmwaresDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the list firmwares default response
func (o *ListFirmwaresDefault) Code() int {
	return o._statusCode
}

func (o *ListFirmwaresDefault) Error() string {
	return fmt.Sprintf("[GET /v1/firmware][%d] listFirmwares default  %+v", o._statusCode, o.Payload)
}

func (o *ListFirmwaresDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ListFirmwaresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}