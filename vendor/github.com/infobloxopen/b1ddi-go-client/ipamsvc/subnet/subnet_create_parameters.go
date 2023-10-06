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

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// NewSubnetCreateParams creates a new SubnetCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubnetCreateParams() *SubnetCreateParams {
	return &SubnetCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubnetCreateParamsWithTimeout creates a new SubnetCreateParams object
// with the ability to set a timeout on a request.
func NewSubnetCreateParamsWithTimeout(timeout time.Duration) *SubnetCreateParams {
	return &SubnetCreateParams{
		timeout: timeout,
	}
}

// NewSubnetCreateParamsWithContext creates a new SubnetCreateParams object
// with the ability to set a context for a request.
func NewSubnetCreateParamsWithContext(ctx context.Context) *SubnetCreateParams {
	return &SubnetCreateParams{
		Context: ctx,
	}
}

// NewSubnetCreateParamsWithHTTPClient creates a new SubnetCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubnetCreateParamsWithHTTPClient(client *http.Client) *SubnetCreateParams {
	return &SubnetCreateParams{
		HTTPClient: client,
	}
}

/* SubnetCreateParams contains all the parameters to send to the API endpoint
   for the subnet create operation.

   Typically these are written to a http.Request.
*/
type SubnetCreateParams struct {

	// Body.
	Body *models.IpamsvcSubnet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subnet create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetCreateParams) WithDefaults() *SubnetCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subnet create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubnetCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subnet create params
func (o *SubnetCreateParams) WithTimeout(timeout time.Duration) *SubnetCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subnet create params
func (o *SubnetCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subnet create params
func (o *SubnetCreateParams) WithContext(ctx context.Context) *SubnetCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subnet create params
func (o *SubnetCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subnet create params
func (o *SubnetCreateParams) WithHTTPClient(client *http.Client) *SubnetCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subnet create params
func (o *SubnetCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subnet create params
func (o *SubnetCreateParams) WithBody(body *models.IpamsvcSubnet) *SubnetCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subnet create params
func (o *SubnetCreateParams) SetBody(body *models.IpamsvcSubnet) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SubnetCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
