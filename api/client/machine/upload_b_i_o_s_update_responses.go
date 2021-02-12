// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-lib/httperrors"
)

// UploadBIOSUpdateReader is a Reader for the UploadBIOSUpdate structure.
type UploadBIOSUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadBIOSUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadBIOSUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUploadBIOSUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadBIOSUpdateOK creates a UploadBIOSUpdateOK with default headers values
func NewUploadBIOSUpdateOK() *UploadBIOSUpdateOK {
	return &UploadBIOSUpdateOK{}
}

/*UploadBIOSUpdateOK handles this case with default header values.

OK
*/
type UploadBIOSUpdateOK struct {
}

func (o *UploadBIOSUpdateOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/upload/bios/{vendor}/{board}/{revision}][%d] uploadBIOSUpdateOK ", 200)
}

func (o *UploadBIOSUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadBIOSUpdateDefault creates a UploadBIOSUpdateDefault with default headers values
func NewUploadBIOSUpdateDefault(code int) *UploadBIOSUpdateDefault {
	return &UploadBIOSUpdateDefault{
		_statusCode: code,
	}
}

/*UploadBIOSUpdateDefault handles this case with default header values.

Error
*/
type UploadBIOSUpdateDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the upload b i o s update default response
func (o *UploadBIOSUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UploadBIOSUpdateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/upload/bios/{vendor}/{board}/{revision}][%d] uploadBIOSUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *UploadBIOSUpdateDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UploadBIOSUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
