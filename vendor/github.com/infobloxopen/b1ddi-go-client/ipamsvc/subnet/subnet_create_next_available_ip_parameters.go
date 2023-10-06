// Code generated by go-swagger; DO NOT EDIT.

package subnet

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

// NewSubnetCreateNextAvailableIPParams creates a new SubnetCreateNextAvailableIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubnetCreateNextAvailableIPParams() *SubnetCreateNextAvailableIPParams {
	return &SubnetCreateNextAvailableIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubnetCreateNextAvailableIPParamsWithTimeout creates a new SubnetCreateNextAvailableIPParams object
// with the ability to set a timeout on a request.
func NewSubnetCreateNextAvailableIPParamsWithTimeout(timeout time.Duration) *SubnetCreateNextAvailableIPParams {
	return &SubnetCreateNextAvailableIPParams{
		timeout: timeout,
	}
}

// NewSubnetCreateNextAvailableIPParamsWithContext creates a new SubnetCreateNextAvailableIPParams object
// with the ability to set a context for a request.
func NewSubnetCreateNextAvailableIPParamsWithContext(ctx context.Context) *SubnetCreateNextAvailableIPParams {
	return &SubnetCreateNextAvailableIPParams{
		Context: ctx,
	}
}

// NewSubnetCreateNextAvailableIPParamsWithHTTPClient creates a new SubnetCreateNextAvailableIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubnetCreateNextAvailableIPParamsWithHTTPClient(client *http.Client) *SubnetCreateNextAvailableIPParams {
	return &SubnetCreateNextAvailableIPParams{
		HTTPClient: client,
	}
}

/* SubnetCreateNextAvailableIPParams contains all the parameters to send to the API endpoint
   for the subnet create next available IP operation.

   Typically these are written to a http.Request.
*/
type SubnetCreateNextAvailableIPParams struct {

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subnet create next available IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetCreateNextAvailableIPParams) WithDefaults() *SubnetCreateNextAvailableIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subnet create next available IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetCreateNextAvailableIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subnet create next available IP params
func (o *SubnetCreateNextAvailableIPParams) WithTimeout(timeout time.Duration) *SubnetCreateNextAvailableIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subnet create next available IP params
func (o *SubnetCreateNextAvailableIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subnet create next available IP params
func (o *SubnetCreateNextAvailableIPParams) WithContext(ctx context.Context) *SubnetCreateNextAvailableIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subnet create next available IP params
func (o *SubnetCreateNextAvailableIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subnet create next available IP params
func (o *SubnetCreateNextAvailableIPParams) WithHTTPClient(client *http.Client) *SubnetCreateNextAvailableIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subnet create next available IP params
func (o *SubnetCreateNextAvailableIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the subnet create next available IP params
func (o *SubnetCreateNextAvailableIPParams) WithID(id string) *SubnetCreateNextAvailableIPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subnet create next available IP params
func (o *SubnetCreateNextAvailableIPParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubnetCreateNextAvailableIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
