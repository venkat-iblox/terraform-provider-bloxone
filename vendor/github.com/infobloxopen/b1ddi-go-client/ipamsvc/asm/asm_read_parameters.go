// Code generated by go-swagger; DO NOT EDIT.

package asm

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

// NewAsmReadParams creates a new AsmReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAsmReadParams() *AsmReadParams {
	return &AsmReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAsmReadParamsWithTimeout creates a new AsmReadParams object
// with the ability to set a timeout on a request.
func NewAsmReadParamsWithTimeout(timeout time.Duration) *AsmReadParams {
	return &AsmReadParams{
		timeout: timeout,
	}
}

// NewAsmReadParamsWithContext creates a new AsmReadParams object
// with the ability to set a context for a request.
func NewAsmReadParamsWithContext(ctx context.Context) *AsmReadParams {
	return &AsmReadParams{
		Context: ctx,
	}
}

// NewAsmReadParamsWithHTTPClient creates a new AsmReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewAsmReadParamsWithHTTPClient(client *http.Client) *AsmReadParams {
	return &AsmReadParams{
		HTTPClient: client,
	}
}

/* AsmReadParams contains all the parameters to send to the API endpoint
   for the asm read operation.

   Typically these are written to a http.Request.
*/
type AsmReadParams struct {

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

// WithDefaults hydrates default values in the asm read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AsmReadParams) WithDefaults() *AsmReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the asm read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AsmReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the asm read params
func (o *AsmReadParams) WithTimeout(timeout time.Duration) *AsmReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the asm read params
func (o *AsmReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the asm read params
func (o *AsmReadParams) WithContext(ctx context.Context) *AsmReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the asm read params
func (o *AsmReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the asm read params
func (o *AsmReadParams) WithHTTPClient(client *http.Client) *AsmReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the asm read params
func (o *AsmReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the asm read params
func (o *AsmReadParams) WithFields(fields *string) *AsmReadParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the asm read params
func (o *AsmReadParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the asm read params
func (o *AsmReadParams) WithID(id string) *AsmReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the asm read params
func (o *AsmReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AsmReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
