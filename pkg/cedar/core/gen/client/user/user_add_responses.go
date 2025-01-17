// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// UserAddReader is a Reader for the UserAdd structure.
type UserAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserAddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserAddOK creates a UserAddOK with default headers values
func NewUserAddOK() *UserAddOK {
	return &UserAddOK{}
}

/* UserAddOK describes a response with status code 200, with default header values.

OK
*/
type UserAddOK struct {
	Payload *models.Response
}

func (o *UserAddOK) Error() string {
	return fmt.Sprintf("[POST /user][%d] userAddOK  %+v", 200, o.Payload)
}
func (o *UserAddOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *UserAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAddBadRequest creates a UserAddBadRequest with default headers values
func NewUserAddBadRequest() *UserAddBadRequest {
	return &UserAddBadRequest{}
}

/* UserAddBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserAddBadRequest struct {
	Payload *models.Response
}

func (o *UserAddBadRequest) Error() string {
	return fmt.Sprintf("[POST /user][%d] userAddBadRequest  %+v", 400, o.Payload)
}
func (o *UserAddBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *UserAddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAddUnauthorized creates a UserAddUnauthorized with default headers values
func NewUserAddUnauthorized() *UserAddUnauthorized {
	return &UserAddUnauthorized{}
}

/* UserAddUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type UserAddUnauthorized struct {
	Payload *models.Response
}

func (o *UserAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user][%d] userAddUnauthorized  %+v", 401, o.Payload)
}
func (o *UserAddUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *UserAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAddInternalServerError creates a UserAddInternalServerError with default headers values
func NewUserAddInternalServerError() *UserAddInternalServerError {
	return &UserAddInternalServerError{}
}

/* UserAddInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UserAddInternalServerError struct {
	Payload *models.Response
}

func (o *UserAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user][%d] userAddInternalServerError  %+v", 500, o.Payload)
}
func (o *UserAddInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *UserAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
