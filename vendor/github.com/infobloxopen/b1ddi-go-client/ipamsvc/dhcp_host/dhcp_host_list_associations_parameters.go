// Code generated by go-swagger; DO NOT EDIT.

package dhcp_host

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

// NewDhcpHostListAssociationsParams creates a new DhcpHostListAssociationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDhcpHostListAssociationsParams() *DhcpHostListAssociationsParams {
	return &DhcpHostListAssociationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDhcpHostListAssociationsParamsWithTimeout creates a new DhcpHostListAssociationsParams object
// with the ability to set a timeout on a request.
func NewDhcpHostListAssociationsParamsWithTimeout(timeout time.Duration) *DhcpHostListAssociationsParams {
	return &DhcpHostListAssociationsParams{
		timeout: timeout,
	}
}

// NewDhcpHostListAssociationsParamsWithContext creates a new DhcpHostListAssociationsParams object
// with the ability to set a context for a request.
func NewDhcpHostListAssociationsParamsWithContext(ctx context.Context) *DhcpHostListAssociationsParams {
	return &DhcpHostListAssociationsParams{
		Context: ctx,
	}
}

// NewDhcpHostListAssociationsParamsWithHTTPClient creates a new DhcpHostListAssociationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDhcpHostListAssociationsParamsWithHTTPClient(client *http.Client) *DhcpHostListAssociationsParams {
	return &DhcpHostListAssociationsParams{
		HTTPClient: client,
	}
}

/* DhcpHostListAssociationsParams contains all the parameters to send to the API endpoint
   for the dhcp host list associations operation.

   Typically these are written to a http.Request.
*/
type DhcpHostListAssociationsParams struct {

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dhcp host list associations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DhcpHostListAssociationsParams) WithDefaults() *DhcpHostListAssociationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dhcp host list associations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DhcpHostListAssociationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dhcp host list associations params
func (o *DhcpHostListAssociationsParams) WithTimeout(timeout time.Duration) *DhcpHostListAssociationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dhcp host list associations params
func (o *DhcpHostListAssociationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dhcp host list associations params
func (o *DhcpHostListAssociationsParams) WithContext(ctx context.Context) *DhcpHostListAssociationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dhcp host list associations params
func (o *DhcpHostListAssociationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dhcp host list associations params
func (o *DhcpHostListAssociationsParams) WithHTTPClient(client *http.Client) *DhcpHostListAssociationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dhcp host list associations params
func (o *DhcpHostListAssociationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dhcp host list associations params
func (o *DhcpHostListAssociationsParams) WithID(id string) *DhcpHostListAssociationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dhcp host list associations params
func (o *DhcpHostListAssociationsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DhcpHostListAssociationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
