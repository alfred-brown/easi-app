// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// NewSystemDetailAddParams creates a new SystemDetailAddParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemDetailAddParams() *SystemDetailAddParams {
	return &SystemDetailAddParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemDetailAddParamsWithTimeout creates a new SystemDetailAddParams object
// with the ability to set a timeout on a request.
func NewSystemDetailAddParamsWithTimeout(timeout time.Duration) *SystemDetailAddParams {
	return &SystemDetailAddParams{
		timeout: timeout,
	}
}

// NewSystemDetailAddParamsWithContext creates a new SystemDetailAddParams object
// with the ability to set a context for a request.
func NewSystemDetailAddParamsWithContext(ctx context.Context) *SystemDetailAddParams {
	return &SystemDetailAddParams{
		Context: ctx,
	}
}

// NewSystemDetailAddParamsWithHTTPClient creates a new SystemDetailAddParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemDetailAddParamsWithHTTPClient(client *http.Client) *SystemDetailAddParams {
	return &SystemDetailAddParams{
		HTTPClient: client,
	}
}

/* SystemDetailAddParams contains all the parameters to send to the API endpoint
   for the system detail add operation.

   Typically these are written to a http.Request.
*/
type SystemDetailAddParams struct {

	/* Body.

	   System information to be added to Alfabet.
	*/
	Body *models.SystemDetail

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system detail add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemDetailAddParams) WithDefaults() *SystemDetailAddParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system detail add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemDetailAddParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system detail add params
func (o *SystemDetailAddParams) WithTimeout(timeout time.Duration) *SystemDetailAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system detail add params
func (o *SystemDetailAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system detail add params
func (o *SystemDetailAddParams) WithContext(ctx context.Context) *SystemDetailAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system detail add params
func (o *SystemDetailAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system detail add params
func (o *SystemDetailAddParams) WithHTTPClient(client *http.Client) *SystemDetailAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system detail add params
func (o *SystemDetailAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the system detail add params
func (o *SystemDetailAddParams) WithBody(body *models.SystemDetail) *SystemDetailAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the system detail add params
func (o *SystemDetailAddParams) SetBody(body *models.SystemDetail) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SystemDetailAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
