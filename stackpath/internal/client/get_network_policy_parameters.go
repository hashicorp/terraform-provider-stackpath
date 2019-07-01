// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNetworkPolicyParams creates a new GetNetworkPolicyParams object
// with the default values initialized.
func NewGetNetworkPolicyParams() *GetNetworkPolicyParams {
	var ()
	return &GetNetworkPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkPolicyParamsWithTimeout creates a new GetNetworkPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkPolicyParamsWithTimeout(timeout time.Duration) *GetNetworkPolicyParams {
	var ()
	return &GetNetworkPolicyParams{

		timeout: timeout,
	}
}

// NewGetNetworkPolicyParamsWithContext creates a new GetNetworkPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkPolicyParamsWithContext(ctx context.Context) *GetNetworkPolicyParams {
	var ()
	return &GetNetworkPolicyParams{

		Context: ctx,
	}
}

// NewGetNetworkPolicyParamsWithHTTPClient creates a new GetNetworkPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkPolicyParamsWithHTTPClient(client *http.Client) *GetNetworkPolicyParams {
	var ()
	return &GetNetworkPolicyParams{
		HTTPClient: client,
	}
}

/*GetNetworkPolicyParams contains all the parameters to send to the API endpoint
for the get network policy operation typically these are written to a http.Request
*/
type GetNetworkPolicyParams struct {

	/*NetworkPolicyID
	  The ID the network policy to retrieve

	*/
	NetworkPolicyID string
	/*StackID
	  The ID of the stack containing the network policy to retrieve

	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network policy params
func (o *GetNetworkPolicyParams) WithTimeout(timeout time.Duration) *GetNetworkPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network policy params
func (o *GetNetworkPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network policy params
func (o *GetNetworkPolicyParams) WithContext(ctx context.Context) *GetNetworkPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network policy params
func (o *GetNetworkPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network policy params
func (o *GetNetworkPolicyParams) WithHTTPClient(client *http.Client) *GetNetworkPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network policy params
func (o *GetNetworkPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkPolicyID adds the networkPolicyID to the get network policy params
func (o *GetNetworkPolicyParams) WithNetworkPolicyID(networkPolicyID string) *GetNetworkPolicyParams {
	o.SetNetworkPolicyID(networkPolicyID)
	return o
}

// SetNetworkPolicyID adds the networkPolicyId to the get network policy params
func (o *GetNetworkPolicyParams) SetNetworkPolicyID(networkPolicyID string) {
	o.NetworkPolicyID = networkPolicyID
}

// WithStackID adds the stackID to the get network policy params
func (o *GetNetworkPolicyParams) WithStackID(stackID string) *GetNetworkPolicyParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get network policy params
func (o *GetNetworkPolicyParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_policy_id
	if err := r.SetPathParam("network_policy_id", o.NetworkPolicyID); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
