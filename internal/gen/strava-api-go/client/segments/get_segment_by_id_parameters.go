// Code generated by go-swagger; DO NOT EDIT.

package segments

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
	"github.com/go-openapi/swag"
)

// NewGetSegmentByIDParams creates a new GetSegmentByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSegmentByIDParams() *GetSegmentByIDParams {
	return &GetSegmentByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSegmentByIDParamsWithTimeout creates a new GetSegmentByIDParams object
// with the ability to set a timeout on a request.
func NewGetSegmentByIDParamsWithTimeout(timeout time.Duration) *GetSegmentByIDParams {
	return &GetSegmentByIDParams{
		timeout: timeout,
	}
}

// NewGetSegmentByIDParamsWithContext creates a new GetSegmentByIDParams object
// with the ability to set a context for a request.
func NewGetSegmentByIDParamsWithContext(ctx context.Context) *GetSegmentByIDParams {
	return &GetSegmentByIDParams{
		Context: ctx,
	}
}

// NewGetSegmentByIDParamsWithHTTPClient creates a new GetSegmentByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSegmentByIDParamsWithHTTPClient(client *http.Client) *GetSegmentByIDParams {
	return &GetSegmentByIDParams{
		HTTPClient: client,
	}
}

/*
GetSegmentByIDParams contains all the parameters to send to the API endpoint

	for the get segment by Id operation.

	Typically these are written to a http.Request.
*/
type GetSegmentByIDParams struct {
	/* ID.

	   The identifier of the segment.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get segment by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSegmentByIDParams) WithDefaults() *GetSegmentByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get segment by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSegmentByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get segment by Id params
func (o *GetSegmentByIDParams) WithTimeout(timeout time.Duration) *GetSegmentByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get segment by Id params
func (o *GetSegmentByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get segment by Id params
func (o *GetSegmentByIDParams) WithContext(ctx context.Context) *GetSegmentByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get segment by Id params
func (o *GetSegmentByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get segment by Id params
func (o *GetSegmentByIDParams) WithHTTPClient(client *http.Client) *GetSegmentByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get segment by Id params
func (o *GetSegmentByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get segment by Id params
func (o *GetSegmentByIDParams) WithID(id int64) *GetSegmentByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get segment by Id params
func (o *GetSegmentByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSegmentByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}