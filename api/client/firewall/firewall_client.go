// Code generated by go-swagger; DO NOT EDIT.

package firewall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new firewall API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for firewall API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AllocateFirewall allocates a firewall
*/
func (a *Client) AllocateFirewall(params *AllocateFirewallParams, authInfo runtime.ClientAuthInfoWriter) (*AllocateFirewallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocateFirewallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "allocateFirewall",
		Method:             "POST",
		PathPattern:        "/v1/firewall/allocate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocateFirewallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllocateFirewallOK), nil

}

/*
FindFirewall gets firewall by id
*/
func (a *Client) FindFirewall(params *FindFirewallParams, authInfo runtime.ClientAuthInfoWriter) (*FindFirewallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindFirewallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findFirewall",
		Method:             "GET",
		PathPattern:        "/v1/firewall/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindFirewallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindFirewallOK), nil

}

/*
FindFirewalls finds firewalls by multiple criteria
*/
func (a *Client) FindFirewalls(params *FindFirewallsParams, authInfo runtime.ClientAuthInfoWriter) (*FindFirewallsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindFirewallsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findFirewalls",
		Method:             "POST",
		PathPattern:        "/v1/firewall/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindFirewallsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindFirewallsOK), nil

}

/*
ListFirewalls gets all known firewalls
*/
func (a *Client) ListFirewalls(params *ListFirewallsParams, authInfo runtime.ClientAuthInfoWriter) (*ListFirewallsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFirewallsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listFirewalls",
		Method:             "GET",
		PathPattern:        "/v1/firewall",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListFirewallsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListFirewallsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
