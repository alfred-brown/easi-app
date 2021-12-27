// Code generated by go-swagger; DO NOT EDIT.

package exchange

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

// NewExchangeFindByIDParams creates a new ExchangeFindByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExchangeFindByIDParams() *ExchangeFindByIDParams {
	return &ExchangeFindByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExchangeFindByIDParamsWithTimeout creates a new ExchangeFindByIDParams object
// with the ability to set a timeout on a request.
func NewExchangeFindByIDParamsWithTimeout(timeout time.Duration) *ExchangeFindByIDParams {
	return &ExchangeFindByIDParams{
		timeout: timeout,
	}
}

// NewExchangeFindByIDParamsWithContext creates a new ExchangeFindByIDParams object
// with the ability to set a context for a request.
func NewExchangeFindByIDParamsWithContext(ctx context.Context) *ExchangeFindByIDParams {
	return &ExchangeFindByIDParams{
		Context: ctx,
	}
}

// NewExchangeFindByIDParamsWithHTTPClient creates a new ExchangeFindByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewExchangeFindByIDParamsWithHTTPClient(client *http.Client) *ExchangeFindByIDParams {
	return &ExchangeFindByIDParams{
		HTTPClient: client,
	}
}

/* ExchangeFindByIDParams contains all the parameters to send to the API endpoint
   for the exchange find by Id operation.

   Typically these are written to a http.Request.
*/
type ExchangeFindByIDParams struct {

	/* ID.

	   ID of exchange to retrieve.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the exchange find by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExchangeFindByIDParams) WithDefaults() *ExchangeFindByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the exchange find by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExchangeFindByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the exchange find by Id params
func (o *ExchangeFindByIDParams) WithTimeout(timeout time.Duration) *ExchangeFindByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exchange find by Id params
func (o *ExchangeFindByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exchange find by Id params
func (o *ExchangeFindByIDParams) WithContext(ctx context.Context) *ExchangeFindByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exchange find by Id params
func (o *ExchangeFindByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exchange find by Id params
func (o *ExchangeFindByIDParams) WithHTTPClient(client *http.Client) *ExchangeFindByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exchange find by Id params
func (o *ExchangeFindByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the exchange find by Id params
func (o *ExchangeFindByIDParams) WithID(id string) *ExchangeFindByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the exchange find by Id params
func (o *ExchangeFindByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExchangeFindByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}