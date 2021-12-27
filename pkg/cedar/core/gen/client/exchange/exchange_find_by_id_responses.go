// Code generated by go-swagger; DO NOT EDIT.

package exchange

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// ExchangeFindByIDReader is a Reader for the ExchangeFindByID structure.
type ExchangeFindByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExchangeFindByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExchangeFindByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExchangeFindByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExchangeFindByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExchangeFindByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExchangeFindByIDOK creates a ExchangeFindByIDOK with default headers values
func NewExchangeFindByIDOK() *ExchangeFindByIDOK {
	return &ExchangeFindByIDOK{}
}

/* ExchangeFindByIDOK describes a response with status code 200, with default header values.

OK
*/
type ExchangeFindByIDOK struct {
	Payload *models.Exchange
}

func (o *ExchangeFindByIDOK) Error() string {
	return fmt.Sprintf("[GET /exchange/{id}][%d] exchangeFindByIdOK  %+v", 200, o.Payload)
}
func (o *ExchangeFindByIDOK) GetPayload() *models.Exchange {
	return o.Payload
}

func (o *ExchangeFindByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Exchange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeFindByIDBadRequest creates a ExchangeFindByIDBadRequest with default headers values
func NewExchangeFindByIDBadRequest() *ExchangeFindByIDBadRequest {
	return &ExchangeFindByIDBadRequest{}
}

/* ExchangeFindByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExchangeFindByIDBadRequest struct {
	Payload *models.Response
}

func (o *ExchangeFindByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /exchange/{id}][%d] exchangeFindByIdBadRequest  %+v", 400, o.Payload)
}
func (o *ExchangeFindByIDBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *ExchangeFindByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeFindByIDUnauthorized creates a ExchangeFindByIDUnauthorized with default headers values
func NewExchangeFindByIDUnauthorized() *ExchangeFindByIDUnauthorized {
	return &ExchangeFindByIDUnauthorized{}
}

/* ExchangeFindByIDUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type ExchangeFindByIDUnauthorized struct {
	Payload *models.Response
}

func (o *ExchangeFindByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /exchange/{id}][%d] exchangeFindByIdUnauthorized  %+v", 401, o.Payload)
}
func (o *ExchangeFindByIDUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *ExchangeFindByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeFindByIDInternalServerError creates a ExchangeFindByIDInternalServerError with default headers values
func NewExchangeFindByIDInternalServerError() *ExchangeFindByIDInternalServerError {
	return &ExchangeFindByIDInternalServerError{}
}

/* ExchangeFindByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExchangeFindByIDInternalServerError struct {
	Payload *models.Response
}

func (o *ExchangeFindByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /exchange/{id}][%d] exchangeFindByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *ExchangeFindByIDInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *ExchangeFindByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}