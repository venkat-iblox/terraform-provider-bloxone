// Code generated by go-swagger; DO NOT EDIT.

package host

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

// NewHostListParams creates a new HostListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostListParams() *HostListParams {
	return &HostListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostListParamsWithTimeout creates a new HostListParams object
// with the ability to set a timeout on a request.
func NewHostListParamsWithTimeout(timeout time.Duration) *HostListParams {
	return &HostListParams{
		timeout: timeout,
	}
}

// NewHostListParamsWithContext creates a new HostListParams object
// with the ability to set a context for a request.
func NewHostListParamsWithContext(ctx context.Context) *HostListParams {
	return &HostListParams{
		Context: ctx,
	}
}

// NewHostListParamsWithHTTPClient creates a new HostListParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostListParamsWithHTTPClient(client *http.Client) *HostListParams {
	return &HostListParams{
		HTTPClient: client,
	}
}

/* HostListParams contains all the parameters to send to the API endpoint
   for the host list operation.

   Typically these are written to a http.Request.
*/
type HostListParams struct {

	/* Fields.



	A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.

	Specify this parameter as a comma-separated list of JSON tag names.


	*/
	Fields *string

	/* Filter.



	A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.

	Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'. The following operators are commonly used in filter expressions:

	|  Op   |  Description               |
	|  --   |  -----------               |
	|  ==   |  Equal                     |
	|  !=   |  Not Equal                 |
	|  >    |  Greater Than              |
	|   >=  |  Greater Than or Equal To  |
	|  <    |  Less Than                 |
	|  <=   |  Less Than or Equal To     |
	|  and  |  Logical AND               |
	|  ~    |  Matches Regex             |
	|  !~   |  Does Not Match Regex      |
	|  or   |  Logical OR                |
	|  not  |  Logical NOT               |
	|  ()   |  Groupping Operators       |


	*/
	Filter *string

	/* Limit.



	The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.


	*/
	Limit *int64

	/* Offset.



	The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.


	*/
	Offset *int64

	/* OrderBy.



	A collection of response resources can be sorted by their JSON tags. For a 'flat' resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)

	Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix 'asc' sorts the data in ascending order. The suffix 'desc' sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.


	*/
	OrderBy *string

	/* PageToken.



	The service-defined string used to identify a page of resources. A null value indicates the first page.


	*/
	PageToken *string

	/* Tfilter.

	   This parameter is used for filtering by tags.
	*/
	Tfilter *string

	/* TorderBy.

	   This parameter is used for sorting by tags.
	*/
	TorderBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the host list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostListParams) WithDefaults() *HostListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the host list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the host list params
func (o *HostListParams) WithTimeout(timeout time.Duration) *HostListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the host list params
func (o *HostListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the host list params
func (o *HostListParams) WithContext(ctx context.Context) *HostListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the host list params
func (o *HostListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the host list params
func (o *HostListParams) WithHTTPClient(client *http.Client) *HostListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the host list params
func (o *HostListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the host list params
func (o *HostListParams) WithFields(fields *string) *HostListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the host list params
func (o *HostListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the host list params
func (o *HostListParams) WithFilter(filter *string) *HostListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the host list params
func (o *HostListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the host list params
func (o *HostListParams) WithLimit(limit *int64) *HostListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the host list params
func (o *HostListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the host list params
func (o *HostListParams) WithOffset(offset *int64) *HostListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the host list params
func (o *HostListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the host list params
func (o *HostListParams) WithOrderBy(orderBy *string) *HostListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the host list params
func (o *HostListParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithPageToken adds the pageToken to the host list params
func (o *HostListParams) WithPageToken(pageToken *string) *HostListParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the host list params
func (o *HostListParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithTfilter adds the tfilter to the host list params
func (o *HostListParams) WithTfilter(tfilter *string) *HostListParams {
	o.SetTfilter(tfilter)
	return o
}

// SetTfilter adds the tfilter to the host list params
func (o *HostListParams) SetTfilter(tfilter *string) {
	o.Tfilter = tfilter
}

// WithTorderBy adds the torderBy to the host list params
func (o *HostListParams) WithTorderBy(torderBy *string) *HostListParams {
	o.SetTorderBy(torderBy)
	return o
}

// SetTorderBy adds the torderBy to the host list params
func (o *HostListParams) SetTorderBy(torderBy *string) {
	o.TorderBy = torderBy
}

// WriteToRequest writes these params to a swagger request
func (o *HostListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Filter != nil {

		// query param _filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("_filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param _limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("_limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param _offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("_offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// query param _order_by
		var qrOrderBy string

		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {

			if err := r.SetQueryParam("_order_by", qOrderBy); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param _page_token
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("_page_token", qPageToken); err != nil {
				return err
			}
		}
	}

	if o.Tfilter != nil {

		// query param _tfilter
		var qrTfilter string

		if o.Tfilter != nil {
			qrTfilter = *o.Tfilter
		}
		qTfilter := qrTfilter
		if qTfilter != "" {

			if err := r.SetQueryParam("_tfilter", qTfilter); err != nil {
				return err
			}
		}
	}

	if o.TorderBy != nil {

		// query param _torder_by
		var qrTorderBy string

		if o.TorderBy != nil {
			qrTorderBy = *o.TorderBy
		}
		qTorderBy := qrTorderBy
		if qTorderBy != "" {

			if err := r.SetQueryParam("_torder_by", qTorderBy); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
