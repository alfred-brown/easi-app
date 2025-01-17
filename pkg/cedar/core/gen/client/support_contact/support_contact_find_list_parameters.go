// Code generated by go-swagger; DO NOT EDIT.

package support_contact

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
)

// NewSupportContactFindListParams creates a new SupportContactFindListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSupportContactFindListParams() *SupportContactFindListParams {
	return &SupportContactFindListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSupportContactFindListParamsWithTimeout creates a new SupportContactFindListParams object
// with the ability to set a timeout on a request.
func NewSupportContactFindListParamsWithTimeout(timeout time.Duration) *SupportContactFindListParams {
	return &SupportContactFindListParams{
		timeout: timeout,
	}
}

// NewSupportContactFindListParamsWithContext creates a new SupportContactFindListParams object
// with the ability to set a context for a request.
func NewSupportContactFindListParamsWithContext(ctx context.Context) *SupportContactFindListParams {
	return &SupportContactFindListParams{
		Context: ctx,
	}
}

// NewSupportContactFindListParamsWithHTTPClient creates a new SupportContactFindListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSupportContactFindListParamsWithHTTPClient(client *http.Client) *SupportContactFindListParams {
	return &SupportContactFindListParams{
		HTTPClient: client,
	}
}

/* SupportContactFindListParams contains all the parameters to send to the API endpoint
   for the support contact find list operation.

   Typically these are written to a http.Request.
*/
type SupportContactFindListParams struct {

	/* Application.

	   The applications id.
	*/
	Application string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the support contact find list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportContactFindListParams) WithDefaults() *SupportContactFindListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the support contact find list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportContactFindListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the support contact find list params
func (o *SupportContactFindListParams) WithTimeout(timeout time.Duration) *SupportContactFindListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the support contact find list params
func (o *SupportContactFindListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the support contact find list params
func (o *SupportContactFindListParams) WithContext(ctx context.Context) *SupportContactFindListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the support contact find list params
func (o *SupportContactFindListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the support contact find list params
func (o *SupportContactFindListParams) WithHTTPClient(client *http.Client) *SupportContactFindListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the support contact find list params
func (o *SupportContactFindListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplication adds the application to the support contact find list params
func (o *SupportContactFindListParams) WithApplication(application string) *SupportContactFindListParams {
	o.SetApplication(application)
	return o
}

// SetApplication adds the application to the support contact find list params
func (o *SupportContactFindListParams) SetApplication(application string) {
	o.Application = application
}

// WriteToRequest writes these params to a swagger request
func (o *SupportContactFindListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param application
	qrApplication := o.Application
	qApplication := qrApplication
	if qApplication != "" {

		if err := r.SetQueryParam("application", qApplication); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
