// Code generated by go-swagger; DO NOT EDIT.

package athletes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

// GetStatsReader is a Reader for the GetStats structure.
type GetStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStatsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStatsOK creates a GetStatsOK with default headers values
func NewGetStatsOK() *GetStatsOK {
	return &GetStatsOK{}
}

/*
GetStatsOK describes a response with status code 200, with default header values.

Activity stats of the athlete.
*/
type GetStatsOK struct {
	Payload *models.ActivityStats
}

// IsSuccess returns true when this get stats o k response has a 2xx status code
func (o *GetStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get stats o k response has a 3xx status code
func (o *GetStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stats o k response has a 4xx status code
func (o *GetStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stats o k response has a 5xx status code
func (o *GetStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get stats o k response a status code equal to that given
func (o *GetStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get stats o k response
func (o *GetStatsOK) Code() int {
	return 200
}

func (o *GetStatsOK) Error() string {
	return fmt.Sprintf("[GET /athletes/{id}/stats][%d] getStatsOK  %+v", 200, o.Payload)
}

func (o *GetStatsOK) String() string {
	return fmt.Sprintf("[GET /athletes/{id}/stats][%d] getStatsOK  %+v", 200, o.Payload)
}

func (o *GetStatsOK) GetPayload() *models.ActivityStats {
	return o.Payload
}

func (o *GetStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ActivityStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatsDefault creates a GetStatsDefault with default headers values
func NewGetStatsDefault(code int) *GetStatsDefault {
	return &GetStatsDefault{
		_statusCode: code,
	}
}

/*
GetStatsDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetStatsDefault struct {
	_statusCode int

	Payload *models.Fault
}

// IsSuccess returns true when this get stats default response has a 2xx status code
func (o *GetStatsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get stats default response has a 3xx status code
func (o *GetStatsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get stats default response has a 4xx status code
func (o *GetStatsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get stats default response has a 5xx status code
func (o *GetStatsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get stats default response a status code equal to that given
func (o *GetStatsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get stats default response
func (o *GetStatsDefault) Code() int {
	return o._statusCode
}

func (o *GetStatsDefault) Error() string {
	return fmt.Sprintf("[GET /athletes/{id}/stats][%d] getStats default  %+v", o._statusCode, o.Payload)
}

func (o *GetStatsDefault) String() string {
	return fmt.Sprintf("[GET /athletes/{id}/stats][%d] getStats default  %+v", o._statusCode, o.Payload)
}

func (o *GetStatsDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetStatsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}