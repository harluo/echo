package kernel

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/harluo/echo/internal/internal/constant"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewContextUsesRequestLifecycle(t *testing.T) {
	requestContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := httptest.NewRequest("GET", "/", nil).WithContext(requestContext)
	source := echo.New().NewContext(request, httptest.NewRecorder())
	actual := NewContext(source)

	assert.NoError(t, actual.Err())
	cancel()
	assert.ErrorIs(t, actual.Err(), context.Canceled)

	select {
	case <-actual.Done():
	default:
		t.Fatal("请求取消后业务上下文未结束")
	}
}

func TestNewContextWithoutRequest(t *testing.T) {
	actual := NewContext(nil)

	assert.Nil(t, actual.Echo())
	assert.NoError(t, actual.Err())
}

func TestContextDerivesValuesFromRequestContext(t *testing.T) {
	actual := NewContext(nil).With("key", "value")

	assert.Equal(t, "value", actual.Value("key"))
	actual.Unresponsive()
	assert.Equal(t, false, actual.Value(constant.ContextResponse))
}

func TestContextRequestCancellationIsObservable(t *testing.T) {
	requestContext, cancel := context.WithCancel(context.Background())
	request := httptest.NewRequest("GET", "/", nil).WithContext(requestContext)
	actual := NewContext(echo.New().NewContext(request, httptest.NewRecorder()))

	cancel()
	assert.True(t, errors.Is(actual.Err(), context.Canceled))
}
