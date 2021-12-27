// Code generated by go-swagger; DO NOT EDIT.

package budget

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

// NewBudgetUpdateParams creates a new BudgetUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBudgetUpdateParams() *BudgetUpdateParams {
	return &BudgetUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBudgetUpdateParamsWithTimeout creates a new BudgetUpdateParams object
// with the ability to set a timeout on a request.
func NewBudgetUpdateParamsWithTimeout(timeout time.Duration) *BudgetUpdateParams {
	return &BudgetUpdateParams{
		timeout: timeout,
	}
}

// NewBudgetUpdateParamsWithContext creates a new BudgetUpdateParams object
// with the ability to set a context for a request.
func NewBudgetUpdateParamsWithContext(ctx context.Context) *BudgetUpdateParams {
	return &BudgetUpdateParams{
		Context: ctx,
	}
}

// NewBudgetUpdateParamsWithHTTPClient creates a new BudgetUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBudgetUpdateParamsWithHTTPClient(client *http.Client) *BudgetUpdateParams {
	return &BudgetUpdateParams{
		HTTPClient: client,
	}
}

/* BudgetUpdateParams contains all the parameters to send to the API endpoint
   for the budget update operation.

   Typically these are written to a http.Request.
*/
type BudgetUpdateParams struct {

	/* Body.

	   Budgets to be updated in CEDAR. This required input in a list of Budget documents (id, projectId, systemId, fundingId and funding).
	*/
	Body *models.BudgetUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the budget update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BudgetUpdateParams) WithDefaults() *BudgetUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the budget update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BudgetUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the budget update params
func (o *BudgetUpdateParams) WithTimeout(timeout time.Duration) *BudgetUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the budget update params
func (o *BudgetUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the budget update params
func (o *BudgetUpdateParams) WithContext(ctx context.Context) *BudgetUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the budget update params
func (o *BudgetUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the budget update params
func (o *BudgetUpdateParams) WithHTTPClient(client *http.Client) *BudgetUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the budget update params
func (o *BudgetUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the budget update params
func (o *BudgetUpdateParams) WithBody(body *models.BudgetUpdateRequest) *BudgetUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the budget update params
func (o *BudgetUpdateParams) SetBody(body *models.BudgetUpdateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BudgetUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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