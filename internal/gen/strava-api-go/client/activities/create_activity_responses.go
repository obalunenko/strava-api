// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

// CreateActivityReader is a Reader for the CreateActivity structure.
type CreateActivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateActivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateActivityCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateActivityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateActivityCreated creates a CreateActivityCreated with default headers values
func NewCreateActivityCreated() *CreateActivityCreated {
	return &CreateActivityCreated{}
}

/*
CreateActivityCreated describes a response with status code 201, with default header values.

The activity's detailed representation.
*/
type CreateActivityCreated struct {
	Payload *models.DetailedActivity
}

// IsSuccess returns true when this create activity created response has a 2xx status code
func (o *CreateActivityCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create activity created response has a 3xx status code
func (o *CreateActivityCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create activity created response has a 4xx status code
func (o *CreateActivityCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create activity created response has a 5xx status code
func (o *CreateActivityCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create activity created response a status code equal to that given
func (o *CreateActivityCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create activity created response
func (o *CreateActivityCreated) Code() int {
	return 201
}

func (o *CreateActivityCreated) Error() string {
	return fmt.Sprintf("[POST /activities][%d] createActivityCreated  %+v", 201, o.Payload)
}

func (o *CreateActivityCreated) String() string {
	return fmt.Sprintf("[POST /activities][%d] createActivityCreated  %+v", 201, o.Payload)
}

func (o *CreateActivityCreated) GetPayload() *models.DetailedActivity {
	return o.Payload
}

func (o *CreateActivityCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DetailedActivity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActivityDefault creates a CreateActivityDefault with default headers values
func NewCreateActivityDefault(code int) *CreateActivityDefault {
	return &CreateActivityDefault{
		_statusCode: code,
	}
}

/*
CreateActivityDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type CreateActivityDefault struct {
	_statusCode int

	Payload *models.Fault
}

// IsSuccess returns true when this create activity default response has a 2xx status code
func (o *CreateActivityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create activity default response has a 3xx status code
func (o *CreateActivityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create activity default response has a 4xx status code
func (o *CreateActivityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create activity default response has a 5xx status code
func (o *CreateActivityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create activity default response a status code equal to that given
func (o *CreateActivityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create activity default response
func (o *CreateActivityDefault) Code() int {
	return o._statusCode
}

func (o *CreateActivityDefault) Error() string {
	return fmt.Sprintf("[POST /activities][%d] createActivity default  %+v", o._statusCode, o.Payload)
}

func (o *CreateActivityDefault) String() string {
	return fmt.Sprintf("[POST /activities][%d] createActivity default  %+v", o._statusCode, o.Payload)
}

func (o *CreateActivityDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *CreateActivityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}