// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package connectormgmtclient

import (
	_context "context"
	_nethttp "net/http"
	"sync"
)

// Ensure, that ConnectorTypesApiMock does implement ConnectorTypesApi.
// If this is not the case, regenerate this file with moq.
var _ ConnectorTypesApi = &ConnectorTypesApiMock{}

// ConnectorTypesApiMock is a mock implementation of ConnectorTypesApi.
//
// 	func TestSomethingThatUsesConnectorTypesApi(t *testing.T) {
//
// 		// make and configure a mocked ConnectorTypesApi
// 		mockedConnectorTypesApi := &ConnectorTypesApiMock{
// 			GetConnectorTypeByIDFunc: func(ctx _context.Context, connectorTypeId string) ApiGetConnectorTypeByIDRequest {
// 				panic("mock out the GetConnectorTypeByID method")
// 			},
// 			GetConnectorTypeByIDExecuteFunc: func(r ApiGetConnectorTypeByIDRequest) (ConnectorType, *_nethttp.Response, error) {
// 				panic("mock out the GetConnectorTypeByIDExecute method")
// 			},
// 			ListConnectorTypesFunc: func(ctx _context.Context) ApiListConnectorTypesRequest {
// 				panic("mock out the ListConnectorTypes method")
// 			},
// 			ListConnectorTypesExecuteFunc: func(r ApiListConnectorTypesRequest) (ConnectorTypeList, *_nethttp.Response, error) {
// 				panic("mock out the ListConnectorTypesExecute method")
// 			},
// 		}
//
// 		// use mockedConnectorTypesApi in code that requires ConnectorTypesApi
// 		// and then make assertions.
//
// 	}
type ConnectorTypesApiMock struct {
	// GetConnectorTypeByIDFunc mocks the GetConnectorTypeByID method.
	GetConnectorTypeByIDFunc func(ctx _context.Context, connectorTypeId string) ApiGetConnectorTypeByIDRequest

	// GetConnectorTypeByIDExecuteFunc mocks the GetConnectorTypeByIDExecute method.
	GetConnectorTypeByIDExecuteFunc func(r ApiGetConnectorTypeByIDRequest) (ConnectorType, *_nethttp.Response, error)

	// ListConnectorTypesFunc mocks the ListConnectorTypes method.
	ListConnectorTypesFunc func(ctx _context.Context) ApiListConnectorTypesRequest

	// ListConnectorTypesExecuteFunc mocks the ListConnectorTypesExecute method.
	ListConnectorTypesExecuteFunc func(r ApiListConnectorTypesRequest) (ConnectorTypeList, *_nethttp.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetConnectorTypeByID holds details about calls to the GetConnectorTypeByID method.
		GetConnectorTypeByID []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ConnectorTypeId is the connectorTypeId argument value.
			ConnectorTypeId string
		}
		// GetConnectorTypeByIDExecute holds details about calls to the GetConnectorTypeByIDExecute method.
		GetConnectorTypeByIDExecute []struct {
			// R is the r argument value.
			R ApiGetConnectorTypeByIDRequest
		}
		// ListConnectorTypes holds details about calls to the ListConnectorTypes method.
		ListConnectorTypes []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// ListConnectorTypesExecute holds details about calls to the ListConnectorTypesExecute method.
		ListConnectorTypesExecute []struct {
			// R is the r argument value.
			R ApiListConnectorTypesRequest
		}
	}
	lockGetConnectorTypeByID        sync.RWMutex
	lockGetConnectorTypeByIDExecute sync.RWMutex
	lockListConnectorTypes          sync.RWMutex
	lockListConnectorTypesExecute   sync.RWMutex
}

// GetConnectorTypeByID calls GetConnectorTypeByIDFunc.
func (mock *ConnectorTypesApiMock) GetConnectorTypeByID(ctx _context.Context, connectorTypeId string) ApiGetConnectorTypeByIDRequest {
	if mock.GetConnectorTypeByIDFunc == nil {
		panic("ConnectorTypesApiMock.GetConnectorTypeByIDFunc: method is nil but ConnectorTypesApi.GetConnectorTypeByID was just called")
	}
	callInfo := struct {
		Ctx             _context.Context
		ConnectorTypeId string
	}{
		Ctx:             ctx,
		ConnectorTypeId: connectorTypeId,
	}
	mock.lockGetConnectorTypeByID.Lock()
	mock.calls.GetConnectorTypeByID = append(mock.calls.GetConnectorTypeByID, callInfo)
	mock.lockGetConnectorTypeByID.Unlock()
	return mock.GetConnectorTypeByIDFunc(ctx, connectorTypeId)
}

