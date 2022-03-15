// Code generated by go-swagger; DO NOT EDIT.

package health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/intake/gen/models"
)

// HealthCheckReader is a Reader for the HealthCheck structure.
type HealthCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHealthCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHealthCheckUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHealthCheckOK creates a HealthCheckOK with default headers values
func NewHealthCheckOK() *HealthCheckOK {
	return &HealthCheckOK{}
}

/* HealthCheckOK describes a response with status code 200, with default header values.

OK
*/
type HealthCheckOK struct {
	Payload *models.HealthCheckResponse
}

func (o *HealthCheckOK) Error() string {
	return fmt.Sprintf("[GET /healthCheck][%d] healthCheckOK  %+v", 200, o.Payload)
}
func (o *HealthCheckOK) GetPayload() *models.HealthCheckResponse {
	return o.Payload
}

func (o *HealthCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHealthCheckUnauthorized creates a HealthCheckUnauthorized with default headers values
func NewHealthCheckUnauthorized() *HealthCheckUnauthorized {
	return &HealthCheckUnauthorized{}
}

/* HealthCheckUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type HealthCheckUnauthorized struct {
}

func (o *HealthCheckUnauthorized) Error() string {
	return fmt.Sprintf("[GET /healthCheck][%d] healthCheckUnauthorized ", 401)
}

func (o *HealthCheckUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}