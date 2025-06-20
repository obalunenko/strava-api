// Code generated by go-swagger; DO NOT EDIT.

package streams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

// GetRouteStreamsReader is a Reader for the GetRouteStreams structure.
type GetRouteStreamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRouteStreamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRouteStreamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRouteStreamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRouteStreamsOK creates a GetRouteStreamsOK with default headers values
func NewGetRouteStreamsOK() *GetRouteStreamsOK {
	return &GetRouteStreamsOK{}
}

/*
GetRouteStreamsOK describes a response with status code 200, with default header values.

The set of requested streams.
*/
type GetRouteStreamsOK struct {
	Payload *models.StreamSet
}

// IsSuccess returns true when this get route streams o k response has a 2xx status code
func (o *GetRouteStreamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get route streams o k response has a 3xx status code
func (o *GetRouteStreamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get route streams o k response has a 4xx status code
func (o *GetRouteStreamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get route streams o k response has a 5xx status code
func (o *GetRouteStreamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get route streams o k response a status code equal to that given
func (o *GetRouteStreamsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get route streams o k response
func (o *GetRouteStreamsOK) Code() int {
	return 200
}

func (o *GetRouteStreamsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routes/{id}/streams][%d] getRouteStreamsOK %s", 200, payload)
}

func (o *GetRouteStreamsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routes/{id}/streams][%d] getRouteStreamsOK %s", 200, payload)
}

func (o *GetRouteStreamsOK) GetPayload() *models.StreamSet {
	return o.Payload
}

func (o *GetRouteStreamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StreamSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRouteStreamsDefault creates a GetRouteStreamsDefault with default headers values
func NewGetRouteStreamsDefault(code int) *GetRouteStreamsDefault {
	return &GetRouteStreamsDefault{
		_statusCode: code,
	}
}

/*
GetRouteStreamsDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetRouteStreamsDefault struct {
	_statusCode int

	Payload *models.Fault
}

// IsSuccess returns true when this get route streams default response has a 2xx status code
func (o *GetRouteStreamsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get route streams default response has a 3xx status code
func (o *GetRouteStreamsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get route streams default response has a 4xx status code
func (o *GetRouteStreamsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get route streams default response has a 5xx status code
func (o *GetRouteStreamsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get route streams default response a status code equal to that given
func (o *GetRouteStreamsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get route streams default response
func (o *GetRouteStreamsDefault) Code() int {
	return o._statusCode
}

func (o *GetRouteStreamsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routes/{id}/streams][%d] getRouteStreams default %s", o._statusCode, payload)
}

func (o *GetRouteStreamsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routes/{id}/streams][%d] getRouteStreams default %s", o._statusCode, payload)
}

func (o *GetRouteStreamsDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetRouteStreamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
