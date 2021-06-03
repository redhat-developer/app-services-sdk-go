// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package kafkamgmtclient

import (
	_context "context"
	_nethttp "net/http"
	"sync"
)

// Ensure, that DefaultApiMock does implement DefaultApi.
// If this is not the case, regenerate this file with moq.
var _ DefaultApi = &DefaultApiMock{}

// DefaultApiMock is a mock implementation of DefaultApi.
//
// 	func TestSomethingThatUsesDefaultApi(t *testing.T) {
//
// 		// make and configure a mocked DefaultApi
// 		mockedDefaultApi := &DefaultApiMock{
// 			CreateKafkaFunc: func(ctx _context.Context) ApiCreateKafkaRequest {
// 				panic("mock out the CreateKafka method")
// 			},
// 			CreateKafkaExecuteFunc: func(r ApiCreateKafkaRequest) (KafkaRequest, *_nethttp.Response, error) {
// 				panic("mock out the CreateKafkaExecute method")
// 			},
// 			DeleteKafkaByIdFunc: func(ctx _context.Context, id string) ApiDeleteKafkaByIdRequest {
// 				panic("mock out the DeleteKafkaById method")
// 			},
// 			DeleteKafkaByIdExecuteFunc: func(r ApiDeleteKafkaByIdRequest) (Error, *_nethttp.Response, error) {
// 				panic("mock out the DeleteKafkaByIdExecute method")
// 			},
// 			GetCloudProviderRegionsFunc: func(ctx _context.Context, id string) ApiGetCloudProviderRegionsRequest {
// 				panic("mock out the GetCloudProviderRegions method")
// 			},
// 			GetCloudProviderRegionsExecuteFunc: func(r ApiGetCloudProviderRegionsRequest) (CloudRegionList, *_nethttp.Response, error) {
// 				panic("mock out the GetCloudProviderRegionsExecute method")
// 			},
// 			GetCloudProvidersFunc: func(ctx _context.Context) ApiGetCloudProvidersRequest {
// 				panic("mock out the GetCloudProviders method")
// 			},
// 			GetCloudProvidersExecuteFunc: func(r ApiGetCloudProvidersRequest) (CloudProviderList, *_nethttp.Response, error) {
// 				panic("mock out the GetCloudProvidersExecute method")
// 			},
// 			GetKafkaByIdFunc: func(ctx _context.Context, id string) ApiGetKafkaByIdRequest {
// 				panic("mock out the GetKafkaById method")
// 			},
// 			GetKafkaByIdExecuteFunc: func(r ApiGetKafkaByIdRequest) (KafkaRequest, *_nethttp.Response, error) {
// 				panic("mock out the GetKafkaByIdExecute method")
// 			},
// 			GetKafkasFunc: func(ctx _context.Context) ApiGetKafkasRequest {
// 				panic("mock out the GetKafkas method")
// 			},
// 			GetKafkasExecuteFunc: func(r ApiGetKafkasRequest) (KafkaRequestList, *_nethttp.Response, error) {
// 				panic("mock out the GetKafkasExecute method")
// 			},
// 			GetMetricsByInstantQueryFunc: func(ctx _context.Context, id string) ApiGetMetricsByInstantQueryRequest {
// 				panic("mock out the GetMetricsByInstantQuery method")
// 			},
// 			GetMetricsByInstantQueryExecuteFunc: func(r ApiGetMetricsByInstantQueryRequest) (MetricsInstantQueryList, *_nethttp.Response, error) {
// 				panic("mock out the GetMetricsByInstantQueryExecute method")
// 			},
// 			GetMetricsByRangeQueryFunc: func(ctx _context.Context, id string) ApiGetMetricsByRangeQueryRequest {
// 				panic("mock out the GetMetricsByRangeQuery method")
// 			},
// 			GetMetricsByRangeQueryExecuteFunc: func(r ApiGetMetricsByRangeQueryRequest) (MetricsRangeQueryList, *_nethttp.Response, error) {
// 				panic("mock out the GetMetricsByRangeQueryExecute method")
// 			},
// 			GetServiceStatusFunc: func(ctx _context.Context) ApiGetServiceStatusRequest {
// 				panic("mock out the GetServiceStatus method")
// 			},
// 			GetServiceStatusExecuteFunc: func(r ApiGetServiceStatusRequest) (ServiceStatus, *_nethttp.Response, error) {
// 				panic("mock out the GetServiceStatusExecute method")
// 			},
// 			GetVersionMetadataFunc: func(ctx _context.Context) ApiGetVersionMetadataRequest {
// 				panic("mock out the GetVersionMetadata method")
// 			},
// 			GetVersionMetadataExecuteFunc: func(r ApiGetVersionMetadataRequest) (VersionMetadata, *_nethttp.Response, error) {
// 				panic("mock out the GetVersionMetadataExecute method")
// 			},
// 		}
//
// 		// use mockedDefaultApi in code that requires DefaultApi
// 		// and then make assertions.
//
// 	}
type DefaultApiMock struct {
	// CreateKafkaFunc mocks the CreateKafka method.
	CreateKafkaFunc func(ctx _context.Context) ApiCreateKafkaRequest

	// CreateKafkaExecuteFunc mocks the CreateKafkaExecute method.
	CreateKafkaExecuteFunc func(r ApiCreateKafkaRequest) (KafkaRequest, *_nethttp.Response, error)

	// DeleteKafkaByIdFunc mocks the DeleteKafkaById method.
	DeleteKafkaByIdFunc func(ctx _context.Context, id string) ApiDeleteKafkaByIdRequest

	// DeleteKafkaByIdExecuteFunc mocks the DeleteKafkaByIdExecute method.
	DeleteKafkaByIdExecuteFunc func(r ApiDeleteKafkaByIdRequest) (Error, *_nethttp.Response, error)

	// GetCloudProviderRegionsFunc mocks the GetCloudProviderRegions method.
	GetCloudProviderRegionsFunc func(ctx _context.Context, id string) ApiGetCloudProviderRegionsRequest

	// GetCloudProviderRegionsExecuteFunc mocks the GetCloudProviderRegionsExecute method.
	GetCloudProviderRegionsExecuteFunc func(r ApiGetCloudProviderRegionsRequest) (CloudRegionList, *_nethttp.Response, error)

	// GetCloudProvidersFunc mocks the GetCloudProviders method.
	GetCloudProvidersFunc func(ctx _context.Context) ApiGetCloudProvidersRequest

	// GetCloudProvidersExecuteFunc mocks the GetCloudProvidersExecute method.
	GetCloudProvidersExecuteFunc func(r ApiGetCloudProvidersRequest) (CloudProviderList, *_nethttp.Response, error)

	// GetKafkaByIdFunc mocks the GetKafkaById method.
	GetKafkaByIdFunc func(ctx _context.Context, id string) ApiGetKafkaByIdRequest

	// GetKafkaByIdExecuteFunc mocks the GetKafkaByIdExecute method.
	GetKafkaByIdExecuteFunc func(r ApiGetKafkaByIdRequest) (KafkaRequest, *_nethttp.Response, error)

	// GetKafkasFunc mocks the GetKafkas method.
	GetKafkasFunc func(ctx _context.Context) ApiGetKafkasRequest

	// GetKafkasExecuteFunc mocks the GetKafkasExecute method.
	GetKafkasExecuteFunc func(r ApiGetKafkasRequest) (KafkaRequestList, *_nethttp.Response, error)

	// GetMetricsByInstantQueryFunc mocks the GetMetricsByInstantQuery method.
	GetMetricsByInstantQueryFunc func(ctx _context.Context, id string) ApiGetMetricsByInstantQueryRequest

	// GetMetricsByInstantQueryExecuteFunc mocks the GetMetricsByInstantQueryExecute method.
	GetMetricsByInstantQueryExecuteFunc func(r ApiGetMetricsByInstantQueryRequest) (MetricsInstantQueryList, *_nethttp.Response, error)

	// GetMetricsByRangeQueryFunc mocks the GetMetricsByRangeQuery method.
	GetMetricsByRangeQueryFunc func(ctx _context.Context, id string) ApiGetMetricsByRangeQueryRequest

	// GetMetricsByRangeQueryExecuteFunc mocks the GetMetricsByRangeQueryExecute method.
	GetMetricsByRangeQueryExecuteFunc func(r ApiGetMetricsByRangeQueryRequest) (MetricsRangeQueryList, *_nethttp.Response, error)

	// GetServiceStatusFunc mocks the GetServiceStatus method.
	GetServiceStatusFunc func(ctx _context.Context) ApiGetServiceStatusRequest

	// GetServiceStatusExecuteFunc mocks the GetServiceStatusExecute method.
	GetServiceStatusExecuteFunc func(r ApiGetServiceStatusRequest) (ServiceStatus, *_nethttp.Response, error)

	// GetVersionMetadataFunc mocks the GetVersionMetadata method.
	GetVersionMetadataFunc func(ctx _context.Context) ApiGetVersionMetadataRequest

	// GetVersionMetadataExecuteFunc mocks the GetVersionMetadataExecute method.
	GetVersionMetadataExecuteFunc func(r ApiGetVersionMetadataRequest) (VersionMetadata, *_nethttp.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateKafka holds details about calls to the CreateKafka method.
		CreateKafka []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// CreateKafkaExecute holds details about calls to the CreateKafkaExecute method.
		CreateKafkaExecute []struct {
			// R is the r argument value.
			R ApiCreateKafkaRequest
		}
		// DeleteKafkaById holds details about calls to the DeleteKafkaById method.
		DeleteKafkaById []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ID is the id argument value.
			ID string
		}
		// DeleteKafkaByIdExecute holds details about calls to the DeleteKafkaByIdExecute method.
		DeleteKafkaByIdExecute []struct {
			// R is the r argument value.
			R ApiDeleteKafkaByIdRequest
		}
		// GetCloudProviderRegions holds details about calls to the GetCloudProviderRegions method.
		GetCloudProviderRegions []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ID is the id argument value.
			ID string
		}
		// GetCloudProviderRegionsExecute holds details about calls to the GetCloudProviderRegionsExecute method.
		GetCloudProviderRegionsExecute []struct {
			// R is the r argument value.
			R ApiGetCloudProviderRegionsRequest
		}
		// GetCloudProviders holds details about calls to the GetCloudProviders method.
		GetCloudProviders []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetCloudProvidersExecute holds details about calls to the GetCloudProvidersExecute method.
		GetCloudProvidersExecute []struct {
			// R is the r argument value.
			R ApiGetCloudProvidersRequest
		}
		// GetKafkaById holds details about calls to the GetKafkaById method.
		GetKafkaById []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ID is the id argument value.
			ID string
		}
		// GetKafkaByIdExecute holds details about calls to the GetKafkaByIdExecute method.
		GetKafkaByIdExecute []struct {
			// R is the r argument value.
			R ApiGetKafkaByIdRequest
		}
		// GetKafkas holds details about calls to the GetKafkas method.
		GetKafkas []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetKafkasExecute holds details about calls to the GetKafkasExecute method.
		GetKafkasExecute []struct {
			// R is the r argument value.
			R ApiGetKafkasRequest
		}
		// GetMetricsByInstantQuery holds details about calls to the GetMetricsByInstantQuery method.
		GetMetricsByInstantQuery []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ID is the id argument value.
			ID string
		}
		// GetMetricsByInstantQueryExecute holds details about calls to the GetMetricsByInstantQueryExecute method.
		GetMetricsByInstantQueryExecute []struct {
			// R is the r argument value.
			R ApiGetMetricsByInstantQueryRequest
		}
		// GetMetricsByRangeQuery holds details about calls to the GetMetricsByRangeQuery method.
		GetMetricsByRangeQuery []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ID is the id argument value.
			ID string
		}
		// GetMetricsByRangeQueryExecute holds details about calls to the GetMetricsByRangeQueryExecute method.
		GetMetricsByRangeQueryExecute []struct {
			// R is the r argument value.
			R ApiGetMetricsByRangeQueryRequest
		}
		// GetServiceStatus holds details about calls to the GetServiceStatus method.
		GetServiceStatus []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetServiceStatusExecute holds details about calls to the GetServiceStatusExecute method.
		GetServiceStatusExecute []struct {
			// R is the r argument value.
			R ApiGetServiceStatusRequest
		}
		// GetVersionMetadata holds details about calls to the GetVersionMetadata method.
		GetVersionMetadata []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetVersionMetadataExecute holds details about calls to the GetVersionMetadataExecute method.
		GetVersionMetadataExecute []struct {
			// R is the r argument value.
			R ApiGetVersionMetadataRequest
		}
	}
	lockCreateKafka                     sync.RWMutex
	lockCreateKafkaExecute              sync.RWMutex
	lockDeleteKafkaById                 sync.RWMutex
	lockDeleteKafkaByIdExecute          sync.RWMutex
	lockGetCloudProviderRegions         sync.RWMutex
	lockGetCloudProviderRegionsExecute  sync.RWMutex
	lockGetCloudProviders               sync.RWMutex
	lockGetCloudProvidersExecute        sync.RWMutex
	lockGetKafkaById                    sync.RWMutex
	lockGetKafkaByIdExecute             sync.RWMutex
	lockGetKafkas                       sync.RWMutex
	lockGetKafkasExecute                sync.RWMutex
	lockGetMetricsByInstantQuery        sync.RWMutex
	lockGetMetricsByInstantQueryExecute sync.RWMutex
	lockGetMetricsByRangeQuery          sync.RWMutex
	lockGetMetricsByRangeQueryExecute   sync.RWMutex
	lockGetServiceStatus                sync.RWMutex
	lockGetServiceStatusExecute         sync.RWMutex
	lockGetVersionMetadata              sync.RWMutex
	lockGetVersionMetadataExecute       sync.RWMutex
}

