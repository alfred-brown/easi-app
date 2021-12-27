// Code generated by go-swagger; DO NOT EDIT.

package support_contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// SupportContactAddReader is a Reader for the SupportContactAdd structure.
type SupportContactAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupportContactAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupportContactAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSupportContactAddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSupportContactAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSupportContactAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSupportContactAddOK creates a SupportContactAddOK with default headers values
func NewSupportContactAddOK() *SupportContactAddOK {
	return &SupportContactAddOK{}
}

/* SupportContactAddOK describes a response with status code 200, with default header values.

OK
*/
type SupportContactAddOK struct {
	Payload *models.Response
}

func (o *SupportContactAddOK) Error() string {
	return fmt.Sprintf("[POST /supportContact][%d] supportContactAddOK  %+v", 200, o.Payload)
}
func (o *SupportContactAddOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *SupportContactAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSupportContactAddBadRequest creates a SupportContactAddBadRequest with default headers values
func NewSupportContactAddBadRequest() *SupportContactAddBadRequest {
	return &SupportContactAddBadRequest{}
}

/* SupportContactAddBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SupportContactAddBadRequest struct {
	Payload *models.Response
}

func (o *SupportContactAddBadRequest) Error() string {
	return fmt.Sprintf("[POST /supportContact][%d] supportContactAddBadRequest  %+v", 400, o.Payload)
}
func (o *SupportContactAddBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *SupportContactAddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSupportContactAddUnauthorized creates a SupportContactAddUnauthorized with default headers values
func NewSupportContactAddUnauthorized() *SupportContactAddUnauthorized {
	return &SupportContactAddUnauthorized{}
}

/* SupportContactAddUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type SupportContactAddUnauthorized struct {
	Payload *models.Response
}

func (o *SupportContactAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /supportContact][%d] supportContactAddUnauthorized  %+v", 401, o.Payload)
}
func (o *SupportContactAddUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *SupportContactAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSupportContactAddInternalServerError creates a SupportContactAddInternalServerError with default headers values
func NewSupportContactAddInternalServerError() *SupportContactAddInternalServerError {
	return &SupportContactAddInternalServerError{}
}

/* SupportContactAddInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SupportContactAddInternalServerError struct {
	Payload *models.Response
}

func (o *SupportContactAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /supportContact][%d] supportContactAddInternalServerError  %+v", 500, o.Payload)
}
func (o *SupportContactAddInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *SupportContactAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}