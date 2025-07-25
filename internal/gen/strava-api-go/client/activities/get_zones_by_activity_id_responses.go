// Code generated by go-swagger; DO NOT EDIT.

package activities

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

// GetZonesByActivityIDReader is a Reader for the GetZonesByActivityID structure.
type GetZonesByActivityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZonesByActivityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZonesByActivityIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetZonesByActivityIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetZonesByActivityIDOK creates a GetZonesByActivityIDOK with default headers values
func NewGetZonesByActivityIDOK() *GetZonesByActivityIDOK {
	return &GetZonesByActivityIDOK{}
}

/*
GetZonesByActivityIDOK describes a response with status code 200, with default header values.

Activity Zones.
*/
type GetZonesByActivityIDOK struct {
	Payload []*models.ActivityZone
}

// IsSuccess returns true when this get zones by activity Id o k response has a 2xx status code
func (o *GetZonesByActivityIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get zones by activity Id o k response has a 3xx status code
func (o *GetZonesByActivityIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get zones by activity Id o k response has a 4xx status code
func (o *GetZonesByActivityIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get zones by activity Id o k response has a 5xx status code
func (o *GetZonesByActivityIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get zones by activity Id o k response a status code equal to that given
func (o *GetZonesByActivityIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get zones by activity Id o k response
func (o *GetZonesByActivityIDOK) Code() int {
	return 200
}

func (o *GetZonesByActivityIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /activities/{id}/zones][%d] getZonesByActivityIdOK %s", 200, payload)
}

func (o *GetZonesByActivityIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /activities/{id}/zones][%d] getZonesByActivityIdOK %s", 200, payload)
}

func (o *GetZonesByActivityIDOK) GetPayload() []*models.ActivityZone {
	return o.Payload
}

func (o *GetZonesByActivityIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZonesByActivityIDDefault creates a GetZonesByActivityIDDefault with default headers values
func NewGetZonesByActivityIDDefault(code int) *GetZonesByActivityIDDefault {
	return &GetZonesByActivityIDDefault{
		_statusCode: code,
	}
}

/*
GetZonesByActivityIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetZonesByActivityIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// IsSuccess returns true when this get zones by activity Id default response has a 2xx status code
func (o *GetZonesByActivityIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get zones by activity Id default response has a 3xx status code
func (o *GetZonesByActivityIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get zones by activity Id default response has a 4xx status code
func (o *GetZonesByActivityIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get zones by activity Id default response has a 5xx status code
func (o *GetZonesByActivityIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get zones by activity Id default response a status code equal to that given
func (o *GetZonesByActivityIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get zones by activity Id default response
func (o *GetZonesByActivityIDDefault) Code() int {
	return o._statusCode
}

func (o *GetZonesByActivityIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /activities/{id}/zones][%d] getZonesByActivityId default %s", o._statusCode, payload)
}

func (o *GetZonesByActivityIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /activities/{id}/zones][%d] getZonesByActivityId default %s", o._statusCode, payload)
}

func (o *GetZonesByActivityIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetZonesByActivityIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
