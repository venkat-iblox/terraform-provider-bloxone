// Code generated by go-swagger; DO NOT EDIT.

package address_block

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

// NewAddressBlockCopyParams creates a new AddressBlockCopyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressBlockCopyParams() *AddressBlockCopyParams {
	return &AddressBlockCopyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressBlockCopyParamsWithTimeout creates a new AddressBlockCopyParams object
// with the ability to set a timeout on a request.
func NewAddressBlockCopyParamsWithTimeout(timeout time.Duration) *AddressBlockCopyParams {
	return &AddressBlockCopyParams{
		timeout: timeout,
	}
}

// NewAddressBlockCopyParamsWithContext creates a new AddressBlockCopyParams object
// with the ability to set a context for a request.
func NewAddressBlockCopyParamsWithContext(ctx context.Context) *AddressBlockCopyParams {
	return &AddressBlockCopyParams{
		Context: ctx,
	}
}

// NewAddressBlockCopyParamsWithHTTPClient creates a new AddressBlockCopyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressBlockCopyParamsWithHTTPClient(client *http.Client) *AddressBlockCopyParams {
	return &AddressBlockCopyParams{
		HTTPClient: client,
	}
}

/* AddressBlockCopyParams contains all the parameters to send to the API endpoint
   for the address block copy operation.

   Typically these are written to a http.Request.
*/
type AddressBlockCopyParams struct {

	// Body.
	Body *models.IpamsvcCopyAddressBlock

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address block copy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockCopyParams) WithDefaults() *AddressBlockCopyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address block copy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockCopyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address block copy params
func (o *AddressBlockCopyParams) WithTimeout(timeout time.Duration) *AddressBlockCopyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address block copy params
func (o *AddressBlockCopyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address block copy params
func (o *AddressBlockCopyParams) WithContext(ctx context.Context) *AddressBlockCopyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address block copy params
func (o *AddressBlockCopyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address block copy params
func (o *AddressBlockCopyParams) WithHTTPClient(client *http.Client) *AddressBlockCopyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address block copy params
func (o *AddressBlockCopyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the address block copy params
func (o *AddressBlockCopyParams) WithBody(body *models.IpamsvcCopyAddressBlock) *AddressBlockCopyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the address block copy params
func (o *AddressBlockCopyParams) SetBody(body *models.IpamsvcCopyAddressBlock) {
	o.Body = body
}

// WithID adds the id to the address block copy params
func (o *AddressBlockCopyParams) WithID(id string) *AddressBlockCopyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the address block copy params
func (o *AddressBlockCopyParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddressBlockCopyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
