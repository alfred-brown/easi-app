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
	"github.com/go-openapi/swag"
)

// NewSystemSummaryFindByIDParams creates a new SystemSummaryFindByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemSummaryFindByIDParams() *SystemSummaryFindByIDParams {
	return &SystemSummaryFindByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemSummaryFindByIDParamsWithTimeout creates a new SystemSummaryFindByIDParams object
// with the ability to set a timeout on a request.
func NewSystemSummaryFindByIDParamsWithTimeout(timeout time.Duration) *SystemSummaryFindByIDParams {
	return &SystemSummaryFindByIDParams{
		timeout: timeout,
	}
}

// NewSystemSummaryFindByIDParamsWithContext creates a new SystemSummaryFindByIDParams object
// with the ability to set a context for a request.
func NewSystemSummaryFindByIDParamsWithContext(ctx context.Context) *SystemSummaryFindByIDParams {
	return &SystemSummaryFindByIDParams{
		Context: ctx,
	}
}

// NewSystemSummaryFindByIDParamsWithHTTPClient creates a new SystemSummaryFindByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemSummaryFindByIDParamsWithHTTPClient(client *http.Client) *SystemSummaryFindByIDParams {
	return &SystemSummaryFindByIDParams{
		HTTPClient: client,
	}
}

/* SystemSummaryFindByIDParams contains all the parameters to send to the API endpoint
   for the system summary find by Id operation.

   Typically these are written to a http.Request.
*/
type SystemSummaryFindByIDParams struct {

	/* BelongsTo.

	   Return only sub-systems of the system ID provided
	*/
	BelongsTo *string

	/* ID.

	   ID of system to retrieve.
	*/
	ID string

	/* IdsOnly.

	   Return only system ids and names
	*/
	IdsOnly *bool

	/* IncludeInSurvey.

	   Include only system census eligible systems
	*/
	IncludeInSurvey *bool

	/* State.

	   System state
	*/
	State *string

	/* Status.

	   System status
	*/
	Status *string

	/* Version.

	   System versions
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system summary find by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemSummaryFindByIDParams) WithDefaults() *SystemSummaryFindByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system summary find by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemSummaryFindByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithTimeout(timeout time.Duration) *SystemSummaryFindByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithContext(ctx context.Context) *SystemSummaryFindByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithHTTPClient(client *http.Client) *SystemSummaryFindByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBelongsTo adds the belongsTo to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithBelongsTo(belongsTo *string) *SystemSummaryFindByIDParams {
	o.SetBelongsTo(belongsTo)
	return o
}

// SetBelongsTo adds the belongsTo to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetBelongsTo(belongsTo *string) {
	o.BelongsTo = belongsTo
}

// WithID adds the id to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithID(id string) *SystemSummaryFindByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetID(id string) {
	o.ID = id
}

// WithIdsOnly adds the idsOnly to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithIdsOnly(idsOnly *bool) *SystemSummaryFindByIDParams {
	o.SetIdsOnly(idsOnly)
	return o
}

// SetIdsOnly adds the idsOnly to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetIdsOnly(idsOnly *bool) {
	o.IdsOnly = idsOnly
}

// WithIncludeInSurvey adds the includeInSurvey to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithIncludeInSurvey(includeInSurvey *bool) *SystemSummaryFindByIDParams {
	o.SetIncludeInSurvey(includeInSurvey)
	return o
}

// SetIncludeInSurvey adds the includeInSurvey to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetIncludeInSurvey(includeInSurvey *bool) {
	o.IncludeInSurvey = includeInSurvey
}

// WithState adds the state to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithState(state *string) *SystemSummaryFindByIDParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetState(state *string) {
	o.State = state
}

// WithStatus adds the status to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithStatus(status *string) *SystemSummaryFindByIDParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetStatus(status *string) {
	o.Status = status
}

// WithVersion adds the version to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) WithVersion(version *string) *SystemSummaryFindByIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the system summary find by Id params
func (o *SystemSummaryFindByIDParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SystemSummaryFindByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BelongsTo != nil {

		// query param belongsTo
		var qrBelongsTo string

		if o.BelongsTo != nil {
			qrBelongsTo = *o.BelongsTo
		}
		qBelongsTo := qrBelongsTo
		if qBelongsTo != "" {

			if err := r.SetQueryParam("belongsTo", qBelongsTo); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IdsOnly != nil {

		// query param idsOnly
		var qrIdsOnly bool

		if o.IdsOnly != nil {
			qrIdsOnly = *o.IdsOnly
		}
		qIdsOnly := swag.FormatBool(qrIdsOnly)
		if qIdsOnly != "" {

			if err := r.SetQueryParam("idsOnly", qIdsOnly); err != nil {
				return err
			}
		}
	}

	if o.IncludeInSurvey != nil {

		// query param includeInSurvey
		var qrIncludeInSurvey bool

		if o.IncludeInSurvey != nil {
			qrIncludeInSurvey = *o.IncludeInSurvey
		}
		qIncludeInSurvey := swag.FormatBool(qrIncludeInSurvey)
		if qIncludeInSurvey != "" {

			if err := r.SetQueryParam("includeInSurvey", qIncludeInSurvey); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}