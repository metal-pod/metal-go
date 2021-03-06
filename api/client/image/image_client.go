// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new image API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for image API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateImage(params *CreateImageParams, authInfo runtime.ClientAuthInfoWriter) (*CreateImageCreated, error)

	DeleteImage(params *DeleteImageParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteImageOK, error)

	FindImage(params *FindImageParams, authInfo runtime.ClientAuthInfoWriter) (*FindImageOK, error)

	FindLatestImage(params *FindLatestImageParams, authInfo runtime.ClientAuthInfoWriter) (*FindLatestImageOK, error)

	ListImages(params *ListImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ListImagesOK, error)

	UpdateImage(params *UpdateImageParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateImageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateImage creates an image if the given ID already exists a conflict is returned
*/
func (a *Client) CreateImage(params *CreateImageParams, authInfo runtime.ClientAuthInfoWriter) (*CreateImageCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createImage",
		Method:             "PUT",
		PathPattern:        "/v1/image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateImageCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteImage deletes an image and returns the deleted entity
*/
func (a *Client) DeleteImage(params *DeleteImageParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteImage",
		Method:             "DELETE",
		PathPattern:        "/v1/image/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindImage gets image by id
*/
func (a *Client) FindImage(params *FindImageParams, authInfo runtime.ClientAuthInfoWriter) (*FindImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findImage",
		Method:             "GET",
		PathPattern:        "/v1/image/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindLatestImage finds latest image by id
*/
func (a *Client) FindLatestImage(params *FindLatestImageParams, authInfo runtime.ClientAuthInfoWriter) (*FindLatestImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindLatestImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findLatestImage",
		Method:             "GET",
		PathPattern:        "/v1/image/{id}/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindLatestImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindLatestImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindLatestImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListImages gets all images
*/
func (a *Client) ListImages(params *ListImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ListImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listImages",
		Method:             "GET",
		PathPattern:        "/v1/image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateImage updates an image if the image was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateImage(params *UpdateImageParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateImage",
		Method:             "POST",
		PathPattern:        "/v1/image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
