// Code generated by go-swagger; DO NOT EDIT.

package auth_nsg

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

// NewAuthNsgReadParams creates a new AuthNsgReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthNsgReadParams() *AuthNsgReadParams {
	return &AuthNsgReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthNsgReadParamsWithTimeout creates a new AuthNsgReadParams object
// with the ability to set a timeout on a request.
func NewAuthNsgReadParamsWithTimeout(timeout time.Duration) *AuthNsgReadParams {
	return &AuthNsgReadParams{
		timeout: timeout,
	}
}

// NewAuthNsgReadParamsWithContext creates a new AuthNsgReadParams object
// with the ability to set a context for a request.
func NewAuthNsgReadParamsWithContext(ctx context.Context) *AuthNsgReadParams {
	return &AuthNsgReadParams{
		Context: ctx,
	}
}

// NewAuthNsgReadParamsWithHTTPClient creates a new AuthNsgReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthNsgReadParamsWithHTTPClient(client *http.Client) *AuthNsgReadParams {
	return &AuthNsgReadParams{
		HTTPClient: client,
	}
}

/* AuthNsgReadParams contains all the parameters to send to the API endpoint
   for the auth nsg read operation.

   Typically these are written to a http.Request.
*/
type AuthNsgReadParams struct {

	/* Fields.



	A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.

	Specify this parameter as a comma-separated list of JSON tag names.


	*/
	Fields *string

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth nsg read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthNsgReadParams) WithDefaults() *AuthNsgReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth nsg read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthNsgReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth nsg read params
func (o *AuthNsgReadParams) WithTimeout(timeout time.Duration) *AuthNsgReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth nsg read params
func (o *AuthNsgReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth nsg read params
func (o *AuthNsgReadParams) WithContext(ctx context.Context) *AuthNsgReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth nsg read params
func (o *AuthNsgReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth nsg read params
func (o *AuthNsgReadParams) WithHTTPClient(client *http.Client) *AuthNsgReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth nsg read params
func (o *AuthNsgReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the auth nsg read params
func (o *AuthNsgReadParams) WithFields(fields *string) *AuthNsgReadParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the auth nsg read params
func (o *AuthNsgReadParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the auth nsg read params
func (o *AuthNsgReadParams) WithID(id string) *AuthNsgReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the auth nsg read params
func (o *AuthNsgReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AuthNsgReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param _fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("_fields", qFields); err != nil {
				return err
			}
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
