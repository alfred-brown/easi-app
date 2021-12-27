// Code generated by go-swagger; DO NOT EDIT.

package enumeration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// EnumerationFindListReader is a Reader for the EnumerationFindList structure.
type EnumerationFindListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumerationFindListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnumerationFindListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnumerationFindListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEnumerationFindListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnumerationFindListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnumerationFindListOK creates a EnumerationFindListOK with default headers values
func NewEnumerationFindListOK() *EnumerationFindListOK {
	return &EnumerationFindListOK{}
}

/* EnumerationFindListOK describes a response with status code 200, with default header values.

OK
*/
type EnumerationFindListOK struct {
	Payload *models.EnumerationFindResponse
}

func (o *EnumerationFindListOK) Error() string {
	return fmt.Sprintf("[GET /enumeration][%d] enumerationFindListOK  %+v", 200, o.Payload)
}
func (o *EnumerationFindListOK) GetPayload() *models.EnumerationFindResponse {
	return o.Payload
}

func (o *EnumerationFindListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnumerationFindResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumerationFindListBadRequest creates a EnumerationFindListBadRequest with default headers values
func NewEnumerationFindListBadRequest() *EnumerationFindListBadRequest {
	return &EnumerationFindListBadRequest{}
}

/* EnumerationFindListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EnumerationFindListBadRequest struct {
	Payload *models.Response
}

func (o *EnumerationFindListBadRequest) Error() string {
	return fmt.Sprintf("[GET /enumeration][%d] enumerationFindListBadRequest  %+v", 400, o.Payload)
}
func (o *EnumerationFindListBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *EnumerationFindListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumerationFindListUnauthorized creates a EnumerationFindListUnauthorized with default headers values
func NewEnumerationFindListUnauthorized() *EnumerationFindListUnauthorized {
	return &EnumerationFindListUnauthorized{}
}

/* EnumerationFindListUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type EnumerationFindListUnauthorized struct {
	Payload *models.Response
}

func (o *EnumerationFindListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /enumeration][%d] enumerationFindListUnauthorized  %+v", 401, o.Payload)
}
func (o *EnumerationFindListUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *EnumerationFindListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumerationFindListInternalServerError creates a EnumerationFindListInternalServerError with default headers values
func NewEnumerationFindListInternalServerError() *EnumerationFindListInternalServerError {
	return &EnumerationFindListInternalServerError{}
}

/* EnumerationFindListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EnumerationFindListInternalServerError struct {
	Payload *models.Response
}

func (o *EnumerationFindListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /enumeration][%d] enumerationFindListInternalServerError  %+v", 500, o.Payload)
}
func (o *EnumerationFindListInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *EnumerationFindListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}