// CreateKafka calls CreateKafkaFunc.
func (mock *DefaultApiMock) CreateKafka(ctx _context.Context) ApiCreateKafkaRequest {
	if mock.CreateKafkaFunc == nil {
		panic("DefaultApiMock.CreateKafkaFunc: method is nil but DefaultApi.CreateKafka was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCreateKafka.Lock()
	mock.calls.CreateKafka = append(mock.calls.CreateKafka, callInfo)
	mock.lockCreateKafka.Unlock()
	return mock.CreateKafkaFunc(ctx)
}

// CreateKafkaCalls gets all the calls that were made to CreateKafka.
// Check the length with:
//     len(mockedDefaultApi.CreateKafkaCalls())
func (mock *DefaultApiMock) CreateKafkaCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockCreateKafka.RLock()
	calls = mock.calls.CreateKafka
	mock.lockCreateKafka.RUnlock()
	return calls
}

// CreateKafkaExecute calls CreateKafkaExecuteFunc.
func (mock *DefaultApiMock) CreateKafkaExecute(r ApiCreateKafkaRequest) (KafkaRequest, *_nethttp.Response, error) {
	if mock.CreateKafkaExecuteFunc == nil {
		panic("DefaultApiMock.CreateKafkaExecuteFunc: method is nil but DefaultApi.CreateKafkaExecute was just called")
	}
	callInfo := struct {
		R ApiCreateKafkaRequest
	}{
		R: r,
	}
	mock.lockCreateKafkaExecute.Lock()
	mock.calls.CreateKafkaExecute = append(mock.calls.CreateKafkaExecute, callInfo)
	mock.lockCreateKafkaExecute.Unlock()
	return mock.CreateKafkaExecuteFunc(r)
}

// CreateKafkaExecuteCalls gets all the calls that were made to CreateKafkaExecute.
// Check the length with:
//     len(mockedDefaultApi.CreateKafkaExecuteCalls())
func (mock *DefaultApiMock) CreateKafkaExecuteCalls() []struct {
	R ApiCreateKafkaRequest
} {
	var calls []struct {
		R ApiCreateKafkaRequest
	}
	mock.lockCreateKafkaExecute.RLock()
	calls = mock.calls.CreateKafkaExecute
	mock.lockCreateKafkaExecute.RUnlock()
	return calls
}

// DeleteKafkaById calls DeleteKafkaByIdFunc.
func (mock *DefaultApiMock) DeleteKafkaById(ctx _context.Context, id string) ApiDeleteKafkaByIdRequest {
	if mock.DeleteKafkaByIdFunc == nil {
		panic("DefaultApiMock.DeleteKafkaByIdFunc: method is nil but DefaultApi.DeleteKafkaById was just called")
	}
	callInfo := struct {
		Ctx _context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteKafkaById.Lock()
	mock.calls.DeleteKafkaById = append(mock.calls.DeleteKafkaById, callInfo)
	mock.lockDeleteKafkaById.Unlock()
	return mock.DeleteKafkaByIdFunc(ctx, id)
}

// DeleteKafkaByIdCalls gets all the calls that were made to DeleteKafkaById.
// Check the length with:
//     len(mockedDefaultApi.DeleteKafkaByIdCalls())
func (mock *DefaultApiMock) DeleteKafkaByIdCalls() []struct {
	Ctx _context.Context
	ID  string
} {
	var calls []struct {
		Ctx _context.Context
		ID  string
	}
	mock.lockDeleteKafkaById.RLock()
	calls = mock.calls.DeleteKafkaById
	mock.lockDeleteKafkaById.RUnlock()
	return calls
}

// DeleteKafkaByIdExecute calls DeleteKafkaByIdExecuteFunc.
func (mock *DefaultApiMock) DeleteKafkaByIdExecute(r ApiDeleteKafkaByIdRequest) (Error, *_nethttp.Response, error) {
	if mock.DeleteKafkaByIdExecuteFunc == nil {
		panic("DefaultApiMock.DeleteKafkaByIdExecuteFunc: method is nil but DefaultApi.DeleteKafkaByIdExecute was just called")
	}
	callInfo := struct {
		R ApiDeleteKafkaByIdRequest
	}{
		R: r,
	}
	mock.lockDeleteKafkaByIdExecute.Lock()
	mock.calls.DeleteKafkaByIdExecute = append(mock.calls.DeleteKafkaByIdExecute, callInfo)
	mock.lockDeleteKafkaByIdExecute.Unlock()
	return mock.DeleteKafkaByIdExecuteFunc(r)
}

// DeleteKafkaByIdExecuteCalls gets all the calls that were made to DeleteKafkaByIdExecute.
// Check the length with:
//     len(mockedDefaultApi.DeleteKafkaByIdExecuteCalls())
func (mock *DefaultApiMock) DeleteKafkaByIdExecuteCalls() []struct {
	R ApiDeleteKafkaByIdRequest
} {
	var calls []struct {
		R ApiDeleteKafkaByIdRequest
	}
	mock.lockDeleteKafkaByIdExecute.RLock()
	calls = mock.calls.DeleteKafkaByIdExecute
	mock.lockDeleteKafkaByIdExecute.RUnlock()
	return calls
}

// GetCloudProviderRegions calls GetCloudProviderRegionsFunc.
func (mock *DefaultApiMock) GetCloudProviderRegions(ctx _context.Context, id string) ApiGetCloudProviderRegionsRequest {
	if mock.GetCloudProviderRegionsFunc == nil {
		panic("DefaultApiMock.GetCloudProviderRegionsFunc: method is nil but DefaultApi.GetCloudProviderRegions was just called")
	}
	callInfo := struct {
		Ctx _context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetCloudProviderRegions.Lock()
	mock.calls.GetCloudProviderRegions = append(mock.calls.GetCloudProviderRegions, callInfo)
	mock.lockGetCloudProviderRegions.Unlock()
	return mock.GetCloudProviderRegionsFunc(ctx, id)
}

// GetCloudProviderRegionsCalls gets all the calls that were made to GetCloudProviderRegions.
// Check the length with:
//     len(mockedDefaultApi.GetCloudProviderRegionsCalls())
func (mock *DefaultApiMock) GetCloudProviderRegionsCalls() []struct {
	Ctx _context.Context
	ID  string
} {
	var calls []struct {
		Ctx _context.Context
		ID  string
	}
	mock.lockGetCloudProviderRegions.RLock()
	calls = mock.calls.GetCloudProviderRegions
	mock.lockGetCloudProviderRegions.RUnlock()
	return calls
}

// GetCloudProviderRegionsExecute calls GetCloudProviderRegionsExecuteFunc.
func (mock *DefaultApiMock) GetCloudProviderRegionsExecute(r ApiGetCloudProviderRegionsRequest) (CloudRegionList, *_nethttp.Response, error) {
	if mock.GetCloudProviderRegionsExecuteFunc == nil {
		panic("DefaultApiMock.GetCloudProviderRegionsExecuteFunc: method is nil but DefaultApi.GetCloudProviderRegionsExecute was just called")
	}
	callInfo := struct {
		R ApiGetCloudProviderRegionsRequest
	}{
		R: r,
	}
	mock.lockGetCloudProviderRegionsExecute.Lock()
	mock.calls.GetCloudProviderRegionsExecute = append(mock.calls.GetCloudProviderRegionsExecute, callInfo)
	mock.lockGetCloudProviderRegionsExecute.Unlock()
	return mock.GetCloudProviderRegionsExecuteFunc(r)
}

// GetCloudProviderRegionsExecuteCalls gets all the calls that were made to GetCloudProviderRegionsExecute.
// Check the length with:
//     len(mockedDefaultApi.GetCloudProviderRegionsExecuteCalls())
func (mock *DefaultApiMock) GetCloudProviderRegionsExecuteCalls() []struct {
	R ApiGetCloudProviderRegionsRequest
} {
	var calls []struct {
		R ApiGetCloudProviderRegionsRequest
	}
	mock.lockGetCloudProviderRegionsExecute.RLock()
	calls = mock.calls.GetCloudProviderRegionsExecute
	mock.lockGetCloudProviderRegionsExecute.RUnlock()
	return calls
}

// GetCloudProviders calls GetCloudProvidersFunc.
func (mock *DefaultApiMock) GetCloudProviders(ctx _context.Context) ApiGetCloudProvidersRequest {
	if mock.GetCloudProvidersFunc == nil {
		panic("DefaultApiMock.GetCloudProvidersFunc: method is nil but DefaultApi.GetCloudProviders was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetCloudProviders.Lock()
	mock.calls.GetCloudProviders = append(mock.calls.GetCloudProviders, callInfo)
	mock.lockGetCloudProviders.Unlock()
	return mock.GetCloudProvidersFunc(ctx)
}

// GetCloudProvidersCalls gets all the calls that were made to GetCloudProviders.
// Check the length with:
//     len(mockedDefaultApi.GetCloudProvidersCalls())
func (mock *DefaultApiMock) GetCloudProvidersCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetCloudProviders.RLock()
	calls = mock.calls.GetCloudProviders
	mock.lockGetCloudProviders.RUnlock()
	return calls
}

// GetCloudProvidersExecute calls GetCloudProvidersExecuteFunc.
func (mock *DefaultApiMock) GetCloudProvidersExecute(r ApiGetCloudProvidersRequest) (CloudProviderList, *_nethttp.Response, error) {
	if mock.GetCloudProvidersExecuteFunc == nil {
		panic("DefaultApiMock.GetCloudProvidersExecuteFunc: method is nil but DefaultApi.GetCloudProvidersExecute was just called")
	}
	callInfo := struct {
		R ApiGetCloudProvidersRequest
	}{
		R: r,
	}
	mock.lockGetCloudProvidersExecute.Lock()
	mock.calls.GetCloudProvidersExecute = append(mock.calls.GetCloudProvidersExecute, callInfo)
	mock.lockGetCloudProvidersExecute.Unlock()
	return mock.GetCloudProvidersExecuteFunc(r)
}

// GetCloudProvidersExecuteCalls gets all the calls that were made to GetCloudProvidersExecute.
// Check the length with:
//     len(mockedDefaultApi.GetCloudProvidersExecuteCalls())
func (mock *DefaultApiMock) GetCloudProvidersExecuteCalls() []struct {
	R ApiGetCloudProvidersRequest
} {
	var calls []struct {
		R ApiGetCloudProvidersRequest
	}
	mock.lockGetCloudProvidersExecute.RLock()
	calls = mock.calls.GetCloudProvidersExecute
	mock.lockGetCloudProvidersExecute.RUnlock()
	return calls
}

// GetKafkaById calls GetKafkaByIdFunc.
func (mock *DefaultApiMock) GetKafkaById(ctx _context.Context, id string) ApiGetKafkaByIdRequest {
	if mock.GetKafkaByIdFunc == nil {
		panic("DefaultApiMock.GetKafkaByIdFunc: method is nil but DefaultApi.GetKafkaById was just called")
	}
	callInfo := struct {
		Ctx _context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetKafkaById.Lock()
	mock.calls.GetKafkaById = append(mock.calls.GetKafkaById, callInfo)
	mock.lockGetKafkaById.Unlock()
	return mock.GetKafkaByIdFunc(ctx, id)
}

// GetKafkaByIdCalls gets all the calls that were made to GetKafkaById.
// Check the length with:
//     len(mockedDefaultApi.GetKafkaByIdCalls())
func (mock *DefaultApiMock) GetKafkaByIdCalls() []struct {
	Ctx _context.Context
	ID  string
} {
	var calls []struct {
		Ctx _context.Context
		ID  string
	}
	mock.lockGetKafkaById.RLock()
	calls = mock.calls.GetKafkaById
	mock.lockGetKafkaById.RUnlock()
	return calls
}

// GetKafkaByIdExecute calls GetKafkaByIdExecuteFunc.
func (mock *DefaultApiMock) GetKafkaByIdExecute(r ApiGetKafkaByIdRequest) (KafkaRequest, *_nethttp.Response, error) {
	if mock.GetKafkaByIdExecuteFunc == nil {
		panic("DefaultApiMock.GetKafkaByIdExecuteFunc: method is nil but DefaultApi.GetKafkaByIdExecute was just called")
	}
	callInfo := struct {
		R ApiGetKafkaByIdRequest
	}{
		R: r,
	}
	mock.lockGetKafkaByIdExecute.Lock()
	mock.calls.GetKafkaByIdExecute = append(mock.calls.GetKafkaByIdExecute, callInfo)
	mock.lockGetKafkaByIdExecute.Unlock()
	return mock.GetKafkaByIdExecuteFunc(r)
}

// GetKafkaByIdExecuteCalls gets all the calls that were made to GetKafkaByIdExecute.
// Check the length with:
//     len(mockedDefaultApi.GetKafkaByIdExecuteCalls())
func (mock *DefaultApiMock) GetKafkaByIdExecuteCalls() []struct {
	R ApiGetKafkaByIdRequest
} {
	var calls []struct {
		R ApiGetKafkaByIdRequest
	}
	mock.lockGetKafkaByIdExecute.RLock()
	calls = mock.calls.GetKafkaByIdExecute
	mock.lockGetKafkaByIdExecute.RUnlock()
	return calls
}

// GetKafkas calls GetKafkasFunc.
func (mock *DefaultApiMock) GetKafkas(ctx _context.Context) ApiGetKafkasRequest {
	if mock.GetKafkasFunc == nil {
		panic("DefaultApiMock.GetKafkasFunc: method is nil but DefaultApi.GetKafkas was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetKafkas.Lock()
	mock.calls.GetKafkas = append(mock.calls.GetKafkas, callInfo)
	mock.lockGetKafkas.Unlock()
	return mock.GetKafkasFunc(ctx)
}

// GetKafkasCalls gets all the calls that were made to GetKafkas.
// Check the length with:
//     len(mockedDefaultApi.GetKafkasCalls())
func (mock *DefaultApiMock) GetKafkasCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetKafkas.RLock()
	calls = mock.calls.GetKafkas
	mock.lockGetKafkas.RUnlock()
	return calls
}

// GetKafkasExecute calls GetKafkasExecuteFunc.
func (mock *DefaultApiMock) GetKafkasExecute(r ApiGetKafkasRequest) (KafkaRequestList, *_nethttp.Response, error) {
	if mock.GetKafkasExecuteFunc == nil {
		panic("DefaultApiMock.GetKafkasExecuteFunc: method is nil but DefaultApi.GetKafkasExecute was just called")
	}
	callInfo := struct {
		R ApiGetKafkasRequest
	}{
		R: r,
	}
	mock.lockGetKafkasExecute.Lock()
	mock.calls.GetKafkasExecute = append(mock.calls.GetKafkasExecute, callInfo)
	mock.lockGetKafkasExecute.Unlock()
	return mock.GetKafkasExecuteFunc(r)
}

// GetKafkasExecuteCalls gets all the calls that were made to GetKafkasExecute.
// Check the length with:
//     len(mockedDefaultApi.GetKafkasExecuteCalls())
func (mock *DefaultApiMock) GetKafkasExecuteCalls() []struct {
	R ApiGetKafkasRequest
} {
	var calls []struct {
		R ApiGetKafkasRequest
	}
	mock.lockGetKafkasExecute.RLock()
	calls = mock.calls.GetKafkasExecute
	mock.lockGetKafkasExecute.RUnlock()
	return calls
}

// GetMetricsByInstantQuery calls GetMetricsByInstantQueryFunc.
func (mock *DefaultApiMock) GetMetricsByInstantQuery(ctx _context.Context, id string) ApiGetMetricsByInstantQueryRequest {
	if mock.GetMetricsByInstantQueryFunc == nil {
		panic("DefaultApiMock.GetMetricsByInstantQueryFunc: method is nil but DefaultApi.GetMetricsByInstantQuery was just called")
	}
	callInfo := struct {
		Ctx _context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetMetricsByInstantQuery.Lock()
	mock.calls.GetMetricsByInstantQuery = append(mock.calls.GetMetricsByInstantQuery, callInfo)
	mock.lockGetMetricsByInstantQuery.Unlock()
	return mock.GetMetricsByInstantQueryFunc(ctx, id)
}

// GetMetricsByInstantQueryCalls gets all the calls that were made to GetMetricsByInstantQuery.
// Check the length with:
//     len(mockedDefaultApi.GetMetricsByInstantQueryCalls())
func (mock *DefaultApiMock) GetMetricsByInstantQueryCalls() []struct {
	Ctx _context.Context
	ID  string
} {
	var calls []struct {
		Ctx _context.Context
		ID  string
	}
	mock.lockGetMetricsByInstantQuery.RLock()
	calls = mock.calls.GetMetricsByInstantQuery
	mock.lockGetMetricsByInstantQuery.RUnlock()
	return calls
}

// GetMetricsByInstantQueryExecute calls GetMetricsByInstantQueryExecuteFunc.
func (mock *DefaultApiMock) GetMetricsByInstantQueryExecute(r ApiGetMetricsByInstantQueryRequest) (MetricsInstantQueryList, *_nethttp.Response, error) {
	if mock.GetMetricsByInstantQueryExecuteFunc == nil {
		panic("DefaultApiMock.GetMetricsByInstantQueryExecuteFunc: method is nil but DefaultApi.GetMetricsByInstantQueryExecute was just called")
	}
	callInfo := struct {
		R ApiGetMetricsByInstantQueryRequest
	}{
		R: r,
	}
	mock.lockGetMetricsByInstantQueryExecute.Lock()
	mock.calls.GetMetricsByInstantQueryExecute = append(mock.calls.GetMetricsByInstantQueryExecute, callInfo)
	mock.lockGetMetricsByInstantQueryExecute.Unlock()
	return mock.GetMetricsByInstantQueryExecuteFunc(r)
}

// GetMetricsByInstantQueryExecuteCalls gets all the calls that were made to GetMetricsByInstantQueryExecute.
// Check the length with:
//     len(mockedDefaultApi.GetMetricsByInstantQueryExecuteCalls())
func (mock *DefaultApiMock) GetMetricsByInstantQueryExecuteCalls() []struct {
	R ApiGetMetricsByInstantQueryRequest
} {
	var calls []struct {
		R ApiGetMetricsByInstantQueryRequest
	}
	mock.lockGetMetricsByInstantQueryExecute.RLock()
	calls = mock.calls.GetMetricsByInstantQueryExecute
	mock.lockGetMetricsByInstantQueryExecute.RUnlock()
	return calls
}

// GetMetricsByRangeQuery calls GetMetricsByRangeQueryFunc.
func (mock *DefaultApiMock) GetMetricsByRangeQuery(ctx _context.Context, id string) ApiGetMetricsByRangeQueryRequest {
	if mock.GetMetricsByRangeQueryFunc == nil {
		panic("DefaultApiMock.GetMetricsByRangeQueryFunc: method is nil but DefaultApi.GetMetricsByRangeQuery was just called")
	}
	callInfo := struct {
		Ctx _context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetMetricsByRangeQuery.Lock()
	mock.calls.GetMetricsByRangeQuery = append(mock.calls.GetMetricsByRangeQuery, callInfo)
	mock.lockGetMetricsByRangeQuery.Unlock()
	return mock.GetMetricsByRangeQueryFunc(ctx, id)
}

// GetMetricsByRangeQueryCalls gets all the calls that were made to GetMetricsByRangeQuery.
// Check the length with:
//     len(mockedDefaultApi.GetMetricsByRangeQueryCalls())
func (mock *DefaultApiMock) GetMetricsByRangeQueryCalls() []struct {
	Ctx _context.Context
	ID  string
} {
	var calls []struct {
		Ctx _context.Context
		ID  string
	}
	mock.lockGetMetricsByRangeQuery.RLock()
	calls = mock.calls.GetMetricsByRangeQuery
	mock.lockGetMetricsByRangeQuery.RUnlock()
	return calls
}

// GetMetricsByRangeQueryExecute calls GetMetricsByRangeQueryExecuteFunc.
func (mock *DefaultApiMock) GetMetricsByRangeQueryExecute(r ApiGetMetricsByRangeQueryRequest) (MetricsRangeQueryList, *_nethttp.Response, error) {
	if mock.GetMetricsByRangeQueryExecuteFunc == nil {
		panic("DefaultApiMock.GetMetricsByRangeQueryExecuteFunc: method is nil but DefaultApi.GetMetricsByRangeQueryExecute was just called")
	}
	callInfo := struct {
		R ApiGetMetricsByRangeQueryRequest
	}{
		R: r,
	}
	mock.lockGetMetricsByRangeQueryExecute.Lock()
	mock.calls.GetMetricsByRangeQueryExecute = append(mock.calls.GetMetricsByRangeQueryExecute, callInfo)
	mock.lockGetMetricsByRangeQueryExecute.Unlock()
	return mock.GetMetricsByRangeQueryExecuteFunc(r)
}

// GetMetricsByRangeQueryExecuteCalls gets all the calls that were made to GetMetricsByRangeQueryExecute.
// Check the length with:
//     len(mockedDefaultApi.GetMetricsByRangeQueryExecuteCalls())
func (mock *DefaultApiMock) GetMetricsByRangeQueryExecuteCalls() []struct {
	R ApiGetMetricsByRangeQueryRequest
} {
	var calls []struct {
		R ApiGetMetricsByRangeQueryRequest
	}
	mock.lockGetMetricsByRangeQueryExecute.RLock()
	calls = mock.calls.GetMetricsByRangeQueryExecute
	mock.lockGetMetricsByRangeQueryExecute.RUnlock()
	return calls
}

// GetServiceStatus calls GetServiceStatusFunc.
func (mock *DefaultApiMock) GetServiceStatus(ctx _context.Context) ApiGetServiceStatusRequest {
	if mock.GetServiceStatusFunc == nil {
		panic("DefaultApiMock.GetServiceStatusFunc: method is nil but DefaultApi.GetServiceStatus was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetServiceStatus.Lock()
	mock.calls.GetServiceStatus = append(mock.calls.GetServiceStatus, callInfo)
	mock.lockGetServiceStatus.Unlock()
	return mock.GetServiceStatusFunc(ctx)
}

// GetServiceStatusCalls gets all the calls that were made to GetServiceStatus.
// Check the length with:
//     len(mockedDefaultApi.GetServiceStatusCalls())
func (mock *DefaultApiMock) GetServiceStatusCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetServiceStatus.RLock()
	calls = mock.calls.GetServiceStatus
	mock.lockGetServiceStatus.RUnlock()
	return calls
}

// GetServiceStatusExecute calls GetServiceStatusExecuteFunc.
func (mock *DefaultApiMock) GetServiceStatusExecute(r ApiGetServiceStatusRequest) (ServiceStatus, *_nethttp.Response, error) {
	if mock.GetServiceStatusExecuteFunc == nil {
		panic("DefaultApiMock.GetServiceStatusExecuteFunc: method is nil but DefaultApi.GetServiceStatusExecute was just called")
	}
	callInfo := struct {
		R ApiGetServiceStatusRequest
	}{
		R: r,
	}
	mock.lockGetServiceStatusExecute.Lock()
	mock.calls.GetServiceStatusExecute = append(mock.calls.GetServiceStatusExecute, callInfo)
	mock.lockGetServiceStatusExecute.Unlock()
	return mock.GetServiceStatusExecuteFunc(r)
}

// GetServiceStatusExecuteCalls gets all the calls that were made to GetServiceStatusExecute.
// Check the length with:
//     len(mockedDefaultApi.GetServiceStatusExecuteCalls())
func (mock *DefaultApiMock) GetServiceStatusExecuteCalls() []struct {
	R ApiGetServiceStatusRequest
} {
	var calls []struct {
		R ApiGetServiceStatusRequest
	}
	mock.lockGetServiceStatusExecute.RLock()
	calls = mock.calls.GetServiceStatusExecute
	mock.lockGetServiceStatusExecute.RUnlock()
	return calls
}

// GetVersionMetadata calls GetVersionMetadataFunc.
func (mock *DefaultApiMock) GetVersionMetadata(ctx _context.Context) ApiGetVersionMetadataRequest {
	if mock.GetVersionMetadataFunc == nil {
		panic("DefaultApiMock.GetVersionMetadataFunc: method is nil but DefaultApi.GetVersionMetadata was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetVersionMetadata.Lock()
	mock.calls.GetVersionMetadata = append(mock.calls.GetVersionMetadata, callInfo)
	mock.lockGetVersionMetadata.Unlock()
	return mock.GetVersionMetadataFunc(ctx)
}

// GetVersionMetadataCalls gets all the calls that were made to GetVersionMetadata.
// Check the length with:
//     len(mockedDefaultApi.GetVersionMetadataCalls())
func (mock *DefaultApiMock) GetVersionMetadataCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetVersionMetadata.RLock()
	calls = mock.calls.GetVersionMetadata
	mock.lockGetVersionMetadata.RUnlock()
	return calls
}

// GetVersionMetadataExecute calls GetVersionMetadataExecuteFunc.
func (mock *DefaultApiMock) GetVersionMetadataExecute(r ApiGetVersionMetadataRequest) (VersionMetadata, *_nethttp.Response, error) {
	if mock.GetVersionMetadataExecuteFunc == nil {
		panic("DefaultApiMock.GetVersionMetadataExecuteFunc: method is nil but DefaultApi.GetVersionMetadataExecute was just called")
	}
	callInfo := struct {
		R ApiGetVersionMetadataRequest
	}{
		R: r,
	}
	mock.lockGetVersionMetadataExecute.Lock()
	mock.calls.GetVersionMetadataExecute = append(mock.calls.GetVersionMetadataExecute, callInfo)
	mock.lockGetVersionMetadataExecute.Unlock()
	return mock.GetVersionMetadataExecuteFunc(r)
}

// GetVersionMetadataExecuteCalls gets all the calls that were made to GetVersionMetadataExecute.
// Check the length with:
//     len(mockedDefaultApi.GetVersionMetadataExecuteCalls())
func (mock *DefaultApiMock) GetVersionMetadataExecuteCalls() []struct {
	R ApiGetVersionMetadataRequest
} {
	var calls []struct {
		R ApiGetVersionMetadataRequest
	}
	mock.lockGetVersionMetadataExecute.RLock()
	calls = mock.calls.GetVersionMetadataExecute
	mock.lockGetVersionMetadataExecute.RUnlock()
	return calls
}
