// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"context"
	metricsapi "github.com/keptn/lifecycle-toolkit/metrics-operator/api/v1"
	"net/http"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sync"
)

// IRoundTripperMock is a mock implementation of prometheus.IRoundTripper.
//
//	func TestSomethingThatUsesIRoundTripper(t *testing.T) {
//
//		// make and configure a mocked prometheus.IRoundTripper
//		mockedIRoundTripper := &IRoundTripperMock{
//			GetRoundTripperFunc: func(contextMoqParam context.Context, keptnMetricsProvider metricsapi.KeptnMetricsProvider, clientMoqParam client.Client) (http.RoundTripper, error) {
//				panic("mock out the GetRoundTripper method")
//			},
//		}
//
//		// use mockedIRoundTripper in code that requires prometheus.IRoundTripper
//		// and then make assertions.
//
//	}
type IRoundTripperMock struct {
	// GetRoundTripperFunc mocks the GetRoundTripper method.
	GetRoundTripperFunc func(contextMoqParam context.Context, keptnMetricsProvider metricsapi.KeptnMetricsProvider, clientMoqParam client.Client) (http.RoundTripper, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetRoundTripper holds details about calls to the GetRoundTripper method.
		GetRoundTripper []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// KeptnMetricsProvider is the keptnMetricsProvider argument value.
			KeptnMetricsProvider metricsapi.KeptnMetricsProvider
			// ClientMoqParam is the clientMoqParam argument value.
			ClientMoqParam client.Client
		}
	}
	lockGetRoundTripper sync.RWMutex
}

// GetRoundTripper calls GetRoundTripperFunc.
func (mock *IRoundTripperMock) GetRoundTripper(contextMoqParam context.Context, keptnMetricsProvider metricsapi.KeptnMetricsProvider, clientMoqParam client.Client) (http.RoundTripper, error) {
	if mock.GetRoundTripperFunc == nil {
		panic("IRoundTripperMock.GetRoundTripperFunc: method is nil but IRoundTripper.GetRoundTripper was just called")
	}
	callInfo := struct {
		ContextMoqParam      context.Context
		KeptnMetricsProvider metricsapi.KeptnMetricsProvider
		ClientMoqParam       client.Client
	}{
		ContextMoqParam:      contextMoqParam,
		KeptnMetricsProvider: keptnMetricsProvider,
		ClientMoqParam:       clientMoqParam,
	}
	mock.lockGetRoundTripper.Lock()
	mock.calls.GetRoundTripper = append(mock.calls.GetRoundTripper, callInfo)
	mock.lockGetRoundTripper.Unlock()
	return mock.GetRoundTripperFunc(contextMoqParam, keptnMetricsProvider, clientMoqParam)
}

// GetRoundTripperCalls gets all the calls that were made to GetRoundTripper.
// Check the length with:
//
//	len(mockedIRoundTripper.GetRoundTripperCalls())
func (mock *IRoundTripperMock) GetRoundTripperCalls() []struct {
	ContextMoqParam      context.Context
	KeptnMetricsProvider metricsapi.KeptnMetricsProvider
	ClientMoqParam       client.Client
} {
	var calls []struct {
		ContextMoqParam      context.Context
		KeptnMetricsProvider metricsapi.KeptnMetricsProvider
		ClientMoqParam       client.Client
	}
	mock.lockGetRoundTripper.RLock()
	calls = mock.calls.GetRoundTripper
	mock.lockGetRoundTripper.RUnlock()
	return calls
}
