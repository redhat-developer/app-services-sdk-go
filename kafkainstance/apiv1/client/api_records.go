/*
 * Kafka Instance API
 *
 * API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs
 *
 * API version: 0.12.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

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

type RecordsApi interface {

	/*
	 * ConsumeRecords Consume records from a topic
	 * Consume a limited number of records from a topic, optionally specifying a partition and an absolute offset or timestamp as the starting point for message retrieval.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param topicName Topic name
	 * @return ApiConsumeRecordsRequest
	 */
	ConsumeRecords(ctx _context.Context, topicName string) ApiConsumeRecordsRequest

	/*
	 * ConsumeRecordsExecute executes the request
	 * @return RecordList
	 */
	ConsumeRecordsExecute(r ApiConsumeRecordsRequest) (RecordList, *_nethttp.Response, error)

	/*
	 * ProduceRecord Send a record to a topic
	 * Produce (write) a single record to a topic.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param topicName Topic name
	 * @return ApiProduceRecordRequest
	 */
	ProduceRecord(ctx _context.Context, topicName string) ApiProduceRecordRequest

	/*
	 * ProduceRecordExecute executes the request
	 * @return Record
	 */
	ProduceRecordExecute(r ApiProduceRecordRequest) (Record, *_nethttp.Response, error)
}

// RecordsApiService RecordsApi service
type RecordsApiService service

type ApiConsumeRecordsRequest struct {
	ctx _context.Context
	ApiService RecordsApi
	topicName string
	include *[]RecordIncludedProperty
	limit *int32
	maxValueLength *int32
	offset *int32
	partition *int32
	timestamp *interface{}
}

func (r ApiConsumeRecordsRequest) Include(include []RecordIncludedProperty) ApiConsumeRecordsRequest {
	r.include = &include
	return r
}
func (r ApiConsumeRecordsRequest) Limit(limit int32) ApiConsumeRecordsRequest {
	r.limit = &limit
	return r
}
func (r ApiConsumeRecordsRequest) MaxValueLength(maxValueLength int32) ApiConsumeRecordsRequest {
	r.maxValueLength = &maxValueLength
	return r
}
func (r ApiConsumeRecordsRequest) Offset(offset int32) ApiConsumeRecordsRequest {
	r.offset = &offset
	return r
}
func (r ApiConsumeRecordsRequest) Partition(partition int32) ApiConsumeRecordsRequest {
	r.partition = &partition
	return r
}
func (r ApiConsumeRecordsRequest) Timestamp(timestamp interface{}) ApiConsumeRecordsRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiConsumeRecordsRequest) Execute() (RecordList, *_nethttp.Response, error) {
	return r.ApiService.ConsumeRecordsExecute(r)
}

/*
 * ConsumeRecords Consume records from a topic
 * Consume a limited number of records from a topic, optionally specifying a partition and an absolute offset or timestamp as the starting point for message retrieval.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param topicName Topic name
 * @return ApiConsumeRecordsRequest
 */
func (a *RecordsApiService) ConsumeRecords(ctx _context.Context, topicName string) ApiConsumeRecordsRequest {
	return ApiConsumeRecordsRequest{
		ApiService: a,
		ctx: ctx,
		topicName: topicName,
	}
}

/*
 * Execute executes the request
 * @return RecordList
 */
func (a *RecordsApiService) ConsumeRecordsExecute(r ApiConsumeRecordsRequest) (RecordList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RecordList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.ConsumeRecords")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/topics/{topicName}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"topicName"+"}", _neturl.PathEscape(parameterToString(r.topicName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, "csv"))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.maxValueLength != nil {
		localVarQueryParams.Add("maxValueLength", parameterToString(*r.maxValueLength, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.partition != nil {
		localVarQueryParams.Add("partition", parameterToString(*r.partition, ""))
	}
	if r.timestamp != nil {
		localVarQueryParams.Add("timestamp", parameterToString(*r.timestamp, ""))
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
		if localVarHTTPResponse.StatusCode == 403 {
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiProduceRecordRequest struct {
	ctx _context.Context
	ApiService RecordsApi
	topicName string
	record *Record
}

func (r ApiProduceRecordRequest) Record(record Record) ApiProduceRecordRequest {
	r.record = &record
	return r
}

func (r ApiProduceRecordRequest) Execute() (Record, *_nethttp.Response, error) {
	return r.ApiService.ProduceRecordExecute(r)
}

/*
 * ProduceRecord Send a record to a topic
 * Produce (write) a single record to a topic.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param topicName Topic name
 * @return ApiProduceRecordRequest
 */
func (a *RecordsApiService) ProduceRecord(ctx _context.Context, topicName string) ApiProduceRecordRequest {
	return ApiProduceRecordRequest{
		ApiService: a,
		ctx: ctx,
		topicName: topicName,
	}
}

/*
 * Execute executes the request
 * @return Record
 */
func (a *RecordsApiService) ProduceRecordExecute(r ApiProduceRecordRequest) (Record, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Record
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.ProduceRecord")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/topics/{topicName}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"topicName"+"}", _neturl.PathEscape(parameterToString(r.topicName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.record == nil {
		return localVarReturnValue, nil, reportError("record is required and must be specified")
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
	localVarPostBody = r.record
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
		if localVarHTTPResponse.StatusCode == 403 {
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
