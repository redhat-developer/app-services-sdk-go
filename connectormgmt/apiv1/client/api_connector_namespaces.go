/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type ConnectorNamespacesApi interface {

	/*
	 * CreateConnectorNamespace Create a new connector namespace
	 * Create a new connector namespace
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiCreateConnectorNamespaceRequest
	 */
	CreateConnectorNamespace(ctx _context.Context) ApiCreateConnectorNamespaceRequest

	/*
	 * CreateConnectorNamespaceExecute executes the request
	 * @return ConnectorNamespace
	 */
	CreateConnectorNamespaceExecute(r ApiCreateConnectorNamespaceRequest) (ConnectorNamespace, *_nethttp.Response, error)

	/*
	 * CreateEvaluationNamespace Create a new short lived evaluation connector namespace
	 * Create a new evaluation connector namespace
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiCreateEvaluationNamespaceRequest
	 */
	CreateEvaluationNamespace(ctx _context.Context) ApiCreateEvaluationNamespaceRequest

	/*
	 * CreateEvaluationNamespaceExecute executes the request
	 * @return ConnectorNamespace
	 */
	CreateEvaluationNamespaceExecute(r ApiCreateEvaluationNamespaceRequest) (ConnectorNamespace, *_nethttp.Response, error)

	/*
	 * DeleteConnectorNamespace Delete a connector namespace
	 * Delete a connector namespace
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param connectorNamespaceId The id of the connector namespace
	 * @return ApiDeleteConnectorNamespaceRequest
	 */
	DeleteConnectorNamespace(ctx _context.Context, connectorNamespaceId string) ApiDeleteConnectorNamespaceRequest

	/*
	 * DeleteConnectorNamespaceExecute executes the request
	 * @return Error
	 */
	DeleteConnectorNamespaceExecute(r ApiDeleteConnectorNamespaceRequest) (Error, *_nethttp.Response, error)

	/*
	 * GetConnectorNamespace Get a connector namespace
	 * Get a connector namespace
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param connectorNamespaceId The id of the connector namespace
	 * @return ApiGetConnectorNamespaceRequest
	 */
	GetConnectorNamespace(ctx _context.Context, connectorNamespaceId string) ApiGetConnectorNamespaceRequest

	/*
	 * GetConnectorNamespaceExecute executes the request
	 * @return ConnectorNamespace
	 */
	GetConnectorNamespaceExecute(r ApiGetConnectorNamespaceRequest) (ConnectorNamespace, *_nethttp.Response, error)

	/*
	 * ListConnectorNamespaces Returns a list of connector namespaces
	 * Returns a list of connector namespaces
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiListConnectorNamespacesRequest
	 */
	ListConnectorNamespaces(ctx _context.Context) ApiListConnectorNamespacesRequest

	/*
	 * ListConnectorNamespacesExecute executes the request
	 * @return ConnectorNamespaceList
	 */
	ListConnectorNamespacesExecute(r ApiListConnectorNamespacesRequest) (ConnectorNamespaceList, *_nethttp.Response, error)

	/*
	 * UpdateConnectorNamespaceById udpate a connector namespace
	 * udpate a connector namespace
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param connectorNamespaceId The id of the connector namespace
	 * @return ApiUpdateConnectorNamespaceByIdRequest
	 */
	UpdateConnectorNamespaceById(ctx _context.Context, connectorNamespaceId string) ApiUpdateConnectorNamespaceByIdRequest

	/*
	 * UpdateConnectorNamespaceByIdExecute executes the request
	 */
	UpdateConnectorNamespaceByIdExecute(r ApiUpdateConnectorNamespaceByIdRequest) (*_nethttp.Response, error)
}

// ConnectorNamespacesApiService ConnectorNamespacesApi service
type ConnectorNamespacesApiService service

type ApiCreateConnectorNamespaceRequest struct {
	ctx _context.Context
	ApiService ConnectorNamespacesApi
	connectorNamespaceRequest *ConnectorNamespaceRequest
}

func (r ApiCreateConnectorNamespaceRequest) ConnectorNamespaceRequest(connectorNamespaceRequest ConnectorNamespaceRequest) ApiCreateConnectorNamespaceRequest {
	r.connectorNamespaceRequest = &connectorNamespaceRequest
	return r
}

func (r ApiCreateConnectorNamespaceRequest) Execute() (ConnectorNamespace, *_nethttp.Response, error) {
	return r.ApiService.CreateConnectorNamespaceExecute(r)
}

