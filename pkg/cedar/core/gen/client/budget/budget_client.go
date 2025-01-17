// Code generated by go-swagger; DO NOT EDIT.

package budget

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new budget API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for budget API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BudgetAdd(params *BudgetAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BudgetAddOK, error)

	BudgetDeleteList(params *BudgetDeleteListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BudgetDeleteListOK, error)

	BudgetFind(params *BudgetFindParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BudgetFindOK, error)

	BudgetUpdate(params *BudgetUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BudgetUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BudgetAdd adds new budget s to the alfabet system this interface takes an array of one or more budget documents id project Id system Id funding Id and funding

  Add new budget(s) to the Alfabet system. This interface takes an array of one or more budgets (id, projectId, systemId, fundingId and funding).
*/
func (a *Client) BudgetAdd(params *BudgetAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BudgetAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBudgetAddParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "budgetAdd",
		Method:             "POST",
		PathPattern:        "/budget",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BudgetAddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BudgetAddOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for budgetAdd: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BudgetDeleteList deletes a list of budgets based on an array of budget ids

  Deletes a list of budgets based on an array of budget ids.
*/
func (a *Client) BudgetDeleteList(params *BudgetDeleteListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BudgetDeleteListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBudgetDeleteListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "budgetDeleteList",
		Method:             "DELETE",
		PathPattern:        "/budget",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BudgetDeleteListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BudgetDeleteListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for budgetDeleteList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BudgetFind retrieves a list of budgets based on query criteria listed in the parameters section passing a system Id will cause the interface to return the budget s for just that system setting only ids to true will only return the id s whereas if not set the response will also include project Id system Id funding Id and funding this interface has a limit of 5000 records

  Retrieve a list of budgets based on query criteria listed in the parameters section. Passing a SystemId will cause the interface to return the budget(s) for just that system. Setting onlyIds to true will only return the id's, whereas if not set, the response will also include projectId, systemId, fundingId and funding. This interface has a limit of 5000 records.
*/
func (a *Client) BudgetFind(params *BudgetFindParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BudgetFindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBudgetFindParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "budgetFind",
		Method:             "GET",
		PathPattern:        "/budget",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BudgetFindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BudgetFindOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for budgetFind: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BudgetUpdate updates a existing budget with alfabet the input requires an array of one or more budget documents id project Id system Id funding Id and funding

  Update a existing budget with Alfabet. The input requires an array of one or more budget documents (id, projectId, systemId, fundingId and funding).
*/
func (a *Client) BudgetUpdate(params *BudgetUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BudgetUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBudgetUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "budgetUpdate",
		Method:             "PUT",
		PathPattern:        "/budget",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BudgetUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BudgetUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for budgetUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