// GetConnectorTypeByIDCalls gets all the calls that were made to GetConnectorTypeByID.
// Check the length with:
//     len(mockedConnectorTypesApi.GetConnectorTypeByIDCalls())
func (mock *ConnectorTypesApiMock) GetConnectorTypeByIDCalls() []struct {
	Ctx             _context.Context
	ConnectorTypeId string
} {
	var calls []struct {
		Ctx             _context.Context
		ConnectorTypeId string
	}
	mock.lockGetConnectorTypeByID.RLock()
	calls = mock.calls.GetConnectorTypeByID
	mock.lockGetConnectorTypeByID.RUnlock()
	return calls
}

// GetConnectorTypeByIDExecute calls GetConnectorTypeByIDExecuteFunc.
func (mock *ConnectorTypesApiMock) GetConnectorTypeByIDExecute(r ApiGetConnectorTypeByIDRequest) (ConnectorType, *_nethttp.Response, error) {
	if mock.GetConnectorTypeByIDExecuteFunc == nil {
		panic("ConnectorTypesApiMock.GetConnectorTypeByIDExecuteFunc: method is nil but ConnectorTypesApi.GetConnectorTypeByIDExecute was just called")
	}
	callInfo := struct {
		R ApiGetConnectorTypeByIDRequest
	}{
		R: r,
	}
	mock.lockGetConnectorTypeByIDExecute.Lock()
	mock.calls.GetConnectorTypeByIDExecute = append(mock.calls.GetConnectorTypeByIDExecute, callInfo)
	mock.lockGetConnectorTypeByIDExecute.Unlock()
	return mock.GetConnectorTypeByIDExecuteFunc(r)
}

// GetConnectorTypeByIDExecuteCalls gets all the calls that were made to GetConnectorTypeByIDExecute.
// Check the length with:
//     len(mockedConnectorTypesApi.GetConnectorTypeByIDExecuteCalls())
func (mock *ConnectorTypesApiMock) GetConnectorTypeByIDExecuteCalls() []struct {
	R ApiGetConnectorTypeByIDRequest
} {
	var calls []struct {
		R ApiGetConnectorTypeByIDRequest
	}
	mock.lockGetConnectorTypeByIDExecute.RLock()
	calls = mock.calls.GetConnectorTypeByIDExecute
	mock.lockGetConnectorTypeByIDExecute.RUnlock()
	return calls
}

// ListConnectorTypes calls ListConnectorTypesFunc.
func (mock *ConnectorTypesApiMock) ListConnectorTypes(ctx _context.Context) ApiListConnectorTypesRequest {
	if mock.ListConnectorTypesFunc == nil {
		panic("ConnectorTypesApiMock.ListConnectorTypesFunc: method is nil but ConnectorTypesApi.ListConnectorTypes was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockListConnectorTypes.Lock()
	mock.calls.ListConnectorTypes = append(mock.calls.ListConnectorTypes, callInfo)
	mock.lockListConnectorTypes.Unlock()
	return mock.ListConnectorTypesFunc(ctx)
}

// ListConnectorTypesCalls gets all the calls that were made to ListConnectorTypes.
// Check the length with:
//     len(mockedConnectorTypesApi.ListConnectorTypesCalls())
func (mock *ConnectorTypesApiMock) ListConnectorTypesCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockListConnectorTypes.RLock()
	calls = mock.calls.ListConnectorTypes
	mock.lockListConnectorTypes.RUnlock()
	return calls
}

// ListConnectorTypesExecute calls ListConnectorTypesExecuteFunc.
func (mock *ConnectorTypesApiMock) ListConnectorTypesExecute(r ApiListConnectorTypesRequest) (ConnectorTypeList, *_nethttp.Response, error) {
	if mock.ListConnectorTypesExecuteFunc == nil {
		panic("ConnectorTypesApiMock.ListConnectorTypesExecuteFunc: method is nil but ConnectorTypesApi.ListConnectorTypesExecute was just called")
	}
	callInfo := struct {
		R ApiListConnectorTypesRequest
	}{
		R: r,
	}
	mock.lockListConnectorTypesExecute.Lock()
	mock.calls.ListConnectorTypesExecute = append(mock.calls.ListConnectorTypesExecute, callInfo)
	mock.lockListConnectorTypesExecute.Unlock()
	return mock.ListConnectorTypesExecuteFunc(r)
}

// ListConnectorTypesExecuteCalls gets all the calls that were made to ListConnectorTypesExecute.
// Check the length with:
//     len(mockedConnectorTypesApi.ListConnectorTypesExecuteCalls())
func (mock *ConnectorTypesApiMock) ListConnectorTypesExecuteCalls() []struct {
	R ApiListConnectorTypesRequest
} {
	var calls []struct {
		R ApiListConnectorTypesRequest
	}
	mock.lockListConnectorTypesExecute.RLock()
	calls = mock.calls.ListConnectorTypesExecute
	mock.lockListConnectorTypesExecute.RUnlock()
	return calls
}