/*
 * CreateConnectorNamespace Create a new connector namespace
 * Create a new connector namespace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateConnectorNamespaceRequest
 */
func (a *ConnectorNamespacesApiService) CreateConnectorNamespace(ctx _context.Context) ApiCreateConnectorNamespaceRequest {
	return ApiCreateConnectorNamespaceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ConnectorNamespace
 */
func (a *ConnectorNamespacesApiService) CreateConnectorNamespaceExecute(r ApiCreateConnectorNamespaceRequest) (ConnectorNamespace, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectorNamespace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectorNamespacesApiService.CreateConnectorNamespace")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connector_mgmt/v1/kafka_connector_namespaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.connectorNamespaceRequest == nil {
		return localVarReturnValue, nil, reportError("connectorNamespaceRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.connectorNamespaceRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateEvaluationNamespaceRequest struct {
	ctx _context.Context
	ApiService ConnectorNamespacesApi
	connectorNamespaceEvalRequest *ConnectorNamespaceEvalRequest
}

func (r ApiCreateEvaluationNamespaceRequest) ConnectorNamespaceEvalRequest(connectorNamespaceEvalRequest ConnectorNamespaceEvalRequest) ApiCreateEvaluationNamespaceRequest {
	r.connectorNamespaceEvalRequest = &connectorNamespaceEvalRequest
	return r
}

func (r ApiCreateEvaluationNamespaceRequest) Execute() (ConnectorNamespace, *_nethttp.Response, error) {
	return r.ApiService.CreateEvaluationNamespaceExecute(r)
}

/*
 * CreateEvaluationNamespace Create a new short lived evaluation connector namespace
 * Create a new evaluation connector namespace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateEvaluationNamespaceRequest
 */
func (a *ConnectorNamespacesApiService) CreateEvaluationNamespace(ctx _context.Context) ApiCreateEvaluationNamespaceRequest {
	return ApiCreateEvaluationNamespaceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ConnectorNamespace
 */
func (a *ConnectorNamespacesApiService) CreateEvaluationNamespaceExecute(r ApiCreateEvaluationNamespaceRequest) (ConnectorNamespace, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectorNamespace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectorNamespacesApiService.CreateEvaluationNamespace")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connector_mgmt/v1/kafka_connector_namespaces/eval"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.connectorNamespaceEvalRequest == nil {
		return localVarReturnValue, nil, reportError("connectorNamespaceEvalRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.connectorNamespaceEvalRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteConnectorNamespaceRequest struct {
	ctx _context.Context
	ApiService ConnectorNamespacesApi
	connectorNamespaceId string
}


func (r ApiDeleteConnectorNamespaceRequest) Execute() (Error, *_nethttp.Response, error) {
	return r.ApiService.DeleteConnectorNamespaceExecute(r)
}

/*
 * DeleteConnectorNamespace Delete a connector namespace
 * Delete a connector namespace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connectorNamespaceId The id of the connector namespace
 * @return ApiDeleteConnectorNamespaceRequest
 */
func (a *ConnectorNamespacesApiService) DeleteConnectorNamespace(ctx _context.Context, connectorNamespaceId string) ApiDeleteConnectorNamespaceRequest {
	return ApiDeleteConnectorNamespaceRequest{
		ApiService: a,
		ctx: ctx,
		connectorNamespaceId: connectorNamespaceId,
	}
}

/*
 * Execute executes the request
 * @return Error
 */
func (a *ConnectorNamespacesApiService) DeleteConnectorNamespaceExecute(r ApiDeleteConnectorNamespaceRequest) (Error, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Error
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectorNamespacesApiService.DeleteConnectorNamespace")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connector_mgmt/v1/kafka_connector_namespaces/{connector_namespace_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector_namespace_id"+"}", _neturl.PathEscape(parameterToString(r.connectorNamespaceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetConnectorNamespaceRequest struct {
	ctx _context.Context
	ApiService ConnectorNamespacesApi
	connectorNamespaceId string
}


func (r ApiGetConnectorNamespaceRequest) Execute() (ConnectorNamespace, *_nethttp.Response, error) {
	return r.ApiService.GetConnectorNamespaceExecute(r)
}

/*
 * GetConnectorNamespace Get a connector namespace
 * Get a connector namespace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connectorNamespaceId The id of the connector namespace
 * @return ApiGetConnectorNamespaceRequest
 */
func (a *ConnectorNamespacesApiService) GetConnectorNamespace(ctx _context.Context, connectorNamespaceId string) ApiGetConnectorNamespaceRequest {
	return ApiGetConnectorNamespaceRequest{
		ApiService: a,
		ctx: ctx,
		connectorNamespaceId: connectorNamespaceId,
	}
}

/*
 * Execute executes the request
 * @return ConnectorNamespace
 */
func (a *ConnectorNamespacesApiService) GetConnectorNamespaceExecute(r ApiGetConnectorNamespaceRequest) (ConnectorNamespace, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectorNamespace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectorNamespacesApiService.GetConnectorNamespace")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connector_mgmt/v1/kafka_connector_namespaces/{connector_namespace_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector_namespace_id"+"}", _neturl.PathEscape(parameterToString(r.connectorNamespaceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListConnectorNamespacesRequest struct {
	ctx _context.Context
	ApiService ConnectorNamespacesApi
	page *string
	size *string
	orderBy *string
	search *string
}

func (r ApiListConnectorNamespacesRequest) Page(page string) ApiListConnectorNamespacesRequest {
	r.page = &page
	return r
}
func (r ApiListConnectorNamespacesRequest) Size(size string) ApiListConnectorNamespacesRequest {
	r.size = &size
	return r
}
func (r ApiListConnectorNamespacesRequest) OrderBy(orderBy string) ApiListConnectorNamespacesRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiListConnectorNamespacesRequest) Search(search string) ApiListConnectorNamespacesRequest {
	r.search = &search
	return r
}

func (r ApiListConnectorNamespacesRequest) Execute() (ConnectorNamespaceList, *_nethttp.Response, error) {
	return r.ApiService.ListConnectorNamespacesExecute(r)
}

/*
 * ListConnectorNamespaces Returns a list of connector namespaces
 * Returns a list of connector namespaces
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListConnectorNamespacesRequest
 */
func (a *ConnectorNamespacesApiService) ListConnectorNamespaces(ctx _context.Context) ApiListConnectorNamespacesRequest {
	return ApiListConnectorNamespacesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ConnectorNamespaceList
 */
func (a *ConnectorNamespacesApiService) ListConnectorNamespacesExecute(r ApiListConnectorNamespacesRequest) (ConnectorNamespaceList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectorNamespaceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectorNamespacesApiService.ListConnectorNamespaces")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connector_mgmt/v1/kafka_connector_namespaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateConnectorNamespaceByIdRequest struct {
	ctx _context.Context
	ApiService ConnectorNamespacesApi
	connectorNamespaceId string
	connectorNamespacePatchRequest *ConnectorNamespacePatchRequest
}

func (r ApiUpdateConnectorNamespaceByIdRequest) ConnectorNamespacePatchRequest(connectorNamespacePatchRequest ConnectorNamespacePatchRequest) ApiUpdateConnectorNamespaceByIdRequest {
	r.connectorNamespacePatchRequest = &connectorNamespacePatchRequest
	return r
}

func (r ApiUpdateConnectorNamespaceByIdRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateConnectorNamespaceByIdExecute(r)
}

/*
 * UpdateConnectorNamespaceById udpate a connector namespace
 * udpate a connector namespace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connectorNamespaceId The id of the connector namespace
 * @return ApiUpdateConnectorNamespaceByIdRequest
 */
func (a *ConnectorNamespacesApiService) UpdateConnectorNamespaceById(ctx _context.Context, connectorNamespaceId string) ApiUpdateConnectorNamespaceByIdRequest {
	return ApiUpdateConnectorNamespaceByIdRequest{
		ApiService: a,
		ctx: ctx,
		connectorNamespaceId: connectorNamespaceId,
	}
}

/*
 * Execute executes the request
 */
func (a *ConnectorNamespacesApiService) UpdateConnectorNamespaceByIdExecute(r ApiUpdateConnectorNamespaceByIdRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectorNamespacesApiService.UpdateConnectorNamespaceById")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/connector_mgmt/v1/kafka_connector_namespaces/{connector_namespace_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector_namespace_id"+"}", _neturl.PathEscape(parameterToString(r.connectorNamespaceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.connectorNamespacePatchRequest == nil {
		return nil, reportError("connectorNamespacePatchRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.connectorNamespacePatchRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
