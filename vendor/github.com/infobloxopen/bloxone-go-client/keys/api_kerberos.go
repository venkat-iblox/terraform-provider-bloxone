/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type KerberosAPI interface {
	/*
			Delete Delete the Kerberos key.

			Use this method to delete a __KerberosKey__ object.
		A __KerberosKey__ object represents a Kerberos key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return KerberosAPIDeleteRequest
	*/
	Delete(ctx context.Context, id string) KerberosAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r KerberosAPIDeleteRequest) (*http.Response, error)
	/*
			List Retrieve Kerberos keys.

			Use this method to retrieve __KerberosKey__ objects.
		A __KerberosKey__ object represents a Kerberos key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return KerberosAPIListRequest
	*/
	List(ctx context.Context) KerberosAPIListRequest

	// ListExecute executes the request
	//  @return ListKerberosKeyResponse
	ListExecute(r KerberosAPIListRequest) (*ListKerberosKeyResponse, *http.Response, error)
	/*
			Read Retrieve the Kerberos key.

			Use this method to retrieve a __KerberosKey__ object.
		A __KerberosKey__ object represents a Kerberos key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return KerberosAPIReadRequest
	*/
	Read(ctx context.Context, id string) KerberosAPIReadRequest

	// ReadExecute executes the request
	//  @return ReadKerberosKeyResponse
	ReadExecute(r KerberosAPIReadRequest) (*ReadKerberosKeyResponse, *http.Response, error)
	/*
			Update Update the Kerberos key.

			Use this method to update a __KerberosKey__ object.
		A __KerberosKey__ object represents a Kerberos key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return KerberosAPIUpdateRequest
	*/
	Update(ctx context.Context, id string) KerberosAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateKerberosKeyResponse
	UpdateExecute(r KerberosAPIUpdateRequest) (*UpdateKerberosKeyResponse, *http.Response, error)
}

// KerberosAPIService KerberosAPI service
type KerberosAPIService internal.Service

type KerberosAPIDeleteRequest struct {
	ctx        context.Context
	ApiService KerberosAPI
	id         string
}

func (r KerberosAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete the Kerberos key.

Use this method to delete a __KerberosKey__ object.
A __KerberosKey__ object represents a Kerberos key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return KerberosAPIDeleteRequest
*/
func (a *KerberosAPIService) Delete(ctx context.Context, id string) KerberosAPIDeleteRequest {
	return KerberosAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *KerberosAPIService) DeleteExecute(r KerberosAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "KerberosAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/kerberos/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type KerberosAPIListRequest struct {
	ctx        context.Context
	ApiService KerberosAPI
	fields     *string
	filter     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	orderBy    *string
	tfilter    *string
	torderBy   *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r KerberosAPIListRequest) Fields(fields string) KerberosAPIListRequest {
	r.fields = &fields
	return r
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |
func (r KerberosAPIListRequest) Filter(filter string) KerberosAPIListRequest {
	r.filter = &filter
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r KerberosAPIListRequest) Offset(offset int32) KerberosAPIListRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r KerberosAPIListRequest) Limit(limit int32) KerberosAPIListRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r KerberosAPIListRequest) PageToken(pageToken string) KerberosAPIListRequest {
	r.pageToken = &pageToken
	return r
}

// A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.
func (r KerberosAPIListRequest) OrderBy(orderBy string) KerberosAPIListRequest {
	r.orderBy = &orderBy
	return r
}

// This parameter is used for filtering by tags.
func (r KerberosAPIListRequest) Tfilter(tfilter string) KerberosAPIListRequest {
	r.tfilter = &tfilter
	return r
}

// This parameter is used for sorting by tags.
func (r KerberosAPIListRequest) TorderBy(torderBy string) KerberosAPIListRequest {
	r.torderBy = &torderBy
	return r
}

func (r KerberosAPIListRequest) Execute() (*ListKerberosKeyResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve Kerberos keys.

Use this method to retrieve __KerberosKey__ objects.
A __KerberosKey__ object represents a Kerberos key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return KerberosAPIListRequest
*/
func (a *KerberosAPIService) List(ctx context.Context) KerberosAPIListRequest {
	return KerberosAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListKerberosKeyResponse
func (a *KerberosAPIService) ListExecute(r KerberosAPIListRequest) (*ListKerberosKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListKerberosKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "KerberosAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/kerberos"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
	}
	if r.offset != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_offset", r.offset, "")
	}
	if r.limit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_limit", r.limit, "")
	}
	if r.pageToken != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_token", r.pageToken, "")
	}
	if r.orderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_order_by", r.orderBy, "")
	}
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type KerberosAPIReadRequest struct {
	ctx        context.Context
	ApiService KerberosAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r KerberosAPIReadRequest) Fields(fields string) KerberosAPIReadRequest {
	r.fields = &fields
	return r
}

func (r KerberosAPIReadRequest) Execute() (*ReadKerberosKeyResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Retrieve the Kerberos key.

Use this method to retrieve a __KerberosKey__ object.
A __KerberosKey__ object represents a Kerberos key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return KerberosAPIReadRequest
*/
func (a *KerberosAPIService) Read(ctx context.Context, id string) KerberosAPIReadRequest {
	return KerberosAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ReadKerberosKeyResponse
func (a *KerberosAPIService) ReadExecute(r KerberosAPIReadRequest) (*ReadKerberosKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ReadKerberosKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "KerberosAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/kerberos/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type KerberosAPIUpdateRequest struct {
	ctx        context.Context
	ApiService KerberosAPI
	id         string
	body       *KerberosKey
}

func (r KerberosAPIUpdateRequest) Body(body KerberosKey) KerberosAPIUpdateRequest {
	r.body = &body
	return r
}

func (r KerberosAPIUpdateRequest) Execute() (*UpdateKerberosKeyResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update the Kerberos key.

Use this method to update a __KerberosKey__ object.
A __KerberosKey__ object represents a Kerberos key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return KerberosAPIUpdateRequest
*/
func (a *KerberosAPIService) Update(ctx context.Context, id string) KerberosAPIUpdateRequest {
	return KerberosAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return UpdateKerberosKeyResponse
func (a *KerberosAPIService) UpdateExecute(r KerberosAPIUpdateRequest) (*UpdateKerberosKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateKerberosKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "KerberosAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/kerberos/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if len(a.Client.Cfg.DefaultTags) > 0 && r.body != nil {
		if r.body.Tags == nil {
			r.body.Tags = make(map[string]interface{})
		}
		for k, v := range a.Client.Cfg.DefaultTags {
			if _, ok := r.body.Tags[k]; !ok {
				r.body.Tags[k] = v
			}
		}
	}
	// body params
	localVarPostBody = r.body
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
