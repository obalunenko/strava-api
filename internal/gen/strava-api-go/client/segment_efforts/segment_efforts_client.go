// Code generated by go-swagger; DO NOT EDIT.

package segment_efforts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new segment efforts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for segment efforts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetEffortsBySegmentID(params *GetEffortsBySegmentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEffortsBySegmentIDOK, error)

	GetSegmentEffortByID(params *GetSegmentEffortByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSegmentEffortByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetEffortsBySegmentID lists segment efforts

Returns a set of the authenticated athlete's segment efforts for a given segment.  Requires subscription.
*/
func (a *Client) GetEffortsBySegmentID(params *GetEffortsBySegmentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEffortsBySegmentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEffortsBySegmentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEffortsBySegmentId",
		Method:             "GET",
		PathPattern:        "/segment_efforts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEffortsBySegmentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEffortsBySegmentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEffortsBySegmentIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSegmentEffortByID gets segment effort

Returns a segment effort from an activity that is owned by the authenticated athlete. Requires subscription.
*/
func (a *Client) GetSegmentEffortByID(params *GetSegmentEffortByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSegmentEffortByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSegmentEffortByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSegmentEffortById",
		Method:             "GET",
		PathPattern:        "/segment_efforts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSegmentEffortByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSegmentEffortByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSegmentEffortByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}