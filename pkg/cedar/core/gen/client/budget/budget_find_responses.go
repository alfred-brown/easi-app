// Code generated by go-swagger; DO NOT EDIT.

package budget

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// BudgetFindReader is a Reader for the BudgetFind structure.
type BudgetFindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BudgetFindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBudgetFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBudgetFindBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBudgetFindUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBudgetFindInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBudgetFindOK creates a BudgetFindOK with default headers values
func NewBudgetFindOK() *BudgetFindOK {
	return &BudgetFindOK{}
}

/* BudgetFindOK describes a response with status code 200, with default header values.

OK
*/
type BudgetFindOK struct {
	Payload *models.BudgetFindResponse
}

func (o *BudgetFindOK) Error() string {
	return fmt.Sprintf("[GET /budget][%d] budgetFindOK  %+v", 200, o.Payload)
}
func (o *BudgetFindOK) GetPayload() *models.BudgetFindResponse {
	return o.Payload
}

func (o *BudgetFindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BudgetFindResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBudgetFindBadRequest creates a BudgetFindBadRequest with default headers values
func NewBudgetFindBadRequest() *BudgetFindBadRequest {
	return &BudgetFindBadRequest{}
}

/* BudgetFindBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BudgetFindBadRequest struct {
	Payload *models.Response
}

func (o *BudgetFindBadRequest) Error() string {
	return fmt.Sprintf("[GET /budget][%d] budgetFindBadRequest  %+v", 400, o.Payload)
}
func (o *BudgetFindBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *BudgetFindBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBudgetFindUnauthorized creates a BudgetFindUnauthorized with default headers values
func NewBudgetFindUnauthorized() *BudgetFindUnauthorized {
	return &BudgetFindUnauthorized{}
}

/* BudgetFindUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type BudgetFindUnauthorized struct {
	Payload *models.Response
}

func (o *BudgetFindUnauthorized) Error() string {
	return fmt.Sprintf("[GET /budget][%d] budgetFindUnauthorized  %+v", 401, o.Payload)
}
func (o *BudgetFindUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *BudgetFindUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBudgetFindInternalServerError creates a BudgetFindInternalServerError with default headers values
func NewBudgetFindInternalServerError() *BudgetFindInternalServerError {
	return &BudgetFindInternalServerError{}
}

/* BudgetFindInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BudgetFindInternalServerError struct {
	Payload *models.Response
}

func (o *BudgetFindInternalServerError) Error() string {
	return fmt.Sprintf("[GET /budget][%d] budgetFindInternalServerError  %+v", 500, o.Payload)
}
func (o *BudgetFindInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *BudgetFindInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
